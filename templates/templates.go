package templates

func MainGo() string {
	return `package main

&&IMPORTS&&

func main() {
	dbhandler.ConnectToDatabase()
	router.ConfigureRouter()
	router.CreateRouter()
	router.RunRouter()
}
`
}

func MiddlewareGo() string {
	return `package middleware

&&IMPORTS&&

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		token := authentication.GetTokenData(tokenString)

		if token.&&USERNAME_FIELD_UPPER&& == "" || tokenString == "" {
		    c.JSON(401, "Authentication error")
	    	c.Abort()
			return
		}

		c.Set("&&USERNAME_FIELD_LOWER&&", token.&&USERNAME_FIELD_UPPER&&)
	}
}
`
}

func ConfigGo() string {
	return `package config

type Config struct {
	ENV			string
	PORT 		string
	&&CONFIG_AUTHENTICATION_FIELD&&
	DB_TYPE		string
	DB_USERNAME	string
	DB_PASSWORD	string
	DB_HOST		string
	DB_PORT		string
	DB_NAME 	string
}

var instance *Config

func GetConfig() *Config {
	if instance == nil {
		config := newConfigLocal()
		instance = &config
	}
	return instance
}

func newConfigLocal() Config {
	return Config{
		ENV:			"develop",
		PORT:			"&&CONFIG_PORT&&",
		&&CONFIG_AUTHENTICATION_VALUE&&,
		DB_TYPE:       	"&&CONFIG_DB_TYPE&&",
		DB_USERNAME:    "&&CONFIG_DB_USERNAME&&",
		DB_PASSWORD:    "&&CONFIG_DB_PASSWORD&&",
		DB_HOST:      	"&&CONFIG_DB_HOST&&",
		DB_PORT:       	"&&CONFIG_DB_PORT&&",
		DB_NAME:       	"&&CONFIG_DB_NAME&&",
	}
}
`
}

func RouterGo() string {
	return `package router

&&IMPORTS&&

var router *gin.Engine

func ConfigureRouter() {
	if config.GetConfig().ENV != "develop" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func CreateRouter() {
	router = gin.New()

	&&AUTHENTICATION_ENDPOINTS&&

	api := router.Group("/"&&GIN_MIDDLEWARE_STRING&&)
	{
		&&ENDPOINTS&&
    }
}

func RunRouter() {
	router.Run(":" + config.GetConfig().PORT)
}
`
}

func StructsGo() string {
	return `package structs

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

&&AUTH_STRUCTS&&

&&STRUCTS&&
`
}

func DBHandlerGo() string {
	return `package dbhandler

&&IMPORTS&&

var db *sql.DB

func ConnectToDatabase() {
    var err error
	
	db, err = sql.Open(config.GetConfig().DB_TYPE, config.GetConfig().DB_USERNAME + ":" + config.GetConfig().DB_PASSWORD + "@tcp(" + config.GetConfig().DB_HOST + ":" + config.GetConfig().DB_PORT + ")/" + config.GetConfig().DB_NAME)
	frango.PrintErr(err)
    
    err = db.Ping()
    frango.PrintErr(err)

    createSchema()
}

func GetDatabase() *sql.DB {
    return db
}

func createSchema() {
&&SCHEMA_FUNCTIONS&&
}
`
}

func SchemaGo() string {
	return `package dbhandler

import (
    "github.com/liteByte/frango"
)

&&SCHEMA&&
`
}

func DocumentationMd() string {
	return `&&DOCUMENTATION&&`
}

func AuthenticationGo() string {
	return `package authentication

&&IMPORTS&&

&&CUSTOM_CLAIMS_STRUCT&&

&&TOKEN_STRUCT&&

&&TOKEN_FUNCTIONS&&
`
}

func AuthenticationControllerGo() string {
	return `package authentication

&&IMPORTS&&

&&FUNCTIONS&&
`
}

func ModelGo() string {
	return `&&PACKAGE_NAME&&

&&IMPORTS&&

&&CHECK_LOGIN&&

&&CREATE&&

&&GET_LIST&&

&&GET&&

&&UPDATE&&

&&DELETE&&
`
}

func ControllerGo() string {
	return `&&PACKAGE_NAME&&

&&IMPORTS&&

&&CREATE&&

&&GET_LIST&&

&&GET&&

&&UPDATE&&

&&DELETE&&
`
}
