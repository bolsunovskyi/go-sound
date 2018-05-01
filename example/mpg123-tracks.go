package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"log"

	"github.com/bolsunovskyi/go-sound"
)

func exit(c chan os.Signal) {
	<-c
	fmt.Println("exit")
	os.Exit(0)
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go exit(c)

	tracks := track.MakeMPG123Tracks()
	if err := tracks.AddMultipleTracks("storm", "storm.mp3",
		"rain", "rain.mp3", "fire", "fire.mp3"); err != nil {
		log.Fatalln(err)
	}

	go tracks.Play("storm")
	go tracks.Play("fire")

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("stop")
		tracks.Stop("storm")
	}()

	select {}
}
