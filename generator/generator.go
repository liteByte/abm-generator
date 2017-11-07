package generator

import (
	"github.com/francoFerraguti/go-abm-generator/generator/generatorAuthentication"
	"github.com/francoFerraguti/go-abm-generator/generator/generatorAuthenticationController"
	"github.com/francoFerraguti/go-abm-generator/generator/generatorConfig"
	"github.com/francoFerraguti/go-abm-generator/generator/generatorController"
	"github.com/francoFerraguti/go-abm-generator/generator/generatorDBHandler"
	"github.com/francoFerraguti/go-abm-generator/generator/generatorDocumentation"
	"github.com/francoFerraguti/go-abm-generator/generator/generatorMain"
	"github.com/francoFerraguti/go-abm-generator/generator/generatorMiddleware"
	"github.com/francoFerraguti/go-abm-generator/generator/generatorModel"
	"github.com/francoFerraguti/go-abm-generator/generator/generatorRouter"
	"github.com/francoFerraguti/go-abm-generator/generator/generatorSchema"
	"github.com/francoFerraguti/go-abm-generator/generator/generatorStructs"

	"github.com/francoFerraguti/go-abm-generator/structs"
)

func GetMain(projectPath string) string {
	return generatorMain.Get(projectPath)
}

func GetMiddleware(projectPath string, models []structs.ModelStruct) string {
	return generatorMiddleware.Get(projectPath, models)
}

func GetConfig(projectPath string, needAuthentication bool, config structs.ConfigStruct) string {
	return generatorConfig.Get(projectPath, needAuthentication, config)
}

func GetRouter(projectPath string, needAuthentication bool, models []structs.ModelStruct) string {
	return generatorRouter.Get(projectPath, needAuthentication, models)
}

func GetStructs(projectPath string, needAuthentication bool, models []structs.ModelStruct) string {
	return generatorStructs.Get(projectPath, needAuthentication, models)
}

func GetDBHandler(projectPath string, models []structs.ModelStruct) string {
	return generatorDBHandler.Get(projectPath, models)
}

func GetSchema(models []structs.ModelStruct) string {
	return generatorSchema.Get(models)
}

func GetDocumentation(needAuthentication bool, models []structs.ModelStruct) string {
	return generatorDocumentation.Get(needAuthentication, models)
}

func GetAuthentication(projectPath string, models []structs.ModelStruct) string {
	return generatorAuthentication.Get(projectPath, models)
}

func GetAuthenticationController(projectPath string, models []structs.ModelStruct) string {
	return generatorAuthenticationController.Get(projectPath, models)
}

func GetModel(projectPath string, needAuthentication bool, model structs.ModelStruct) string {
	return generatorModel.Get(projectPath, needAuthentication, model)
}

func GetController(projectPath string, model structs.ModelStruct) string {
	return generatorController.Get(projectPath, model)
}
