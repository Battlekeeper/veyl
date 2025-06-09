package wg

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/Battlekeeper/veil/internal/stun"
	"github.com/Battlekeeper/veil/internal/veil"
	"golang.zx2c4.com/wireguard/conn"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/tun"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type Connection struct {
	Tunnel tun.Device
	Device *device.Device
	Config veil.WgConfig
}

func (wgconn *Connection) TunnelUp() error {
	var err error
	wgconn.Tunnel, err = tun.CreateTUN("veiltun", 1420)
	if err != nil {
		log.Fatal(err)
	}
	wgconn.Device = device.NewDevice(wgconn.Tunnel, conn.NewDefaultBind(), device.NewLogger(device.LogLevelVerbose, "wireguard: "))
	err = wgconn.GenerateKeys()
	if err != nil {
		return fmt.Errorf("failed to generate keys: %w", err)
	}

	err = wgconn.Device.Up()
	if err != nil {
		return fmt.Errorf("failed to bring up TUN device: %w", err)
	}

	// check if running on linux
	if runtime.GOOS == "linux" {
		exec.Command("ip", "link", "set", "veiltun", "up").Run()
	}

	return nil
}

func (wgconn *Connection) TunnelDown() error {
	if wgconn.Device != nil {
		wgconn.Device.Close()
	}
	if wgconn.Tunnel != nil {
		if err := wgconn.Tunnel.Close(); err != nil {
			return fmt.Errorf("failed to close TUN device: %w", err)
		}
	}
	return nil
}

func (wgconn *Connection) GenerateKeys() error {
	var err error
	wgconn.Config.PrivateKey, err = wgtypes.GeneratePrivateKey()
	if err != nil {
		return err
	}
	return nil
}

func (wgconn *Connection) SetListenPort() error {
	ip, stunPort, err := stun.FetchStun()
	if err != nil {
		return err
	}
	wgconn.Config.ListenPort = stunPort
	wgconn.Config.ListenIP = ip.String()
	return nil
}

func (wgconn *Connection) SetIpcConfig() error {
	if wgconn.Config.ListenPort == 0 {
		if err := wgconn.SetListenPort(); err != nil {
			return fmt.Errorf("failed to set listen port: %w", err)
		}
	}

	configStr := fmt.Sprintf(
		`private_key=%s
listen_port=%d
`,
		veil.Base64ToHex(wgconn.Config.PrivateKey.String()), wgconn.Config.ListenPort)

	for _, peer := range wgconn.Config.Peers {
		configStr += fmt.Sprintf(`public_key=%s
endpoint=%s
persistent_keepalive_interval=25
`, veil.Base64ToHex(peer.PublicKey.String()), peer.Endpoint)
		for _, allowedIP := range peer.AllowedIps {
			configStr += fmt.Sprintf(`allowed_ip=%s
`, allowedIP)
		}
	}

	return wgconn.Device.IpcSet(configStr)
}

func (wgconn *Connection) GetIpcConfig() (string, error) {
	out, err := wgconn.Device.IpcGet()
	if err != nil {
		return "", err
	}
	return out, nil
}

func (wgconn *Connection) PrintCurrentIpcConfig() {
	out, err := wgconn.GetIpcConfig()
	if err != nil {
		log.Fatalln("Failed to get IPC config:", err)
	}
	log.Println("WireGuard device is up with the following configuration:\n" + out)
}

func (wgconn *Connection) AddPeer(peer veil.WgPeer) error {
	wgconn.Config.Peers = append(wgconn.Config.Peers, peer)
	wgconn.SetIpcConfig()
	return nil
}
