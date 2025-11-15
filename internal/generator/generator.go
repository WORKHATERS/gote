package generator

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type TgType struct {
	Link        string
	Name        string
	Description string
	Note        string
	Fields      []TgField
}

type TgField struct {
	Name        string
	TypeField   string
	Required    bool
	Description string
}

func Generate() {
	data, err := os.ReadFile("./input/doca.html")
	if err != nil {
		panic("Не могу прочитать файл")
	}
	dataString := string(data) + "<h4>"

	resultTypes := "package types\n"
	resultParams := "package types\n type ReplyMarkup interface {}\n"
	resultMethods := `package bot
import (
	"context"
	"encoding/json"
	"gote/pkg/types"
)

const URL = "https://api.telegram.org/bot"` + "\n\n"

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

		if strings.Contains(strings.Trim(name.data, " "), " ") {
			dataString = dataString[block.indexEnd:]
			continue
		}

		// описание
		description := getInnerData(block.data, "<p>", "</p>")

		// заметка
		blockquote := getInnerData(block.data, "<blockquote>", "</blockquote>")
		var note string
		if blockquote != nil {
			note = clearString(blockquote.data)
		}

		// база
		base := TgType{
			Link:        link.data,
			Name:        clearString(name.data),
			Description: clearString(description.data),
			Note:        note,
		}

		var isType bool
		if unicode.IsUpper(rune(base.Name[0])) {
			isType = true
		}
		// таблица
		table := getInnerData(block.data, "<tbody>", "</tbody>")
		if table == nil {
			table = &InnerDataResult{}
			// dataString = dataString[block.indexEnd:]
			// continue
		}

		var fields []TgField

		for {
			// строка
			row := getInnerData(table.data, "<tr>", "</tr>")
			if row == nil {
				break
			}
			var cellArr []string
			for {
				// ячейка
				cell := getInnerData(row.data, "<td>", "</td>")
				if cell == nil {
					break
				}
				cellArr = append(cellArr, clearString(cell.data))
				row.data = row.data[cell.indexEnd:]
			}

			fieldName := clearString(cellArr[0])
			fieldType := convertType(convertName(fieldName), clearString(cellArr[1]))
			var fieldRequire bool
			var fieldDescription string

			if isType {
				fieldDescription = clearString(cellArr[2])
				if !strings.Contains(fieldDescription, "Optional.") {
					fieldRequire = true
				}
			} else {
				requiredString := cellArr[2]
				if requiredString == "Yes" {
					fieldRequire = true

				}
				fieldDescription = clearString(cellArr[3])
			}

			fields = append(fields, TgField{
				Name:        fieldName,
				TypeField:   fieldType,
				Required:    fieldRequire,
				Description: fieldDescription,
			})

			table.data = table.data[row.indexEnd:]
		}

		base.Fields = fields

		if isType {
			resultTypes += stringify(base)
		} else {
			resultParams += stringify(base)
			resultMethods += stringifyMethod(base)
		}

		dataString = dataString[block.indexEnd:]
	}

	outputDir := "./pkg/types/"

	filePath := outputDir + "types.go"
	f, err := os.Create(filePath)
	if err != nil {
		panic("Не получилось создать файл")
	}

	_, err = f.WriteString(resultTypes)
	if err != nil {
		panic("Не получилось записать в файл")
	}

	filePath = outputDir + "methods.go"
	f, err = os.Create(filePath)
	if err != nil {
		panic("Не получилось создать файл")
	}

	_, err = f.WriteString(resultParams)
	if err != nil {
		panic("Не получилось записать в файл")
	}

	filePath = "./internal/bot/methods.go"
	f, err = os.Create(filePath)
	if err != nil {
		panic("Не получилось создать файл")
	}

	_, err = f.WriteString(resultMethods)
	if err != nil {
		panic("Не получилось записать в файл")
	}

	// cmd := exec.Command("gofmt", "./output/result.go")
	// r, _ := cmd.Output()

	// _ = f.Truncate(0)
	// _, _ = f.Seek(0, 0)

	// _, err = f.WriteString(r)
	// if err != nil {
	// 	panic("Не получилось записать в файл")
	// }

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

	line = strings.ReplaceAll(line, "\n", "")

	return strings.Trim(line, "\n")
}

func stringify(base TgType) string {
	desc := fmt.Sprintf("// %s\n", base.Description)
	note := "//\n"
	if base.Note != "" {
		note = fmt.Sprintf("//\n// %s\n//\n", base.Note)
	}
	link := fmt.Sprintf("// https://core.telegram.org/bots/api%s\n", base.Link)
	comment := desc + note + link

	name := convertName(base.Name)
	signa := fmt.Sprintf("type %s struct {\n", name)

	var fields string
	if base.Fields != nil {
		for _, f := range base.Fields {

			req := ""
			if !f.Required {
				req = ",omitempty"
			}
			fields += fmt.Sprintf("\t%s\t%s\t`json:\"%s%s\"`\n", convertName(f.Name), f.TypeField, f.Name, req)
		}

	}

	return comment + signa + fields + "}\n\n"
}

func stringifyMethod(base TgType) string {
	comment := fmt.Sprintf("// https://core.telegram.org/bots/api%s\n", base.Link)
	name := convertName(base.Name)
	signa := fmt.Sprintf("func (bot *Bot) %s(ctx context.Context, param types.%s) (*types.User, error) {", name, name)
	body := fmt.Sprintf(`
	data, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	url := URL + bot.Token + "/%s"
	resp, err := requestWithContext(ctx, url, data)
	if err != nil {
		return nil, err
	}

	var result TGResponse[types.User]
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if err := resp.Body.Close(); err != nil {
		return &result.Result, err
	}

	return &result.Result, nil
}`, name)
	return comment + signa + body + "\n\n"
}

func convertName(name string) string {
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

func convertType(n, t string) string {
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
	case "Integer":
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
