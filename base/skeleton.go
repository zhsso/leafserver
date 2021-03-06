package base

import (
	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/module"
	"github.com/zhsso/leafserver/conf"
)

func NewSkeleton() *module.Skeleton {
	skeleton := &module.Skeleton{
		GoLen:              conf.GoLen,
		TimerDispatcherLen: conf.TimerDispatcherLen,
		ChanRPCServer:      chanrpc.NewServer(conf.ChanRPCLen),
	}
	skeleton.Init()
	return skeleton
}
