package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for i, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("Ошибка парсинга данных в позиции %d ('%s'): %v", i, data, err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка формирования информации в позиции %d: %v", i, err)
			continue
		}

		fmt.Println(info)
	}
}
