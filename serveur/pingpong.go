package main

import (
        "context"
        "github.com/paul604/pingpong/pingpong"
)

type server struct{}

//<editor-fold desc="impl PingpongServer">
func (s *server) Ping(context.Context, *pingpong.PingMessage) (*pingpong.PingReply, error) {
        return nil, nil
}
func (s *server) Pong(context.Context, *pingpong.PongMessage) (*pingpong.PongReply, error) {
        return nil, nil
}
func (s *server) Start(context.Context, *pingpong.StartMessage) (*pingpong.StartReply, error) {
        return nil, nil
}
func (s *server) Stop(context.Context, *pingpong.StopMessage) (*pingpong.StopReply, error) {
        return nil, nil
}
//</editor-fold>
