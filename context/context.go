package context

type Provider map[string]interface{}

var Env Provider

func New() Provider {
	return Env
}

func Add(name string, value interface{}) {
	Env[name] = value
}

func Get(name string) interface{} {
	return Env[name]
}
