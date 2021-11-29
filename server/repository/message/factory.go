package message

import (
	"github.com/jucardi/go-beans/beans"
	"github.com/jucardi/go-strings/stringx"
	"github.com/jucardi/go-titan/logx"
)

// BeanRepositoryDefault is the name of the bean implementation of IRepository
const BeanRepositoryDefault = "message-repository-default"

func init() {
	logx.WithObj(
		beans.RegisterFunc((*IRepository)(nil), BeanRepositoryDefault, func() interface{} {
			instance := &repository{}
			return instance.init()
		}, true),
	).Fatal("Unable to register repository bean " + BeanRepositoryDefault)

	logx.WithObj(
		beans.SetPrimary((*IRepository)(nil), BeanRepositoryDefault),
	).Fatal("Unable to set repository primary bean " + BeanRepositoryDefault)
}

// Repository returns the repository implementation registered by the given name. If no name is provided, returns the primary bean
func Repository(name ...string) IRepository {
	return beans.Resolve((*IRepository)(nil), stringx.GetOrDefault("", name...)).(IRepository)
}
