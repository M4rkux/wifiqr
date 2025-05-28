package wifi

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func ShowQRCode(ssid, password string) error {
	// Use the standard format for Wi-Fi QR codes
	qrContent := fmt.Sprintf("WIFI:T:WPA;S:%s;P:%s;;", ssid, password)

	// Generate the QR code
	qr, err := qrcode.New(qrContent, qrcode.Low)
	if err != nil {
		return fmt.Errorf("failed to generate QR code: %w", err)
	}

	fmt.Printf("\nSSID: %s\n", ssid)

	// Print as ASCII to terminal
	fmt.Println(qr.ToString(false)) // false = do not invert, prints light on dark
	fmt.Println("Password: ", password)

	return nil
}
