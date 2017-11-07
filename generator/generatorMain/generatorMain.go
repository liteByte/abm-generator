package generatorMain

import (
	"github.com/francoFerraguti/go-abm-generator/common"
	"github.com/francoFerraguti/go-abm-generator/templates"
	"strings"
)

func Get(projectPath string) string {
	template := templates.MainGo()

	imports := common.GetImports(projectPath+"/dbhandler", projectPath+"/router")

	fileContent := strings.Replace(template, "&&IMPORTS&&", imports, -1)

	return fileContent
}
