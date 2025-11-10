package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Printf("Ошибка парсинга - %s", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Smthng went wrong - %s", err)
		}
		fmt.Println(info)
	}
}
