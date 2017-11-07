package generatorAuthentication

import (
	"github.com/francoFerraguti/go-abm-generator/common"
	"github.com/francoFerraguti/go-abm-generator/structs"
	"github.com/francoFerraguti/go-abm-generator/templates"
	"github.com/liteByte/frango"
	"strings"
)

func Get(projectPath string, models []structs.ModelStruct) string {
	template := templates.AuthenticationGo()

	usernameField := getUsernameField(models)

	imports := common.GetImports("time", "github.com/dgrijalva/jwt-go", projectPath+"/config")
	customClaimsStruct := getCustomClaimsStruct(frango.FirstLetterToUpper(usernameField.Name) + " " + usernameField.Type)
	tokenStruct := getTokenStruct(frango.FirstLetterToUpper(usernameField.Name) + " " + usernameField.Type)
	tokenFunctions := getTokenFunctions(usernameField)

	fileContent := strings.Replace(template, "&&IMPORTS&&", imports, -1)
	fileContent = strings.Replace(fileContent, "&&CUSTOM_CLAIMS_STRUCT&&", customClaimsStruct, -1)
	fileContent = strings.Replace(fileContent, "&&TOKEN_STRUCT&&", tokenStruct, -1)
	fileContent = strings.Replace(fileContent, "&&TOKEN_FUNCTIONS&&", tokenFunctions, -1)

	return fileContent
}

func getTokenFunctions(field structs.FieldStruct) string {
	return `func CreateToken(` + frango.FirstLetterToLower(field.Name) + " " + field.Type + `) string {
    claims := CustomClaims {
        ` + frango.FirstLetterToLower(field.Name) + `,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
            Issuer:    "TODO change to application name",
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, _ := token.SignedString([]byte(config.GetConfig().JWT_SECRET))

    return tokenString
}

func GetTokenData(tokenString string) Token {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JWT_SECRET), nil
	})
	if err != nil {
		return Token{}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return Token{
			claims.` + frango.FirstLetterToUpper(field.Name) + `,
		}
	} else {
		return Token{}
	}
}
`
}

func getUsernameField(models []structs.ModelStruct) structs.FieldStruct {
	usernameField := structs.FieldStruct{}

	for _, model := range models {
		for _, field := range model.Fields {
			if field.AuthenticationUsername {
				usernameField = field
			}
		}
	}

	return usernameField
}

func getCustomClaimsStruct(field string) string {
	return `type CustomClaims struct {
	` + field + `
	jwt.StandardClaims
}
`
}

func getTokenStruct(field string) string {
	return `type Token struct {
	` + field + `
}
`
}
