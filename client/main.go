package main

import (
	"fmt"
	"log"
	"net/http"

	"orchestrator-servic/proto"

	"github.com/gin-gonic/gin"

	// "github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := proto.NewServiceClient(conn)

	g := gin.Default()

	g.get("/get-user", func(ctx *gin.Context) {
		if res, err := client.GetUserByName(ctx, ""); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"user": fmt.Sprint(res),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run client: %v", err)
	}
}
