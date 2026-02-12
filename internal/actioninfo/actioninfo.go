package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {

	for _, value := range dataset {

		err := dp.Parse(value)

		if err != nil {
			log.Println("parsing error", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("parsing error", err)
			continue
		}

		fmt.Println(info)
	}
}
