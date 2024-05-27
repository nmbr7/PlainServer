package server

import (
	. "DesignSphere_Server/src/server/model"
	. "DesignSphere_Server/src/server/service"
	"DesignSphere_Server/src/server/storage"
	"DesignSphere_Server/src/utils"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var backendGlobalContext = make(map[string]any)

type APIController struct {
	StackService *StackService
	Count        int64
}

// func (_controllerStatus *APIController) ValidateState() bool {
// 	return len(_controllerStatus.ProjectConfig) > 0
// }

func (_controllerStatus *APIController) AddCORSHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
}

func StartServer() {

	// TODO: Read the db path from a config file
	store, err := storage.Init(ROOTDIR + "DS_Store")
	if err != nil {
		utils.Log(utils.ERROR, "Stating server failed with error: "+err.Error())
		panic(err)
	}

	stack_service := StackService{DataStore: *store}

	apiController := &APIController{
		StackService: &stack_service,
		Count:        0,
	}

	apiController.InitServer()

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20

	r.Use(apiController.AddCORSHeader)

	r.POST("/deploy", apiController.DeployStack)
	r.POST("/state", apiController.SaveState)
	r.GET("/state", apiController.GetState)
	r.GET("/stack", apiController.DestroyStack)

	r.POST("/uploadProjectConfig", apiController.UploadProjectConfig)
	r.GET("/getProjectConfig", apiController.GetProjectConfig)

	r.Run()
}

func (_controllerStatus *APIController) InitServer() {
	utils.Log(utils.INFO, "Initialializing the server.")
	data, err := os.ReadFile("./.localStore/ServerState")
	if err == nil {
		utils.Log(utils.INFO, "Syncing server state")
		json.Unmarshal(data, &backendGlobalContext)

		// _controllerStatus.ProjectConfig["GCP"] = ProjectData{
		// 	ProjectName:        backendGlobalContext["GCP_ProjectName"].(string),
		// 	GCP_APIKeyFileName: backendGlobalContext["GCP_APIFileName"].(string),
		// }

		// _controllerStatus.ProjectConfig["AWS"] = ProjectData{
		// 	AWS_AccessKeyId:     backendGlobalContext["AWS_AccessKeyId"].(string),
		// 	AWS_SecretAccessKey: backendGlobalContext["AWS_SecretKey"].(string),
		// }
		// utils.Log(utils.DEBUG, _controllerStatus.ProjectConfig)
	}

	//var depResource []DeploymentResource
	data, err = os.ReadFile("./.localStore/Database")
	// if err == nil {
	// 	utils.Log(utils.INFO, "Syncing resource state")
	// 	json.Unmarshal(data, &depResource)
	// 	_controllerStatus.DataStore.Set() = depResource
	// }

}

func (_controllerStatus *APIController) UploadProjectConfig(c *gin.Context) {
	file, _ := c.FormFile("file")
	projectName := c.Request.Form.Get("projectName")

	awsKeyid := c.Request.Form.Get("awsKeyid")
	awsSecretKey := c.Request.Form.Get("awsSecretKey")

	if file != nil {
		filename := url.PathEscape(strings.Replace(projectName, ".", "", -1)) + ".json"
		if filename != "" {
			c.SaveUploadedFile(file, ROOTDIR+filename)
			backendGlobalContext["GCP_APIFileName"] = filename
		}
	}

	if projectName != "" {
		backendGlobalContext["GCP_ProjectName"] = projectName
	}
	if awsKeyid != "" {
		backendGlobalContext["AWS_AccessKeyId"] = awsKeyid
	}
	if awsSecretKey != "" {
		backendGlobalContext["AWS_SecretKey"] = awsSecretKey
	}

	var project_config ProjectConfig

	_, c1 := backendGlobalContext["GCP_ProjectName"]
	_, c2 := backendGlobalContext["GCP_APIFileName"]

	if c1 {
		project_config.GCP_ProjectName = backendGlobalContext["GCP_ProjectName"].(string)
	}

	if c2 {
		project_config.GCP_APIKeyFileName = backendGlobalContext["GCP_APIFileName"].(string)
	}

	_, c1 = backendGlobalContext["AWS_AccessKeyId"]
	_, c2 = backendGlobalContext["AWS_SecretKey"]

	if c1 {
		project_config.AWS_AccessKeyId = backendGlobalContext["AWS_AccessKeyId"].(string)
	}
	if c2 {
		project_config.AWS_SecretAccessKey = backendGlobalContext["AWS_SecretKey"].(string)
	}

	err := _controllerStatus.StackService.SetProjectConfig(project_config)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.Status(http.StatusOK)
}

func (_controllerStatus *APIController) GetProjectConfig(c *gin.Context) {
	config, err := _controllerStatus.StackService.GetProjectConfig()

	if err != nil {
		if err == storage.DBKeyNotFound {
			c.JSON(http.StatusOK, ProjectConfig{})
			return
		}

		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, config)
}

// Deploy and cloud resource stack
//
// TODO: Currently we deploy all the resources togther, we will still need this feature
// but will also require API to deploy the resources separatly.
// TODO: Fix deployment order related bugs
func (_controllerStatus *APIController) DeployStack(c *gin.Context) {

	err := _controllerStatus.StackService.DeployResources()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.Status(200)
}

// API to bring down the deployed stack
//
// TODO: Will also need to create an API to destroy the resources separately
func (_controllerStatus *APIController) DestroyStack(c *gin.Context) {

	err := _controllerStatus.StackService.DestroyResources()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.Status(200)
}

// Save UI rendering state and config changes
func (_controllerStatus *APIController) SaveState(c *gin.Context) {
	var depResource []DeploymentResource
	var jsonData []byte

	if c.Request.Body == nil {
		c.Status(http.StatusBadRequest)
	}

	jsonData, _ = io.ReadAll(c.Request.Body)
	if err := json.Unmarshal(jsonData, &depResource); err != nil {
		utils.Log(utils.ERROR, err)
		return
	}

	utils.Log(utils.DEBUG, "Saving state")

	if err := _controllerStatus.StackService.SaveState(depResource); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(200)
}

// Get UI state with resource config
//
// TODO: Split this API to manage the UI rendering state and resource config changes separately
func (_controllerStatus *APIController) GetState(c *gin.Context) {
	utils.Log(utils.DEBUG, "Returning Stored state")

	resources, err := _controllerStatus.StackService.GetState()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	var resString, _ = json.Marshal(resources)
	utils.Log(utils.DEBUG, "Getting state"+string(resString))

	c.JSON(200, resources)
}
