package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gocolly/colly/v2"
)

type Downloader struct {
	link     string
	password int
	size     int
}

func Download() {
	fmt.Println("hehe")
	c := colly.NewCollector()
	c.SetClient(&http.Client{
		Timeout: 30 * time.Second,
	})
	// domain := "https://www.ovagames.com"
	//Stil working for auto download the game, currently busy....
	//will complete soon....
}
