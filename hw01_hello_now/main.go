package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

const ntpHost = "0.beevik-ntp.pool.ntp.org"
const timeLayout = "2006-01-02 15:04:05 -0700 MST"

func main() {
	exactTime, err := ntp.Time(ntpHost)
	if err != nil {
		log.Fatal(err)
	}
	currentTime := time.Now()

	fmt.Printf("current time: %s\n", currentTime.Format(timeLayout))
	fmt.Printf("exact time: %s\n", exactTime.Format(timeLayout))
}
