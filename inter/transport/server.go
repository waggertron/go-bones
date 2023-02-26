package inter

import "context"

type Server interface {
	Start() (context.CancelFunc, error)
	Stop() error
}
