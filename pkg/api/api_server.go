package api

import (
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	addr  string
	egine *gin.Engine
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr:  addr,
		egine: gin.Default(),
	}
}

func (a *APIServer) Run() error {
	if err := a.egine.Run(a.addr); err != nil {
		return err
	}
	return nil
}

func (a *APIServer) Engine() *gin.Engine {
	return a.egine
}
