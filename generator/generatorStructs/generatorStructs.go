package generatorStructs

import (
	"github.com/francoFerraguti/go-abm-generator/structs"
	"github.com/francoFerraguti/go-abm-generator/templates"
	"github.com/liteByte/frango"
	"strings"
)

func Get(projectPath string, needAuthentication bool, models []structs.ModelStruct) string {
	structsString := getStructsString(models)
	authStructString := getAuthStructString(models, needAuthentication)

	template := templates.StructsGo()

	fileContent := strings.Replace(template, "&&STRUCTS&&", structsString, -1)
	fileContent = strings.Replace(fileContent, "&&AUTH_STRUCTS&&", authStructString, -1)

	return fileContent
}

func getAuthStructString(models []structs.ModelStruct, needAuthentication bool) string {
	authStructString := ""

	if needAuthentication {
		authStructString += "type LoginStruct struct {\n"

		for _, model := range models {
			for _, field := range model.Fields {
				if field.AuthenticationUsername || field.AuthenticationPassword {
					authStructString += "	" + frango.FirstLetterToUpper(field.Name) + " " + field.Type + "\n"
				}
			}
		}

		authStructString += "}\n\n"

		authStructString += "func ParseBodyIntoLoginStruct(body io.ReadCloser) (LoginStruct, error) {\n"
		authStructString += "    bodyBytes, _ := ioutil.ReadAll(body)\n"
		authStructString += "    loginStruct := LoginStruct{}\n"
		authStructString += "    err := json.Unmarshal(bodyBytes, &loginStruct)\n"
		authStructString += "    return loginStruct, err\n"
		authStructString += "}\n\n"
	}

	return authStructString
}

func getStructsString(models []structs.ModelStruct) string {
	structsString := ""

	for _, model := range models {
		structsString += getStructFirstPart(model)
		structsString += getStructSecondPart(model)
	}

	return structsString
}

func getStructFirstPart(model structs.ModelStruct) string {
	structsString := ""

	structsString += "type " + model.Name + "Struct struct {\n"
	for _, field := range model.Fields {

		if field.Type == "float" {
			field.Type = "float64"
		}

		structsString += "	" + frango.FirstLetterToUpper(field.Name) + " " + field.Type + "\n"
	}
	structsString += "}\n\n"

	return structsString
}

func getStructSecondPart(model structs.ModelStruct) string {
	structsString := ""

	structsString += "func ParseBodyInto" + model.Name + "Struct(body io.ReadCloser) (" + model.Name + "Struct, error) {\n"
	structsString += "    bodyBytes, _ := ioutil.ReadAll(body)\n"
	structsString += "    " + frango.FirstLetterToLower(model.Name) + "Struct := " + model.Name + "Struct{}\n"
	structsString += "    err := json.Unmarshal(bodyBytes, &" + frango.FirstLetterToLower(model.Name) + "Struct)\n"
	structsString += "    return " + frango.FirstLetterToLower(model.Name) + "Struct, err\n"
	structsString += "}\n\n"

	return structsString
}
