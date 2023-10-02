package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	_apiController := &APIControllerImp{
		DataStore:    nil,
		CurrentStack: nil,
		Count:        0,
	}

	data, _ := os.ReadFile("./Database")
	json.Unmarshal(data, &_apiController.DataStore)

	// Middleware
	r.Use(_apiController.AddCORSHeader)

	// Route Handlers
	r.POST("/deploy", _apiController.DeployStack)
	r.POST("/state", _apiController.SaveState)
	r.GET("/state", _apiController.GetState)
	// r.POST("/preview", DeploymentHandler)
	r.GET("/stack", _apiController.DestroyStack)

	// listen and serve on 0.0.0.0:8080
	r.Run()

}
