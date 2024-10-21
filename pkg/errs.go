package pkg

import "errors"

var (
	ErrConfigNotLoad = errors.New("config is not loaded")
	ErrDialDB        = errors.New("dial db with error")
	ErrDbNotSupport  = errors.New("not supported db kind")
	ErrGinEginnil    = errors.New("gin engin with err")
)
