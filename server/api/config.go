package api

import "context"

// Conf Configuration of api handler
type Conf struct {
	ctx     context.Context
	eventCb func(result string)
}

// NewConf ...
func NewConf() *Conf {
	return &Conf{ctx: context.Background()}
}

// SetContext context setter
func (c *Conf) SetContext(ctx context.Context) {
	c.ctx = ctx
}

// SetEventCallback event callback setter
func (c *Conf) SetEventCallback(eventCb func(result string)) {
	c.eventCb = eventCb
}
