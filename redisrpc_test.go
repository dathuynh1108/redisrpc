package redisrpc

import (
	"context"
	"fmt"
	"redisrpc/testgrpc"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func Test(t *testing.T) {
	ctx := context.Background()
	r := redis.NewClient(
		&redis.Options{
			Addr:         "localhost:6379", // use default Addr
			Password:     "",               // no password set
			DB:           0,                // use default DB
			DialTimeout:  3 * time.Second,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		})
	_, err := r.Ping(ctx).Result()
	if err != nil {
		t.Errorf("Cannot connect to redis: %v", err)
		return
	}

	service := NewServer(r, "node_01")
	testServer := &testgrpc.Server{}
	testgrpc.RegisterTestServerServer(service, testServer)

	cli := testgrpc.NewTestServerClient(NewClient(r, "node_01", "node_01"))
	res, err := cli.MakeRequest(ctx, &testgrpc.Request{
		Message: "Dat",
	})
	fmt.Println(res)
}
