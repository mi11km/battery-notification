package charge

import (
	"log"

	"github.com/distatus/battery"
)

func Fetch(batteryCharges chan float64) {
	batteries, err := battery.GetAll()
	if err != nil {
		log.Printf("action=fetch battery info, err=%v\n", err)
		return
	}
	for _, b := range batteries {
		batteryCharges <- (b.Current / b.Full) * 100
	}
}
