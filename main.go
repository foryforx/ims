package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/karuppaiah/ims/models"
)

var db *gorm.DB

// func setupRouter() *gin.Engine {
// 	// Disable Console Color
// 	// gin.DisableConsoleColor()
// 	r := gin.Default()

// 	// Ping test
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.String(http.StatusOK, "pong")
// 	})

// 	// Get user value
// 	r.GET("/user/:name", func(c *gin.Context) {
// 		user := c.Params.ByName("name")
// 		value, ok := DB[user]
// 		if ok {
// 			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
// 		} else {
// 			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
// 		}
// 	})

// 	// Authorized group (uses gin.BasicAuth() middleware)
// 	// Same than:
// 	// authorized := r.Group("/")
// 	// authorized.Use(gin.BasicAuth(gin.Credentials{
// 	//	  "foo":  "bar",
// 	//	  "manu": "123",
// 	//}))
// 	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
// 		"foo":  "bar", // user:foo password:bar
// 		"manu": "123", // user:manu password:123
// 	}))

// 	authorized.POST("admin", func(c *gin.Context) {
// 		user := c.MustGet(gin.AuthUserKey).(string)

// 		// Parse JSON
// 		var json struct {
// 			Value string `json:"value" binding:"required"`
// 		}

// 		if c.Bind(&json) == nil {
// 			DB[user] = json.Value
// 			c.JSON(http.StatusOK, gin.H{"status": "ok"})
// 		}
// 	})

// 	return r
// }

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:@/ims?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&models.TodoModel{})
}
func main() {
	//r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	//r.Run(":8080")
	router := gin.Default()
	v1 := router.Group("/api/v1/ims")
	{
		v1.POST("/", createTodo)
		v1.GET("/", fetchAllTodo)
		v1.GET("/:id", fetchSingleTodo)
		v1.PUT("/:id", updateTodo)
		v1.DELETE("/:id", deleteTodo)
	}
	router.Run()

}
