//go:build softdevice

package main

import (
	"time"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	must("enable BLE stack", adapter.Enable())

	deviceName := make([]byte, 16)
	must("get SD name", bluetooth.DeviceName(deviceName))
	println("SoftDevice GAP Device Name:", deviceName)

	advertisedName := "Go Bluetooth"
	adv := adapter.DefaultAdvertisement()
	must("config adv", adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: advertisedName,
	}))
	must("start adv", adv.Start())

	println("advertising...")
	address, _ := adapter.Address()
	for {
		println("GAP Device:", deviceName, "is advertising wth LocalName:", advertisedName, "with address", address.MAC.String())
		time.Sleep(time.Second)
	}
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
