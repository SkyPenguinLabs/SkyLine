package SkyLin_Backend_SkyStandardLib_Image_Forensics

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

const (
	GEN_Signature = '\x50'
)

var (
	ZIP_Signature = []byte{'\x4b', '\x03', '\x04'}
)

// Controller function | Verifies the factor of a ZIP signature existing within the image data
func Image_Controller_Function_External_1_Verify_Archive(imagefile string) (string, bool) {
	f, err := os.Open(imagefile)
	if err != nil {
		log.Fatal(err)
		return "", false
	}
	defer f.Close()

	buffer := bufio.NewReader(f)
	offset := int64(0)
	for {
		byter, err := buffer.ReadByte()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			log.Fatal(err)
			return "", false
		}

		if byter == GEN_Signature {
			B_Mark := make([]byte, len(ZIP_Signature))
			_, err := buffer.Read(B_Mark)
			if err != nil {
				log.Fatal(err)
				return "", false
			}
			if bytes.Equal(B_Mark, ZIP_Signature) {
				fmt.Println(B_Mark)
				offsetHex := fmt.Sprintf("0x%X", offset)
				return offsetHex, true
			}
		}
		offset++
	}
	return "", false
}
