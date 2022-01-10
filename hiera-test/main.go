package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/knasan/structs" // forzen library (readonly) - project is cool, werde ich übernehmen und bei fehler weiterpflegen! // TODO: project ist geforkt, code anlalysieren und dann ggf. weiterpflegen
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"io"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/lyraproj/dgo/vf"
	"github.com/lyraproj/hiera/api"
	"github.com/lyraproj/hiera/hiera"
	"github.com/lyraproj/hiera/provider"
	sdk "github.com/lyraproj/hierasdk/hiera"
)

const (
	MaxSnapshots           = 20
	MinSpaceLeft           = "50%"
	DefaultCfgPath         = "./examples"
	FilePermission         = 0644 // FileExecutePermission = 0776
	DirectoryPermission    = 0776
	ApplicationFileName    = "application.yaml"
	ApplicationName        = "hiera-test"
	SyslogAddress          = "127.0.0.1:514"
	PupptFileName          = "puppet.yaml"
	workaroundYamlTempPath = "/tmp/hiera-test"
)

// Config
type Config struct {
	Destination  string          `json:"destination,omitempty" yaml:"destination,omitempty"`
	MaxWorker    int             `json:"maxWorker" yaml:"maxWorker,omitempty"`
	Hosts        map[string]Host `json:"hosts,omitempty" yaml:"hosts,omitempty"`
	Devices      []Device        `json:"devices,omitempty" yaml:"devices,omitempty"`
	RsyncOptions RsyncOptions    `json:"rsyncOptions" yaml:"rsyncOptions,omitempty"`
}

// RsyncOptions
type RsyncOptions struct {
	Options           []string `json:"options,omitempty" yaml:"options,omitempty"`
	ContinueOnError   bool     `json:"continueOnError,omitempty" yaml:"continueOnError,omitempty"`
	IgnoreStatusCodes []int    `json:"ignoreStatusCodes,omitempty" yaml:"ignoreStatusCodes,omitempty"`
}

// Host
type Host struct {
	Address string            `json:"address,omitempty" yaml:"address,omitempty"`
	Port    int               `json:"port,omitempty" yaml:"port,omitempty"`
	Rsyncd  bool              `json:"rsyncd,omitempty" yaml:"rsyncd,omitempty"`
	Plugins map[string]Plugin `json:"plugins" yaml:"plugins"`
}

// Plugin
type Plugin struct {
	Include  string   `json:"include" yaml:"include"`
	Exclude  []string `json:"exclude,omitempty" yaml:"exclude,omitempty"`
	Commands Commands `json:"commands,omitempty" yaml:"commands,omitempty"`
}

// Commands
type Commands struct {
	Before []Cmd `json:"before,omitempty" yaml:"before,omitempty"`
	After  []Cmd `json:"after,omitempty" yaml:"after,omitempty"`
}

// Cmd
type Cmd struct {
	File                 string   `json:"file,omitempty" yaml:"file,omitempty"`
	RunAs                string   `json:"runas,omitempty" yaml:"runas,omitempty"`
	Location             string   `json:"location,omitempty" yaml:"location,omitempty"`
	ScriptBlock          string   `json:"scriptBlock,omitempty" yaml:"scriptBlock,omitempty"`
	AdditionalRessources []string `json:"additionalRessources,omitempty" yaml:"additionalRessources,omitempty"`
	Args                 []string `json:"args,omitempty" yaml:"args,omitempty"`
}

// Device
type Device struct {
	Label        string `json:"label,omitempty" yaml:"label,omitempty"`
	Fstype       string `json:"fstype,omitempty" yaml:"fstype,omitempty"`
	Encrypt      bool   `json:"encrypt,omitempty" yaml:"encrypt,omitempty"`
	MinSpaceLeft string `json:"minSpaceLeft,omitempty" yaml:"minSpaceLeft,omitempty"`
	MaxSnapshots int    `json:"maxSnapshots,omitempty" yaml:"maxSnapshots,omitempty"`
}

// hieraConfig
type hieraConfig struct {
	Hierarchy []hierarchy
}

// hierarchy
type hierarchy struct {
	Name string
	Path string
}

// MakeFirstCharToLowerCase hiera helper, find the result when "MaxWorker => maxWorker, RsyncOptions => ryncOptions"
func makeFirstCharToLowerCase(s string) string {
	log.Println("makeFirstCharToLowerCase")
	if len(s) < 2 {
		return strings.ToLower(s)
	}

	bts := []byte(s)
	lc := bytes.ToLower([]byte{bts[0]})
	rest := bts[1:]

	return string(bytes.Join([][]byte{lc, rest}, nil))
}

// prepareString trim strings
func prepareString(s string) string {
	log.Println("prepareString")
	s = strings.TrimSpace(s)
	s = strings.Trim(s, "\n\"")
	return s
}

