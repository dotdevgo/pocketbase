package pocketbase

import (
	"github.com/defval/di"
	"github.com/pocketbase/pocketbase/core"
)

type (
	Controller interface {
		New(ctx *core.ServeEvent)
	}
)

// NewController godoc
func NewController(provideFn di.Constructor) di.Option {
	return di.Provide(provideFn, di.As(new(Controller)))
}

// type controllerExtension struct {
// 	Extension
// }

// func (controllerExtension) Boot(pb *PocketBase) error {
// 	pb.Logger().Info("PocketBase: controllers")

// 	return nil
// }
