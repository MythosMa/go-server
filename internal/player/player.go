package player

import (
	"github.com/fasthttp/websocket"
)

type Player struct {
	Conn *websocket.Conn
	ID   string
}

func NewPlayer(playerID string, conn *websocket.Conn) *Player {
	return &Player{Conn: conn, ID: playerID}
}
