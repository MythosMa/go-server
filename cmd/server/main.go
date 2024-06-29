package main

import (
	"fmt"

	"github.com/valyala/fasthttp"

	"game-server/internal/websocket"
)

func main() {
	wsHandler := websocket.NewHandler()
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			fmt.Fprintf(ctx, "Welcome to the game server!")
		case "/ws":
			wsHandler.ServeHTTP(ctx)
		default:
			ctx.Error("404 page not found", fasthttp.StatusNotFound)
		}
	}

	address := "0.0.0.0:9000"
	fmt.Printf("Server started on %s\n", address)
	if err := fasthttp.ListenAndServe(address, requestHandler); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
