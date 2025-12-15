package modules

import (
	"fmt"
	"os"
)

type TCCFullDiskAccess struct{}

func (t *TCCFullDiskAccess) Name() string {
	return "tcc_fda"
}

func (t *TCCFullDiskAccess) Description() string {
	return "Attempts Full Disk Access to protected system paths"
}

func (t *TCCFullDiskAccess) Generate(target string, port string) error {
	fmt.Println("[*] Attempting Full Disk Access to protected paths...")

	paths := []string{
		"/Library/Application Support/com.apple.TCC/TCC.db",
		os.Getenv("HOME") + "/Library/Safari/History.db",
		os.Getenv("HOME") + "/Library/Mail/V10/MailData/Envelope Index",
		os.Getenv("HOME") + "/Library/Messages/chat.db",
	}

	accessDenied := true
	for _, path := range paths {
		fmt.Printf("[*] Checking: %s\n", path)
		if _, err := os.Stat(path); err == nil {
			f, err := os.Open(path)
			if err != nil {
				fmt.Printf("[*] FDA required for: %s (Access Denied)\n", path)
			} else {
				fmt.Printf("[+] FDA access successful: %s\n", path)
				f.Close()
				accessDenied = false
				break
			}
		}
	}

	if accessDenied {
		fmt.Println("[!] Full Disk Access denied or not granted - telemetry generated")
	}

	return nil
}

func (t *TCCFullDiskAccess) Cleanup() error {
	return nil
}

type TCCDocuments struct{}

func (t *TCCDocuments) Name() string {
	return "tcc_documents"
}

func (t *TCCDocuments) Description() string {
	return "Attempts to access Documents/Downloads/Desktop folders"
}

func (t *TCCDocuments) Generate(target string, port string) error {
	fmt.Println("[*] Attempting to access protected user folders...")

	folders := []string{
		os.Getenv("HOME") + "/Documents",
		os.Getenv("HOME") + "/Downloads",
		os.Getenv("HOME") + "/Desktop",
	}

	for _, folder := range folders {
		fmt.Printf("[*] Attempting to access: %s\n", folder)

		entries, err := os.ReadDir(folder)
		if err != nil {
			fmt.Printf("[*] Access denied to %s: %v\n", folder, err)
			continue
		}

		fmt.Printf("[+] Successfully accessed %s (%d items)\n", folder, len(entries))
	}

	return nil
}

func (t *TCCDocuments) Cleanup() error {
	return nil
}

type TCCPhotos struct{}

func (t *TCCPhotos) Name() string {
	return "tcc_photos"
}

func (t *TCCPhotos) Description() string {
	return "Attempts to access Photos Library"
}

func (t *TCCPhotos) Generate(target string, port string) error {
	photosPath := os.Getenv("HOME") + "/Pictures/Photos Library.photoslibrary"

	fmt.Printf("[*] Attempting to access Photos Library at: %s\n", photosPath)

	entries, err := os.ReadDir(photosPath)
	if err != nil {
		fmt.Printf("[*] Photos access denied (telemetry generated): %v\n", err)
		return nil
	}

	fmt.Printf("[+] Photos Library access successful (%d items)\n", len(entries))
	return nil
}

func (t *TCCPhotos) Cleanup() error {
	return nil
}

type TCCCalendar struct{}

func (t *TCCCalendar) Name() string {
	return "tcc_calendar"
}

func (t *TCCCalendar) Description() string {
	return "Attempts to access Calendar data"
}

func (t *TCCCalendar) Generate(target string, port string) error {
	fmt.Println("[*] Attempting to access Calendar data...")

	calendarPath := os.Getenv("HOME") + "/Library/Calendars"

	entries, err := os.ReadDir(calendarPath)
	if err != nil {
		fmt.Printf("[*] Calendar access denied (telemetry generated): %v\n", err)
		return nil
	}

	fmt.Printf("[+] Calendar data access successful (%d items)\n", len(entries))
	return nil
}

func (t *TCCCalendar) Cleanup() error {
	return nil
}

type TCCReminders struct{}

func (t *TCCReminders) Name() string {
	return "tcc_reminders"
}

func (t *TCCReminders) Description() string {
	return "Attempts to access Reminders data"
}

func (t *TCCReminders) Generate(target string, port string) error {
	fmt.Println("[*] Attempting to access Reminders data...")

	remindersPath := os.Getenv("HOME") + "/Library/Reminders"

	entries, err := os.ReadDir(remindersPath)
	if err != nil {
		fmt.Printf("[*] Reminders access denied (telemetry generated): %v\n", err)
		return nil
	}

	fmt.Printf("[+] Reminders data access successful (%d items)\n", len(entries))
	return nil
}

func (t *TCCReminders) Cleanup() error {
	return nil
}

func init() {
	Register(&TCCFullDiskAccess{})
	Register(&TCCDocuments{})
	Register(&TCCPhotos{})
	Register(&TCCCalendar{})
	Register(&TCCReminders{})

	// TODO: DarwinKit modules
	// Register(&TCCCamera{})
	// Register(&TCCMicrophone{})
	// Register(&TCCContacts{})
	// Register(&TCCScreenCapture{})
	// Register(&TCCLocation{})
	// Register(&TCCAccessibility{})
}
