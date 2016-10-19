package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/phedoreanu/counter/mockdb"
)

func TestStatusNoHits(t *testing.T) {
	env := NewEnv(mockdb.NewMiniRedis())
	env.db.InitCounter()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/status", nil)
	http.HandlerFunc(env.handleStatus).ServeHTTP(rec, req)

	expected := "0"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestOneHit(t *testing.T) {
	env := NewEnv(mockdb.NewMiniRedis())
	env.db.InitCounter()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get", nil)
	http.HandlerFunc(env.handleGet).ServeHTTP(rec, req)

	req, _ = http.NewRequest("GET", "/status", nil)
	rec = httptest.NewRecorder()
	http.HandlerFunc(env.handleStatus).ServeHTTP(rec, req)

	expected := "1"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func Test10Hits(t *testing.T) {
	env := NewEnv(mockdb.NewMiniRedis())
	env.db.InitCounter()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get", nil)

	for i := 0; i < 10; i++ {
		http.HandlerFunc(env.handleGet).ServeHTTP(rec, req)
	}

	req, _ = http.NewRequest("GET", "/status", nil)
	rec = httptest.NewRecorder()
	http.HandlerFunc(env.handleStatus).ServeHTTP(rec, req)

	expected := "10"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func Test19Hits(t *testing.T) {
	env := NewEnv(mockdb.NewMiniRedis())
	env.db.InitCounter()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get", nil)

	for i := 0; i < 19; i++ {
		http.HandlerFunc(env.handleGet).ServeHTTP(rec, req)
	}

	req, _ = http.NewRequest("GET", "/status", nil)
	rec = httptest.NewRecorder()
	http.HandlerFunc(env.handleStatus).ServeHTTP(rec, req)

	expected := "1"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func Test20Hits(t *testing.T) {
	env := NewEnv(mockdb.NewMiniRedis())
	env.db.InitCounter()

	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/get", nil)

	for i := 0; i < 20; i++ {
		http.HandlerFunc(env.handleGet).ServeHTTP(rec, req)
	}

	req, _ = http.NewRequest("GET", "/status", nil)
	rec = httptest.NewRecorder()
	http.HandlerFunc(env.handleStatus).ServeHTTP(rec, req)

	expected := "0"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

/*func TestReadCounter(t *testing.T) {
	env := NewEnv(db.NewPool())
	env.db.InitCounter()

	c := env.db.ReadCounter()
	expected := int64(0)
	if expected != c {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, c)
	}
}

func TestIncrementCounter(t *testing.T) {
	env := NewEnv(db.NewPool())
	env.db.InitCounter()

	env.db.IncrementCounter()
	c := env.db.ReadCounter()
	expected := int64(1)
	if expected != c {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, c)
	}
}

func TestDecrementCounter(t *testing.T) {
	env := NewEnv(db.NewPool())
	env.db.InitCounter()

	env.db.IncrementCounter()
	env.db.DecrementCounter()
	c := env.db.ReadCounter()
	expected := int64(0)
	if expected != c {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, c)
	}
}*/
