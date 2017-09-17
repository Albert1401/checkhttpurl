package main

import (
	"net/http"
	"flag"
	"fmt"
	"time"
)

const usage = "Usage:\n\tcheckurl [-t timeout] [-i interval] url" +
	"\n" +
	"Examples:\n\tcheckurl -t 1m http://ya.ru" +
	"\n" +
	"\tcheckurl -t 45s -i 2s http://mail.ru" +
	"\n" +
	"\tcheckurl https://www.jetbrains.com/go/"



	func ping(url string, interval time.Duration, ch chan bool) {
	for {
		resp, err := http.Head(url)
		if (err != nil) {
			fmt.Println(err.Error())
		} else {
			if (resp.StatusCode != 200) {
				fmt.Println("Bad response status code:", resp.StatusCode)
			} else {
				ch <- true
			}
		}
		time.Sleep(interval)
	}
}

func main() {
	timeout := flag.Duration("t", time.Minute * 10, "")
	interval := flag.Duration("i", time.Second, "")
	flag.Usage = func() {
		fmt.Println(usage)
	}
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println(usage)
		return
	}
	url := flag.Arg(0)

	fmt.Println("Checking", url)
	ticker := time.NewTicker(*timeout)

	pingch := make(chan bool)
	go ping(url, *interval, pingch)

	select {
	case <-pingch:
		fmt.Println("Status code 200. URL is available!")
	case <-ticker.C:
		fmt.Println("Timeout!")
	}
}
