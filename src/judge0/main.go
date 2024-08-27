// import nessasory packages
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func make_submission(filePath string, fileContent string, expectedOutputFile string, id string) (any, int) {
	//convert id which is a string to int
	if expectedOutputFile == "" {
		return 0, 400
	}
	if fileContent == "" && filePath == "" {
		//return err
		return 0, 400
	}
	if fileContent == "" && filePath != "" {
		inputFile, err := ioutil.ReadFile(filePath)
		fileContent = string(inputFile)
		if err != nil {
			fmt.Print(err.Error())
			return 0, 400
		}
	}
	expectedOutput, err := ioutil.ReadFile(expectedOutputFile)
	expectedOutputContent := string(expectedOutput)
	fmt.Println(expectedOutputContent, fileContent)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	payload := map[string]string{
		"language_id":     id,
		"source_code":     fileContent,
		"expected_output": expectedOutputContent,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	// Make the POST request
	response, err := http.Post(
		"http://localhost:2358/submissions/",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		log.Fatalf("Error making POST request: %v", err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObj map[string]interface{}
	err = json.Unmarshal(responseData, &responseObj)
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("Error parsing response")
		return 0, response.StatusCode
	}

	token, ok := responseObj["token"]
	if !ok {
		fmt.Println("Token not found in response or not a number")
		return 0, response.StatusCode
	}

	return token, response.StatusCode
}

func main() {
	// check if the file is provided
	if len(os.Args) < 5 {
		fmt.Println("Please provide the file path, file content and expected output file")
		os.Exit(1)
	}
	filePath := os.Args[1]
	fileContent := os.Args[2]
	expectedOutputFile := os.Args[3]
	id := os.Args[4]
	token, statusCode := make_submission(filePath, fileContent, expectedOutputFile, id)
	if token == 0 {
		fmt.Println("Error in making submission")
		os.Exit(1)
	}

	fmt.Println(token)
	fmt.Println(statusCode)
}
