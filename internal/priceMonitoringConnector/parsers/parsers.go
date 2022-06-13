package parsers

import (
	"fmt"
	"time"
)

func ParseDate(stringDate string) time.Time {
	//Mon Jun 13 01:03:13 MSK 2022
	date, _ := time.Parse("Mon Jan 2 15:04:05 MST 2006", stringDate)
	fmt.Println(date)
	return date
}
