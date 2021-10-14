package interfaces

import (
	"stonks/internal/config"
)

type INewsHandler interface {
	GetConfig() *config.Config
}
