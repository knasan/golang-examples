package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "/dev/sdb1"
	str2 := "/dev/sdb"

	last := string(str[len(str)-1:])
	last2 := string(str2[len(str2)-1:])

	fmt.Printf("last characters %s - from %s\n", last, str)
	fmt.Printf("last characters %s - from %s\n", last2, str2)

	if v, err := strconv.Atoi(last); err != nil {
		fmt.Printf("no number found at the end of the string: %s\n", str)
	} else {
		fmt.Println(v)
	}

	if v, err := strconv.Atoi(last2); err != nil {
		fmt.Printf("no number found at the end of the string: %s\n", str2)
	} else {
		fmt.Println(v)
	}

}
