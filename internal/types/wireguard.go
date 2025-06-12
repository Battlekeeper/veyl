package types

import "golang.zx2c4.com/wireguard/wgctrl/wgtypes"

type WgPeer struct {
	PublicKey  wgtypes.Key
	Endpoint   string
	AllowedIps []string
}

type WgConfig struct {
	PrivateKey wgtypes.Key
	ListenPort int
	ListenIP   string
	Peers      []WgPeer
}
