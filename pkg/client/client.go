package client

import (
	"github.com/atomix/atomix-go/pkg/client/_map"
	"github.com/atomix/atomix-go/pkg/client/lock"
	"github.com/atomix/atomix-go/pkg/client/protocol"
	"github.com/atomix/atomix-go/pkg/client/session"
	"google.golang.org/grpc"
)

func NewClient(address string, opts ...grpc.DialOption) (*Client, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
	}, nil
}

type Client struct {
	conn *grpc.ClientConn
}

func (c *Client) NewMap(name string, protocol *protocol.Protocol, opts ...session.Option) (*_map.Map, error) {
	return _map.NewMap(c.conn, name, protocol, opts...)
}

func (c *Client) NewLock(name string, protocol *protocol.Protocol, opts ...session.Option) (*lock.Lock, error) {
	return lock.NewLock(c.conn, name, protocol, opts...)
}

func (c *Client) Close() error {
	return c.conn.Close()
}