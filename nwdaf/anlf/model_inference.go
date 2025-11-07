package anlf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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

func requestModelInference(c *gin.Context) { //TODO: Change input data 'data' to appropriate attribute

	reqBody, err := c.GetRawData()
	if err != nil {
		log.Println(err)
	}

	jsonBody := map[string]interface{}{}
	json.Unmarshal(reqBody, &jsonBody)
	log.Println(jsonBody)

	jsonStr, _ := json.Marshal(jsonBody)
	transport := &http.Transport{
		ForceAttemptHTTP2: false,
	}
	http := &http.Client{Transport: transport}
	resp, err := http.Post("http://localhost:9537", "application/json; charset=UTF-8", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Println("error: %v", err)
	} else {
		fmt.Println(resp.Header)
		fmt.Println("************")
		respBody, _ := ioutil.ReadAll(resp.Body)
		jsonData := map[string]interface{}{}
		json.Unmarshal(respBody, &jsonData)
		fmt.Println(jsonData)

		c.JSON(200, gin.H{
			"nfService":     "test-nwdaf",
			"reqNFInstance": "test-anlf",
			"reqTime":       jsonData["reqTime"],
			"data":          jsonData["data"],
		})
	}

}

func AddService(engine *gin.Engine) *gin.RouterGroup {
	group := engine.Group("/nwdaf-anlf/v1")
	group.POST("/", Index)
	group.POST("/:inference", requestModelInference)
	return group
}

func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}
