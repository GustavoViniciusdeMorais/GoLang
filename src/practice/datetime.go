package main

import (
	"fmt"
	"time"
)

func main() {
	var now = time.Now()
	fmt.Println("Date is ", now)

	formattedDate := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("Go release at ", formattedDate)

	nowTime := time.Now()
	nowTimeFormatted := fmt.Sprintf("Now time is %d-%d-%02d \n", nowTime.Year(), nowTime.Month(), nowTime.Day())
	fmt.Println(nowTimeFormatted)
}
