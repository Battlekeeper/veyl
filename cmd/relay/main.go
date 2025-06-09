package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"os/exec"

	"github.com/Battlekeeper/veil/internal/stun"
	"github.com/Battlekeeper/veil/internal/veil"
	"github.com/Battlekeeper/veil/internal/wg"
	"github.com/google/nftables"
	"github.com/google/nftables/expr"
	"github.com/gorilla/websocket"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func main() {
	wgconn := &wg.Connection{}

	if err := wgconn.TunnelUp(); err != nil {
		log.Fatalln("Failed to create TUN device:", err)
	}

	defer wgconn.TunnelDown()

	if err := wgconn.SetIpcConfig(); err != nil {
		log.Fatalln("Failed to set IPC config:", err)
	}

	wgconn.PrintCurrentIpcConfig()

	ipcmd := exec.Command("ip", "addr", "add", "100.64.0.1/10", "dev", "veiltun")
	data, err := ipcmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Failed to add IP address: %s, error: %v", data, err)
	}
	log.Printf("IP address added to TUN device: %s\n", string(data))

	u := url.URL{Scheme: "ws", Host: "192.168.1.49:8080", Path: "/relay"}
	log.Printf("connecting to %s", u.String())

	// Connect to the server
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer c.Close()

	stunIp, stunPort, err := stun.FetchStun()
	if err != nil {
		log.Fatalf("Failed to fetch STUN server: %v", err)
	}

	auth := veil.RelayAuth{
		RelayID:   "relay1",
		PublicKey: wgconn.Config.PrivateKey.PublicKey().String(),
		IP:        stunIp.String(),
		Port:      stunPort,
	}

	c.WriteJSON(auth)

	for {
		// decode to veil.RelayConnection
		var relayConnection veil.RelayConnection
		if err := c.ReadJSON(&relayConnection); err != nil {
			log.Println("read error:", err)
			break
		}
		pubkey, err := wgtypes.ParseKey(relayConnection.PublicKey)
		if err != nil {
			log.Println("Failed to parse public key:", err)
			continue
		}
		peer := veil.WgPeer{
			PublicKey:  pubkey,
			Endpoint:   fmt.Sprintf("%s:%d", relayConnection.IP, relayConnection.Port),
			AllowedIps: []string{"100.64.0.50/32"},
		}
		wgconn.AddPeer(peer)
		wgconn.PrintCurrentIpcConfig()
		route()
	}
}

func route() {
	c := &nftables.Conn{}

	table := &nftables.Table{
		Family: nftables.TableFamilyIPv4,
		Name:   "nat",
	}
	c.AddTable(table)

	chain := &nftables.Chain{
		Name:     "postrouting",
		Table:    table,
		Type:     nftables.ChainTypeNAT,
		Hooknum:  nftables.ChainHookPostrouting,
		Priority: nftables.ChainPriorityNATSource,
	}
	c.AddChain(chain)

	// Source and destination IPs to match
	srcNet := &net.IPNet{
		IP:   net.IPv4(100, 64, 0, 0),
		Mask: net.CIDRMask(10, 32),
	}
	dstNet := &net.IPNet{
		IP:   net.IPv4(10, 10, 10, 0),
		Mask: net.CIDRMask(24, 32),
	}

	// Create the rule
	rule := &nftables.Rule{
		Table: table,
		Chain: chain,
		Exprs: []expr.Any{
			// [ match source IP ]
			&expr.Payload{
				DestRegister: 1,
				Base:         expr.PayloadBaseNetworkHeader,
				Offset:       12, // source IP in IPv4 header
				Len:          4,
			},
			&expr.Cmp{
				Register: 1,
				Op:       expr.CmpOpEq,
				Data:     srcNet.IP.To4(),
			},
			// [ match destination IP ]
			&expr.Payload{
				DestRegister: 1,
				Base:         expr.PayloadBaseNetworkHeader,
				Offset:       16, // destination IP in IPv4 header
				Len:          4,
			},
			&expr.Cmp{
				Register: 1,
				Op:       expr.CmpOpEq,
				Data:     dstNet.IP.To4(),
			},
			// [ masquerade ]
			&expr.Masq{},
		},
	}
	c.AddRule(rule)

	if err := c.Flush(); err != nil {
		panic(err)
	}
}
