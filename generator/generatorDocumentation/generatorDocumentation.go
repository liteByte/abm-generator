package generatorDocumentation

import (
	"github.com/francoFerraguti/go-abm-generator/structs"
	"github.com/francoFerraguti/go-abm-generator/templates"
	"github.com/liteByte/frango"
	"strings"
)

func Get(needAuthentication bool, models []structs.ModelStruct) string {
	template := templates.DocumentationMd()

	documentation := getDocumentation(needAuthentication, models)

	fileContent := strings.Replace(template, "&&DOCUMENTATION&&", documentation, -1)

	return fileContent
}

func getDocumentation(needAuthentication bool, models []structs.ModelStruct) string {
	s := ""

	if needAuthentication {

		authModel := structs.ModelStruct{}
		usernameField := structs.FieldStruct{}
		passwordField := structs.FieldStruct{}

		for _, model := range models {
			for _, field := range model.Fields {
				if field.AuthenticationUsername {
					usernameField = field
					authModel = model
				}
				if field.AuthenticationPassword {
					passwordField = field
					authModel = model
				}
			}
		}

		title := "Signup"
		description := "Creates a user on the database"
		endpoint := "/signup"
		method := "POST"
		url_params := ""
		headers := ""
		body := getEndpointBody(authModel, true)
		code200 := "Signup successful"
		code400 := "Request body is wrong"
		code401 := ""
		code500 := "Username already in use"

		s += getEndpointDocumentation(title, description, endpoint, method, url_params, headers, body, code200, code400, code401, code500)

		title = "Login"
		description = "Returns an access token for the provided user"
		endpoint = "/login"
		method = "POST"
		url_params = ""
		headers = ""
		body = `{
	"` + usernameField.Name + `": "` + usernameField.Type + `",
	"` + passwordField.Name + `": "` + passwordField.Type + `"
}`
		code200 = "a63ab36162a4f4ee6622ccd787b0a048c26b93acfc05c6b1843659b253c3c00b //authentication token"
		code400 = "Request body is wrong"
		code401 = ""
		code500 = "Wrong username or password"

		s += getEndpointDocumentation(title, description, endpoint, method, url_params, headers, body, code200, code400, code401, code500)
	}

	for _, model := range models {

		//Create
		title := "Create " + model.Name
		description := "Creates a " + frango.FirstLetterToLower(model.Name) + " on the database"
		endpoint := "/" + frango.FirstLetterToLower(model.Name)
		method := "POST"
		url_params := ""
		headers := "Authorization: Token"
		body := getEndpointBody(model, true)
		code200 := model.Name + " created successfully"
		code400 := "Request body is wrong"
		code401 := "Unauthorized"
		code500 := "Server error"

		s += getEndpointDocumentation(title, description, endpoint, method, url_params, headers, body, code200, code400, code401, code500)

		//GetList
		title = "Get " + model.Name + " list"
		description = "Gets a list of every " + frango.FirstLetterToLower(model.Name) + " on the database"
		endpoint = "/" + frango.FirstLetterToLower(model.Name) + "/list"
		method = "GET"
		url_params = ""
		headers = "Authorization: Token"
		body = ""
		code200 = getEndpointBody(model, false)
		code400 = ""
		code401 = "Unauthorized"
		code500 = "Server error"

		s += getEndpointDocumentation(title, description, endpoint, method, url_params, headers, body, code200, code400, code401, code500)

		for _, field := range model.Fields {
			if !field.Unique {
				continue
			}

			//Get
			title = "Get " + model.Name + " by " + frango.FirstLetterToUpper(field.Name)
			description = "Returns a " + frango.FirstLetterToLower(model.Name) + " using the provided " + field.Name
			endpoint = "/" + frango.FirstLetterToLower(model.Name) + "/" + field.Name + "/:" + field.Name
			method = "GET"
			url_params = ""
			headers = "Authorization: Token"
			body = ""
			code200 = getEndpointBody(model, false)
			code400 = ""
			code401 = "Unauthorized"
			code500 = "Server error"

			s += getEndpointDocumentation(title, description, endpoint, method, url_params, headers, body, code200, code400, code401, code500)

			//Update
			title = "Update " + model.Name + " by " + frango.FirstLetterToUpper(field.Name)
			description = "Updates a " + frango.FirstLetterToLower(model.Name) + " using the provided " + field.Name
			endpoint = "/" + frango.FirstLetterToLower(model.Name) + "/" + field.Name + "/:" + field.Name
			method = "PUT"
			url_params = ""
			headers = "Authorization: Token"
			body = getEndpointBody(model, true)
			code200 = model.Name + " updated successfully"
			code400 = "Request body is wrong"
			code401 = "Unauthorized"
			code500 = "Server error"

			s += getEndpointDocumentation(title, description, endpoint, method, url_params, headers, body, code200, code400, code401, code500)

			//Delete
			title = "Delete " + model.Name + " by " + frango.FirstLetterToUpper(field.Name)
			description = "Deletes a " + frango.FirstLetterToLower(model.Name) + " using the provided " + field.Name
			endpoint = "/" + frango.FirstLetterToLower(model.Name) + "/" + field.Name + "/:" + field.Name
			method = "DELETE"
			url_params = ""
			headers = "Authorization: Token"
			body = ""
			code200 = model.Name + " deleted successfully"
			code400 = ""
			code401 = "Unauthorized"
			code500 = "Server error"

			s += getEndpointDocumentation(title, description, endpoint, method, url_params, headers, body, code200, code400, code401, code500)
		}
	}

	return s
}

func getEndpointDocumentation(title, description, endpoint, method, url_params, headers, body, code200, code400, code401, code500 string) string {
	if url_params == "" {
		url_params = "-"
	}
	if headers == "" {
		headers = "-"
	}
	if body == "" {
		body = "-"
	}
	if code400 == "" {
		code400 = "-"
	}
	if code401 == "" {
		code401 = "-"
	}

	return `**` + title + `**
` + description + `

**` + endpoint + `** [` + method + `]

**URL Parameters**
` + url_params + `

**Headers**
` + headers + `

**Request Body**
` + body + `

**Success 200 Response**
` + code200 + `

**Bad Request 400 Response**
` + code400 + `

**Unauthorized 401 Response**
` + code401 + `

**Internal Server Error 500 Response**
` + code500 + `

-----------------------------------------------------
`
}

func getEndpointBody(model structs.ModelStruct, onlySend bool) string {
	s := "{\n"

	for i, field := range model.Fields {

		if field.AutoGenerated && onlySend {
			continue
		}

		fieldTypeString := ""

		if field.Type == "string" {
			fieldTypeString = "`string`"
		} else {
			fieldTypeString = field.Type
		}

		if i != len(model.Fields)-1 {
			fieldTypeString += ","
		}

		s += "	`" + frango.FirstLetterToLower(field.Name) + "`: " + fieldTypeString + "\n"
	}

	s += "}"

	s = strings.Replace(s, "`", `"`, -1)

	return s
}
