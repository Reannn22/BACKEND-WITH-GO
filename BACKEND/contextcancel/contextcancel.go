package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	c := context.Background()
	ctx, cancel := context.WithTimeout(c, time.Second*2) //change this second value, and see the different.
	defer cancel()
	fbreceiver := make(chan string)
	dbreceiver := make(chan string)
	go GetDataFromFacebook(ctx, fbreceiver)
	go GetDataFromDatabase(ctx, dbreceiver)
	defer close(fbreceiver)
	defer close(dbreceiver)
	for i := 0; i < 2; i++ {
		select {
		case fb := <-fbreceiver:
			fmt.Println(">>>>>> Data Received From: ", fb)
		case db := <-dbreceiver:
			fmt.Println(">>>>>> Data Received From: ", db)
		case <-ctx.Done():
			fmt.Println(">>>>> Timeout to process")
		}
	}
	time.Sleep(time.Second * 50) // Just To simulate the context
}

func GetDataFromDatabase(ctx context.Context, datareceiver chan string) {
	startTime := time.Now()
	ticker := time.NewTicker(time.Second * 1)
	for _ = range ticker.C {
		fmt.Println("Still Processing From DB")
		if time.Since(startTime).Seconds() >= 10 { // Example: Fetch Data needs 10 detik
			ticker.Stop()
			break
		}
	}
	if ctx.Err() == nil { // ctx.Err() will filled when the ctx is timeout or the ctx canceled
		datareceiver <- "database"
	}
	return
}

func GetDataFromFacebook(ctx context.Context, datareceiver chan string) {
	startTime := time.Now()
	ticker := time.NewTicker(time.Second * 1)
	for _ = range ticker.C {
		fmt.Println("Still Processing From FB")
		if time.Since(startTime).Seconds() >= 5 { // Ex:Fetch Data needs 5 seconds
			ticker.Stop()
			break
		}
	}
	if ctx.Err() == nil { // ctx.Err() will filled when the ctx is timeout or the ctx canceled
		datareceiver <- "facebook"
	}
	return
}