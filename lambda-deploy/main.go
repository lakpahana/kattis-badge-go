package main

import (
	// "context"
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	"lakpahana.me/api/controllers"
	"lakpahana.me/db"
)

var ginLambda *ginadapter.GinLambda

func init() {
	r := gin.Default()

	db.ConnectToDB()

	// Existing routes
	r.GET("/", controllers.Hello)
	r.GET("/allUsers", controllers.GetAllUsers)
	r.GET("/rank/:id", controllers.GetRankByCountry)

	// Test routes
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Test route reached",
		})
	})
	r.GET("/test/echo/:message", func(c *gin.Context) {
		message := c.Param("message")
		c.JSON(200, gin.H{
			"echo": message,
		})
	})
	r.POST("/test/post", func(c *gin.Context) {
		var body map[string]interface{}
		if err := c.BindJSON(&body); err != nil {
			c.JSON(400, gin.H{"error": "Could not read body"})
			return
		}
		c.JSON(200, gin.H{
			"received": body,
		})
	})

	ginLambda = ginadapter.New(r)
}


// Handler is your Lambda function handler
// It uses the ginLambda Proxy method to convert the API Gateway proxy event into an HTTP request,
// and then creates an HTTP response from the Gin application's response
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received request for %s\n", req.Path)
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	// Start the Lambda function handler
	lambda.Start(Handler)
}
