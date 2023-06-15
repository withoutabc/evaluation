package utils

import (
	"github.com/bwmarrin/snowflake"
	"log"
)

func GenerateId(n int64) int64 {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(n)
	if err != nil {
		log.Println(err)
		return 0
	}
	return node.Generate().Int64()
}