// configHelper,
func (c *Config) configHelper(section string) (err error) {
	log.Println("configHelper")
	section = makeFirstCharToLowerCase(section)

	var cmdOpts hiera.CommandOptions
	cmdOpts.RenderAs = "json"

	cmdOpts.Default = nil

	cfgOpts := vf.MutableMap()
	cfgOpts.Put(api.HieraRoot, workaroundYamlTempPath)
	cfgOpts.Put(api.HieraDialect, "pcore")
	cfgOpts.Put(provider.LookupKeyFunctions, []sdk.LookupKey{provider.ConfigLookupKey, provider.Environment})

	buf := bytes.Buffer{}
	args := []string{section}

	sectionlow := strings.ToLower(section)

	err = hiera.TryWithParent(context.Background(), provider.MuxLookupKey, cfgOpts, func(a api.Session) (err error) {
		ok := hiera.LookupAndRender(a, &cmdOpts, args, &buf)
		if !ok {
			msg := fmt.Sprintf("lookupAndRender failed: %v: ", section)
			if _, err = os.Stat(filepath.Join(DefaultCfgPath, ApplicationFileName)); err != nil {
				return nil
			}
			log.Error(msg)
			err = fmt.Errorf(msg)
		}
		return err
	})

	if err != nil {
		log.Error(err)
		fmt.Println(err)
		return err
	}

	switch sectionlow {
	case "destination":
		c.Destination = prepareString(buf.String())
	case "maxworker":
		strNumber := strings.Split(prepareString(buf.String()), "\n")[0]
		var maxint int
		maxint, err = strconv.Atoi(strNumber)
		if err != nil {
			log.Error(err)
			return err
		}
		c.MaxWorker = maxint
	case "devices":
		nc := new(Config)
		err = yaml.Unmarshal(buf.Bytes(), &nc.Devices)
		if err != nil {
			log.Error(err)
			return err
		}
		d := Device{}
		for _, dev := range nc.Devices {
			d.Label = prepareString(dev.Label)
			d.Fstype = prepareString(dev.Fstype)
			d.Encrypt = dev.Encrypt
			d.MaxSnapshots = dev.MaxSnapshots
			d.MinSpaceLeft = prepareString(dev.MinSpaceLeft)
			c.Devices = append(c.Devices, d)
		}
	case "rsyncoptions":
		nc := new(Config)
		err = yaml.Unmarshal(buf.Bytes(), &nc.RsyncOptions)
		if err != nil {
			log.Error(err)
			return err
		}
		c.RsyncOptions = nc.RsyncOptions
	case "hosts":
		err = yaml.Unmarshal(buf.Bytes(), &c.Hosts)
		if err != nil {
			log.Error(err)
			fmt.Printf("%#v\n", c.Hosts)
			return err
		}
	default:
		err = yaml.Unmarshal(buf.Bytes(), c)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return err
}

// Load config with hiera
func (c *Config) Load() (err error) {
	log.Println("Load")
	err = workaround()
	if err != nil {
		return err
	}
	for _, sections := range structs.Names(c) {
		err = c.configHelper(sections)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

// workaround - here you can't get along with yaml anchor, so every yaml file has to be loaded and written to the structure once.
// This is how the anchors are dissolved
func workaround() error {
	log.Println("workaround")
	var hc hieraConfig
	b, err := ioutil.ReadFile(filepath.Join(DefaultCfgPath, "hiera.yaml"))
	if err != nil {
		log.Error(err)
		return err
	}
	if err = yaml.Unmarshal(b, &hc); err != nil {
		log.Error(err)
		return err
	}

	if _, err = os.Stat(workaroundYamlTempPath); err == nil {
		if err = os.RemoveAll(workaroundYamlTempPath); err != nil {
			log.Error(err)
			return err
		}
	}

	if err = os.MkdirAll(workaroundYamlTempPath, DirectoryPermission); err != nil {
		log.Error(err)
		return err
	}

	for _, p := range hc.Hierarchy {
		if _, err = os.Stat(filepath.Join(DefaultCfgPath, p.Path)); err != nil {
			continue
		}
		if err = writeWorker(p.Path); err != nil {
			return err
		}
	}

	// copy the hiera file to workaroundYamlTempPath
	source, err := os.Open(filepath.Join(DefaultCfgPath, "hiera.yaml"))
	if err != nil {
		log.Error()
		return err
	}
	defer source.Close()

	dst, err := os.Create(filepath.Join(workaroundYamlTempPath, "hiera.yaml"))
	if err != nil {
		log.Error(err)
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, source); err != nil {
		log.Error(err)
		return err
	}

	return err
}

func writeWorker(filename string) error {
	log.Println("writeWorker")
	// d := new(Config) // wenn mapper dann ist __type, __value usw. in byte drin.
	d := make(map[string]interface{})

	file := filepath.Join(DefaultCfgPath, filename)

	b, err := ioutil.ReadFile(file)
	if err != nil {
		log.Error(err)
		return err
	}

	if err = yaml.Unmarshal(b, &d); err != nil {
		erm := fmt.Errorf("%v in: %s", err, filename)
		log.Error(erm)
		return erm
	}

	m, err := yaml.Marshal(&d)
	if err != nil {
		log.Error(err)
		return err
	}

	filecontent := []byte("---\n")
	filecontent = append(filecontent, m...)

	if err = ioutil.WriteFile(filepath.Join(workaroundYamlTempPath, filename), filecontent, FilePermission); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

// main
func main() {
	cfg := Config{}
	if err := cfg.Load(); err != nil {
		panic(err)
	}
	fmt.Printf("Hosts: %v\n", cfg.Hosts)
}
