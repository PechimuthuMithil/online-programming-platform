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

func make_submission(filePath string, fileContent string, expectedOutputFile string, id string) (string, int) {
	//convert id which is a string to int
	if expectedOutputFile == "" {
		return "No expected output file provided", 400
	}
	if fileContent == "" && filePath == "" {
		//return err
		return "No file content or file path provided", 400
	}
	if fileContent == "" && filePath != "" {
		inputFile, err := ioutil.ReadFile(filePath)
		fileContent = string(inputFile)
		if err != nil {
			fmt.Print(err.Error())
			return "Error reading file", 400
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
	fmt.Println(response)

	if err != nil {
		log.Fatal(err)
	}
	return string(responseData), response.StatusCode
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
	fmt.Println(filePath)
	fmt.Println(fileContent)
	fmt.Println(expectedOutputFile)
	fmt.Println(id)

	response, statusCode := make_submission(filePath, fileContent, expectedOutputFile, id)
	fmt.Println(response)
	fmt.Println(statusCode)
}
