package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type jobApiResponse struct {
	Name   string
	Age    int
	Passed bool
	Skills []string
}

const candidates = "https://run.mocky.io/v3/f7ceece5-47ee-4955-b974-438982267dc8"

func fetchContent() []jobApiResponse {
	httpResponse, err := http.Get(candidates)
	//httpClient := &http.Client{}
	//
	//httpResponse, err := httpClient.Get(candidates)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "HTTP get towards job_api API"),
		)
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "reading body of candidates API response"),
		)
	}

	var decodedContent []jobApiResponse
	err = json.Unmarshal(bodyContent, &decodedContent)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}
	return decodedContent
}

func getFile(path string) *os.File {
	f, err := os.Create(path)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}
	return f
}

func formatSkills(slice []string) string {
	var str string
	for _, item := range slice {
		str += ", " + item
	}
	return str[2:]
}

func writeToFile(f *os.File, content []jobApiResponse) {
	for _, val := range content {
		if !val.Passed {
			continue
		}

		skills := formatSkills(val.Skills)
		if strings.Contains(skills, "Go") || strings.Contains(skills, "Java") {
			f.WriteString(val.Name + " - " + skills + "\n")
		}
	}
}

func main() {
	decodedContent := fetchContent()

	f := getFile("output.txt")

	defer f.Close()
	writeToFile(f, decodedContent)
	f.Sync()
	log.Printf("Response from candidates api: %v", decodedContent)
}
