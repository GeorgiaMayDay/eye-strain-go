package main

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

var errTest = errors.New("called correctly")

type FakeNotif struct {
	count int
}

func (fn *FakeNotif) increaseCount() {
	fn.count++
}

func (fn *FakeNotif) Notif(header, body string) error {
	fn.increaseCount()
	fmt.Printf(" Inside the create notif %d", fn.count)
	return errTest
}

func TestNotif(t *testing.T) {

	t.Run("Timer works correctly", func(t *testing.T) {
		mockNotif := FakeNotif{count: 0}
		mockNotif.count += 2
		go eye_strain_ticker(3*time.Millisecond, 1*time.Microsecond, &mockNotif)

		fmt.Printf("Within Test:%d", mockNotif.count)
	})

	t.Run("create notification works", func(t *testing.T) {
		expected_err := createNotif(&FakeNotif{count: 0}, "", "")

		if expected_err == nil {
			t.Errorf("No Error")
		}

		if expected_err.Error() != "called correctly" {
			t.Errorf("Wrong Error")
		}

	})

}
