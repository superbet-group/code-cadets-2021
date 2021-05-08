package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
)

type application struct {
	Name string
	Age int
	Passed  bool
	Skills []string
}

const applications = "https://run.mocky.io/v3/f7ceece5-47ee-4955-b974-438982267dc8"

func linearBackoff(retry int) time.Duration {
	return time.Duration(retry) * time.Second
}

func containsSkills(Skills []string) bool{
	for i := 0; i<len(Skills); i++{
		if Skills[i]=="Go" || Skills[i]=="Java" {
			return true
		}
	}
	return false
}

func writeApplications(f *os.File, content []application) {
	for i := 0; i < len(content); i++{
		if content[i].Passed && containsSkills(content[i].Skills){
			allSkills := content[i].Skills[0]
			for j := 1; j< len(content[i].Skills); j++{
				allSkills = allSkills + ", " + content[i].Skills[j]
			}

			defer f.WriteString(fmt.Sprint(content[i].Name) + " - " + fmt.Sprint(allSkills) + "\n")
		}
	}
}



func main() {
	httpClient := pester.New()
	httpClient.Backoff = linearBackoff

	httpResponse, err := httpClient.Get(applications)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "HTTP get towards yesno API"),
		)
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "reading body of yesno API response"),
		)
	}

	var decodedContent []application
	err = json.Unmarshal(bodyContent, &decodedContent)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshalling the JSON body content"),
		)
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}
	defer f.Close()
	writeApplications(f, decodedContent)
}