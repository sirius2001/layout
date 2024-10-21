package web

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirius2001/layout/pkg/log"
)

var routers = []RouterInner{}

type WebServer struct {
	addr   string
	egine  *gin.Engine
	server *http.Server
}

// ServiceAddr implements core.ServiceInner.
func (a *WebServer) ServiceAddr() string {
	return a.addr
}

// ServiceName implements core.ServiceInner.
func (a *WebServer) ServiceName() string {
	return "WebService"
}

// NewWebService implements core.ServiceInner.
func NewWebService(addr string) (*WebServer, error) {
	return &WebServer{
		addr:  addr,
		egine: gin.Default(),
	}, nil
}

// StartService implements core.ServiceInner.
func (a *WebServer) StartService() error {
	for _, router := range routers {
		router.Route(a.egine)
	}

	// 创建一个 HTTP 服务器
	a.server = &http.Server{
		Addr:    a.addr,  // 设置监听地址
		Handler: a.egine, // 设置路由
	}

	if err := a.server.ListenAndServe(); err != nil {
		panic(err)
	}

	return nil
}

// StopService implements core.ServiceInner.
func (a *WebServer) StopService() error {
	if err := a.server.Shutdown(context.Background()); err != nil {
		log.Error("webServer stop failed", "err", err)
		return err
	}
	return nil
}
