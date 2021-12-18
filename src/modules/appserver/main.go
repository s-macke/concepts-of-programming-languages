package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"plugin"
)

var serviceServer *http.Server = nil
var serviceEngine *gin.Engine = nil

func RunStoppableService(engine *gin.Engine) *http.Server {
	srv := &http.Server{
		Addr:    ":8081",
		Handler: engine,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()
	return srv
}

func RegisterNewService(c *gin.Context) {
	filename := c.Query("filename")
	p, err := plugin.Open(filename)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	if serviceServer != nil {
		serviceServer.Shutdown(context.Background())
	}

	serviceEngine = gin.Default()
	serviceEngine.SetTrustedProxies(nil)
	serviceServer = RunStoppableService(serviceEngine)

	regService, err := p.Lookup("RegisterService")
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	regService.(func(router *gin.Engine))(serviceEngine)
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/register", RegisterNewService)
	r.Run(":8080") // listen and serve on
}
