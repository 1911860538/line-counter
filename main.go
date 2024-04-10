package main

import (
	"log"
	"time"

	"github.com/1911860538/line-counter/count"
)

func main() {
	start := time.Now()
	count.Run()
	cost := time.Now().Sub(start).Milliseconds()
	log.Printf("Done! Successfully calculated statistic in %d milliseconds.", cost)
}
