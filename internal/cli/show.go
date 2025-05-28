package cli

import (
	"github.com/m4rkux/wifiqr/internal/wifi"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the QR codes of a saved Wi-Fi network by SSID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			cmd.Help()
			return
		}
		ssid := args[0]
		err := wifi.ShowQRCode(ssid, "")
		if err != nil {
			cmd.PrintErrf("Error showing QR code for SSID %s: %v\n", ssid, err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
