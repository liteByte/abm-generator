package generatorAuthenticationController

import (
	"github.com/francoFerraguti/go-abm-generator/common"
	"github.com/francoFerraguti/go-abm-generator/structs"
	"github.com/francoFerraguti/go-abm-generator/templates"
	"github.com/liteByte/frango"
	"strings"
)

func Get(projectPath string, models []structs.ModelStruct) string {
	template := templates.AuthenticationControllerGo()

	authenticationModel, usernameField, passwordField := getData(models)

	imports := common.GetImports("github.com/liteByte/frango", "github.com/gin-gonic/gin", projectPath+"/structs", projectPath+"/authentication", projectPath+"/models/"+frango.FirstLetterToLower(authenticationModel.Name))
	functions := getFunctions(authenticationModel, usernameField, passwordField)

	fileContent := strings.Replace(template, "&&IMPORTS&&", imports, -1)
	fileContent = strings.Replace(fileContent, "&&FUNCTIONS&&", functions, -1)

	return fileContent
}

func getData(models []structs.ModelStruct) (structs.ModelStruct, structs.FieldStruct, structs.FieldStruct) {
	authenticationModel := structs.ModelStruct{}
	usernameField := structs.FieldStruct{}
	passwordField := structs.FieldStruct{}

	for _, model := range models {
		for _, field := range model.Fields {
			if field.AuthenticationUsername {
				authenticationModel = model
				usernameField = field
			}

			if field.AuthenticationPassword {
				authenticationModel = model
				passwordField = field
			}
		}
	}

	return authenticationModel, usernameField, passwordField
}

func getFunctions(authenticationModel structs.ModelStruct, usernameField, passwordField structs.FieldStruct) string {
	return `func Login(c *gin.Context) {
	loginStruct, err := structs.ParseBodyIntoLoginStruct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	loginStruct.` + frango.FirstLetterToUpper(passwordField.Name) + ` = frango.Hash(loginStruct.` + frango.FirstLetterToUpper(usernameField.Name) + `, loginStruct.` + frango.FirstLetterToUpper(passwordField.Name) + `)

	if err = ` + frango.FirstLetterToLower(authenticationModel.Name) + `.CheckLogin(loginStruct); err != nil {
		c.JSON(500, err.Error())
		return
	}

	token := authentication.CreateToken(loginStruct.` + frango.FirstLetterToUpper(usernameField.Name) + `)

	c.JSON(200, token)
}

func Signup(c *gin.Context) {
	` + frango.FirstLetterToLower(authenticationModel.Name) + `Struct, err := structs.ParseBodyInto` + frango.FirstLetterToUpper(authenticationModel.Name) + `Struct(c.Request.Body)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	` + frango.FirstLetterToLower(authenticationModel.Name) + `Struct.` + frango.FirstLetterToUpper(passwordField.Name) + ` = frango.Hash(` + frango.FirstLetterToLower(authenticationModel.Name) + `Struct.` + frango.FirstLetterToUpper(usernameField.Name) + `, ` + frango.FirstLetterToLower(authenticationModel.Name) + `Struct.` + frango.FirstLetterToUpper(passwordField.Name) + `)

	if err = ` + frango.FirstLetterToLower(authenticationModel.Name) + `.Create(` + frango.FirstLetterToLower(authenticationModel.Name) + `Struct); err != nil {
		c.JSON(500, err.Error())
		return		
	}

	c.JSON(200, "Signup successful")
}`
}
