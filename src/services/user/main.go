package main

import (
	"go-video/src/common/log4video"
	"go-video/src/common/rpcinterface/userrpcinterface"
	"go-video/src/services/user/config"
	userrpc "go-video/src/services/user/rpc"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	reg := etcd.NewRegistry(func(o *registry.Options) {
		o.Addrs = []string{
			config.EtcdConf.Addr,
		}
	})

	srv := micro.NewService(
		micro.Registry(reg),
		micro.Name("go.video.service.user"),
	)

	srv.Init()

	err := userrpcinterface.RegisterUserServiceHandler(srv.Server(), new(userrpc.UserHandler))
	if err != nil {
		panic(err)
	}

	//listen in, hard code port now
	log4video.D("user service start register to etcd")
	err = srv.Run()
	if err != nil {
		panic(err)
	}
}
