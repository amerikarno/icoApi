package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	t1 := 5
	test := fmt.Sprintf("test = %d%%", t1)
	log.Println(test)
	log.Println(convertDateInThai(time.Now().Local()))
}

func convertDateInThai(d time.Time) (thaiDate string) {
	var monthStr string
	date := d.Day()
	month := int(d.Month())
	year := d.Year() + 543
	switch month {
	case 1:
		monthStr = "มกราคม"
	case 2:
		monthStr = "กุมภาพันธ์"
	case 3:
		monthStr = "มีนาคม"
	case 4:
		monthStr = "เมษายน"
	case 5:
		monthStr = "พฤษภาคม"
	case 6:
		monthStr = "มิถุนายม"
	case 7:
		monthStr = "กรกฎาคม"
	case 8:
		monthStr = "สิงหาคม"
	case 9:
		monthStr = "กันยายน"
	case 10:
		monthStr = "ตุลาคม"
	case 11:
		monthStr = "พฤษจิกายน"
	case 12:
		monthStr = "ธันวาคม"
	}
	return fmt.Sprintf("%d %s %d", date, monthStr, year)
}
