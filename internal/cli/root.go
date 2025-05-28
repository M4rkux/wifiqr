package cli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "wifiqr",
	Short: "A CLI tool to extract Wi-Fi passwords and generate QR codes",
	Long:  "wifiqr is a tool to list saved Wi-Fi networks and export them as QR codes.",
}

// Execute runs the root command and handles any errors.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}
