package generator

import (
	"log"
	"os"
	"strings"
	"text/template"
	"unicode"
)

type TgObject struct {
	Link               string
	Name               string
	NameUpperCamelCase string
	Description        string
	List               []string
	Note               string
	ReturnType         string
	IsPrimitiveType    bool
	ReturnValue        string
	Fields             []TgField
}

type TgField struct {
	NameSnakeCase      string
	NameUpperCamelCase string
	TypeField          string
	Required           bool
	Description        string
}

func Generate() {
	types := []TgObject{}
	params := []TgObject{}

	data, err := os.ReadFile("./input/doca.html")
	if err != nil {
		panic("Не могу прочитать файл")
	}
	dataString := string(data) + "<h4>"

	// перебор блоков от h4 до h4
	for {

		// блок от названия до названия
		block := getTagData(dataString, "<h4>", "<h4>")
		if block == nil {
			break
		}

		// проверка на наличие в блоке h3
		indexH3 := strings.Index(block.data, "<h3>")
		if indexH3 != -1 {
			block.data = block.data[:indexH3]
		}

		// ссылка
		link := getAttributeValue(block.data, "href")

		// название
		name := getTagData(block.data, "</a>", "</h4>")
		if strings.Contains(strings.Trim(name.data, " "), " ") {
			dataString = dataString[block.indexEnd:]
			continue
		}

		// описание
		description := getTagData(block.data, "<p>", "</p>")

		// список
		var list []string
		listHtml := getTagData(block.data, "<ul>", "</ul>")
		if listHtml != nil {
			list = makeListFromTags(listHtml.data)
		}

		// заметка
		blockquote := getTagData(block.data, "<blockquote>", "</blockquote>")
		var note string
		if blockquote != nil {
			note = clearString(blockquote.data)
		}

		// Telegram объект
		tgObject := TgObject{
			Link:               link,
			Name:               clearString(name.data),
			NameUpperCamelCase: toUpperCamelCase(clearString(name.data)),
			Description:        clearString(description.data),
			Note:               note,
			List:               list,
		}

		// проверка типа по первой букве имени
		isType := unicode.IsUpper(rune(tgObject.Name[0]))

		// если метод, то добавить возвращаемый тип "ReturnType"
		if !isType {
			tgObject.ReturnType = searchReturnType(description.data)
			tgObject.IsPrimitiveType = unicode.IsLower(rune(tgObject.ReturnType[0]))
			returnValue := "nil"
			switch tgObject.ReturnType {
			case "bool":
				{
					returnValue = "false"
				}
			case "string":
				{
					returnValue = "\"\""
				}
			case "int64":
				{
					returnValue = "0"
				}
			}
			tgObject.ReturnValue = returnValue

		}

		// таблица
		table := getTagData(block.data, "<tbody>", "</tbody>")
		if table == nil {
			table = &TagDataResult{}
		}

		// перебор строк
		var fields []TgField
		for {
			row := getTagData(table.data, "<tr>", "</tr>")
			if row == nil {
				break
			}

			// перебор ячеек
			var cells []string
			for {
				cell := getTagData(row.data, "<td>", "</td>")
				if cell == nil {
					break
				}
				cells = append(cells, clearString(cell.data))

				// срез строки
				row.data = row.data[cell.indexEnd:]
			}

			// составление поля
			fieldNameSnakeCase := clearString(cells[0])
			fieldNameUpperCamelCase := toUpperCamelCase(fieldNameSnakeCase)
			fieldType := convertTypeTgToGo(fieldNameUpperCamelCase, clearString(cells[1]))
			var fieldRequire bool
			var fieldDescription string

			if isType {
				fieldDescription = clearString(cells[2])
				if !strings.Contains(fieldDescription, "Optional.") {
					fieldRequire = true
				}
			} else {
				requiredString := cells[2]
				if requiredString == "Yes" {
					fieldRequire = true
				}
				fieldDescription = clearString(cells[3])
			}

			fields = append(fields, TgField{
				NameSnakeCase:      fieldNameSnakeCase,
				NameUpperCamelCase: fieldNameUpperCamelCase,
				TypeField:          fieldType,
				Required:           fieldRequire,
				Description:        fieldDescription,
			})

			// срез таблицы
			table.data = table.data[row.indexEnd:]
		}

		tgObject.Fields = fields

		if isType {
			types = append(types, tgObject)
		} else {
			params = append(params, tgObject)
		}

		// срез всего файла
		dataString = dataString[block.indexEnd:]
	}

	type TemplateData struct {
		Name       string
		Path       string
		OutputPath string
		Data       []TgObject
	}

	tamplatesPath := "./internal/generator/"
	templatesData := []TemplateData{
		{Name: "types", Path: tamplatesPath, OutputPath: "./pkg/types/", Data: types},
		{Name: "params", Path: tamplatesPath, OutputPath: "./pkg/types/", Data: params},
		{Name: "methods", Path: tamplatesPath, OutputPath: "./internal/bot/", Data: params},
	}

	for _, td := range templatesData {
		// создание шаблона
		tmpl := createTamplate(td.Path + td.Name + ".txt")

		// создание файла для записи
		f, err := os.Create(td.OutputPath + td.Name + ".go")
		if err != nil {
			panic("Не получилось создать файл")
		}

		// запись шаблона в файл
		err = tmpl.Execute(f, td.Data)
		if err != nil {
			panic(err)
		}
	}
}

