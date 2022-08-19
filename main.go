package main

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type student struct {
    ID     	int     `json:"id"`
    Name 	string  `json:"name"`
    Program string  `json:"program"`
    GPA  	float64 `json:"gpa"`
}

var class = []student {
    {ID: 123, Name: "Kris Gaudel", Program: "CSBBA", GPA: 4.0},
    {ID: 233, Name: "Prasad Sharma", Program: "ECE", GPA: 2.9},
    {ID: 233, Name: "Hasan Khan", Program: "BBA", GPA: 3.4},
}

func getStudents(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, class)
}

func postStudents(c *gin.Context) {
    var newStudent student

    if err := c.BindJSON(&newStudent); err != nil {
        return
    }

    class = append(class, newStudent)
    c.IndentedJSON(http.StatusCreated, newStudent)
}

func getStudentByID(c *gin.Context) {
    id := c.Param("id")

    for _, s := range class {
        if strconv.Itoa(s.ID) == id {
            c.IndentedJSON(http.StatusOK, s)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}

func main() {
    router := gin.Default()
    router.GET("/students", getStudents)
    router.GET("/students/:id", getStudentByID)
    router.POST("/students", postStudents)

    router.Run("localhost:9090")
}
