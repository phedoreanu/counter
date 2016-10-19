package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
	"golang.org/x/net/context"
)

// Datastore is an interface helper that allows mocking a test database.
type Datastore interface {
	ReadCounter() int64
	IncrementCounter()
	DecrementCounter()
	InitCounter()
}

// DB is a wrapper over the connection pool.
type DB struct {
	*pools.ResourcePool
}

// ResourceConn adapts a Redigo connection to a Vitess Resource.
type ResourceConn struct {
	redis.Conn
}

// Close closes a connection owned by a Vitess Resource.
func (r ResourceConn) Close() {
	r.Conn.Close()
}

// NewPool returns a new Redis connection pool.
func NewPool() *DB {
	p := pools.NewResourcePool(func() (pools.Resource, error) {
		c, err := redis.Dial("tcp", fmt.Sprintf("%s:6379", os.Getenv("REDIS_HOST")))
		return ResourceConn{c}, err
	}, 4, 100, time.Minute)
	//defer p.Close()
	return &DB{p}
}

// ReadCounter returns the current contents of the counter in the database.
func (p *DB) ReadCounter() int64 {
	ctx := context.TODO()
	r, err := p.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer p.Put(r)
	c := r.(ResourceConn)
	i64, err := redis.Int64(c.Do("GET", "counter"))
	if err != nil {
		log.Panicln(err)
	}
	return i64
}

// IncrementCounter increments the counter by one.
func (p *DB) IncrementCounter() {
	ctx := context.TODO()
	r, err := p.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer p.Put(r)
	c := r.(ResourceConn)
	c.Send("MULTI")
	c.Send("INCR", "counter")
	//c.Send("EXEC")
	_, err = c.Do("EXEC")
	if err != nil {
		log.Panicln(err)
	}
}

// DecrementCounter decrements the counter by one.
func (p *DB) DecrementCounter() {
	ctx := context.TODO()
	r, err := p.Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer p.Put(r)
	c := r.(ResourceConn)
	c.Send("MULTI")
	c.Send("DECR", "counter")
	//c.Send("EXEC")
	_, err = c.Do("EXEC")
	if err != nil {
		log.Panicln(err)
	}
}

// InitCounter sets the `counter` key to zero.
func (p *DB) InitCounter() {
	ctx := context.TODO()
	r, err := p.Get(ctx)
	if err != nil {
		log.Panicln(err)
	}
	defer p.Put(r)
	c := r.(ResourceConn)
	c.Do("SET", "counter", int64(0))
}
