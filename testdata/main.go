package main

import (
	"context"
	"fmt"
	"github.com/zigen/go-missing-type-generator/testdata/go-sample-api"
)

func main() {
	cfg := go_sample_api.NewConfiguration()
	c := go_sample_api.NewAPIClient(cfg)
	ctx := context.Background()
	result, response, err := c.DefaultApi.GetState(ctx)
	if err != nil {
		fmt.Printf("error: %#v\n", err)
	}
	println(result, response, err)
}
