package core

type ServiceInner interface {
	StartService() error
	StopService() error
	ServiceName() string
	ServiceAddr() string
}
