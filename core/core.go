package core

import (
	"sync"

	"github.com/sirius2001/layout/config"
	"github.com/sirius2001/layout/internal/router"
	"github.com/sirius2001/layout/pkg/api"
	"github.com/sirius2001/layout/pkg/db"
	"github.com/sirius2001/layout/pkg/log"
)

type Core struct {
	apiServer *api.APIServer
	serviceWg sync.WaitGroup
}

// 将会注册的路由组
var routes = []router.RouterInner{}

func NewCore(confPath string) (*Core, error) {
	if err := config.LoadConfig(confPath); err != nil {
		return nil, err
	}

	log.SetupLogger(log.Config{
		Dir:      config.Conf().Dir,
		Level:    config.Conf().Level,
		Duration: config.Conf().Duration,
		MaxAge:   config.Conf().MaxAge,
		MaxSize:  config.Conf().MaxSize,
	})

	if err := db.NewDB(config.Conf().Merge, config.Conf().DSN); err != nil {
		log.Error("NewCore", "err", err)
		return nil, err
	}

	apiService := api.NewAPIServer("")
	for _, v := range routes {
		v.Route(apiService.Engine())
	}

	return &Core{apiServer: apiService}, nil
}

func (c *Core) Run() {
	c.serviceWg.Add(1)
	go func() {
		if err := c.apiServer.Run(); err != nil {
			log.Error("pp start api serivice failed", "err", err)
		}
		c.serviceWg.Done()
	}()
	log.Info("app start api serivice successfully")
	c.serviceWg.Wait()
}
