package main

import (
	"fmt"
	"regexp"
)

func main() {
	var re = regexp.MustCompile(`name:\w*`)
	var str = `{
		{
		  name:
			system,
		  exclude:  {
					  /proc/,
					  /sys,
					  /dev,
					  /var/log/*.gz,
					  /var/log/*.old,
					  /var/log/*.[0-9],
					  /var/log/**/*.gz,
					  /var/log/**/*.old,
					  /var/log/**/*.[0-9],
					  /var/lib/mlocate/mlocate.db
					},
		  include:  {
					  /,
					  myinclude2
					},
		  command:  {
					  {
						cmd:
						  mycommandCustom,
							options: {
							  arg1,
							  arg2,
							  arg3
							},
							before:true
						},
						{
						  cmd:
							mycommand2Custom,
							options:  {
							  arg1,
							  arg2,
							  arg3
							},
							before:false
						}
					  }
		},
		{
		  name:
			lxc,
		  exclude: {
			/var/lib/lxc,
			/var/lib/lxc.data
		  }
		}
	  }
`

	for i, match := range re.FindAllString(str, -1) {
		fmt.Println(match, "found at index", i)
	}
}
