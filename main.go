package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func main() {
	// Get Cloudinary credentials from user input
	cloudName := getInput("Enter Cloudinary cloud name: ")
	apiKey := getInput("Enter Cloudinary API key: ")
	apiSecret := getInput("Enter Cloudinary API secret: ")
	// Get folder name to upload to
	folderName := getInput("Enter folder name to upload to: ")

	// Initialize Cloudinary
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		fmt.Printf("Error initializing Cloudinary: %v\n", err)
		return
	}

	// Get all files in the current directory
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}

	// Upload each file
	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(".", file.Name())
			publicID := filepath.Join(folderName, strings.ReplaceAll(strings.TrimSuffix(file.Name(), ".jpg"), "'", ""))

			// Create boolean pointers
			useFilename := true
			uniqueFilename := false

			// Upload file to Cloudinary
			_, err := cld.Upload.Upload(context.Background(), filePath, uploader.UploadParams{
				PublicID:       publicID,
				UseFilename:    &useFilename,
				UniqueFilename: &uniqueFilename,
			})
			if err != nil {
				fmt.Printf("Error uploading %s: %v\n", file.Name(), err)
			} else {
				fmt.Printf("Uploaded %s successfully\n", file.Name())
			}
		}
	}

	fmt.Println("Upload process completed.")
}

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
