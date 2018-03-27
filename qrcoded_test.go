package main

import (
	"testing"
)

// a test function in Golang will generally have the form of the following
// func TestXXX(t *testing.T)
// XXX := as the name of the testing function
// this is a unit test
func TestGenerateQRCodeReturnsValue(t *testing.T) {
	result := GenerateQRCode("555-2368") // essence of the test
	if result == nil { // assertions of the test
		t.Errorf("Generated QRCode is nil")
	}
	if len(result) == 0 {
		t.Errorf("Generated QRCode has no data")
	}

}

// Create the minimum thing necessary to pass the test 
func GenerateQRCode(code string) []byte {
	return []byte{0xFF }
}
