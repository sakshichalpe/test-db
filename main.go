package main

import (
	"Test/db"
	"Test/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db.ConnectDB()
	r.POST("/addOne", handler.AddOneRecord)
	r.GET("/GetAll", handler.GetAll)
	r.GET("/GetOne/:id", handler.GetOne)
	r.GET("/deleteOne/:id", handler.DeleteOne)
	r.POST("/update/:id", handler.Update)
	r.Run()

}

// func RemoveDuplicateChild(chOrders []string) []string {
// 	encountered := make(map[string]bool)
// 	var result []string

// 	for _, val := range chOrders {

// 		if !encountered[val] {
// 			encountered[val] = true
// 			result = append(result, val)
// 		}
// 	}
// }
