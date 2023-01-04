package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Record struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var RecordOne []Record = []Record{
	{Id: "1", Name: "ashwin", Email: "ashwin@gmail.com"},
	{Id: "2", Name: "anuj", Email: "anuj@gmail.com"},
}

func main() {
	r := gin.Default()
	r.GET("/", print)
	r.GET("/records", records)
	r.GET("/records/:id", getById)
	r.POST("records", addRecord)
	r.PUT("/records/:id", editRecord)
	r.DELETE("/records/:id", deleteRecord)

	r.Run()
}

func records(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, RecordOne)
}

func getById(c *gin.Context) {
	Id := c.Param("id")

	for _, val := range RecordOne {
		if Id == val.Id {
			c.IndentedJSON(http.StatusOK, val)
			return
		}
	}
}

func addRecord(c *gin.Context) {
	var newUser Record

	if err := c.BindJSON(&newUser); err != nil {
		panic(err)
	}

	RecordOne = append(RecordOne, newUser)
	c.IndentedJSON(http.StatusOK, RecordOne)
}

func deleteRecord(c *gin.Context) {
	id := c.Param("id")

	for i := 0; i < len(RecordOne); i++ {
		if RecordOne[i].Id == id {
			RecordOne = append(RecordOne[:i], RecordOne[i+1:]...)
			c.IndentedJSON(http.StatusOK, RecordOne)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, "No results found")
}

func editRecord(c *gin.Context) {
	id := c.Param("id")

	var newValue Record

	if err := c.BindJSON(&newValue); err != nil {
		panic(err)
	}

	for i, val := range RecordOne {
		if val.Id == id {
			RecordOne = append(RecordOne[:i], RecordOne[i+1:]...)
			RecordOne = append(RecordOne, newValue)
			c.IndentedJSON(http.StatusOK, RecordOne)
			return
		}
	}

	c.IndentedJSON(http.StatusOK, "No id found")
}

func print(c *gin.Context) {
	fmt.Fprintln(c.Writer, "Working")
}
