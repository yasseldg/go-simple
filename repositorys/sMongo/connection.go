package sMongo

import (
	"fmt"

	"github.com/yasseldg/go-simple/configs/sEnv"
	"github.com/yasseldg/go-simple/logs/sLog"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection, Databases access data
type Connection struct {
	Environment      string `yaml: "environment"`
	Username         string `yaml: "username"`
	Password         string `yaml: "password"`
	Host             string `yaml: "host"`
	Port             string `yaml: "port"`
	AuthDatabase     string `yaml: "authdatabase"`
	Tls              bool   `yaml: "tls"`
	Protocol         string `yaml: "protocol"`
	AuthMechanism    string `yaml: "authmechanism"`
	ReadPreference   string `yaml: "readpreference"`
	DirectConnection bool   `yaml: "directconnection"`
}

// GetConnection, Databases access data predefined
func newConnection(name string) (*Connection, error) {
	var conn Connection
	err := sEnv.LoadYaml(fmt.Sprint(".env/mongodb/", name, ".yaml"), &conn)
	if err != nil {
		sLog.Error("sMongo: getConnection: can't load env file %s: %s", name, err)
		return nil, err
	}
	return &conn, nil
}

// getUri, return (Uri, Credentials)
func (conn Connection) getUri() (string, options.Credential) {

	optCredential := options.Credential{AuthMechanism: conn.AuthMechanism, AuthSource: conn.AuthDatabase, Username: conn.Username, Password: conn.Password}

	if conn.DirectConnection && conn.Tls && len(conn.ReadPreference) > 0 {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/?directConnection=%t&tls=%t&readPreference=%s",
			conn.Username, conn.Password, conn.Host, conn.Port, conn.DirectConnection, conn.Tls, conn.ReadPreference), optCredential
	}

	if len(conn.ReadPreference) > 0 && len(conn.AuthMechanism) > 0 && conn.DirectConnection {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/?readPreference=%s&authMechanism=%s&directConnection=%t",
			conn.Username, conn.Password, conn.Host, conn.Port, conn.ReadPreference, conn.AuthMechanism, conn.DirectConnection), optCredential
	}

	if conn.Environment == "prod" {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s", conn.Username, conn.Password, conn.Host, conn.Port), optCredential
	}

	return fmt.Sprintf("mongodb://%s:%s", conn.Host, conn.Port), optCredential
}

func (conn Connection) getClientOpt() *options.ClientOptions {

	Uri, Credentials := conn.getUri()

	sLog.Debug("sMongo: Uri: %s  ..  Credentials: %#v", Uri, Credentials)

	switch conn.Environment {
	case "prod":
		return options.Client().ApplyURI(Uri)

	default:
		return options.Client().ApplyURI(Uri).SetAuth(Credentials)
	}
}
