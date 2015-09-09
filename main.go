package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"github.com/zhsso/leafserver/conf"
	"github.com/zhsso/leafserver/game"
	"github.com/zhsso/leafserver/gate"
	"github.com/zhsso/leafserver/login"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath

	leaf.Run(
		game.Module,
		gate.Module,
		login.Module,
	)
}
