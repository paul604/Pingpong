package main

import (
        "github.com/paul604/pingpong/pingpong"
        "log"
        "google.golang.org/grpc"
        "github.com/paul604/pingpong/score"
        "google.golang.org/grpc/codes"
        "net"
        "google.golang.org/grpc/reflection"
        "golang.org/x/net/context"
        "fmt"
)

type server struct{}

const (
        port = ":50050"
)

var connectionScore score.ScoreClient

//<editor-fold desc="impl PingpongServer">
func (s *server) Ping(ctx context.Context, msg *pingpong.PingMessage) (*pingpong.PingReply, error) {
        fmt.Println("Ping")
        connectionScore.Set(ctx, &score.SetMessage{"ping", 1})
        return &pingpong.PingReply{}, nil
}

func (s *server) Pong(ctx context.Context, msg *pingpong.PongMessage) (*pingpong.PongReply, error) {
        fmt.Println("Pong")
        connectionScore.Set(ctx, &score.SetMessage{"pong", 1})
        return &pingpong.PongReply{}, nil
}

func (s *server) Start(ctx context.Context, msg *pingpong.StartMessage) (*pingpong.StartReply, error) {
        fmt.Println("Start")

        scorePing, _ := connectionScore.Get(ctx, &score.GetMessage{"ping"})
        scorePong, _ := connectionScore.Get(ctx, &score.GetMessage{"pong"})

        if scorePing.Score != 0 || scorePong.Score != 0 {
                return nil, grpc.Errorf(codes.AlreadyExists, "partie run")
        }
        return &pingpong.StartReply{}, nil
}

func (s *server) Stop(ctx context.Context, msg *pingpong.StopMessage) (*pingpong.StopReply, error) {
        fmt.Println("Stop")
        connectionScore.Reset(ctx, &score.ResetMessage{})

        scorePing, _ := connectionScore.Get(ctx, &score.GetMessage{"ping"})
        scorePong, _ := connectionScore.Get(ctx, &score.GetMessage{"pong"})
        if scorePing.Score != 0 || scorePong.Score != 0 {
                return nil, grpc.Errorf(codes.Unknown, "error")
        }

        return &pingpong.StopReply{}, nil
}
//</editor-fold>

func main() {
        fmt.Println("init")

        lis, err := net.Listen("tcp", port)
        if err != nil {
                log.Fatalf("failed to listen: %v", err)
        }
        s := grpc.NewServer()
        pingpong.RegisterPingpongServer(s, &server{})
        // Register reflection service on gRPC server.
        reflection.Register(s)


        //connection score
        conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()
        connectionScore = score.NewScoreClient(conn)
        //fin conn

        if err := s.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %v", err)
        }

        fmt.Println("init OK")
}
