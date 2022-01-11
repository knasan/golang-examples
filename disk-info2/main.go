package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Info Struct for lsblk
type Info struct {
	Blockdevices []Blockdevice
}

type Child struct {
	Children []Blockdevice
}

type Blockdevice struct {
	Name        string `json:"name"`
	Fstype      string `json:"fstype"`
	Label       string `json:"label"`
	Size        string `json:"size"`
	MountPoint  string `json:"mountpoint"`
	CryptDevice string `json:"-"`
	Child
}

func loadJson() Info {
	// lsblk -I 8,179,252,259 -f -J -o fstype,name,label,size,mountpoint
	jsonFile, err := os.Open("disk2.json")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Json File opened")
	defer jsonFile.Close()

	// to byte Array
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	// initialize sturct
	var diskinfo Info

	if err = json.Unmarshal(byteValue, &diskinfo); err != nil {
		log.Fatal(err)
	}
	return diskinfo
}

func main() {
	info := loadJson()
	// fmt.Printf("%+v\n", info)
	di := info.deviceInfo("/dev/sdc")
	fmt.Printf("%+v\n", di)
}

func (i Info) deviceInfo(device string) Blockdevice {
	var blockdevice Blockdevice
	for _, bd := range i.Blockdevices {
		if bd.Name == device {
			blockdevice.Name = bd.Name
			blockdevice.Label = bd.Label
			blockdevice.Size = bd.Size
			blockdevice.Fstype = bd.Fstype
			blockdevice.MountPoint = bd.MountPoint
			for _, bdc := range bd.Children {
				blockdevice.Name = bdc.Name
				blockdevice.Label = bdc.Label
				blockdevice.Size = bdc.Size
				blockdevice.Fstype = bdc.Fstype
				blockdevice.MountPoint = bdc.MountPoint
				if bdc.Fstype == "crypto_LUKS" {
					for _, bdc1 := range bdc.Children {
						blockdevice.CryptDevice = bdc1.Name
						blockdevice.Label = bdc1.Label
						blockdevice.Size = bdc1.Size
						blockdevice.Fstype = bdc1.Fstype
						blockdevice.MountPoint = bdc1.MountPoint
					}
				}
			}
		}

	}

	return blockdevice
}
