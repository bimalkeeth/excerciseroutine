package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type ConnectionPool struct {
	mutex sync.RWMutex
	list  map[int]net.Conn
}

func NewConnectionPool() *ConnectionPool {
	pool := &ConnectionPool{
		list: make(map[int]net.Conn),
	}
	return pool
}
func (pool *ConnectionPool) Add(c net.Conn) int {
	pool.mutex.Lock()
	nextConnectionId := len(pool.list)
	pool.list[nextConnectionId] = c
	pool.mutex.Unlock()
	return nextConnectionId
}
func (pool *ConnectionPool) Get(connectionId int) net.Conn {
	pool.mutex.RLock()
	connection := pool.list[connectionId]
	pool.mutex.RUnlock()
	return connection
}
func (pool *ConnectionPool) Remove(connectionId int) {
	pool.mutex.Lock()
	delete(pool.list, connectionId)
	pool.mutex.Unlock()
}
func (pool *ConnectionPool) Size() int {
	return len(pool.list)
}
func (pool *ConnectionPool) Range(callback func(net.Conn, int)) {
	pool.mutex.RLock()
	for connectionId, connection := range pool.list {
		callback(connection, connectionId)
	}
	pool.mutex.RUnlock()
}

func main() {

	socket, err := net.Listen("tcp", "127.0.0.1")
	if err != nil {
		fmt.Println(err)
	}
	connPool := NewConnectionPool()
	go func(pool *ConnectionPool) {
		for {
			c, _ := socket.Accept()
			fmt.Println(c)
			cid := pool.Add(c)
			fmt.Println("New client id ", cid)

			pool.Range(func(con net.Conn, id int) {
				writer := bufio.NewWriter(con)
				_, _ = writer.WriteString("Git new Connection")
				_ = writer.Flush()
			})
		}
	}(connPool)
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGILL, syscall.SIGQUIT)
	for {
		<-c
		fmt.Println("killing")
		break
	}
}
