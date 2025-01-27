package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func main() {
        if len(os.Args) < 3 {
                fmt.Println("ব্যবহার: go run main.go <ইনপুট এক্সিকিউটেবল> <আউটপুট ইমেজ>")
                return
        }

        executablePath := os.Args[1]
        imagePath := os.Args[2]

        // এক্সিকিউটেবল ফাইল পড়া
        executableData, err := ioutil.ReadFile(executablePath)
        if err != nil {
                fmt.Println("এক্সিকিউটেবল ফাইল পড়তে সমস্যা:", err)
                return
        }

        // ইমেজ ফাইল পড়া (বেস ইমেজ)
        imageData, err := ioutil.ReadFile("base.png") // একটা বেস PNG ইমেজ লাগবে
        if err != nil {
                fmt.Println("বেস ইমেজ ফাইল পড়তে সমস্যা:", err)
                return
        }

        // এক্সিকিউটেবল ডেটা ইমেজের শেষে যুক্ত করা
        outputData := append(imageData, executableData...)

        // নতুন ফাইল তৈরি এবং ডেটা লেখা
        err = ioutil.WriteFile(imagePath, outputData, 0644)
        if err != nil {
                fmt.Println("আউটপুট ফাইল লিখতে সমস্যা:", err)
                return
        }

        fmt.Println("সফলভাবে তৈরি হয়েছে:", imagePath)

        // যদি ফাইলটি সরাসরি এক্সিকিউট করা হয় (উইন্ডোজের জন্য)
        if strings.HasSuffix(imagePath, ".exe") || strings.HasSuffix(imagePath, ".bat") {
                cmd := exec.Command(imagePath)
                err = cmd.Start()
                if err != nil {
                        fmt.Println("এক্সিকিউট করতে সমস্যা:", err)
                }
        }

        // লিনাক্স/ম্যাকে এক্সিকিউটেবল করার জন্য পারমিশন দেওয়া:
        if !strings.HasSuffix(imagePath, ".exe") && !strings.HasSuffix(imagePath, ".bat"){
                err = os.Chmod(imagePath, 777)
                if err != nil {
                        fmt.Println("এক্সিকিউটেবল পারমিশন দিতে সমস্যা:", err)
                }
        }
}