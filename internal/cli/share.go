package cli

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/m4rkux/wifiqr/internal/wifi"
	"github.com/spf13/cobra"
)

var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Share the QR code of a saved Wi-Fi networks",
	Run: func(cmd *cobra.Command, args []string) {
		networks := wifi.GetNetworks()
		if len(networks) < 1 {
			fmt.Println("No stored Wi-Fi networks found.")
			return
		}

		ssids := []string{}
		for _, net := range networks {
			ssids = append(ssids, net.SSID)
		}

		var selected string
		prompt := &survey.Select{
			Message: "Select a Wi-Fi network to share",
			Options: ssids,
		}

		err := survey.AskOne(prompt, &selected)
		if err != nil {
			fmt.Printf("Error selecting Wi-Fi network: %v\n", err)
		}

		var chosen wifi.WiFiNetwork
		for _, net := range networks {
			if net.SSID == selected {
				chosen = net
				break
			}
		}

		err = wifi.ShowQRCode(chosen.SSID, chosen.Password)
		if err != nil {
			log.Fatalf("Failed to generate QR code: %v", err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(shareCmd)
}
