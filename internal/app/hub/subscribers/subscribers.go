package subscribers

import (
	"sync"
	"time"
)

type Subscribers struct {
	lock sync.RWMutex
	list map[string]*Subscriber
}

func New() *Subscribers {
	return &Subscribers{
		list: make(map[string]*Subscriber),
	}
}

func (s *Subscribers) Exists(id string) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	_, exists := s.list[id]
	return exists
}

func (s *Subscribers) List() []string {
	s.lock.RLock()
	defer s.lock.RUnlock()

	resp := make([]string, 0, len(s.list))

	for id := range s.list {
		resp = append(resp, id)
	}

	return resp
}

func (s *Subscribers) Add(row *Subscriber) {
	s.lock.Lock()
	defer s.lock.Unlock()

	row.ConnectedTime = time.Now()

	s.list[row.ID] = row
}

func (s *Subscribers) Delete(id string) {
	s.lock.Lock()
	defer s.lock.Unlock()

	delete(s.list, id)
}
