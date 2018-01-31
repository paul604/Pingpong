package main

import (
        "github.com/paul604/pingpong/score"
        "net"
        "log"
        "google.golang.org/grpc/reflection"
        "google.golang.org/grpc"
        "golang.org/x/net/context"
        "fmt"
)

type server struct{}

const (
        port = ":50051"
)

var scores = map[string]int32{}

//<editor-fold desc="impl SetpongServer">
func (s *server) Set(ctx context.Context, msg *score.SetMessage) (*score.SetReply, error) {
        scores[msg.Nom] = msg.Score
        fmt.Printf("Set: %s %v\n", msg.Nom, msg.Score)
        return &score.SetReply{}, nil
}
func (s *server) Get(ctx context.Context, msg *score.GetMessage) (*score.GetReply, error) {
        fmt.Printf("Get: %s %v\n", msg.Nom,  scores[msg.Nom])
        return &score.GetReply{Score: scores[msg.Nom]}, nil
}
func (s *server) Reset(context.Context, *score.ResetMessage) (*score.ResetReply, error) {
        fmt.Println("Reset")
        for key := range scores {
                scores[key] = 0
        }
        return &score.ResetReply{}, nil
}
//</editor-fold>

func main() {
        lis, err := net.Listen("tcp", port)
        if err != nil {
                log.Fatalf("failed to listen: %v", err)
        }
        s := grpc.NewServer()
        score.RegisterScoreServer(s, &server{})
        // Register reflection service on gRPC server.
        reflection.Register(s)
        if err := s.Serve(lis); err != nil {
                log.Fatalf("failed to serve: %v", err)
        }
}
