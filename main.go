package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/liteByte/frango"

	"github.com/francoFerraguti/go-abm-generator/generator"
	"github.com/francoFerraguti/go-abm-generator/structs"
)

var router *gin.Engine

func main() {
	router = gin.New()

	router.POST("/create", create)

	router.Run(":8000")
}

func create(c *gin.Context) {
	data, err := structs.ParseBody(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	projectName := getProjectName(data.ProjectPath)
	needAuthentication := data.NeedAuthentication
	parentFolder := "temp/" + projectName

	createFolderStructure(parentFolder, needAuthentication)
	createFiles(parentFolder, data.ProjectPath, needAuthentication, data.Config, data.Models)

	c.JSON(200, projectName+" created successfully")
}

func createFolderStructure(parentFolder string, needAuthentication bool) {
	frango.CreateFolder("temp")
	frango.CreateFolder(parentFolder)
	frango.CreateFolder(parentFolder + "/config")
	frango.CreateFolder(parentFolder + "/router")
	frango.CreateFolder(parentFolder + "/controllers")
	frango.CreateFolder(parentFolder + "/models")
	frango.CreateFolder(parentFolder + "/dbhandler")
	frango.CreateFolder(parentFolder + "/structs")

	if needAuthentication {
		frango.CreateFolder(parentFolder + "/authentication")
		frango.CreateFolder(parentFolder + "/middleware")
		frango.CreateFolder(parentFolder + "/controllers/authentication")
	}
}

func createFiles(parentFolder string, projectPath string, needAuthentication bool, config structs.ConfigStruct, models []structs.ModelStruct) {
	frango.CreateFile(parentFolder+"/main.go", generator.GetMain(projectPath))
	frango.CreateFile(parentFolder+"/config/config.go", generator.GetConfig(projectPath, needAuthentication, config))
	frango.CreateFile(parentFolder+"/dbhandler/dbhandler.go", generator.GetDBHandler(projectPath, models))
	frango.CreateFile(parentFolder+"/dbhandler/schema.go", generator.GetSchema(models))
	frango.CreateFile(parentFolder+"/structs/structs.go", generator.GetStructs(projectPath, needAuthentication, models))
	frango.CreateFile(parentFolder+"/router/router.go", generator.GetRouter(projectPath, needAuthentication, models))
	frango.CreateFile(parentFolder+"/documentation.md", generator.GetDocumentation(needAuthentication, models))

	createModelsAndControllers(parentFolder, projectPath, needAuthentication, models)

	if needAuthentication {
		frango.CreateFile(parentFolder+"/authentication/authentication.go", generator.GetAuthentication(projectPath, models))
		frango.CreateFile(parentFolder+"/middleware/middleware.go", generator.GetMiddleware(projectPath, models))
		frango.CreateFile(parentFolder+"/controllers/authentication/authentication.go", generator.GetAuthenticationController(projectPath, models))
	}
}

func createModelsAndControllers(parentFolder string, projectPath string, needAuthentication bool, models []structs.ModelStruct) {
	for _, model := range models {
		frango.CreateFolder(parentFolder + "/models/" + frango.FirstLetterToLower(model.Name))
		frango.CreateFolder(parentFolder + "/controllers/" + frango.FirstLetterToLower(model.Name))

		frango.CreateFile(parentFolder+"/models/"+frango.FirstLetterToLower(model.Name)+"/"+frango.FirstLetterToLower(model.Name)+".go", generator.GetModel(projectPath, needAuthentication, model))
		frango.CreateFile(parentFolder+"/controllers/"+frango.FirstLetterToLower(model.Name)+"/"+frango.FirstLetterToLower(model.Name)+".go", generator.GetController(projectPath, model))
	}
}

func getProjectName(projectPath string) string {
	substringArray := strings.Split(projectPath, "/")
	return substringArray[len(substringArray)-1]
}
