package main

import (
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./input/doca.html")
	if err != nil {
		panic("Не могу прочитать файл")
	}
	dataString := string(data)

	var result string

	for {
		tagOpen := "<h4>"
		tagClose := "</h4>"

		indexStart := strings.Index(dataString, tagOpen)
		if indexStart == -1 {
			break
		}

		indexEndOffsetByStart := strings.Index(dataString[indexStart:], tagClose)
		indexEnd := indexStart + indexEndOffsetByStart + len(tagClose)

		result += dataString[indexStart:indexEnd] + "\n"
		dataString = dataString[indexEnd:]
	}

	f, err := os.Create("./output/result.html")
	if err != nil {
		panic("Не получилось создать файл")
	}

	_, err = f.WriteString(result)
	if err != nil {
		panic("Не получилось записать в файл")
	}
}
