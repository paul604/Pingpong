package main

import (
        "context"
        "github.com/paul604/pingpong/pingpong"
        "log"
        "google.golang.org/grpc"
        "github.com/paul604/pingpong/score"
        "google.golang.org/grpc/codes"
)

type server struct{}

const (
        Port = 50050
)

var connectionScore score.ScoreClient

//<editor-fold desc="impl PingpongServer">
func (s *server) Ping(context.Context, *pingpong.PingMessage) (*pingpong.PingReply, error) {
        return nil, nil
}

func (s *server) Pong(context.Context, *pingpong.PongMessage) (*pingpong.PongReply, error) {
        return nil, nil
}

func (s *server) Start(ctx context.Context, msg *pingpong.StartMessage) (*pingpong.StartReply, error) {

        if connectionScore.Get(ctx, score.GetMessage{"ping"}) != 0 || connectionScore.Get(ctx, score.GetMessage{"pong"}) != 0 {
                return nil, grpc.Errorf(codes.AlreadyExists, "partie run")
        }

        return pingpong.StartReply{}, nil
}
func (s *server) Stop(ctx context.Context, msg *pingpong.StopMessage) (*pingpong.StopReply, error) {
        connectionScore.Reset(ctx, score.ResetMessage{})
        return pingpong.StopReply{}, nil
}
//</editor-fold>

func main() {
        conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()
        connectionScore = score.NewScoreClient(conn)
}
