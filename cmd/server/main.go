package main

import (
	"fmt"
	"game-server/internal/db"
	"game-server/internal/user"
	"game-server/internal/websocket"
	"log"

	"github.com/valyala/fasthttp"

)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.DB.Close()

	wsHandler := websocket.NewHandler()
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			fmt.Fprintf(ctx, "Welcome to the game server!")
		case "/register":
			user.Register(ctx)
		case "/login":
			user.Login(ctx)
		case "/ws":
			wsHandler.ServeHTTP(ctx)
		default:
			ctx.Error("404 page not found", fasthttp.StatusNotFound)
		}
	}

	address := "0.0.0.0:3000"
	fmt.Printf("Server started on %s\n", address)
	if err := fasthttp.ListenAndServe(address, requestHandler); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
