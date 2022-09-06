package main

import (
	"context"
	"fmt"
	"github.com/karthick-workspace/grpc-workshop/api"
	"google.golang.org/grpc"
	"math/rand"
	"os"
	"time"
)

var port = ":8080"

func AskingDateTime(ctx context.Context, m api.RandomClient) (*api.DateTime, error) {
	request := &api.RequestDateTime{
		Value: "Please send me the date and time",
	}

	return m.GetDate(ctx, request)
}

func AskPass(ctx context.Context, m api.RandomClient, seed int64, length int64) (*api.RandomPass, error) {
	request := &api.RequestPass{
		Seed:   seed,
		Length: length,
	}

	return m.GetRandomPass(ctx, request)
}

func AskRandom(ctx context.Context, m api.RandomClient, seed int64, place int64) (*api.RandomInt, error) {
	request := &api.RandomParams{
		Seed:  seed,
		Place: place,
	}

	return m.GetRandom(ctx, request)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default port:", port)
	} else {
		port = os.Args[1]
	}

	conn, err := grpc.Dial(port, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Dial:", err)
		return
	}

	rand.Seed(time.Now().Unix())
	seed := int64(rand.Intn(100))

	client := api.NewRandomClient(conn)

	r, err := AskingDateTime(context.Background(), client)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server Date and Time:", r.Value)

	length := int64(rand.Intn(20))
	p, err := AskPass(context.Background(), client, 100, length+1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Random Password:", p.Password)

	place := int64(rand.Intn(100))
	i, err := AskRandom(context.Background(), client, seed, place)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Random Integer 1:", i.Value)

	k, err := AskRandom(context.Background(), client, seed, place-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Random Integer 1:", k.Value)
}
