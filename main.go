package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"passthrough/pkg/openai"

	"golang.org/x/sync/errgroup"
)

func main() {

	ctx := context.Background()
	eg, _ := errgroup.WithContext(ctx)

	eg.Go(func() error {
		fmt.Println("Starting OpenAI proxy")
		err := http.ListenAndServe(":9999", &openai.OpenAIProxy{})
		return err
	})

	err := eg.Wait()
	log.Panic(err)

}
