package sNet

import (
	"fmt"
	"path"

	"github.com/yasseldg/go-simple/configs/sEnv"
	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sBool"
	"github.com/yasseldg/go-simple/types/sInts"
)

type Service struct {
	env         string
	url         string
	secure      bool
	port        int
	protocol    string
	path_prefix string
}

// NewService
func NewService(env, file_path string) (*Service, error) {
	if len(file_path) == 0 {
		name := sEnv.Get(env, "DEV")
		file_path = fmt.Sprint(".env/services/", name, ".yaml")
	}

	m := new(model)
	err := sEnv.LoadYaml(file_path, m)
	if err != nil {
		return nil, fmt.Errorf("sNet: getConf: can't load env file %s: %s", file_path, err)
	}

	conf := m.Service()
	conf.env = env
	conf.update()

	return conf, nil
}

func (c *Service) Log() {
	sLog.Info("%s: %s ", c.env, c.GetUrl())
}

func (c *Service) Port() int {
	return c.port
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

// private methods

func (c *Service) update() {
	c.url = sEnv.Get(fmt.Sprintf("%s_Url", c.env), c.url)
	c.port = sInts.Get(sEnv.Get(fmt.Sprintf("%s_Port", c.env), sInts.ToString(int64(c.port))))
	c.secure = sBool.Get(sEnv.Get(fmt.Sprintf("%s_Secure", c.env), sBool.ToString(c.secure)))
	c.protocol = sEnv.Get(fmt.Sprintf("%s_Protocol", c.env), c.protocol)
	c.path_prefix = sEnv.Get(fmt.Sprintf("%s_Path_Prefix", c.env), c.path_prefix)
}

// model for yaml

type model struct {
	Url        string
	Secure     bool
	Port       int
	Network    string
	PathPrefix string
}

func (c *model) Service() *Service {
	return &Service{
		url:         c.Url,
		secure:      c.Secure,
		port:        c.Port,
		protocol:    c.Network,
		path_prefix: c.PathPrefix,
	}
}
