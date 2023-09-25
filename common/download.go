package common

import (
	"os"
	"fmt"
	"net/http"
	"io"
)

func DownloadFile(filepath string, url string) (err error) {

	// Create file
	outfile, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("Error while creating output file: %s", err)
	}	
	defer outfile.Close()

	// Get data
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error while fetching URL: %s", err)
	}
	defer response.Body.Close()

	// Check server response
	if response.StatusCode != http.StatusOK {
	 	return fmt.Errorf("Error: bad HTTP status: %s, URL: %s", response.Status, response.Request.URL)
	}

	// Write data to file
	_, err = io.Copy(outfile, response.Body)
	if err != nil {
		return fmt.Errorf("Error while saving to file: %s", err)
	}

	return nil
}
