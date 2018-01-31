package main

import (
        "log"
        "google.golang.org/grpc"
        "github.com/paul604/pingpong/pingpong"
        "golang.org/x/net/context"
)

func main() {
        conn, err := grpc.Dial("localhost:50050", grpc.WithInsecure())
        if err != nil {
                log.Fatalf("did not connect: %v", err)
        }
        defer conn.Close()
        connectionPingPong := pingpong.NewPingpongClient(conn)
        connectionPingPong.Start(context.Background(), &pingpong.StartMessage{})
        connectionPingPong.Ping(context.Background(), &pingpong.PingMessage{})
        connectionPingPong.Pong(context.Background(), &pingpong.PongMessage{})
        connectionPingPong.Ping(context.Background(), &pingpong.PingMessage{})
}
