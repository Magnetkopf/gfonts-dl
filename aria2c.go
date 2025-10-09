package main

import (
	"log"
	"os/exec"
	"time"
)

func downloadFile(url string, savePath string) bool {
	maxRetries := 5
	for attempts := 0; attempts < maxRetries; attempts++ {
		cmd := exec.Command("aria2c", "--allow-overwrite=true", "-o", savePath, url)
		cmd.Stdout = nil
		cmd.Stderr = nil

		err := cmd.Run()
		if err == nil {
			log.Printf("Downloaded: %s", savePath)
			return true
		} else {
			log.Printf("Failed to download %s (attempt %d/%d), retrying in 1s... (%v)", url, attempts+1, maxRetries, err)
			time.Sleep(1 * time.Second)
		}
	}
	log.Fatalf("Failed to download after %d attempts: %s", maxRetries, url)
	return false
}
