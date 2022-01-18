package singleton

import (
	"context"
	"sync"

	"peeral.com/proxy-libp2p/domain/interfaces"
)

var (
	once     sync.Once
	instance *Configuration
)

//Configuration ...
type Configuration struct {
	ctx     context.Context
	eventCb func(result string)
}

// NewConfiguration ...
func NewConfiguration(ctx context.Context, eventCb func(result string)) *Configuration {
	once.Do(func() {
		instance = &Configuration{ctx: ctx, eventCb: eventCb}
	})

	return instance
}

func SetContext(ctx context.Context) {
	instance.ctx = ctx
}

func GetContext() context.Context {
	return instance.ctx
}

func LogEvent(messageToLog interfaces.Serializable) {
	instance.eventCb(messageToLog.Serialize())
}
