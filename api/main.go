package main

import (
	"context"
	"log"

	"github.com/ei-sugimoto/technews/api/cmd"
	"github.com/ei-sugimoto/technews/api/internal/infra"
)

func main() {
	client := infra.NewClient()
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	cmd.Serve()
}
