package main

import (
	"os"
	"strings"
)

var index int

type Type struct {
	name        string
	description string
	note        string
	fields      []Field
}

type Field struct {
	name        string
	typeField   string
	description string
}

type Method struct {
	name        string
	description string
	note        string
	parameters  []Parameter
}

type Parameter struct {
	name        string
	typeField   string
	required    bool
	description string
}

func main() {
	data, err := os.ReadFile("./input/doca.html")
	if err != nil {
		panic("Не могу прочитать файл")
	}
	dataString := string(data) + "<h4>"

	var result string

	for {
		// блок от названия до названия
		block := getInnerData(dataString, "<h4>", "<h4>")
		if block == nil {
			break
		}

		indexH3 := strings.Index(block.data, "<h3>")
		if indexH3 != -1 {
			block.data = block.data[:indexH3]
		}

		// ссылка
		link := getInnerData(block.data, "href=\"", "\"")
		result += link.data + "\n"

		// название
		name := getInnerData(block.data, "</a>", "</h4>")
		result += clearString(name.data) + "\n"

		// описание
		description := getInnerData(block.data, "<p>", "</p>")
		result += clearString(description.data) + "\n"

		// таблица
		table := getInnerData(block.data, "<tbody>", "</tbody>")
		if table == nil {
			dataString = dataString[block.indexEnd:]
			continue
		}

		for {
			// строка
			row := getInnerData(table.data, "<tr>", "</tr>")
			if row == nil {
				break
			}

			var rowString string
			for {
				// ячейка
				cell := getInnerData(row.data, "<td>", "</td>")
				if cell == nil {
					break
				}
				rowString += clearString(cell.data) + "\n"
				row.data = row.data[cell.indexEnd:]
			}
			result += rowString + "\n"

			table.data = table.data[row.indexEnd:]
		}

		// заметка
		blockquote := getInnerData(block.data, "<blockquote>", "</blockquote>")
		if blockquote != nil {
			result += clearString(blockquote.data) + "\n"
		}

		dataString = dataString[block.indexEnd:]
	}

	indexStart := strings.Index(result, "#update")
	result = result[indexStart:]

	f, err := os.Create("./output/result.html")
	if err != nil {
		panic("Не получилось создать файл")
	}

	_, err = f.WriteString(result)
	if err != nil {
		panic("Не получилось записать в файл")
	}
}

type InnerDataResult struct {
	indexEnd int
	data     string
}

func getInnerData(dataString, tagStart, tagEnd string) *InnerDataResult {
	indexStart := strings.Index(dataString, tagStart)
	if indexStart == -1 {
		return nil
	}
	indexStart += len(tagStart)

	var indexEnd int
	if tagEnd == "" {
		indexEnd = len(dataString)
	} else {
		indexOffset := strings.Index(dataString[indexStart:], tagEnd)
		if indexOffset == -1 {
			return nil
		}

		indexEnd = indexStart + indexOffset
	}

	result := &InnerDataResult{
		indexEnd: indexEnd,
		data:     dataString[indexStart:indexEnd],
	}

	return result
}

func clearString(line string) string {
	for {
		indexStart := strings.Index(line, "<")
		if indexStart == -1 {
			break
		}
		indexEnd := strings.Index(line, ">")
		if indexEnd == -1 {
			break
		}

		line = line[:indexStart] + line[indexEnd+1:]
	}

	return line
}
