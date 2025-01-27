package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
        port      = "8080" // Default port
        shareDir  = "."    // Default sharing directory
        baseURL   string
        startTime time.Time
        fileList  []string
        mu        sync.Mutex
)

func main() {
        startTime = time.Now()

        if len(os.Args) > 1 {
                port = os.Args[1]
        }
        if len(os.Args) > 2 {
                shareDir = os.Args[2]
        }

        absPath, err := filepath.Abs(shareDir)
        if err != nil {
                log.Fatal("Error getting absolute path:", err)
        }
        shareDir = absPath

        fileList, err = listFiles(shareDir)
        if err != nil {
                log.Fatal("Error listing files:", err)
        }

        baseURL = fmt.Sprintf("http://%s:%s/", getLocalIP(), port)

        fmt.Println("Sharing files from:", shareDir)
        fmt.Println("Server started at:", baseURL)
        fmt.Println("Use Ctrl+C to stop.")

        http.HandleFunc("/", fileListHandler)
        http.HandleFunc("/download/", downloadHandler)

        err = http.ListenAndServe(":"+port, nil)
        if err != nil {
                log.Fatal("ListenAndServe:", err)
        }
}

func fileListHandler(w http.ResponseWriter, r *http.Request) {
        mu.Lock()
        defer mu.Unlock()
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        fmt.Fprintf(w, "<h1>Shared Files</h1><ul>")
        for _, file := range fileList {
                escapedFile := strings.ReplaceAll(file, " ", "%20") // URL encode spaces
                fmt.Fprintf(w, `<li><a href="/download/%s">%s</a></li>`, escapedFile, file)
        }
        fmt.Fprintf(w, "</ul><p>Server started at: %s<br>Uptime: %s</p>", baseURL, time.Since(startTime))
}

func downloadHandler(w http.ResponseWriter, r *http.Request) {
        filename := strings.TrimPrefix(r.URL.Path, "/download/")
        filename = strings.ReplaceAll(filename, "%20", " ") // URL decode spaces
        filepath := filepath.Join(shareDir, filename)

        mu.Lock()
        defer mu.Unlock()

        if !fileExists(filepath) {
                http.NotFound(w, r)
                return
        }

        fileInfo, err := os.Stat(filepath)
        if err != nil {
                http.Error(w, "Error getting file info", http.StatusInternalServerError)
                return
        }
    w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))
        w.Header().Set("Content-Disposition", "attachment; filename=\""+filename+"\"")
        http.ServeFile(w, r, filepath)
}

func listFiles(dir string) ([]string, error) {
        var files []string
        err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
                if !info.IsDir() {
                        relPath, _ := filepath.Rel(dir, path)
                        files = append(files, relPath)
                }
                return nil
        })
        return files, err
}

func fileExists(filename string) bool {
        info, err := os.Stat(filename)
        if os.IsNotExist(err) {
                return false
        }
        return !info.IsDir()
}

func getLocalIP() string {
        addrs, err := net.InterfaceAddrs()
        if err != nil {
                return "localhost"
        }
        for _, address := range addrs {
                if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
                        if ipnet.IP.To4() != nil {
                                return ipnet.IP.String()
                        }
                }
        }
        return "localhost"
}