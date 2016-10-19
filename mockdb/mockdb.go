package mockdb

import (
	"fmt"

	"github.com/alicebob/miniredis"
	"github.com/garyburd/redigo/redis"
)

// MockDB is a wrapper over the miniredis server.
type MockDB struct {
	*miniredis.Miniredis
}

// NewMiniRedis returns a miniredis server.
func NewMiniRedis() *MockDB {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	return &MockDB{s}
}

// ReadCounter returns the current contents of the counter in the database.
func (s *MockDB) ReadCounter() int64 {
	c, err := redis.Dial("tcp", s.Addr())
	if err != nil {
		panic(err)
	}
	i64, err := redis.Int64(c.Do("GET", "counter"))
	if err != nil {
		fmt.Println(err)
	}
	return i64
}

// IncrementCounter increments the counter by one.
func (s *MockDB) IncrementCounter() {
	c, err := redis.Dial("tcp", s.Addr())
	if err != nil {
		panic(err)
	}
	c.Send("MULTI")
	c.Send("INCR", "counter")
	_, err = c.Do("EXEC")
	if err != nil {
		fmt.Println(err)
	}
}

// DecrementCounter decrements the counter by one.
func (s *MockDB) DecrementCounter() {
	c, err := redis.Dial("tcp", s.Addr())
	if err != nil {
		panic(err)
	}
	c.Send("MULTI")
	c.Send("DECR", "counter")
	_, err = c.Do("EXEC")
	if err != nil {
		fmt.Println(err)
	}
}

// InitCounter sets the `counter` key to zero.
func (s *MockDB) InitCounter() {
	c, err := redis.Dial("tcp", s.Addr())
	if err != nil {
		panic(err)
	}
	c.Do("SET", "counter", int64(0))
}
