package rpc

import (
	"go-video/src/common/rpcinterface/userrpcinterface"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

var UserServiceCli userrpcinterface.UserService

func init() {
	initUserServiceCli()
}

func initUserServiceCli() {

	//TODO: hard code for now, need univeal config services
	reg := etcd.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{
			"http://localhost:2379",
		}
	})

	etcd.NewRegistry()
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.video.service.user"),
	)

	service.Init()

	//
	UserServiceCli = userrpcinterface.NewUserService("go.video.service.user", service.Client())
}
