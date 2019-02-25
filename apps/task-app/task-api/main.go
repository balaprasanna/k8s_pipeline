package main

import (
	"fmt"
	"os"
	"time"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var OK int = 200
var NOT_OK int = 400
var ERROR int= 500


type Task struct {
	Id int64 `gorm:"primary_key,column:id"`
	Name string `gorm:"column:name" json:"Name" binding:"required"`
	Type string `gorm:"column:type" json:"Type" binding:"required"`
	Status uint `gorm:"column:status"`
	Content string `gorm:"column:content" json:"Content" binding:"required"`
}

func (Task) TableName() string {
	return "task"
}

type Any struct {}

var db *gorm.DB

func main() {
	// Setup PORT
	PORT := os.Getenv("PORT")
	if PORT == "" { PORT = "8080" }

	// DB Config
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")

	if (user == "") { user = "root" }
	if (password == "") { password = "root" }
	if (host == "") { host = "localhost" }
	if (dbname== "") { dbname = "taskstore" }
	
	ConnectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8&loc=Local", 
		user, 
		password,
		host,
		dbname,
	)
	fmt.Println(ConnectionString)
	db_, err := gorm.Open("mysql", ConnectionString)

	if (err != nil) {
		fmt.Println(err)
		os.Exit(0)
	}	

	// Get Reference to DB outside
	db = db_
	defer db_.Close()

	// Disable tablename pluralization	
	db.SingularTable(true)

	// GIN Config
	gin.SetMode(gin.ReleaseMode)

	// GIN Routes
	router := gin.Default()

	v1 := router.Group("/v1/")
	v1.GET("/task", GetAllTask)
	v1.GET("/task/:id", GetTask)
	v1.POST("/task", PostTask)
	v1.PUT("/task/:id", UpdateTask)
	v1.DELETE("/task/:id", DeleteTask)

	router.GET("/", func(c *gin.Context ) {
		c.JSON(OK, gin.H{ "status": "ok" })
	})
	router.GET("/time", TimeHandler)

	router.GET("/health", func (c *gin.Context)  { c.String(OK, "ok") })

	// Start the APP
	fmt.Printf("Starting Gin APP Server on PORT NO %s \n", PORT)
	router.Run( fmt.Sprintf(":%s", PORT) )
}

// Handlers for API Endpoints
func TimeHandler (c *gin.Context)  {
	resp := gin.H{
			"status": "ok",
			"time": time.Now().String(),
	}
	c.JSON(OK, resp);
}

func  GetAllTask(c *gin.Context)  {
	var tasks []Task
	db.Find(&tasks)
	resp := gin.H{
			"status": "ok",
			"data": tasks,
	}
	c.JSON(OK, resp);
}

func  GetTask(c *gin.Context)  {

	id := c.Param("id")
	id_int, err := strconv.ParseInt(id, 10,64)
	if err != nil  {
		c.AbortWithError(NOT_OK, err)
	}

	var tasks []Task
	query := Task{Id: id_int}

	db.Where( &query ).Find(&tasks)
	
	var resp gin.H

	if len(tasks) > 0 {
		resp = gin.H{
			"status": "ok",
			"data": tasks[0],
		}
	} else {
		resp = gin.H{
			"status": "ok",
			"data": gin.H{},
			"message": "no records found",
		}
	}
	c.JSON(OK, resp);
}

func  PostTask(c *gin.Context)  {

	var task Task
	c.BindJSON(&task)

	fmt.Println(":>> ",task.Name, task.Type, task.Content)
	if task.Name == "" || task.Type == "" || task.Content == "" {
		c.AbortWithStatusJSON(NOT_OK, gin.H{
			"status": "ok",
			"message": "Name, Type, Content is mandatory",
		})
		return
	}

	taskToCreate := &Task{Name: task.Name, Type: task.Type, Content: task.Content}
	db.Create(taskToCreate)

	resp := gin.H{
			"status": "ok",
			"data": taskToCreate,
	}
	c.JSON(OK, resp);
}

func  UpdateTask(c *gin.Context)  {

	id := c.Param("id")
	id_int, err := strconv.ParseInt(id, 10,64)
	if err != nil  {
		c.AbortWithError(NOT_OK, err)
		return;
	}

	var task Task
	c.BindJSON(&task)

	var taskToUpdate []Task
	db.Find(&taskToUpdate, &Task{Id: id_int})
	if len(taskToUpdate) == 0 {
		c.AbortWithStatusJSON(
			NOT_OK,
			gin.H{
				"status": "ok",
				"message": "no records found",
			},
		)
		return
	}

	task_to_save := taskToUpdate[0]

	if task.Name != "" {
		task_to_save.Name = task.Name
	}

	if task.Content != "" {
		task_to_save.Content = task.Content
	}

	if task.Type != "" {
		task_to_save.Type = task.Type
	}

	db.Save(&task_to_save)

	resp := gin.H{
			"status": "ok",
			"data": task_to_save,
			"message": "record updated successfully",
	}
	c.JSON(OK, resp);
}

func  DeleteTask(c *gin.Context)  {

	id := c.Param("id")
	id_int, err := strconv.ParseInt(id, 10,64)
	if err != nil  {
		c.AbortWithError(NOT_OK, err)
		return
	}
	
	//Delete id from table
	var task []Task
	query := Task{Id: id_int}
	
	// FIND
	db.Where(&query).First(&task)

	if len(task) == 0 {
		c.JSON( OK, gin.H{ "status": "ok", "message": "no records found"} )
		return
	}

	// db.Delete(&task)
	db.Delete( &query )
	
	resp := gin.H{
			"status": "ok",
			"data": task[0],
			"message": "record deleted successfully",
	}
	c.JSON(OK, resp);
}
