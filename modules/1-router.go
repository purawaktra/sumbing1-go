package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/sumbing1-go/utils"
)

type Sumbing1Router struct {
	engine *gin.Engine
	rh     Sumbing1RequestHandler
}

func CreateSumbing1Router(engine *gin.Engine, rh Sumbing1RequestHandler) Sumbing1Router {
	return Sumbing1Router{
		engine: engine,
		rh:     rh,
	}
}

func (r Sumbing1Router) Init(path string) {
	pathGroup := r.engine.Group(path, utils.CheckBasicAuth)
	pathGroup.POST("/convert/json-to-bson", r.rh.ConvertJSONtoBSON)
	pathGroup.POST("/convert/image-to-bson-with-account-id", r.rh.ConvertImagetoBSONWithAccountId)
	pathGroup.POST("/convert/bson-to-json", r.rh.ConvertBSONtoJSON)
	pathGroup.GET("/generate/uuid", r.rh.GenerateGoogleUUID)
}
