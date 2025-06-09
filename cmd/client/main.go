package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/Battlekeeper/veil/internal/veil"
	"github.com/Battlekeeper/veil/internal/wg"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	wgconn := &wg.Connection{}

	if err := wgconn.TunnelUp(); err != nil {
		log.Fatalln("Failed to create TUN device:", err)
	}

	defer wgconn.TunnelDown()

	if err := wgconn.SetIpcConfig(); err != nil {
		log.Fatalln("Failed to set IPC config:", err)
	}

	wgconn.PrintCurrentIpcConfig()

	params := url.Values{}
	params.Add("public_key", wgconn.Config.PrivateKey.PublicKey().String())
	params.Add("relay_id", "relay1")
	params.Add("port", fmt.Sprintf("%d", wgconn.Config.ListenPort))
	params.Add("ip", wgconn.Config.ListenIP)

	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/register?%s", params.Encode()))
	if err != nil {
		log.Fatalln("Failed to register relay:", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to register relay, status code: %d", resp.StatusCode)
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Failed to read response body:", err)
	}
	auth, err := veil.DecodeRelayAuth(bytes)
	if err != nil {
		log.Fatalln("Failed to decode relay auth:", err)
	}
	pubkey, err := wgtypes.ParseKey(auth.PublicKey)
	if err != nil {
		log.Fatalln("Failed to parse public key:", err)
	}
	wgconn.AddPeer(veil.WgPeer{
		PublicKey:  pubkey,
		Endpoint:   fmt.Sprintf("%s:%d", auth.IP, auth.Port),
		AllowedIps: []string{"100.64.0.0/10", "10.10.10.0/24"},
	})

	exec.Command("netsh", "interface", "ip", "set", "address", "name=\"veiltun\"", "static", "100.64.0.50", "255.192.0.0", "none").Run()
	exec.Command("route", "ADD", "100.64.0.1", "mask", "255.255.255.255", "100.64.0.50", "metric", "5", "IF", "55").Run()
	exec.Command("route", "ADD", "100.64.255.255", "MASK", "255.255.255.255", "100.64.0.50", "METRIC", "261", "IF", "55").Run()
	exec.Command("route", "ADD", "10.10.10.0", "MASK", "255.255.255.0", "100.64.0.50", "METRIC", "5", "IF", "55").Run()

	wgconn.PrintCurrentIpcConfig()

	<-sigs
}
