package sNet

import (
	"bytes"
	"context"
	"fmt"
	"path"
	"time"

	"github.com/yasseldg/go-simple/configs/sEnv"
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sBool"
	"github.com/yasseldg/go-simple/types/sInts"
	"github.com/yasseldg/go-simple/types/sJson"
)

type Service struct {
	env         string
	url         string
	secure      bool
	port        int
	protocol    string
	path_prefix string
	user        string
	secret      string

	debug bool
}

// NewService
func NewService(env, file_path string) (*Service, error) {
	if len(file_path) == 0 {
		name := sEnv.Get(env, "")

		if len(name) == 0 {
			return nil, fmt.Errorf("empty name")
		}

		file_path = fmt.Sprint(".env/services/", name, ".yaml")
	}

	m := new(model)
	err := sEnv.LoadYaml(file_path, m)
	if err != nil {
		return nil, fmt.Errorf("can't load env file %s: %s", file_path, err)
	}

	conf := m.Service()
	conf.env = env
	conf.update()

	conf.SetDebug(sEnv.GetBool(
		fmt.Sprintf("%s_Debug", env), false))

	return conf, nil
}

func (c *Service) Clone() InterService {
	return &Service{
		env:         c.env,
		url:         c.url,
		secure:      c.secure,
		port:        c.port,
		protocol:    c.protocol,
		path_prefix: c.path_prefix,
		user:        c.user,
		secret:      c.secret,
		debug:       c.debug,
	}
}

func (c *Service) SetPathPrefix(path_prefix string) {
	c.path_prefix = path_prefix
}

func (c *Service) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Service) Debug() bool {
	return c.debug
}

func (c *Service) String() string {
	return fmt.Sprintf("%s: %s", c.env, c.GetUrl())
}

func (c *Service) Log() {
	sLog.Info(c.String())
}

func (c *Service) Url() string {
	return c.url
}

func (c *Service) Port() int {
	return c.port
}

func (c *Service) Secret() string {
	return c.secret
}

func (c *Service) User() string {
	return c.user
}

func (c *Service) LocalAddr() string {
	port := 80
	if c.port > 0 {
		port = c.port
	}
	return fmt.Sprintf("0.0.0.0:%d", port)
}

func (c *Service) GetUri() string {
	uri := c.url
	if c.port > 0 {
		uri = fmt.Sprintf("%s:%d", uri, c.port)
	}
	if len(c.path_prefix) > 0 {
		uri = path.Join(uri, c.path_prefix)
	}
	return uri
}

func (c *Service) GetUrl() string {
	url := c.GetUri()
	if c.secure {
		return fmt.Sprintf("https://%s", url)
	}
	return fmt.Sprintf("http://%s", url)
}

func (c *Service) HandlePath(handler string) string {
	if len(c.path_prefix) > 0 {
		return fmt.Sprintf("/%s/%s", c.path_prefix, handler)
	}
	return fmt.Sprintf("/%s", handler)
}

func (c *Service) SendObj(end_point string, obj any) error {

	byteObj, err := sJson.ToByte(obj)
	if err != nil {
		return fmt.Errorf("sJson.ToByte(): %s", err)
	}

	request := NewRequest().MethodPost()
	request.SetBody(bytes.NewReader(byteObj))
	request.SetEndPoint(end_point)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = request.Call(ctx, c, nil)
	if err != nil {
		return fmt.Errorf("request.Call(): %s", err)
	}

	return nil
}

// private methods

func (c *Service) update() {
	c.url = sEnv.Get(fmt.Sprintf("%s_Url", c.env), c.url)
	c.port = sInts.Get(sEnv.Get(fmt.Sprintf("%s_Port", c.env), sInts.ToString(int64(c.port))))
	c.secure = sBool.Get(sEnv.Get(fmt.Sprintf("%s_Secure", c.env), sBool.ToString(c.secure)))
	c.protocol = sEnv.Get(fmt.Sprintf("%s_Protocol", c.env), c.protocol)
	c.path_prefix = sEnv.Get(fmt.Sprintf("%s_Path_Prefix", c.env), c.path_prefix)
	c.user = sEnv.Get(fmt.Sprintf("%s_User", c.env), c.user)
	c.secret = sEnv.Get(fmt.Sprintf("%s_Secret", c.env), c.secret)
}

// model for yaml

type model struct {
	Url        string
	Secure     bool
	Port       int
	Network    string
	PathPrefix string
	User       string
	Secret     string
}

func (c *model) Service() *Service {
	return &Service{
		url:         c.Url,
		secure:      c.Secure,
		port:        c.Port,
		protocol:    c.Network,
		path_prefix: c.PathPrefix,
		user:        c.User,
		secret:      c.Secret,
	}
}