func createTamplate(path string) *template.Template {
	// чтение файла с шаблоном
	dataTemplate, err := os.ReadFile(path)
	if err != nil {
		panic("Не могу прочитать файл")
	}

	// создание шаблона
	tmpl, err := template.New(path).Parse(string(dataTemplate))
	if err != nil {
		panic(err)
	}

	return tmpl
}

type TagDataResult struct {
	indexEnd int
	data     string
}

func getTagData(dataString, tagStart, tagEnd string) *TagDataResult {
	indexStart := strings.Index(dataString, tagStart[:len(tagStart)-1])
	if indexStart == -1 {
		return nil
	}
	indexStart += strings.Index(dataString[indexStart:], ">") + 1
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

	result := &TagDataResult{
		indexEnd: indexEnd,
		data:     dataString[indexStart:indexEnd],
	}

	return result
}

func getAttributeValue(text string, attr string) string {
	indexAttr := strings.Index(text, attr)

	offsetStart := strings.Index(text[indexAttr:], "\"")
	indexStart := offsetStart + indexAttr + 1

	offsetEnd := strings.Index(text[indexStart:], "\"")
	indexEnd := offsetEnd + indexStart

	value := text[indexStart+1 : indexEnd]

	return value
}

func clearString(line string) string {
	for {
		indexStart := strings.Index(line, "<")
		indexEnd := strings.Index(line, ">")
		if indexStart == -1 || indexEnd == -1 {
			break
		}

		line = line[:indexStart] + line[indexEnd+1:]
	}

	line = strings.ReplaceAll(line, "\n", "")

	return strings.Trim(line, "\n")
}

func toUpperCamelCase(name string) string {
	result := strings.ToUpper(string(name[0])) + name[1:]

	for {
		index := strings.Index(result, "_")
		if index == -1 {
			break
		}
		letter := strings.ToUpper(string(result[index+1]))
		result = result[:index] + letter + result[index+2:]
	}

	return result
}

func convertTypeTgToGo(n, t string) string {
	prefix := ""
	for {
		if strings.Contains(t, "Array of") {
			prefix += "[]"
			t = t[9:]
		} else {
			break
		}
	}

	if strings.Count(t, " or ") > 1 {
		return n
	}

	index := strings.Index(t, " or ")
	if index != -1 {
		t = strings.Trim(t[:index], " ")
	}

	if strings.Contains(t, " and ") || strings.Contains(t, ",") {
		return "any"
	}

	result := t

	switch t {
	case "Integer", "Int":
		result = "int64"
	case "Float":
		result = "float64"
	case "String":
		result = "string"
	case "Boolean", "True", "False":
		result = "bool"
	default:
		if len(prefix) == 0 {
			prefix = "*"
		}
	}

	return prefix + result
}

func searchReturnType(text string) string {
	anchorWords := []string{
		"On success",
		"Returns",
	}
	var indexAnchorWord int

	for _, word := range anchorWords {
		indexAnchorWord = strings.LastIndex(text, word)
		if indexAnchorWord != -1 {
			break
		}
	}

	if indexAnchorWord == -1 {
		log.Println("Возвращаемый тип не найден в:", text)
		return ""
	}

	text = text[indexAnchorWord:]

	prefix := ""
	for {
		indexArray := strings.Index(text, "Array of")
		if indexArray == -1 {
			break
		}
		prefix += "[]"
		text = text[indexArray+9:]
	}

	packageName := "types"
	innerTagData := getTagData(text, "<a>", "</a>")
	if innerTagData != nil {
		if len(prefix) == 0 {
			prefix = "*"
		}
		return prefix + packageName + "." + innerTagData.data
	}

	innerTagData = getTagData(text, "<em>", "</em>")
	if innerTagData != nil {
		return prefix + convertTypeTgToGo("", innerTagData.data)
	}

	return ""
}

func makeListFromTags(text string) []string {
	listTags := []string{}
	for {
		li := getTagData(text, "<li>", "</li>")
		if li == nil {
			break
		}
		listTags = append(listTags, clearString(li.data))
		text = text[li.indexEnd:]
	}

	return listTags
}
