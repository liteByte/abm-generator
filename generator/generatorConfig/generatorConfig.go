package generatorConfig

import (
	"github.com/francoFerraguti/go-abm-generator/structs"
	"github.com/francoFerraguti/go-abm-generator/templates"
	"github.com/liteByte/frango"
	"strings"
)

func Get(projectPath string, needAuthentication bool, config structs.ConfigStruct) string {
	configAuthenticationField := ""
	configAuthenticationValue := ""
	if needAuthentication {
		configAuthenticationField = "JWT_SECRET	string"
		configAuthenticationValue = `JWT_SECRET:		"` + frango.GetRandomString(32) + `"`
	}

	template := templates.ConfigGo()

	fileContent := strings.Replace(template, "&&CONFIG_AUTHENTICATION_FIELD&&", configAuthenticationField, -1)
	fileContent = strings.Replace(fileContent, "&&CONFIG_AUTHENTICATION_VALUE&&", configAuthenticationValue, -1)
	fileContent = strings.Replace(fileContent, "&&CONFIG_PORT&&", config.Port, -1)
	fileContent = strings.Replace(fileContent, "&&CONFIG_DB_TYPE&&", config.DB_TYPE, -1)
	fileContent = strings.Replace(fileContent, "&&CONFIG_DB_USERNAME&&", config.DB_USERNAME, -1)
	fileContent = strings.Replace(fileContent, "&&CONFIG_DB_PASSWORD&&", config.DB_PASSWORD, -1)
	fileContent = strings.Replace(fileContent, "&&CONFIG_DB_HOST&&", config.DB_HOST, -1)
	fileContent = strings.Replace(fileContent, "&&CONFIG_DB_PORT&&", config.DB_PORT, -1)
	fileContent = strings.Replace(fileContent, "&&CONFIG_DB_NAME&&", config.DB_NAME, -1)

	return fileContent
}
