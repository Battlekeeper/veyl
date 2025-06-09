package veil

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type RelayAuth struct {
	RelayID   string `json:"relayid"`
	PublicKey string `json:"public_key"`
	IP        string `json:"ip"`
	Port      int    `json:"port"`
}

type RelayClient struct {
	Auth       RelayAuth       `json:"auth"`
	Connection *websocket.Conn `json:"-"`
}

type RelayConnection struct {
	RelayID   string `json:"relayid"`
	PublicKey string `json:"public_key"`
	IP        string `json:"ip"`
	Port      int    `json:"port"`
}

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

// Decode JSON body to veil.RelayAuth
func DecodeRelayAuth(body []byte) (*RelayAuth, error) {
	var auth RelayAuth
	err := json.Unmarshal(body, &auth)
	if err != nil {
		return nil, err
	}
	return &auth, nil
}

func Base64ToHex(base64Key string) string {
	decodedKey, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		log.Panic("Failed to decode base64 key:", err)
	}
	hexKey := hex.EncodeToString(decodedKey)
	return hexKey
}

func HexToBase64(hexKey string) (string, error) {
	decodedKey, err := hex.DecodeString(hexKey)
	if err != nil {
		return "", err
	}
	base64Key := base64.StdEncoding.EncodeToString(decodedKey)
	return base64Key, nil
}
