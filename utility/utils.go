package utility

import (
	"log"

	"github.com/go-routeros/routeros"
)

type M map[string]string

func ParseResponse(response *routeros.Reply) (interface{}, error) {
	var parsedResults interface{};

	if len(response.Re) > 0 {
		var records []M
		for _, re := range response.Re {
			records = append(records, re.Map)
		}
		parsedResults = records
	} else if len(response.Done.Map) > 0 {
		parsedResults = response.Done.Map
	} else {
		parsedResults = []interface{}{}
	}
	return parsedResults, nil
}

func ErrorParser(err error) error {
	if err != nil {
		log.Printf("Error parsing response: %v", err)
	}
	return err
}

func ErrorCommand(err error) error {
	if err != nil {
		log.Printf("Error running command: %v", err)
	}
	return err
}
