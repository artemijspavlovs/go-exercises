package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

type CombinedResult struct {
	Numbers  []int
	Words    []string
	Booleans []bool
}

func main() {
	fd, err := FetchData()
	if err != nil {
		log.Println("fetch data error:", err)
	}

	j, _ := json.MarshalIndent(fd, "", "  ")
	fmt.Println(string(j))
}

func FetchNumbers() []int {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	return []int{1, 6, 4, 3}
}

func fetchNumbers() chan []int {
	c := make(chan []int)
	go func() {
		c <- FetchNumbers()
	}()
	return c
}

func FetchWords() []string {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	return []string{"a", "b", "c", "d", "e", "f"}
}

func fetchWords() chan []string {
	c := make(chan []string)
	go func() {
		c <- FetchWords()
	}()
	return c
}

func FetchBooleans() []bool {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	return []bool{true, false, true, false, true}
}

func fetchBooleans() chan []bool {
	c := make(chan []bool)
	go func() {
		c <- FetchBooleans()
	}()
	return c
}

func FetchData() (*CombinedResult, error) {
	result := &CombinedResult{}
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	errGrp, ctx := errgroup.WithContext(ctx)

	errGrp.Go(func() error {
		select {
		case numbers := <-fetchNumbers():
			result.Numbers = numbers
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	errGrp.Go(func() error {
		select {
		case words := <-fetchWords():
			result.Words = words
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	errGrp.Go(func() error {
		select {
		case booleans := <-fetchBooleans():
			result.Booleans = booleans
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	if err := errGrp.Wait(); err != nil {
		return nil, err
	}

	return result, nil
}
