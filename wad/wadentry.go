package wad

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
	"encoding/hex"
	"io"
	"os"
	"wadverifier/wad/flags"
	"wadverifier/wad/games"
	"wadverifier/wad/status"
)

const (
	IWADbytes = 1145132873
	PWADbytes = 1145132880
)

// WadInfo : all WAD data returned from the program
type Entry struct {
	MD5Hash      string          `json:"md5"`
	SHA1Hash     string          `json:"sha1"` // Unused for checks yet
	Name         string          `json:"name"`
	Version      string          `json:"version"`
	Status       status.Status   `json:"status"` // Is it the final release of this?
	Flags        flags.Flags     `json:"flags"`  // Special Flags
	Patchinfo    games.PatchType `json:"patchinfo"`
	PWADRequires string          `json:"requires"`   // If the official PWAD requires an IWAD to run
	Additional   string          `json:"additional"` // If I need to display an additionnal message for this IWAD.
}

type HashInfo struct {
	MD5  string `json:"md5"`
	SHA1 string `json:"sha1"`
}

/*func LoadDefinitions() [][]Info {


	return [][]
}*/

func (s *Entry) IsFinal() bool {
	return s.Status.IsFinal()
}

func (s *Entry) IsNotFinal() bool {
	return s.Status.IsNotFinal()
}

func GetMD5Hash(filePath string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string

	//Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}

	//Tell the program to call the following function when the current function returns
	defer file.Close()

	//Open a new hash interface to write to
	hash := md5.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]

	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil
}

func GetSHA1Hash(filePath string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnHash string

	//Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}

	//Tell the program to call the following function when the current function returns
	defer file.Close()

	//Open a new hash interface to write to
	hash := sha1.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]

	//Convert the bytes to a string
	returnHash = hex.EncodeToString(hashInBytes)

	return returnHash, nil
}

// Adapted from https://github.com/XerTheSquirrel/go2it/blob/master/wad.go
// We only need the first Long.
func IsValidHeader(header []byte) bool {

	// Need the magic number
	magic := binary.LittleEndian.Uint32(header[0:4])
	if magic != IWADbytes && magic != PWADbytes {
		return false
	}

	return true
}
