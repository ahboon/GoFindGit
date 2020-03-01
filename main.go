package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

func RunIt(domain string) {
	target := "http://" + domain + "/.git/HEAD"
	targetClient := http.Client{Timeout: 2 * time.Second}
	targetResp, err := targetClient.Get(target)
	if err != nil {
		return
	}
	if targetResp.StatusCode != 200 {
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(targetResp.Body)
	newStr := buf.String()
	// fmt.Println(strings.Contains(newStr, "refs/heads"))
	if strings.Contains(newStr, "refs/heads") {
		fmt.Println(domain)
	}
	defer wg.Done()
}

func SecureRunIt(domain string) {
	target := "https://" + domain + "/.git/HEAD"
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	targetClient := http.Client{Timeout: 2 * time.Second}
	targetResp, err := targetClient.Get(target)
	if err != nil {
		return
	}
	if targetResp.StatusCode != 200 {
		return
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(targetResp.Body)
	newStr := buf.String()
	// fmt.Println(strings.Contains(newStr, "refs/heads"))
	if strings.Contains(newStr, "refs/heads") {
		fmt.Println(domain)
	}
	defer wg.Done()
}

func EnvRunIt(domain string) {
	target := "http://" + domain + "/.env"
	targetClient := http.Client{Timeout: 2 * time.Second}
	targetResp, err := targetClient.Get(target)
	if err != nil {
		return
	}
	if targetResp.StatusCode != 200 {
		return
	}
	fmt.Println(domain)
	defer wg.Done()
}

func SecureEnvRunIt(domain string) {
	target := "https://" + domain + "/.env"
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	targetClient := http.Client{Timeout: 2 * time.Second}
	targetResp, err := targetClient.Get(target)
	if err != nil {
		return
	}
	if targetResp.StatusCode != 200 {
		return
	}
	fmt.Println(domain)
	defer wg.Done()
}

func main() {
	banner :=
		`
	 _____        _____      _    _____ _ _   
	/ ____|      / ____|    | |  / ____(_) |  
       | |  __  ___ | |  __  ___| |_| |  __ _| |_ 
       | | |_ |/ _ \| | |_ |/ _ \ __| | |_ | | __|
       | |__| | (_) | |__| |  __/ |_| |__| | | |_ 
	\_____|\___/ \_____|\___|\__|\_____|_|\__|
											  
	`
	fmt.Println(banner)
	helpertext :=
		`
Usage:
	./gofindgit <mode> <domains_text_file_name>

	Modes:
	- git 
		Scans for git repo exposed in web root. http://example.com/.git/ OR https://example.com/.git/
	- env
		Scans for .env exposed in web root. http://example.com/.env/ OR https://example.com/.env/
	`
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		fmt.Println(helpertext)
		return
	} else if os.Args[1] == "help" {
		fmt.Println(helpertext)
		return
	}
	file, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Println("File does not exist")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if os.Args[1] == "git" {
		for scanner.Scan() {
			wg.Add(1)
			go RunIt(scanner.Text())
			wg.Add(1)
			go SecureRunIt(scanner.Text())

		}
	} else if os.Args[1] == "env" {
		for scanner.Scan() {
			wg.Add(1)
			go EnvRunIt(scanner.Text())
			wg.Add(1)
			go SecureEnvRunIt(scanner.Text())

		}
	} else {
		fmt.Println("To get help, run './gofindgit help'")
	}

	wg.Wait()
}
