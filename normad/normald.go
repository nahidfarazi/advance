package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
		var urls string
        // ডাউনলোড করার URL লিঙ্কটি প্রদান করুন
		fmt.Printf("Enter your link : ")
		fmt.Scanf("%s",&urls)
        url := urls 

        // ফাইলটি কোথায় সেভ করতে চান সেটি নির্দিষ্ট করুন
		// fmt.Printf("Enter your folder name: ")
		// fmt.Scanf("%s",&fileName)
        filePath := "./downloads"

        // ফাইলটি সেভ করার জন্য ফোল্ডার তৈরি করুন (যদি না থাকে)
        os.MkdirAll("/home/error/Documents/downloads", os.ModePerm)

        // নতুন ফাইল তৈরি করুন
        outputFile, err := os.Create(filePath)
        if err != nil {
                fmt.Println("Error creating file:", err)
                return
        }
        defer outputFile.Close()

        // URL থেকে ডাটা ফেচ করুন
        response, err := http.Get(url)
        if err != nil {
                fmt.Println("Error fetching URL:", err)
                return
        }
        defer response.Body.Close()

        // ডাটা কপি করুন
        _, err = io.Copy(outputFile, response.Body)
        if err != nil {
                fmt.Println("Error copying data:", err)
                return
        }

        fmt.Println("File downloaded successfully to:", filePath)
}