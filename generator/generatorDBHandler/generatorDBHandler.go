package generatorDBHandler

import (
	"github.com/francoFerraguti/go-abm-generator/common"
	"github.com/francoFerraguti/go-abm-generator/structs"
	"github.com/francoFerraguti/go-abm-generator/templates"
	"strings"
)

func Get(projectPath string, models []structs.ModelStruct) string {
	template := templates.DBHandlerGo()

	imports := common.GetImports("database/sql", "github.com/go-sql-driver/mysql", "github.com/liteByte/frango", projectPath + "/config")
	schemaFunctions := getSchemaFunctions(models)

	fileContent := strings.Replace(template, "&&IMPORTS&&", imports, -1)
	fileContent = strings.Replace(fileContent, "&&SCHEMA_FUNCTIONS&&", schemaFunctions, -1)

	return fileContent
}

func getSchemaFunctions(models []structs.ModelStruct) string {
	schemaString := ""

	for _, model := range models {
		schemaString += "	create" + model.Name + "Table()\n"
	}

	return schemaString
}

