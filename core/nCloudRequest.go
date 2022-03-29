package core

type RequestInterface interface {
	GetPath() string
	GetMethod() string
}

type NCloudRequest struct {
	Path   string
	Method string
}

func (n NCloudRequest) GetPath() string {
	return n.Path
}

func (n NCloudRequest) GetMethod() string {
	return n.Method
}
