package repository

import _model "github.com/ibadsatria/ucantplayalone/game"

// IGameRepository interface of repo
type IGameRepository interface {
	AddGame(m *_model.Game) (bool, error)
}
