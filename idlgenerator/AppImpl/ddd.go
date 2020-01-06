package AppImpl

import "log"

type ILogFactory interface {
	Create() *log.Logger
}

