package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

type tgObject struct {
	Link               string    `json:"link"`
	Name               string    `json:"name"`
	NameUpperCamelCase string    `json:"name_upper_camel_case"`
	Description        string    `json:"description"`
	Note               string    `json:"note"`
	ReturnType         string    `json:"return_type"`
	ReturnValue        string    `json:"return_value"`
	List               []string  `json:"list"`
	Fields             []tgField `json:"fields"`
	IsPrimitiveType    bool      `json:"is_primitive_type"`
}

type tgField struct {
	NameSnakeCase      string `json:"name_snake_case"`
	NameUpperCamelCase string `json:"name_upper_camel_case"`
	TypeField          string `json:"type_ field"`
	Required           bool   `json:"required"`
	Description        string `json:"description"`
}

func main() {
	types := []tgObject{}
	params := []tgObject{}

	data, err := getHTML("https://core.telegram.org/bots/api")
	if err != nil {
		panic("Не могу прочитать файл")
	}

	reSplit := regexp.MustCompile(`(?s)<h4`)
	parts := reSplit.Split(data, -1)

	var blocks []string

	for _, p := range parts {
		indexH3 := strings.Index(p, "<h3")
		if indexH3 != -1 {
			p = p[:indexH3]
		}
		block := "<h4" + strings.TrimSpace(p)
		blocks = append(blocks, block)
	}

	for _, b := range blocks {
		h4Tag := getTag(b, "h4")
		name := getContent(h4Tag)
		nameUpperCamelCase := toUpperCamelCase(name)
		if strings.Contains(name, " ") {
			continue
		}

		isTgType := unicode.IsUpper(rune(name[0]))

		link := getAttributeValue(h4Tag, "href")

		pTag := getTag(b, "p")
		desc := getContent(pTag)

		ulTag := getTag(b, "ul")
		listLiTags := getChildren(ulTag, "li")
		var list []string
		for _, li := range listLiTags {
			list = append(list, getContent(li))
		}

		blockquote := getContent(getTag(b, "blockquote"))

		table := getTag(b, "tbody")
		rows := getChildren(table, "tr")

		fields := []tgField{}
		for _, r := range rows {
			cells := getChildren(r, "td")
			fieldNameSnakeCase := getContent(cells[0])
			fieldNameUpperCase := toUpperCamelCase(fieldNameSnakeCase)
			fieldType := convertTypeTgToGo(fieldNameUpperCase, getContent(cells[1]))
			var fieldDesc string
			var fieldRequired bool

			if isTgType {
				fieldDesc = getContent(cells[2])
				if !strings.Contains(fieldDesc, "Optional.") {
					fieldRequired = true
				}
			} else {
				fieldDesc = getContent(cells[3])
				if cells[2] == "Yes" {
					fieldRequired = true
				}

			}
			fields = append(fields, tgField{
				NameSnakeCase:      fieldNameSnakeCase,
				NameUpperCamelCase: fieldNameUpperCase,
				TypeField:          fieldType,
				Required:           fieldRequired,
				Description:        fieldDesc,
			})
		}

		object := tgObject{
			Link:               link,
			Name:               name,
			NameUpperCamelCase: nameUpperCamelCase,
			Description:        desc,
			Note:               blockquote,
			List:               list,
			Fields:             fields,
		}

		if isTgType {
			types = append(types, object)
		} else {
			returnType := getContent(getReturnType(pTag))
			returnValue := getDefaultValue(returnType)
			object.ReturnType = returnType
			object.ReturnValue = returnValue
			object.IsPrimitiveType = unicode.IsLower(rune(returnType[0]))
			params = append(params, object)
		}
	}

	type TemplateData struct {
		Name       string
		Path       string
		OutputPath string
		Data       []tgObject
	}

	tamplatesPath := "./templates/"
	outputDir := "./pkg/"
	typesDir := "types/"
	methodsDir := "core/"
	templatesData := []TemplateData{
		{Name: "types", Path: tamplatesPath, OutputPath: outputDir + typesDir, Data: types},
		{Name: "params", Path: tamplatesPath, OutputPath: outputDir + typesDir, Data: params},
		{Name: "methods", Path: tamplatesPath, OutputPath: outputDir + methodsDir, Data: params},
	}

	_ = os.Mkdir(outputDir+typesDir, os.ModePerm)
	_ = os.Mkdir(outputDir+methodsDir, os.ModePerm)

	for _, td := range templatesData {
		// создание шаблона
		tmpl := createTamplate(td.Path + td.Name + ".txt")

		// создание файла для записи
		f, err := os.Create(td.OutputPath + "gen_" + td.Name + ".go")
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

// получение тега в виде строки
func getTag(text, tag string) string {
	tagStart := "<" + tag
	tagEnd := "</" + tag + ">"

	indexStart := strings.Index(text, tagStart)
	if indexStart == -1 {
		return ""
	}
	indexEnd := strings.Index(text, tagEnd)
	if indexEnd == -1 {
		return text[indexStart:]
	}

	return text[indexStart : indexEnd+len(tagEnd)]
}

// получение значения атрибута
func getAttributeValue(text string, attr string) string {
	indexAttr := strings.Index(text, attr)
	if indexAttr == -1 {
		return ""
	}

	offsetStart := strings.Index(text[indexAttr:], "\"")
	indexStart := offsetStart + indexAttr

	offsetEnd := strings.Index(text[indexStart+1:], "\"")
	indexEnd := offsetEnd + 1 + indexStart

	if indexStart == -1 || indexEnd == -1 {
		return ""
	}

	value := text[indexStart+1 : indexEnd]

	return value
}

// получение текстового контента тега
func getContent(text string) string {
	for {
		indexStart := strings.Index(text, "<")
		indexEnd := strings.Index(text, ">")
		if indexStart == -1 || indexEnd == -1 {
			break
		}
		emoji := ""
		if strings.Contains(text, "img class=\"emoji\"") {
			emoji = getAttributeValue(text, "alt")
		}
		text = text[:indexStart] + emoji + text[indexEnd+1:]
	}

	text = strings.ReplaceAll(text, "\n", "")

	return strings.Trim(text, "\n")
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

func getReturnType(text string) string {
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
	innerTagData := getTag(text, "a")
	if innerTagData != "" {
		if len(prefix) == 0 {
			prefix = "*"
		}
		return prefix + packageName + "." + innerTagData
	}

	innerTagData = getTag(text, "em")
	if innerTagData != "" {
		return prefix + convertTypeTgToGo("", getContent(innerTagData))
	}

	return ""
}

func getDefaultValue(t string) string {
	returnValue := "nil"
	switch t {
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
	return returnValue
}

func getChildren(text string, tag string) []string {
	tagOpen := "<" + tag
	tagClose := "</" + tag + ">"
	count := strings.Count(text, tagOpen)
	if count == 0 {
		return []string{}
	}

	var tags []string
	for range count {
		indexStart := strings.Index(text, tagOpen)
		text = text[indexStart:]
		before, after, found := strings.Cut(text, tagClose)
		if !found {
			break
		}
		tags = append(tags, before+tagClose)
		text = after
	}

	return tags
}

func getHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
