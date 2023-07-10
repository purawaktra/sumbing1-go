package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/sumbing1-go/modules"
	"github.com/purawaktra/sumbing1-go/utils"
)

func main() {
	utils.InitConfig()
	utils.InitLog()

	utils.Log("============= SUMBING1 RUNTIME STARTED =============")
	// create gin engine
	engine := gin.New()

	// declare architecture
	requestHandler := modules.CreateSumbing1RequestHandler()
	router := modules.CreateSumbing1Router(engine, requestHandler)

	// init routing
	router.Init("/sumbing1/api/v1")

	utils.Log("============= SUMBING1 ENGINE STARTING =============")
	utils.Log(fmt.Sprintf("App Address = %s:%s", utils.AppHost, utils.AppPort))

	// init routing to swagger
	swaggerRouter := utils.CreateSwaggerRouter(engine)
	swaggerRouter.Init("/sumbing1/api/v1/swagger")

	// start http api engine
	err := engine.Run(fmt.Sprintf("%s:%s", utils.AppHost, utils.AppPort))
	if err != nil {
		utils.Fatal(err, "main", "")
		panic(err)
	}

	utils.Log("============= SUMBING1 ENGINE FAILED =============")

}
