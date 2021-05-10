package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

type jobApiResponse struct {
	Name   string
	Age    int
	Passed bool
	Skills []string
}

const candidates = "https://run.mocky.io/v3/f7ceece5-47ee-4955-b974-438982267dc8"

func fetchContent() ([]jobApiResponse, error) {
	httpResponse, err := http.Get(candidates)
	//httpClient := &http.Client{}
	//
	//httpResponse, err := httpClient.Get(candidates)
	if err != nil {
		return nil, err
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var decodedContent []jobApiResponse
	err = json.Unmarshal(bodyContent, &decodedContent)
	if err != nil {
		return decodedContent, err
	}

	return decodedContent, nil
}

func getFile(path string) (*os.File, error) {
	f, err := os.Create(path)
	if err != nil {
		return f, err
	}

	return f, nil
}

func setContains(array []string, strings ...string) bool {
	set := make(map[string]struct{})
	for _, val := range array {
		set[val] = struct{}{}
	}

	for _, str := range strings {
		_, exists := set[str]
		if exists {
			return true
		}
	}

	return false
}

func formatSkills(skills []string) string {
	formattedSkills := ""
	for _, val := range skills {
		formattedSkills += val + ","
	}
	formattedSkills = formattedSkills[:len(formattedSkills)-1]

	return formattedSkills
}

func writeToFile(f *os.File, content []jobApiResponse) {
	for _, val := range content {
		if !val.Passed {
			continue
		}
		skills := val.Skills

		isContained := setContains(skills, "Java", "Go")

		if isContained {
			formattedSkills := formatSkills(skills)

			f.WriteString(val.Name + " - " + formattedSkills + "\n")
		}
	}
}

func main() {
	decodedContent, err := fetchContent()
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "Unable to fetch content"),
		)
	}

	f, err := getFile("output.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "Unable to read file"),
		)
	}

	defer f.Close()
	writeToFile(f, decodedContent)
	f.Sync()
	log.Printf("Response from candidates api: %v", decodedContent)
}
