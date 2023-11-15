package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/huin/goupnp"
)

func main() {
	ctx := context.Background()
	slog.Info("searching for devices")
	if deviceURL, err := findFirstRootDevice(ctx); err != nil {
		slog.Error("find devices", "err", err)
	} else {
		slog.Info("found device", "device", deviceURL.String())
	}
}

func findFirstRootDevice(ctx context.Context) (*url.URL, error) {
	devices, err := goupnp.DiscoverDevicesCtx(ctx, "urn:schemas-upnp-org:device:MediaRenderer:1")
	if err != nil {
		panic(err)
	}

	for _, device := range devices {
		if device.Err != nil {
			slog.Warn("get device", "err", err)
			continue
		}

		if device.Root == nil {
			slog.Warn("device is not root device, skipping", "device", device)
			continue
		}

		if device.Location == nil {
			slog.Warn("device has no location, skipping", "device", device)
			continue
		}

		return device.Location, nil
	}

	return nil, fmt.Errorf("no tv found")
}
