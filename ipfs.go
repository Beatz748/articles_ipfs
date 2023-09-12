package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func uploadToIPFS(apiKey, apiSecret, content string) (string, error) {
	// Create a JSON object with the article content
	articleJSON := struct {
		Content string `json:"content"`
	}{
		Content: content,
	}

	// Convert the JSON object to a string
	articleData, err := json.Marshal(articleJSON)
	if err != nil {
		return "", err
	}

	// Create an HTTP request to the Pinata API
	url := "https://api.pinata.cloud/pinning/pinJSONToIPFS"
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(articleData))
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("pinata_api_key", apiKey)
	request.Header.Set("pinata_secret_api_key", apiSecret)

	// Send the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Handle the response
	if response.StatusCode == http.StatusOK {
		// Parse the response and extract the CID
		var result map[string]interface{}
		if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			return "", err
		}
		cid, ok := result["IpfsHash"].(string)
		if !ok {
			return "", fmt.Errorf("CID not found in the response")
		}
		return cid, nil
	} else {
		return "", fmt.Errorf("Error uploading the article to IPFS")
	}
}
