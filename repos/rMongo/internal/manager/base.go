package manager

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/yasseldg/go-simple/repos/rMongo/internal/client"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/collection"
	"github.com/yasseldg/go-simple/repos/rMongo/internal/connection"

	"github.com/yasseldg/go-simple/repos/rIndex"

	"github.com/yasseldg/go-simple/configs/sEnv"
	"github.com/yasseldg/go-simple/logs/sLog"

	"github.com/yasseldg/mgm/v4"
)

type Base struct {
	mu sync.Mutex

	debug bool

	clients client.Map
}

func New() *Base {
	return &Base{clients: make(client.Map)}
}

func (m *Base) Log() {
	m.mu.Lock()
	defer m.mu.Unlock()

	println()
	for conn_name, client := range m.clients {
		for _, database := range client.Databases() {
			for _, collection := range database.Collections() {
				sLog.Info("client: %s  ..  env: %s  ..  database: %s  ..  coll: %s \n", conn_name, client.Env(), database.Name(), collection.Name())
			}
		}
	}
}

func (m *Base) SetDebug(debug bool) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.debug = debug
}

func (m *Base) GetColl(ctx context.Context, env, conn_name, db_name, coll_name string, indexes ...rIndex.Inter) (collection.Inter, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	client, err := m.getClient(env, conn_name)
	if err != nil {
		return nil, err
	}

	coll, err := client.GetColl(ctx, env, db_name, coll_name, indexes...)
	if err != nil {
		return nil, err
	}

	return collection.NewFull(coll), nil
}

func (m *Base) getClient(env, conn_name string) (*client.Base, error) {

	conn_name = sEnv.Get(fmt.Sprint("CONN_", env), conn_name)

	client := m.clients.Get(conn_name)
	if client != nil {
		return client, nil
	}

	return m.setClient(env, conn_name)
}

func (m *Base) setClient(env, conn_name string) (*client.Base, error) {

	mgm.SetDefaultConfig(getConfig(env))

	conn, err := connection.New(conn_name)
	if err != nil {
		return nil, fmt.Errorf(" newConnection( %s )  ..  err: %s", conn_name, err)
	}

	_client, err := mgm.NewClient(conn.GetClientOpt(m.debug))
	if err != nil {
		return nil, fmt.Errorf(" mgm.NewClient() for env: %s  ..  conn_name: %s  ..  err: %s", env, conn_name, err)
	}

	m.clients.Set(conn_name, client.New(*conn, _client))

	return m.clients[conn_name], nil
}

// private functions

func getConfig(env string) *mgm.Config {
	return &mgm.Config{
		CtxTimeout: time.Duration(sEnv.GetInt(fmt.Sprint("CTX_", env), 10)) * time.Second,
	}
}
