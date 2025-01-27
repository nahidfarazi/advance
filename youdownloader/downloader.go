package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/kkdai/youtube/v2"
)

func main() {
        var videoURL, outputDir string

        fmt.Print("Enter YouTube video URL: ")
        fmt.Scanln(&videoURL)

        fmt.Print("Enter output directory (or press Enter for current directory): ")
        fmt.Scanln(&outputDir)

        if outputDir == "" {
                outputDir = "."
        }

        absPath, err := filepath.Abs(outputDir)
        if err != nil {
                log.Fatal("Error getting absolute path:", err)
        }
        outputDir = absPath

        client := youtube.Client{}

        video, err := client.GetVideo(videoURL)
        if err != nil {
                log.Fatal("Error getting video:", err)
        }

        fmt.Println("Video Title:", video.Title)

        // Sort formats by quality (highest to lowest)
        sort.Slice(video.Formats, func(i, j int) bool {
                q1, _ := strconv.Atoi(strings.TrimSuffix(video.Formats[i].Quality, "p"))
                q2, _ := strconv.Atoi(strings.TrimSuffix(video.Formats[j].Quality, "p"))
                return q1 > q2
        })

        // Ask user to choose quality
        fmt.Println("Available Qualities:")
        for i, format := range video.Formats {
                fmt.Printf("%d: %s (%s, %s)\n", i+1, format.Quality, format.MimeType, format.Container)
        }

        var choice int
        fmt.Print("Enter your choice (1-", len(video.Formats), "): ")
        _, err = fmt.Scanln(&choice)
        if err != nil || choice < 1 || choice > len(video.Formats) {
                log.Fatal("Invalid choice.")
        }

        format := video.Formats[choice-1]

        downloadFormat(client, video, &format, outputDir)
        fmt.Println("Download Complete!")
}

func downloadFormat(client *youtube.Client, video *youtube.Video, format *youtube.Format, outputDir string) {
        stream, err := client.GetStream(video, format) // Correct GetStream usage
        if err != nil {
                log.Printf("Error getting stream: %v", err)
                return
        }
        defer stream.Close()

        filename := sanitizeFilename(video.Title) + "." + format.Container
        filepath := filepath.Join(outputDir, filename)

        file, err := os.Create(filepath)
        if err != nil {
                log.Printf("Error creating file: %v", err)
                return
        }
        defer file.Close()

        fmt.Printf("Downloading %s (%s)...\n", video.Title, format.Quality)

        _, err = io.Copy(file, stream)
        if err != nil {
                log.Printf("Error downloading: %v", err)
                return
        }
}

func sanitizeFilename(filename string) string {
        reg, err := regexp.Compile("[^a-zA-Z0-9\\.\\-]+")
        if err != nil {
                log.Fatal(err)
        }
        return reg.ReplaceAllString(filename, "_")
}