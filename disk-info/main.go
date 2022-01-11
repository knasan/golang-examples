package main

import (
	"fmt"

	"github.com/jaypipes/ghw"
)

func main() {
	block, err := ghw.Block()
	if err != nil {
		fmt.Printf("Error getting block storage info: %v", err)
	}

	//fmt.Printf("Block: %v\n", block)

	for _, disk := range block.Disks {
		//fmt.Printf("Disk: %+v\n", disk)
		for _, part := range disk.Partitions {
			fmt.Printf("Partitions: %+v\n", part)
		}
	}

}
