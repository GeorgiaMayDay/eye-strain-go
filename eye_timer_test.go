package main

import (
	"errors"
	"testing"
)

var errTest = errors.New("called correctly")

type FakeNotif struct{}

func (rn *FakeNotif) Notif(header, body string) error {
	return errTest
}

func TestTicker(t *testing.T) {

	expected_err := createNotif(&FakeNotif{}, "", "")

	if expected_err == nil {
		t.Errorf("No Error")
	}

	if expected_err != errTest {
		t.Errorf("Wrong Error")
	}

}
