package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type resp struct {
	val int
	err error
}

func main() {
	ch := make(chan resp)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go RPCCall(ctx, ch)
	a := <-ch
	fmt.Println(a.val, a.err)
}

func RPCCall(ctx context.Context, ch chan<- resp) {
	select {
	case <-ctx.Done():
		ch <- resp{
			err: errors.New("Error!!!!!"),
		}
	case <-time.After(time.Duration(rand.Intn(5)) * time.Second):
		ch <- resp{
			val: rand.Intn(100),
		}
	}
}
