package main

import (
	"flag"
	"log/slog"

	"golang.org/x/net/websocket"
)

func main() {
	websocketEndpoint := flag.String("ws", "ws://localhost:3000", "tv websocket endpoint")
	flag.Parse()

	slog.Info("connecting to websocket", "endpoint", *websocketEndpoint)
	conn, err := websocket.Dial(*websocketEndpoint, "", "http://192.168.1.241:3001/")
	if err != nil {
		slog.Error("dial", "err", err)
		return
	}
	defer conn.Close()

	slog.Info("registering")
	_, err = conn.Write([]byte(`{"msg", "test"}`))
	if err != nil {
		slog.Error("dial", "err", err)
		return
	}

	slog.Info("reading response")
	buf := make([]byte, 10240)
	_, err = conn.Read(buf)
	if err != nil {
		slog.Error("dial", "err", err)
		return
	}
	slog.Info("read", "buf", string(buf))

	// if deviceURL, err := findFirstRootDevice(ctx); err != nil {
	// 	slog.Error("find devices", "err", err)
	// } else {
	// 	slog.Info("found device", "device", deviceURL.String())
	// }
}

// func findTV(ctx context.Context) (*url.URL, error) {
// 	devices, err := goupnp.DiscoverDevicesCtx(ctx, "ssdp:all")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	for _, device := range devices {
// 		slog.Info("device found", "device", device)
// 		slog.Info("device info", "usn", device.USN, "root", device.Root, "localaddr", device.LocalAddr, "location", device.Location)
// 		if device.Root != nil {
// 			slog.Info("has root", "serial", device.Root.Device.SerialNumber)
// 		}
// 		// if device.Err != nil {
// 		// 	slog.Warn("get device", "err", err)
// 		// 	continue
// 		// }
// 		//
// 		// if device.Root == nil {
// 		// 	slog.Warn("device is not root device, skipping", "device", device)
// 		// 	continue
// 		// }
// 		//
// 		// if device.Location == nil {
// 		// 	slog.Warn("device has no location, skipping", "device", device)
// 		// 	continue
// 		// }
// 		//
// 		// if !strings.EqualFold(device.Root.Device.SerialNumber, "oled48cx6lb") {
// 		// 	slog.Warn("device serial is not the one we want, skipping", "device", device)
// 		// 	continue
// 		// }
// 		//
// 		// return device.Location, nil
// 	}
//
// 	return nil, fmt.Errorf("no tv found")
// }
