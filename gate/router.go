package gate

import (
    "github.com/zhsso/leafserver/game"
    "github.com/zhsso/leafserver/msg"
)

func init() {
    msg.JSONProcessor.SetRouter(&msg.Hello{}, game.ChanRPC)
}
