package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/schollz/progressbar/v3"
	"golang.org/x/sys/windows"
)

func isAdmin() bool {
	var sid *windows.SID
	err := windows.AllocateAndInitializeSid(
		&windows.SECURITY_NT_AUTHORITY,
		2,
		windows.SECURITY_BUILTIN_DOMAIN_RID,
		windows.DOMAIN_ALIAS_RID_ADMINS,
		0, 0, 0, 0, 0, 0,
		&sid)
	if err != nil {
		return false
	}
	defer windows.FreeSid(sid)

	token := windows.Token(0)
	member, err := token.IsMember(sid)
	if err != nil {
		return false
	}
	return member
}

func main() {
	if !isAdmin() {
		fmt.Println("Error: This program requires administrator privileges")
		fmt.Println("Please run the program as administrator")
		os.Exit(1)
	}

	dirPath := flag.String("dir", "", "Directory path to scan for fba_ads files")
	flag.Parse()

	if *dirPath == "" {
		fmt.Println("Error: Please provide a directory path using --dir flag")
		fmt.Println("Usage: FAFRemover --dir=C:")
		os.Exit(1)
	}

	fmt.Println("Scanning directory for fba_ads files...")

	var filesToDelete []string
	files, err := os.ReadDir(*dirPath)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		os.Exit(1)
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(strings.ToLower(file.Name()), "fba_ads") {
			filesToDelete = append(filesToDelete, filepath.Join(*dirPath, file.Name()))
		}
	}

	totalFiles := len(filesToDelete)
	if totalFiles == 0 {
		fmt.Println("No fba_ads files found.")
		return
	}

	fmt.Printf("Found %d files to delete\n", totalFiles)

	bar := progressbar.NewOptions(totalFiles,
		progressbar.OptionSetDescription("Deleting files..."),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowCount(),
		progressbar.OptionSetWidth(15),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionThrottle(65*time.Millisecond),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
	startTime := time.Now()
	deletedCount := 0

	for _, file := range filesToDelete {
		if err := os.Remove(file); err != nil {
			fmt.Printf("\nError deleting %s: %v\n", file, err)
		} else {
			deletedCount++
		}
		bar.Add(1)
	}

	fmt.Printf("\nCompleted: %d/%d files deleted in %v\n", deletedCount, totalFiles, time.Since(startTime).Round(time.Second))
}
