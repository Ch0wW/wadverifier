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
