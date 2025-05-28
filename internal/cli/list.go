package cli

import (
	"fmt"

	"github.com/m4rkux/wifiqr/internal/wifi"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List saved wi-fi networks",
	Run: func(cmd *cobra.Command, args []string) {
		networks := wifi.GetNetworks()
		for _, n := range networks {
			fmt.Println(n.SSID)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
