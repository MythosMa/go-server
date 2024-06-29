package websocket

import (
	"fmt"
	"log"

	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"

	"game-server/internal/player"
)

type Handler struct {
	playerManager *player.Manager
	upgrader      websocket.FastHTTPUpgrader
}

func NewHandler() *Handler {
	return &Handler{
		playerManager: player.NewManager(),
		upgrader: websocket.FastHTTPUpgrader{
			CheckOrigin: func(ctx *fasthttp.RequestCtx) bool {
				return true
			},
		},
	}
}

func (h *Handler) ServeHTTP(ctx *fasthttp.RequestCtx) {
	h.upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
		defer conn.Close()

		playerID := fmt.Sprintf("player-%d", len(h.playerManager.GetAllPlayers())+1)
		newPlayer := player.NewPlayer(playerID, conn)
		h.playerManager.AddPlayer(newPlayer)

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("玩家 %s 断开连接: %v", playerID, err)
				h.playerManager.RemovePlayer(playerID)
				return
			}

			log.Printf("收到玩家 %s 消息: %s", playerID, message)
			h.playerManager.BroadcastMessage(messageType, []byte(playerID+": "+string(message)))
		}
	})
}
