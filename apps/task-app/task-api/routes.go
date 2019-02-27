package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//TimeHandler ...
//Method: GET Route: /time
func TimeHandler(c *gin.Context) {
	resp := gin.H{
		"status": "ok",
		"time":   time.Now().String(),
	}
	c.JSON(OK, resp)
}

//GetAllTask ...
//Method: GET Route: /task
func GetAllTask(c *gin.Context) {
	var tasks []Task
	db.Find(&tasks)
	resp := gin.H{
		"status": "ok",
		"data":   tasks,
	}
	c.JSON(OK, resp)
}

//GetTask ...
//Method: GET Route: /task/:id
func GetTask(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.AbortWithError(NOTOK, err)
	}

	var tasks []Task
	query := Task{Id: idInt}

	db.Where(&query).Find(&tasks)

	var resp gin.H

	if len(tasks) > 0 {
		resp = gin.H{
			"status": "ok",
			"data":   tasks[0],
		}
	} else {
		resp = gin.H{
			"status":  "ok",
			"data":    gin.H{},
			"message": "no records found",
		}
	}
	c.JSON(OK, resp)
}

//PostTask ...
//Method: POST Route: /task
func PostTask(c *gin.Context) {

	var task Task
	c.BindJSON(&task)

	fmt.Println(":>> ", task.Name, task.Type, task.Content)
	if task.Name == "" || task.Type == "" || task.Content == "" {
		c.AbortWithStatusJSON(NOTOK, gin.H{
			"status":  "ok",
			"message": "Name, Type, Content is mandatory",
		})
		return
	}

	taskToCreate := &Task{Name: task.Name, Type: task.Type, Content: task.Content}
	db.Create(taskToCreate)

	resp := gin.H{
		"status": "ok",
		"data":   taskToCreate,
	}
	c.JSON(OK, resp)
}

//UpdateTask ...
//Method: PUT Route: /task/:id
func UpdateTask(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.AbortWithError(NOTOK, err)
		return
	}

	var task Task
	c.BindJSON(&task)

	var taskToUpdate []Task
	db.Find(&taskToUpdate, &Task{Id: idInt})
	if len(taskToUpdate) == 0 {
		c.AbortWithStatusJSON(
			NOTOK,
			gin.H{
				"status":  "ok",
				"message": "no records found",
			},
		)
		return
	}

	taskToSave := taskToUpdate[0]

	if task.Name != "" {
		taskToSave.Name = task.Name
	}

	if task.Content != "" {
		taskToSave.Content = task.Content
	}

	if task.Type != "" {
		taskToSave.Type = task.Type
	}

	db.Save(&taskToSave)

	resp := gin.H{
		"status":  "ok",
		"data":    taskToSave,
		"message": "record updated successfully",
	}
	c.JSON(OK, resp)
}

//DeleteTask ...
//Method: DELETE Route: /task/:id
func DeleteTask(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.AbortWithError(NOTOK, err)
		return
	}

	//Delete id from table
	var task []Task
	query := Task{Id: idInt}

	// FIND
	db.Where(&query).First(&task)

	if len(task) == 0 {
		c.JSON(OK, gin.H{"status": "ok", "message": "no records found"})
		return
	}

	// db.Delete(&task)
	db.Delete(&query)

	resp := gin.H{
		"status":  "ok",
		"data":    task[0],
		"message": "record deleted successfully",
	}
	c.JSON(OK, resp)
}
