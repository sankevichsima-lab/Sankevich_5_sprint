package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(str string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	if len(dataset) == 0 {
		return
	}

	for i := 0; i < len(dataset); i++ {
		err := dp.Parse(dataset[i])

		if err != nil {
			log.Printf("%v", err)
		}

		result, err := dp.ActionInfo()

		if err != nil {
			log.Printf("%v", err)
		}

		fmt.Println(result)
	}
}
