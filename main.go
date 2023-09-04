package main

import (
	"context"
	"fmt"
	"log"

	"github.com/huin/goupnp"
)

func main() {
	ctx := context.Background()
	log.Println("searching for devices")
	find(ctx)
}

func find(ctx context.Context) (string, error) {
	devices, err := goupnp.DiscoverDevicesCtx(ctx, "urn:schemas-upnp-org:device:MediaRenderer:1")
	if err != nil {
		panic(err)
	}

	for _, device := range devices {
		if device.Err != nil {
			log.Println(device.Err)
		}

		if device.Root == nil {
			log.Println("no root")
			continue
		}

		log.Println(device.Root.Device.FriendlyName)

		if device.Location == nil {
			log.Println("no location for", device.Root.Device.FriendlyName)
		}

		return device.Location.String(), nil
	}

	return "", fmt.Errorf("no tv found")
}
