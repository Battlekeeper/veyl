package routing

import (
	"fmt"
	"os/exec"
	"runtime"
)

func SetInterfaceAddress(iface string, addresses string) error {
	// get the OS type
	switch os := runtime.GOOS; os {
	case "windows":
		return _WindowsSetInterfaceAddress(iface, addresses)
	case "linux":
		return _LinuxSetInterfaceAddress(iface, addresses)
	default:
		return fmt.Errorf("unsupported OS: %s", os)
	}
}

func _WindowsSetInterfaceAddress(iface string, address string) error {
	cmd := exec.Command("netsh", "interface", "ip", "set", "address", fmt.Sprintf("name=\"%s\"", iface), "static", address, "255.192.0.0", "none")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set address for interface %s: %w", iface, err)
	}
	return nil
}

func _LinuxSetInterfaceAddress(iface string, address string) error {
	cmd := exec.Command("ip", "addr", "add", address, "dev", iface)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to set address for interface %s: %w", iface, err)
	}
	return nil
}
