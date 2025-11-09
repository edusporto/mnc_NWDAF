package mtlf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

type Routes []Route

type ResponseInfo struct {
	nfService       string `json:"nfService"`
	reqNFInstanceID string `json:"reqNFInstance"`
	reqTime         string `json:"reqTime"`
	data            string `json:"data"`
}

func requestModelTraining(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute
	reqBody := map[string]interface{}{}
	log.Print("##############################3")
	// log.Println(reqBody)
	// replyto, _ := json.Marshal(reqBody)
	// reqData := []byte(`{ "nfService":"training-requseted"}`)
	c.JSON(http.StatusOK, gin.H{
		"nfService":     "test-nwdaf",
		"reqNFInstance": "test-mtlf",
		"reqTime":       reqBody["reqTime"],
		"data":          "finished",
	})

	// c.BindJSON(&replyto)

	jsonBody := map[string]interface{}{}
	jsonBody["reqNFInstanceID"] = "test"
	jsonBody["nfService"] = "training"
	now_t := time.Now().Format("2006-01-02 15:04:05")
	jsonBody["reqTime"] = now_t
	jsonBody["data"] = "None"
	// jsonStr, _ := json.Marshal(jsonBody)
	print("*********")
	transport := &http.Transport{
		ForceAttemptHTTP2: false,
	}
	http := &http.Client{Transport: transport}
	resp, err := http.Get("http://fl-training-manager:9537/training")
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		fmt.Println("************")
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)

	}
}

func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/nwdaf-mtlf/v1")
	group.POST("/", Index)
	group.POST("/:training", requestModelTraining)
	return group
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
