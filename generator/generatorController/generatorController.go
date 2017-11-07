package generatorController

import (
	"github.com/francoFerraguti/go-abm-generator/common"
	"github.com/francoFerraguti/go-abm-generator/structs"
	"github.com/francoFerraguti/go-abm-generator/templates"
	"github.com/liteByte/frango"
	"strings"
)

func Get(projectPath string, model structs.ModelStruct) string {
	template := templates.ControllerGo()

	frangoImport := getFrangoImport(model)
	createString, getListString, getString, updateString, deleteString := getControllerFunctions(model)

	imports := common.GetImports("encoding/json", projectPath+"/models/"+frango.FirstLetterToLower(model.Name), projectPath+"/structs", "github.com/gin-gonic/gin", frangoImport)
	packageName := "package " + frango.FirstLetterToLower(model.Name)

	fileContent := strings.Replace(template, "&&PACKAGE_NAME&&", packageName, -1)
	fileContent = strings.Replace(fileContent, "&&IMPORTS&&", imports, -1)
	fileContent = strings.Replace(fileContent, "&&CREATE&&", createString, -1)
	fileContent = strings.Replace(fileContent, "&&GET_LIST&&", getListString, -1)
	fileContent = strings.Replace(fileContent, "&&UPDATE&&", updateString, -1)
	fileContent = strings.Replace(fileContent, "&&GET&&", getString, -1)
	fileContent = strings.Replace(fileContent, "&&DELETE&&", deleteString, -1)

	return fileContent
}

func getControllerFunctions(model structs.ModelStruct) (string, string, string, string, string) {
	getByString := ""
	deleteByString := ""
	updateByString := ""

	for _, field := range model.Fields {
		if !field.Unique {
			continue
		}

		updateByString += controllerUpdateBy(model, field)
		getByString += controllerGetBy(model, field)
		deleteByString += controllerDeleteBy(model, field)
	}

	return controllerCreate(model), controllerGetList(model), getByString, updateByString, deleteByString
}

func getFrangoImport(model structs.ModelStruct) string {
	frangoImport := ""

	for _, field := range model.Fields {
		if field.Type != "string" && field.Unique {
			frangoImport = "github.com/liteByte/frango"
		}
	}

	return frangoImport
}

//---------

func controllerCreate(model structs.ModelStruct) string {
	createString := ""

	createString += "func Create(c *gin.Context) {\n"
	createString += "	" + frango.FirstLetterToLower(model.Name) + "Struct, err := structs.ParseBodyInto" + model.Name + "Struct(c.Request.Body)\n"
	createString += "	if err != nil {\n"
	createString += "		c.JSON(400, err.Error())\n"
	createString += "		return\n"
	createString += "	}\n\n"
	createString += "	if err = " + frango.FirstLetterToLower(model.Name) + ".Create(" + frango.FirstLetterToLower(model.Name) + "Struct); err != nil {\n"
	createString += "		c.JSON(500, err.Error())\n"
	createString += "		return\n"
	createString += "	}\n\n"
	createString += "	c.JSON(200, `" + model.Name + " created successfully`)\n"
	createString += "}\n\n"

	return createString
}

func controllerGetList(model structs.ModelStruct) string {
	getListString := ""

	getListString += "func GetList(c *gin.Context) {\n"
	getListString += "    " + frango.FirstLetterToLower(model.Name) + "Struct, err := " + frango.FirstLetterToLower(model.Name) + ".GetList()\n"
	getListString += "    if err != nil {\n"
	getListString += "        c.JSON(500, err.Error())\n"
	getListString += "        return\n"
	getListString += "    }\n\n"
	getListString += "    " + frango.FirstLetterToLower(model.Name) + "JSON, _ := json.Marshal(" + frango.FirstLetterToLower(model.Name) + "Struct)\n"
	getListString += "    c.JSON(200, json.RawMessage(string(" + frango.FirstLetterToLower(model.Name) + "JSON)))\n"
	getListString += "}\n\n"

	return getListString
}

