package server

import (
	"time"

	"github.com/NathanBak/go-server/pkg/widget"
)

const (
	defaultPort         = 8080
	defaultReadTimeout  = 3 * time.Second
	defaultWriteTimeout = 3 * time.Second
)

// Config contains information necessary to set up a Server.
type Config struct {
	Port         int           `json:"port" envvar:"PORT"`
	ReadTimeout  time.Duration `json:"readTimeout" envvar:"READ_TIMEOUT"`
	WriteTimeout time.Duration `json:"writeTimeout" envvar:"WRITE_TIMEOUT"`

	Logger Logger `json:"-" envvar:"-"`

	IncludeStatusCodeInMessages bool `json:"-" envvar:"-"`

	Storage Storage `json:"-" envvar:"-"`
}

// The Logger interface defines the methods required by the Server for logging.
type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warning(msg string)
	Error(msg string)
}

// The Storage interface defines the methods required to access backing Widget storage.
type Storage interface {
	Get(string) (widget.Widget, bool, error)
	Set(string, widget.Widget) error
	Delete(string) (widget.Widget, bool, error)
	Keys() ([]string, error)
}

// Init implements the cfgbuild.Config interface and should only be called by a cfgbuild.Builder.
func (c *Config) Init() error {
	if c.Port == 0 {
		c.Port = defaultPort
	}

	if c.ReadTimeout == 0 {
		c.ReadTimeout = defaultReadTimeout
	}

	if c.WriteTimeout == 0 {
		c.WriteTimeout = defaultWriteTimeout
	}

	if c.Logger == nil {
		c.Logger = defaultLogger{}
	}

	return nil
}
