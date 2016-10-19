package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/phedoreanu/counter/db"
	"github.com/phedoreanu/counter/synced"
)

// Env is a dependency injection helper.
type Env struct {
	db     db.Datastore
	synced *synced.Synced
}

// NewEnv returns a new Env from a db.Datastore.
func NewEnv(db db.Datastore) *Env {
	return &Env{
		db:     db,
		synced: synced.NewSynced(),
	}
}

// GetHandler returns the number of hits to /get endpoint.
func (env *Env) handleGet(w http.ResponseWriter, _ *http.Request) {
	env.SyncCounter(env.synced.IncrementHits())
	fmt.Fprint(w, env.synced.Hits())
}

// StatusHandler returns the current contents of the counter in the database.
func (env *Env) handleStatus(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, env.db.ReadCounter())
}

// SyncCounter increments the counter in the database
// and every 10 hits it flips and decrements it.
func (env *Env) SyncCounter(hits int64) {
	switch env.synced.Direction() {
	case true:
		env.db.DecrementCounter()
	case false:
		env.db.IncrementCounter()
	}
	if hits%10 == 0 {
		env.synced.Reverse()
	}
}

func main() {
	env := NewEnv(db.NewPool())
	env.db.InitCounter()

	http.HandleFunc("/get", env.handleGet)
	http.HandleFunc("/status", env.handleStatus)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
