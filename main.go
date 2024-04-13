package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/client"
)

func main() {

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	_, inspectErr := cli.ContainerInspect(context.Background(), "containerId")
	if inspectErr != nil {
		panic(inspectErr)
	}

	fmt.Println("Hello, World!")
}
