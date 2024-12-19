package pocketbase

import (
	"github.com/defval/di"
)

type (
	// Extension godoc
	Extension interface {
		Boot(pb *PocketBase) error
	}
)

// NewExtension godoc
func NewExtension(provideFn di.Constructor) di.Option {
	return di.Provide(provideFn, di.As(new(Extension)))
}
