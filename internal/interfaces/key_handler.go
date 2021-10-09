package interfaces

import "stonks/internal/config"

type INewsKHandler interface {
	GetCfg() (config.Config, error)
}
