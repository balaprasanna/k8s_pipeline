package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// OK ... HTTP Status code for success
var OK = 200

// NOTOK ... HTTP Status code for failure
var NOTOK = 400

// ERROR ... HTTP Status code for error
var ERROR = 500

// db .. reference exposed to routes
var db *gorm.DB

func main() {
	// Setup PORT
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	// DB Config
	envMap := map[string]string{
		"DB_USER":     "root",
		"DB_PASSWORD": "root",
		"DB_NAME":     "taskstore",
		"DB_HOST":     "localhost",
	}

	for k := range envMap {
		value := os.Getenv(k)
		if value != "" {
			envMap[k] = value
		}
	}

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8&loc=Local",
		envMap["DB_USER"],
		envMap["DB_PASSWORD"],
		envMap["DB_HOST"],
		envMap["DB_NAME"],
	)
	fmt.Println(connectionString)
	dbMysql, err := gorm.Open("mysql", connectionString)

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// Get Reference to DB outside
	db = dbMysql
	defer dbMysql.Close()

	// Disable tablename pluralization
	db.SingularTable(true)

	// GINetA Config
	gin.SetMode(gin.ReleaseMode)

	// GIN Routes
	router := gin.Default()

	router.GET("/", func(c *gin.Context) { c.JSON(OK, gin.H{"status": "ok"}) })
	router.GET("/time", TimeHandler)
	router.GET("/health", func(c *gin.Context) { c.String(OK, "ok") })

	v1 := router.Group("/v1/")
	{
		v1.GET("/task", GetAllTask)
		v1.GET("/task/:id", GetTask)
		v1.POST("/task", PostTask)
		v1.PUT("/task/:id", UpdateTask)
		v1.DELETE("/task/:id", DeleteTask)
	}

	// Start the APP
	fmt.Printf("Starting Gin APP Server on PORT NO %s \n", PORT)
	router.Run(fmt.Sprintf(":%s", PORT))
}
