package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode/utf8"

	"github.com/pkg/errors"
	"github.com/sethgrid/pester"
)

type response struct {
	Name   string
	Age    int
	Passed bool
	Skills []string
}

func stringInSlice(a []string, list []string) bool {
	for _, x := range list {
		for _, b := range a {
			if b == x {
				return true
			}
		}
	}
	return false
}

func RemoveLastChar(str string) string {
	for len(str) > 0 {
		_, size := utf8.DecodeLastRuneInString(str)
		return str[:len(str)-size]
	}
	return str
}

func main() {
	httpClient := pester.New()

	httpResponse, err := httpClient.Get("https://run.mocky.io/v3/f7ceece5-47ee-4955-b974-438982267dc8")

	if err != nil {
		log.Fatal(
			errors.WithMessage(err, `HTTP error`),
		)
	}

	bodyContent, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "reading body of yesno API response"),
		)
	}

	var decode []response
	err = json.Unmarshal(bodyContent, &decode)
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "unmarshaling"),
		)
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}

	contained := []string{"Go", "Java"}

	for _, val := range decode {
		if val.Passed {
			if stringInSlice(contained, val.Skills) {
				skills := ""
				for _, x := range val.Skills {
					skills += x + ","
				}

				skills = RemoveLastChar(skills)
				f.WriteString(fmt.Sprint(val.Name) + "-" + skills + "\n")
			}
		}
	}

}
