package rMongo

import (
	"fmt"
	"sync"

	"github.com/yasseldg/go-simple/configs/sEnv"
	"github.com/yasseldg/go-simple/logs/sLog"

	"github.com/yasseldg/mgm/v4"
)

type Inter interface {
	Log()
	GetColl(env, conn_name, db_name, coll_name string, indexes ...Index) (Collection, error)
}

type Manager struct {
	mu sync.Mutex

	clients ClientsMap
}

func NewManager() Manager {
	return Manager{clients: make(ClientsMap)}
}

func (m *Manager) Log() {
	m.mu.Lock()
	defer m.mu.Unlock()

	println()
	for conn_name, client := range m.clients {
		for _, database := range client.databases {
			for _, collection := range database.collections {
				sLog.Info("client: %s  ..  env: %s  ..  database: %s  ..  coll: %s \n", conn_name, client.connection.Environment, database.database.Name(), collection.collection.Name())
			}
		}
	}
}

func (m *Manager) GetColl(env, conn_name, db_name, coll_name string, indexes ...Index) (Collection, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	client, err := m.getClient(env, conn_name)
	if err != nil {
		return Collection{}, err
	}

	return client.getColl(db_name, coll_name, indexes...)
}

func (m *Manager) getClient(env, conn_name string) (*Client, error) {

	conn_name = sEnv.Get(fmt.Sprint("CONN_", env), conn_name)

	client := m.clients.get(conn_name)
	if client != nil {
		return client, nil
	}

	return m.setClient(env, conn_name)
}

func (m *Manager) setClient(env, conn_name string) (*Client, error) {

	mgm.SetDefaultConfig(getConfig(env))

	conn, err := newConnection(conn_name)
	if err != nil {
		err := fmt.Errorf(" newConnection( %s )  ..  err: %s", conn_name, err)
		return nil, err
	}

	client, err := mgm.NewClient(conn.getClientOpt())
	if err != nil {
		err := fmt.Errorf(" mgm.NewClient() for env: %s  ..  conn_name: %s  ..  err: %s", env, conn_name, err)
		return nil, err
	}

	m.clients[conn_name] = &Client{
		connection: *conn,
		client:     client,
		databases:  make(DatabasesMap),
	}
	return m.clients[conn_name], nil
}
