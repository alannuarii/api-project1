package utils

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"strings"
)

func Base64topng(dataURI, outputPath string) error{
	parts := strings.Split(dataURI, ";base64,")
	if len(parts) != 2 {
        return errors.New("Format data tidak valid")
    }
	
	imgData, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil{
		return err
	}

	err = ioutil.WriteFile(outputPath, imgData, 0644)
	if err != nil{
		return err
	}

	return nil
}