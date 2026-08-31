package wadapi

import (
	"encoding/json"
	"fmt"
	"os"
	"wadverifier/wad"
)

const (
	API_VERSION = 2 // 1 was factually speaking pre-0.8.0
)

type JSONGeneratedWADList struct {
	APIVersion int         `json:"version"`
	WadEntries []wad.Entry `json:"wadinfo"`
}

func GenerateJSONFile(list []wad.Entry, resultfile string) {

	buildfile := JSONGeneratedWADList{
		APIVersion: API_VERSION,
		WadEntries: list,
	}

	fmt.Println("Adding", len(list), "WAD entries to the generated JSON file.")

	result, _ := json.MarshalIndent(buildfile, "", "    ")
	os.WriteFile(resultfile, result, os.ModePerm)
}

func LoadCustomPWADFile(filename string) (error, JSONGeneratedWADList) {

	var retvalue JSONGeneratedWADList

	cfg, err := os.Open(filename)
	if err != nil {
		return err, retvalue
	}

	err = json.NewDecoder(cfg).Decode(&retvalue)
	if err != nil {
		return err, JSONGeneratedWADList{}
	}

	cfg.Close()

	// Check its API
	if retvalue.APIVersion < API_VERSION {
		return fmt.Errorf("this file uses an outdated API (Has %d, Requires %d)", retvalue.APIVersion, API_VERSION), JSONGeneratedWADList{}
	}

	if len(retvalue.WadEntries) == 0 {
		return fmt.Errorf("file doesn't contain any PWAD entry"), JSONGeneratedWADList{}
	}

	return nil, retvalue
}
