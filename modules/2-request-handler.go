package modules

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/sumbing1-go/utils"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"net/http"
)

type Sumbing1RequestHandler struct {
}

func CreateSumbing1RequestHandler() Sumbing1RequestHandler {
	return Sumbing1RequestHandler{}
}

type Sumbing1RequestHandlerInterface interface {
}

func (rh Sumbing1RequestHandler) ConvertJSONtoBSON(c *gin.Context) {
	utils.Info("=== New Request ===")

	// get a content type and check for err
	contentType := c.ContentType()
	if contentType != "application/json" {
		utils.Error(errors.New("contentType not match"), "ConvertJSONtoBSON", "")
		c.Data(http.StatusBadRequest, "", nil)
	}

	// get request body and check err
	readData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		utils.Error(err, "ConvertJSONtoBSON", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse JSON raw data to map
	dataJSON := make(map[string]interface{})
	err = json.Unmarshal(readData, &dataJSON)
	if err != nil {
		utils.Error(err, "ConvertJSONtoBSON", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// parse response to bson data
	result, err := bson.Marshal(dataJSON)
	if err != nil {
		utils.Error(err, "ConvertJSONtoBSON", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	c.Header("Content-Type", "application/bson")
	c.Data(http.StatusOK, "", result)
	return

}

func (rh Sumbing1RequestHandler) ConvertFiletoBSON(c *gin.Context) {

}

func (rh Sumbing1RequestHandler) ConvertBSONtoJSON(c *gin.Context) {
	utils.Info("=== New Request ===")

	// get a content type and check for err
	contentType := c.ContentType()
	if contentType != "application/bson" {
		utils.Error(errors.New("contentType not match"), "ConvertBSONtoJSON", "")
		c.Data(http.StatusBadRequest, "", nil)
	}

	// get request body and check err
	readData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		utils.Error(err, "ConvertBSONtoJSON", "")
		c.Data(http.StatusBadRequest, "", nil)
		return
	}

	// parse BSON raw data to map and check err
	var dataBSON map[string]interface{}
	err = bson.Unmarshal(readData, &dataBSON)
	if err != nil {
		utils.Error(err, "ConvertBSONtoJSON", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// parse the map into JSON and check for err
	result, err := json.Marshal(dataBSON)
	if err != nil {
		utils.Error(err, "ConvertBSONtoJSON", "")
		c.Data(http.StatusInternalServerError, "", nil)
		return
	}

	// create success return
	c.Header("Content-Type", "application/json")
	c.Data(http.StatusOK, "", result)
	return
}
