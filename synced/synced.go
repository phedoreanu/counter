package synced

import "sync"

// Synced uses a RWMutex to sync hits and direction.
type Synced struct {
	sync.RWMutex
	hits      int64
	direction bool
}

// NewSynced returns a new Hits
func NewSynced() *Synced {
	return new(Synced)
}

// IncrementHits increases the number hits by 1.
func (h *Synced) IncrementHits() int64 {
	h.Lock()
	h.hits++
	h.Unlock()
	return h.Hits()
}

// Hits returns the number of hits.
func (h *Synced) Hits() int64 {
	h.RLock()
	defer h.RUnlock()
	return h.hits
}

// Reverse reverses the direction of the counter.
func (h *Synced) Reverse() {
	h.Lock()
	h.direction = !h.direction
	h.Unlock()
}

// Direction returns the direction.
func (h *Synced) Direction() bool {
	h.RLock()
	defer h.RUnlock()
	return h.direction
}