func controllerUpdateBy(model structs.ModelStruct, field structs.FieldStruct) string {
	updateString := ""
	beforeConvert := ""
	afterConvert := ""

	if field.Type != "string" {
		beforeConvert = "frango.StringTo" + frango.FirstLetterToUpper(field.Type) + "("
		afterConvert = ")"
	}

	updateString += "func UpdateBy" + frango.FirstLetterToUpper(field.Name) + "(c *gin.Context) {\n"
	updateString += "	" + frango.FirstLetterToLower(model.Name) + "Struct, err := structs.ParseBodyInto" + model.Name + "Struct(c.Request.Body)\n"
	updateString += "	if err != nil {\n"
	updateString += "		c.JSON(400, err.Error())\n"
	updateString += "		return\n"
	updateString += "	}\n\n"
	updateString += "	" + frango.FirstLetterToLower(model.Name) + "Struct." + frango.FirstLetterToUpper(field.Name) + " = " + beforeConvert + "c.Params.ByName(`" + frango.FirstLetterToLower(field.Name) + "`)" + afterConvert + "\n\n"
	updateString += "	if err = " + frango.FirstLetterToLower(model.Name) + ".UpdateBy" + frango.FirstLetterToUpper(field.Name) + "(" + frango.FirstLetterToLower(model.Name) + "Struct); err != nil {\n"
	updateString += "		c.JSON(500, err.Error())\n"
	updateString += "		return\n"
	updateString += "	}\n\n"
	updateString += "	c.JSON(200, `" + model.Name + " updated successfully`)\n"
	updateString += "}\n\n"

	return updateString
}

func controllerGetBy(model structs.ModelStruct, field structs.FieldStruct) string {
	getByString := ""
	beforeConvert := ""
	afterConvert := ""

	if field.Type != "string" {
		beforeConvert = "frango.StringTo" + frango.FirstLetterToUpper(field.Type) + "("
		afterConvert = ")"
	}

	getByString += "func GetBy" + frango.FirstLetterToUpper(field.Name) + "(c *gin.Context) {\n"
	getByString += "    " + frango.FirstLetterToLower(field.Name) + " := " + beforeConvert + "c.Params.ByName(`" + frango.FirstLetterToLower(field.Name) + "`)" + afterConvert + "\n\n"
	getByString += "    " + frango.FirstLetterToLower(model.Name) + "Struct, err := " + frango.FirstLetterToLower(model.Name) + ".GetBy" + frango.FirstLetterToUpper(field.Name) + "(" + frango.FirstLetterToLower(field.Name) + ")\n"
	getByString += "    if err != nil {\n"
	getByString += "        c.JSON(500, err.Error())\n"
	getByString += "        return\n"
	getByString += "    }\n\n"
	getByString += "    " + frango.FirstLetterToLower(model.Name) + "JSON, _ := json.Marshal(" + frango.FirstLetterToLower(model.Name) + "Struct)\n"
	getByString += "    c.JSON(200, json.RawMessage(string(" + frango.FirstLetterToLower(model.Name) + "JSON)))\n"
	getByString += "}\n\n"

	return getByString
}

func controllerDeleteBy(model structs.ModelStruct, field structs.FieldStruct) string {
	deleteByString := ""
	beforeConvert := ""
	afterConvert := ""

	if field.Type != "string" {
		beforeConvert = "frango.StringTo" + frango.FirstLetterToUpper(field.Type) + "("
		afterConvert = ")"
	}

	deleteByString += "func DeleteBy" + frango.FirstLetterToUpper(field.Name) + "(c *gin.Context) {\n"
	deleteByString += "    " + frango.FirstLetterToLower(field.Name) + " := " + beforeConvert + "c.Params.ByName(`" + frango.FirstLetterToLower(field.Name) + "`)" + afterConvert + "\n\n"
	deleteByString += "    if err := " + frango.FirstLetterToLower(model.Name) + ".DeleteBy" + frango.FirstLetterToUpper(field.Name) + "(" + frango.FirstLetterToLower(field.Name) + "); err != nil {\n"
	deleteByString += "        c.JSON(500, err.Error())\n"
	deleteByString += "        return\n"
	deleteByString += "    }\n\n"
	deleteByString += "    c.JSON(200, `" + model.Name + " deleted successfully`)\n"
	deleteByString += "}\n\n"

	return deleteByString
}
