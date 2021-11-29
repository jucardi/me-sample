package message

import (
	"github.com/jucardi/go-beans/beans"
	"github.com/jucardi/go-strings/stringx"
	"github.com/jucardi/go-titan/logx"
)

// BeanServiceDefault is the name of the bean implementation of IService
const BeanServiceDefault = "message-service-default"

func init() {
	logx.WithObj(
		beans.RegisterFunc((*IService)(nil), BeanServiceDefault, func() interface{} {
			instance := &service{}
			return instance.init()
		}, true),
	).Fatal("Unable to register service bean " + BeanServiceDefault)

	logx.WithObj(
		beans.SetPrimary((*IService)(nil), BeanServiceDefault),
	).Fatal("Unable to set service primary bean " + BeanServiceDefault)
}

// Service returns the service implementation registered by the given name. If no name is provided, returns the primary bean
func Service(name ...string) IService {
	return beans.Resolve((*IService)(nil), stringx.GetOrDefault("", name...)).(IService)
}
