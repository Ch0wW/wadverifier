package main

import (
	"errors"
	"io"
	"os"
	"wadverifier/wad"
)

// Adapted from https://github.com/XerTheSquirrel/go2it/blob/master/wad.go
// We only need the first 4 bytes.
func OpenCheckWADValid(filepath string) bool {

	// Opening it AGAIN
	file, err := os.Open(filepath)
	if err != nil {
		return false
	}

	// And AGAIN, don't forget to close it !
	defer file.Close()

	data := make([]byte, 4)
	_, err = io.ReadFull(file, data)
	if err != nil {
		return false
	}

	// Close the file as it is not needed anymore
	file.Close()

	if wad.IsValidHeader(data) {
		return true
	}

	return false
}

func CompIWADData(data wad.Entry, hash wad.HashInfo) (wad.Entry, error) {
	if hash.MD5 != "" && hash.MD5 == data.MD5Hash {
		return data, nil
	}

	// In this case we're going to look for SHA1 detection, IF AND ONLY IF we provide one value
	if hash.SHA1 != "" && hash.SHA1 == data.SHA1Hash {
		return data, nil
	}

	return wad.Entry{}, errors.New("unknown WAD")
}
