package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/cholick/cloud-build-sample/pkg/controller"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

const DefaultPort = 3000

func main() {
	rawLogger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	logger := rawLogger.Sugar()

	router := httprouter.New()
	router.GET("/*path", controller.Index)

	port := getPort()
	logger.Infow("Listening", "port", port)

	err = http.ListenAndServe(fmt.Sprintf(":%v", port), router)
	if err != nil {
		logger.Error(err)
	}
}

func getPort() int {
	portString := os.Getenv("PORT")
	if portString == "" {
		return DefaultPort
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		return DefaultPort
	}

	return port
}
