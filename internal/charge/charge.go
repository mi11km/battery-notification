package charge

import (
	"log"

	"github.com/distatus/battery"
)

func Fetch(batteryInfo chan *battery.Battery) {
	batteries, err := battery.GetAll()
	if err != nil {
		log.Printf("action=fetch battery info, err=%v\n", err)
		return
	}
	for _, b := range batteries {
		batteryInfo <- b
	}
}
