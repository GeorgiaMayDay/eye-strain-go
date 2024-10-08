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
	interval := 20 * time.Minute
	lookAwayInterval := 20 * time.Second
	go eye_strain_ticker(interval, lookAwayInterval, &RealNotif{})

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	//time for cleanup before exit
	fmt.Println("Adios!")
}

// Manages the timings of the notificaton for eye strain
//
// Parameters:
//
//	interval time.Duration
//	lookAwayInterval time.Duration
//	Notif notifier
func eye_strain_ticker(interval, lookAwayInterval time.Duration, Notif notifer) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		createNotif(Notif, "Eye Care Time", "Look Away")
		lookBackTimer := time.NewTimer(lookAwayInterval)
		<-lookBackTimer.C
		createNotif(Notif, "Screen Time", "Look Back")
	}
}

type notifer interface {
	Notif(header, body string) error
}

type RealNotif struct{}

// Wrapper function for Beeep libraries Notify, inserts same photo
// Parameter:

// 	header string
// 	body string

// Return:

// err error
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
