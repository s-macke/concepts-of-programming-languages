package session

import (
	"context"
	"sync"
	"time"
)

// Store keeps the running sessions in memory. Nothing is persisted: restarting
// the server drops all sessions, which is exactly right for a lecture tool.
type Store struct {
	ttl time.Duration

	mu       sync.RWMutex
	sessions map[string]Session
}

// NewStore creates a store whose sessions expire after ttl of inactivity.
func NewStore(ttl time.Duration) *Store {
	return &Store{ttl: ttl, sessions: make(map[string]Session)}
}

// Add stores a session.
func (s *Store) Add(sess Session) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.sessions[sess.ID()] = sess
}

// Get looks a session up by id.
func (s *Store) Get(id string) (Session, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sess, ok := s.sessions[id]
	return sess, ok
}

// Len reports the number of live sessions.
func (s *Store) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.sessions)
}

// Janitor removes expired sessions until ctx is cancelled. Run it in its own
// goroutine; without it an uploaded quiz would be kept forever.
func (s *Store) Janitor(ctx context.Context, every time.Duration) {
	ticker := time.NewTicker(every)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.evict(time.Now().Add(-s.ttl))
		}
	}
}

// evict drops every session that was last used before deadline.
func (s *Store) evict(deadline time.Time) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	removed := 0
	for id, sess := range s.sessions {
		if sess.LastSeen().Before(deadline) {
			delete(s.sessions, id)
			removed++
		}
	}
	return removed
}
