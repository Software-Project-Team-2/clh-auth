package auth_service

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

func GenerateUserId() int64 {
	node, err := snowflake.NewNode(1)

	if err != nil {
		log.Printf("Error creating snowflake node: %v", err)
	}

	return node.Generate().Int64()
}
