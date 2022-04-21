package helloworld

import (
	"github.com/jucardi/go-beans/beans"
	"github.com/jucardi/go-strings/stringx"
	"github.com/jucardi/go-titan/logx"
	"{{.golang.module_path}}/{{.service_name}}/api/helloworld"
)

func init() {
	logx.WithObj(
		beans.RegisterFunc((*helloworld.IApi)(nil), BeanApiDefault, func() interface{} {
			instance := &controller{}
			return instance.init()
		}, true), // This boolean indicates this component will be registered as a singleton instance.
	).Fatal("Unable to register controller bean " + BeanApiDefault)

	logx.WithObj(
		beans.SetPrimary((*helloworld.IApi)(nil), BeanApiDefault),
	).Fatal("Unable to set controller primary bean " + BeanApiDefault)
}

// Controller returns the controller implementation registered by the given name. If no name is provided, returns the primary bean
func Controller(name ...string) helloworld.IApi {
	return beans.Resolve((*helloworld.IApi)(nil), stringx.GetOrDefault("", name...)).(helloworld.IApi)
}
