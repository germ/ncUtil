package main

import (
	"flag"
	"fmt"
	nc "github.com/germ/notCrypto"
)

var encoding, decoding bool
var file, message, key string

func init() {
	flag.BoolVar(&encoding, "e", false, "Pass for encoding")
	flag.BoolVar(&decoding, "d", false, "Pass for decoding")
	flag.StringVar(&file, "f", "", "Location of file to operate on")
	flag.StringVar(&message, "m", "", "Message to be used")
	flag.StringVar(&key, "k", "", "Key to be used for operation")
}

func main() {
	// Error checking and flag exculsivity 
	flag.Parse()
	if !(encoding || decoding) || (encoding && decoding) {
		flag.Usage()
		return
	}
	
	if (file != "") && (message != "") {
		flag.Usage()
		return
	}

	if key == "" {
		flag.Usage()
	}

	// Main logic
	var cipherText 	string
	var err			error

	if encoding && (file != "") {
		cipherText, err = nc.EncodeFile(file, key)
	} else if encoding && (message != "") {
		cipherText = nc.Encode(key, message)
	} else if decoding && (message != "") {
		cipherText = nc.Decode(key, message)
	} else if decoding && (file != "") {
		cipherText, err = nc.DecodeFile(file, key)
	} else {
		flag.Usage()
	}
	
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println(cipherText)
	}
}

