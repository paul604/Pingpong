package main

import (
        "context"
        "github.com/paul604/pingpong/score"
)

type server struct{}

//<editor-fold desc="impl SetpongServer">
func (s *server) Set(context.Context, *score.SetMessage) (*score.SetReply, error) {
        return nil, nil
}
func (s *server) Get(context.Context, *score.GetMessage) (*score.GetReply, error) {
        return nil, nil
}
func (s *server) Reset(context.Context, *score.ResetMessage) (*score.ResetReply, error) {
        return nil, nil
}
//</editor-fold>
