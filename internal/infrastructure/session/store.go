package session

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
)

// Store はインメモリのセッションストア。
type Store struct {
	mu       sync.RWMutex
	sessions map[string]string // sessionID -> userID
}

// NewStore は Store を生成する。
func NewStore() *Store {
	return &Store{sessions: make(map[string]string)}
}

// Create はセッションIDを発行してユーザーIDと紐付ける。
func (s *Store) Create(userID string) (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	id := hex.EncodeToString(b)

	s.mu.Lock()
	s.sessions[id] = userID
	s.mu.Unlock()
	return id, nil
}

// Get はセッションIDに紐付くユーザーIDを返す。
func (s *Store) Get(sessionID string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	userID, ok := s.sessions[sessionID]
	return userID, ok
}

// Delete はセッションを削除する。
func (s *Store) Delete(sessionID string) {
	s.mu.Lock()
	delete(s.sessions, sessionID)
	s.mu.Unlock()
}
