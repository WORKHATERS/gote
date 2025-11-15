package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Base struct {
	Link 		string
	Name        string
	Description string
	Note        string
}

type Type struct {
	Base
	Fields      []Field
}

type Field struct {
	Name        string
	TypeField   string
	Description string
}

type Method struct {
	Base
	Parameters  []Parameter
}

type Parameter struct {
	Name        string
	TypeField   string
	Required    bool
	Description string
}

func main() {
	data, err := os.ReadFile("./input/doca.html")
	if err != nil {
		panic("Не могу прочитать файл")
	}
	dataString := string(data) + "<h4>"

	var base Base
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

		// название
		name := getInnerData(block.data, "</a>", "</h4>")

		// флаг
		var isField bool
		if unicode.IsUpper(rune(name.data[0])) {
			isField = true
		}

		// описание
		description := getInnerData(block.data, "<p>", "</p>")

		// заметка
		blockquote := getInnerData(block.data, "<blockquote>", "</blockquote>")

		// база
		base.Link = link.data
		base.Name = clearString(name.data)
		base.Description = clearString(description.data)
		if blockquote != nil {
			base.Note = clearString(blockquote.data)
		} else { base.Note = "" }

		var fields []Field
		var parameters []Parameter

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
			var cellArr []string
			var rowString string
			for {
				// ячейка
				cell := getInnerData(row.data, "<td>", "</td>")
				if cell == nil {
					break
				}
				rowString += clearString(cell.data) + " | "
				cellArr = append(cellArr, clearString(cell.data))
				row.data = row.data[cell.indexEnd:]
			}

			// заполнение полей | параметров
			if isField {
				fields = append(fields, Field {
					Name:			clearString(cellArr[0]),
					TypeField: 		clearString(cellArr[1]),
					Description: 	clearString(cellArr[2]),
				})
			} else {
				requiredString := cellArr[2]
				var required bool
				if requiredString == "Yes" {
					required = true
				}

				parameters = append(parameters, Parameter {
					Name:			clearString(cellArr[0]),
					TypeField:		clearString(cellArr[1]),
					Required:		required,
					Description:	clearString(cellArr[3]),
				})
			}
			
			table.data = table.data[row.indexEnd:]
		}

		// заполнение структур
		if isField {
			var TGType Type
			TGType.Base = base
			TGType.Fields = fields

			result += writeBase(TGType.Link, TGType.Name, TGType.Description, TGType.Note)
			for _, f := range TGType.Fields {
				result += writeField(f.Name, f.TypeField, f.Description)
			}
		} else {
			var TGMethod Method
			TGMethod.Base = base
			TGMethod.Parameters = parameters

			result += writeBase(TGMethod.Link, TGMethod.Name, TGMethod.Description, TGMethod.Note)
			for _, p := range TGMethod.Parameters {
				result += writeParameter(p.Name, p.TypeField, p.Description, p.Required)
			}
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

func writeBase(link, name, desc, note string) string {
	return fmt.Sprintf("\n%s\n%s\n%s\n%s\n", link, name, desc, note)
}

func writeField(name, typeField, desc string) string {
	return fmt.Sprintf("%s | %s | %s\n", name, typeField, desc)
}

func writeParameter(name, typeField, desc string, required bool) string {
	requiredString := strconv.FormatBool(required)
	return fmt.Sprintf("%s | %s | %s | %s\n", name, typeField, requiredString, desc)
}
