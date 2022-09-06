package main

import (
	"context"
	"fmt"
	"github.com/karthick-workspace/grpc-workshop/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"math/rand"
	"net"
	"os"
	"time"
)

var min = 0
var max = 100

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1

	for {
		// For getting valid ASCII characters
		myRand := random(0, 94)
		newChar := string(startChar[0] + byte(myRand))

		temp = temp + newChar

		if i == len {
			break
		}
		i++
	}
	return temp
}

type RandomServer struct {
	api.UnimplementedRandomServer
}

func (RandomServer) GetDate(ctx context.Context, r *api.RequestDateTime) (*api.DateTime, error) {
	currentTime := time.Now()
	response := &api.DateTime{
		Value: currentTime.String(),
	}

	return response, nil
}

func (RandomServer) GetRandom(ctx context.Context, r *api.RandomParams) (*api.RandomInt, error) {
	rand.Seed(r.GetSeed())
	place := r.GetPlace()

	temp := random(min, max)

	for {
		place--
		if place <= 0 {
			break
		}
		temp = random(min, max)
	}

	response := &api.RandomInt{
		Value: int64(temp),
	}

	return response, nil
}

func (RandomServer) GetRandomPass(ctx context.Context, r *api.RequestPass) (*api.RandomPass, error) {

	rand.Seed(r.GetSeed())
	temp := getString(r.GetLength())

	response := &api.RandomPass{
		Password: temp,
	}

	return response, nil

}

var port = ":8080"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default port:", port)
	} else {
		port = os.Args[1]
	}

	server := grpc.NewServer()
	var randomServer RandomServer
	api.RegisterRandomServer(server, randomServer)

	reflection.Register(server)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Serving requests...")
	server.Serve(listen)
}