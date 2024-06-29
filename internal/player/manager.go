package player

import (
	"log"
	"sync"
)

type Manager struct {
	players map[string]*Player
	mu      sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		players: make(map[string]*Player),
	}
}

func (m *Manager) AddPlayer(player *Player) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.players[player.ID] = player
	log.Printf("Player added: %s", player.ID)
}

func (m *Manager) RemovePlayer(playerID string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.players, playerID)
	log.Printf("Player removed: %s", playerID)
}

func (m *Manager) GetPlayer(playerID string) *Player {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.players[playerID]
}

func (m *Manager) GetAllPlayers() map[string]*Player {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.players
}

func (m *Manager) BroadcastMessage(messageType int, message []byte) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, player := range m.players {
		err := player.Conn.WriteMessage(messageType, message)
		if err != nil {
			log.Printf("Failed to broadcast message to player %s: %v", player.ID, err)
		}
	}
}
