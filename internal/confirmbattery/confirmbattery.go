package confirmbattery

import (
	"battery-notification/internal/charge"
	"battery-notification/internal/slack"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/wailsapp/wails"
)

type ConfirmBattery struct {
	IsLoop         bool
	BatteryCharges chan float64
	runtime        *wails.Runtime
	logger         *wails.CustomLogger
}

func New() (*ConfirmBattery, error) {
	return &ConfirmBattery{IsLoop: false, BatteryCharges: make(chan float64)}, nil
}

func (c *ConfirmBattery) WailsInit(runtime *wails.Runtime) error {
	c.runtime = runtime
	c.logger = c.runtime.Log.New("ConfirmBattery")
	return nil
}

func (c *ConfirmBattery) Stop() {
	c.IsLoop = true
	c.logger.Infof("Stop Loop, IsLoop: %b", c.IsLoop)
}

func (c *ConfirmBattery) Start(name, webhookURL string, confirmationInterval, notificationCondition float64) {
	c.logger.Info("Start Loop")

	var wg sync.WaitGroup
	c.IsLoop = false
	go func() {
		for {
			wg.Add(1)
			charge.Fetch(c.BatteryCharges)
			time.Sleep(time.Second * time.Duration(confirmationInterval))
			if c.IsLoop {
				break
			}
		}
	}()

	for battery := range c.BatteryCharges {
		c.logger.Infof("現在のバッテリー： %f ％", battery)
		if battery < notificationCondition {
			text := fmt.Sprintf("PC：%sのバッテリーが%f以下になりました。現在は%f％です。充電してください！",
				name, notificationCondition, battery)
			message := slack.NewMessage(name, text, ":battery:")
			if err := message.Notify(webhookURL); err != nil {
				log.Printf("action=send message to slack, err=%v\n", err)
			}
		}
		wg.Done()
	}
	wg.Wait()
	close(c.BatteryCharges)
	c.logger.Info("Close Battery Stored Channel")
}
