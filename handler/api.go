package handler

import (
	"Test/db"
	"encoding/json"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type emp1 struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func AddOneRecord(c *gin.Context) {
	var req emp1
	jsonBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("Error while reading data:", err)
	}
	if err = json.Unmarshal(jsonBytes, &req); err != nil {
		fmt.Println("error occured in unmarshalling", err)
	}

	_, err = db.Postgres.Exec("INSERT INTO emp1 VALUES($1, $2)", req.Name, req.Id)
	if err != nil {
		fmt.Println("Insert error:", err)
		return
	}
	fmt.Println("EMP:::", req)
	c.JSON(200, "One man added")
}
func GetAll(c *gin.Context) {
	var req []emp1
	rows, err := db.Postgres.Query("SELECT * FROM EMP1")
	if err != nil {
		fmt.Println("error occured in query:", err)
	}
	defer rows.Close()
	for rows.Next() { //this iteration for each row.
		var ee emp1
		if err := rows.Scan(&ee.Name, &ee.Id); err != nil { //seq in scan should be same as colm present in table
			fmt.Println("error in scanning", err)
		}
		req = append(req, ee)
	}
	fmt.Println("GetAll:", req)
	c.JSON(200, req)
}
func GetOne(c *gin.Context) {
	var req []emp1
	get := c.Param("id")
	rows, err := db.Postgres.Query("Select name, id from EMP1 where id=$1", get)
	if err != nil {
		fmt.Println("error occuured in query", err)
	}
	for rows.Next() {
		var e emp1
		if err := rows.Scan(&e.Name, &e.Id); err != nil {
			fmt.Println("error in scan:", err)
		}
		fmt.Println("e:", e)
		req = append(req, e)
	}
	fmt.Println("req:", req)
	c.JSON(200, req)
}
func DeleteAll(c *gin.Context) {
	c.JSON(200, "")
}
func DeleteOne(c *gin.Context) {
	del := c.Param("id")
	result, err := db.Postgres.Exec("delete from EMP1 where id=$1", del)
	if err != nil {
		fmt.Println("error occured in query", err)
	}
	fmt.Println("result::", result)
	c.JSON(200, "Deleted")
}
func Update(c *gin.Context) {
	fmt.Println("into update")
	getId := c.Param("id")
	fmt.Println(getId)
	var req emp1
	jsonbyte, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("err in reading data", err)
	}
	if err = json.Unmarshal(jsonbyte, &req); err != nil {
		fmt.Println("error occured in unmarshal", err)
	}

	if _, err = db.Postgres.Exec("UPDATE EMP1 SET name = $1 WHERE id = $2", req.Name, getId); err != nil {
		fmt.Println("error in query", err)
	}

	c.JSON(200, "Update")
}
