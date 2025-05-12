package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/letigredununavu/xpath-injection-tool/internal/exploit"
	"github.com/letigredununavu/xpath-injection-tool/internal/httpclient"
)

func main() {

	// Defining flags
	target := flag.String("target", "", "Target URL to scan (e.g. http:/example.com/search?q=)")
	method := flag.String("method", "Get", "HTTP Method to use (e.g. Get, Post, Options, etc. Default: Get)")

	flag.Parse()

	if *target == "" {
		fmt.Println("Error: You must provide a target URL with -target")
		flag.Usage()
		os.Exit(1)
	}

	var response string
	var err error

	if *method == "Get" {
		fmt.Println("[*] Sending Get request to: ", *target)
		response, err = httpclient.SendRequest(*target)

	} else if *method == "Post" {
		payload := "test" // TODO
		contentType := "application/x-www-form-urlencoded"
		fmt.Println("[*] Sending Post request to : ", *target)
		response, err = httpclient.SendPostRequest(*target, payload, contentType)
	}

	if err != nil {
		fmt.Println("[-] Error: ", err)
		os.Exit(1)
	}

	fmt.Println("[+] Response :\n", response)

	fmt.Println("[*] Looking for vulnerable application..")
	vulnerable, err := exploit.Test_BlindXPathInjection(*target)

	if err != nil {
		fmt.Println("Error sending BlindXPathInjection test payload : ", err)
	}

	if vulnerable {
		fmt.Println("[+] Target vulnerable")
	} else {
		fmt.Println("[-] Target not vulnerable")
	}

}
