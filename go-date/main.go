package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func showWeekDay(d int) string {
	t := time.Now()
	ft := t.AddDate(0, 0, d)
	return ft.Weekday().String()
}

func checkMonth(d int) bool {
	t := time.Now()
	ft := t.AddDate(0, 0, d)

	oldMonth := t.Month()
	newMonth := ft.Month()

	return oldMonth == newMonth
}

func main() {

	var addDay int
	var rDay string

	check := false

	flag.IntVar(&addDay, "c", 0, "welcher Wochtentag ist in (x) Tage")
	flag.BoolVar(&check, "m", check, "check is the addDay month")
	flag.Parse()

	if addDay > 0 {
		rDay = showWeekDay(addDay)
	}

	if addDay > 0 && check {
		fmt.Println(checkMonth(addDay))
		if checkMonth(addDay) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

	if addDay > 0 && !check {
		fmt.Printf("in %d tage haben wir den Wochentag %s\n", addDay, rDay)
	}

}
