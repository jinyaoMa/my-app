package cflog

import (
	"errors"
	"log"
)

func New(options Options) (*Cflog, error) {
	return new(Cflog).init(options)
}

type Cflog struct {
	*log.Logger
}

func (c *Cflog) init(options Options) (*Cflog, error) {
	writer, err := NewWriter(options.LogFile, options.EnableConsole)
	if err != nil {
		return nil, errors.Join(errors.New("failed to init cflog"), err)
	}

	c.Logger = log.New(writer, options.LogPrefix, log.LstdFlags)
	return c, nil
}
