package zoggly

import (
	"io"

	"github.com/segmentio/go-loggly"

	"github.com/uber-go/zap"
)

type ZapLoggly struct {
	client *loggly.Client
	io.Writer
}

func NewZoggly(token string, tags ...string) *ZapLoggly {
	client := loggly.New(token, tags...)
	client.Defaults = loggly.Message{}

	z := &ZapLoggly{
		client: client,
	}

	return z
}

func (z *ZapLoggly) GetWriter() zap.WriteSyncer {
	return zap.AddSync(z)
}

func (z *ZapLoggly) Write(p []byte) (n int, err error) {
	return z.client.Write(p)
}

func (z *ZapLoggly) Flush() {
	z.client.Flush()
}
