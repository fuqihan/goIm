package goIm

import (
	"net"
	"sync"
	"time"
)

type roomer interface {
	Join(conn net.Conn, obj JoinRoomApi) error
}

type localRoom struct {
	createDate time.Time
	mode       uint32
	users      []string
	role       map[uint][]string
	mu         sync.Mutex
}

type room struct {
	localRoom map[string]*localRoom
	mu        sync.Mutex
}

func (c *room) Join(conn net.Conn, obj JoinRoomApi) error {

	if data, ok := c.localRoom[obj.RoomName]; ok {
		data.mu.Lock()
		defer data.mu.Unlock()

	} else {
		c.mu.Lock()
		defer c.mu.Unlock()
		c.localRoom[obj.RoomName] = newRoomMap()
		c.localRoom[obj.RoomName].createDate = time.Now()
	}
	users := make([]string, 0)
	if obj.UserId != "" {
		users = append(users, obj.UserId)
	} else {
		users = append(users, obj.UserIds...)
	}
	c.localRoom[obj.RoomName].users = append(c.localRoom[obj.RoomName].users, users...)
	return nil

}

func newRoomMap() *localRoom {
	return &localRoom{}
}

func NewRoomLogic() roomer {
	return &room{}
}
