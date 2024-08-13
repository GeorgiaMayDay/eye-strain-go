package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {
	interval := 30 * time.Second
	go eye_strain_ticker(interval, &RealNotif{})

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	//time for cleanup before exit
	fmt.Println("Adios!")
}

func eye_strain_ticker(interval time.Duration, Notif notifer) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		createNotif(Notif, "Eye Care Time", "Look Away")
	}

}

type notifer interface {
	Notif(header, body string) error
}

type RealNotif struct{}

func (rn *RealNotif) Notif(header, body string) error {
	err := beeep.Notify(header, body, "Eye-Image.png")
	if err != nil {
		panic(err)
	}
	return err
}

func createNotif(Notif notifer, header, body string) error {
	err := Notif.Notif(header, body)
	return err
}
