package wifi

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/ini.v1"
)

var nmPath = "/etc/NetworkManager/system-connections/"

type WiFiNetwork struct {
	SSID     string
	Password string
}

// List and parse the files on /etc/NetworkManager/system-connections/
func GetNetworks() []WiFiNetwork {
	files, err := os.ReadDir(nmPath)
	if err != nil {
		log.Fatalf("Failed to read NetworkManager directory: %v", err)
		return nil
	}
	var networks []WiFiNetwork

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(nmPath, file.Name())
		cfg, err := ini.Load(filePath)
		if err != nil {
			log.Printf("Error parsing the file %s: %v", file.Name(), err)
			continue
		}

		ssid := cfg.Section("wifi").Key("ssid").String()
		psk := cfg.Section("wifi-security").Key("psk").String()
		pskFlags := cfg.Section("wifi-security").Key("psk-flags").String()

		if strings.TrimSpace(ssid) == "" {
			continue
		}

		if strings.TrimSpace(psk) == "" && pskFlags == "1" {
			psk, err = getPasswordFromKeyring(ssid)
			if err != nil {
				log.Printf("Error retrieving password for SSID %s: %v", ssid, err)
				continue
			}
		}

		networks = append(networks, WiFiNetwork{
			SSID:     ssid,
			Password: psk,
		})

	}
	return networks
}

func getPasswordFromKeyring(ssid string) (string, error) {
	// TODO: implement real KWallet access via D-Bus
	return "", nil
}
