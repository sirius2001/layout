package web

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sirius2001/loon/pkg/grpc/pb"
	"github.com/sirius2001/loon/pkg/kaf"
	"github.com/sirius2001/loon/pkg/log"

	"github.com/gin-gonic/gin"
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
		addr: addr,
	}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 确保在函数结束时关闭请求体
	var reocrd pb.AuditRecord
	if err := json.NewDecoder(r.Body).Decode(&reocrd); err != nil {
		log.Error("json decode error", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	go kaf.Message(&reocrd)
	// 返回成功状态
	w.WriteHeader(http.StatusOK)
}

// StartService implements core.ServiceInner.
func (a *WebServer) StartService() error {
	// 初始化 http.Server
	a.server = &http.Server{
		Addr:    a.addr, // 设置监听地址
		Handler: nil,    // 如果没有自定义 handler，使用默认 mux
	}

	http.HandleFunc("/notify", handler)
	if err := a.server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Info("web service stoped ...")
			return nil
		}
		log.Error("listenAndserver err", "err", err)
		return err
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
