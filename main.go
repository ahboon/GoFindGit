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

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		fmt.Println("Usage: gofindgit <domains_text_file_name>")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File does not exist")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		wg.Add(1)
		go RunIt(scanner.Text())
		wg.Add(1)
		go SecureRunIt(scanner.Text())
		//fmt.Println(scanner.Text())  // token in unicode-char

	}
	wg.Wait()
}
