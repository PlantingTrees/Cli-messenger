package notification

import (
	"embed"
	"log"
	"os"

	"github.com/gen2brain/beeep"
)

//go:embed assets/oyeh-logo.png
var iconFS embed.FS

func NotifyUser(title, message string) {
	go PlaySound()

	// extract icon png file to be used in tmp file
	// this tmp file is needed to embed the logo when the app is bundled
	iconPath, err := extractIconToTemp()
	if err != nil {
		log.Println("Could not extract icon:", err)
		iconPath = ""
	}
	// for now the app will render a default logo instead of the app logo, because apple reuires a bundled app ... :(
	if err := beeep.Notify(title, message, iconPath); err != nil {
		log.Println("Notification failed:", err)
	}
}

// Helper: Writes embedded bytes to a temporary file
func extractIconToTemp() (string, error) {
	// Read file from the binary
	data, err := iconFS.ReadFile("assets/oyeh-logo.png")
	if err != nil {
		return "", err
	}

	// tmpfile helps store a tmp file that makes it easier to embed into binary
	tmpFile, err := os.CreateTemp("", "oyeh-logo-*.png")
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()

	// Write data
	if _, err := tmpFile.Write(data); err != nil {
		return "", err
	}

	return tmpFile.Name(), nil
}
