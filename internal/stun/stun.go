package stun

import (
	"net"

	stun "github.com/pion/stun"
)

func FetchStun() (net.IP, int, error) {
	// Connect to a public STUN server
	conn, err := stun.Dial("udp", "stun.l.google.com:19302")
	if err != nil {
		return nil, 0, err
	}
	defer conn.Close()

	var xorAddr stun.XORMappedAddress
	// Build and send a binding request
	message := stun.MustBuild(stun.TransactionID, stun.BindingRequest)

	err = conn.Do(message, func(res stun.Event) {
		if res.Error != nil {
			return
		}
		if err := xorAddr.GetFrom(res.Message); err != nil {
			return
		}
	})

	if err != nil {
		return nil, 0, err
	}

	return xorAddr.IP, xorAddr.Port, nil
}
