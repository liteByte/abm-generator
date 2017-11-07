package generatorMiddleware

import (
	"github.com/francoFerraguti/go-abm-generator/common"
	"github.com/francoFerraguti/go-abm-generator/structs"
	"github.com/francoFerraguti/go-abm-generator/templates"
	"github.com/liteByte/frango"
	"strings"
)

func Get(projectPath string, models []structs.ModelStruct) string {
	usernameFieldLower := ""
	usernameFieldUpper := ""

	for _, model := range models {
		for _, field := range model.Fields {
			if field.AuthenticationUsername {
				usernameFieldLower = frango.FirstLetterToLower(field.Name)
				usernameFieldUpper = frango.FirstLetterToUpper(field.Name)
			}
		}
	}

	template := templates.MiddlewareGo()

	imports := common.GetImports("github.com/gin-gonic/gin", projectPath+"/authentication")

	fileContent := strings.Replace(template, "&&IMPORTS&&", imports, -1)
	fileContent = strings.Replace(fileContent, "&&USERNAME_FIELD_LOWER&&", usernameFieldLower, -1)
	fileContent = strings.Replace(fileContent, "&&USERNAME_FIELD_UPPER&&", usernameFieldUpper, -1)

	return fileContent
}
