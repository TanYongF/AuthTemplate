package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
)

func generateUUID() int64 {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(node)
		return -1
	}
	id := node.Generate()
	return id.Int64()
}
