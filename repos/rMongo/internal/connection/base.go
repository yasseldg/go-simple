package connection

import (
	"context"
	"fmt"

	"github.com/yasseldg/go-simple/configs/sEnv"
	"github.com/yasseldg/go-simple/logs/sLog"

	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection, Databases access data
type Base struct {
	Environment      string `yaml:"environment"`
	Username         string `yaml:"username"`
	Password         string `yaml:"password"`
	Host             string `yaml:"host"`
	Port             string `yaml:"port"`
	AuthDatabase     string `yaml:"authdatabase"`
	Tls              bool   `yaml:"tls"`
	Protocol         string `yaml:"protocol"`
	AuthMechanism    string `yaml:"authmechanism"`
	ReadPreference   string `yaml:"readpreference"`
	DirectConnection bool   `yaml:"directbection"`
	Debug            bool   `yaml:"debug"`
}

// GetConnection, Databases access data predefined
func New(name string) (*Base, error) {
	var b Base
	err := sEnv.LoadYaml(fmt.Sprint(".env/mongodb/", name, ".yaml"), &b)
	if err != nil {
		sLog.Error("sMongo: getConnection: can't load env file %s: %s", name, err)
		return nil, err
	}
	return &b, nil
}

func (b *Base) Env() string {
	return b.Environment
}

// getUri, return (Uri, Credentials)
func (b *Base) GetUri() (string, options.Credential) {

	optCredential := options.Credential{AuthMechanism: b.AuthMechanism, AuthSource: b.AuthDatabase, Username: b.Username, Password: b.Password}

	if b.DirectConnection && b.Tls && len(b.ReadPreference) > 0 {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/?directConnection=%t&tls=%t&readPreference=%s",
			b.Username, b.Password, b.Host, b.Port, b.DirectConnection, b.Tls, b.ReadPreference), optCredential
	}

	if len(b.ReadPreference) > 0 && len(b.AuthMechanism) > 0 && b.DirectConnection {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s/?readPreference=%s&authMechanism=%s&directConnection=%t",
			b.Username, b.Password, b.Host, b.Port, b.ReadPreference, b.AuthMechanism, b.DirectConnection), optCredential
	}

	if b.Environment == "prod" {
		return fmt.Sprintf("mongodb://%s:%s@%s:%s", b.Username, b.Password, b.Host, b.Port), optCredential
	}

	return fmt.Sprintf("mongodb://%s:%s", b.Host, b.Port), optCredential
}

func (b *Base) GetClientOpt(debug bool) *options.ClientOptions {

	options := options.Client()

	if b.Debug || debug {
		options.SetMonitor(&event.CommandMonitor{
			Started: func(_ context.Context, evt *event.CommandStartedEvent) {
				sLog.Warn("Mongo Command: %s  ..  %v", evt.CommandName, evt.Command)
			}})
	}

	Uri, Credentials := b.GetUri()

	// sLog.Debug("sMongo: Uri: %s  ..  Credentials: %#v", Uri, Credentials)

	switch b.Environment {
	case "prod":
		options.ApplyURI(Uri)

	default:
		options.ApplyURI(Uri).SetAuth(Credentials)
	}

	return options
}
