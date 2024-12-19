package pocketbase

import (
	"errors"
	"log"

	"github.com/defval/di"
	pb "github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type (
	// PocketBase godoc
	PocketBase struct {
		*pb.PocketBase
		*di.Container
	}
)

// New Create pocketbase instance
func New(providers ...di.Option) *PocketBase {
	// providers = append(providers, NewExtension(func() *controllerExtension {
	// 	return &controllerExtension{}
	// }))

	container, err := di.New(providers...)
	if err != nil {
		log.Fatal(err)
	}

	app := &PocketBase{
		Container:  container,
		PocketBase: pb.New(),
	}

	app.Provide(func() *PocketBase {
		return app
	})

	return app
}

func (app *PocketBase) Start() error {
	if err := app.Invoke(app.configureExtensions); err != nil {
		if !errors.Is(err, di.ErrTypeNotExists) {
			log.Fatal(err)
		}
	}

	app.OnServe().BindFunc(app.configureRouter)

	return app.PocketBase.Start()
}

// configureExtensions godoc
func (app *PocketBase) configureExtensions(providers []Extension) error {
	app.Logger().Info("PocketBase: boot")

	for _, p := range providers {
		if err := p.Boot(app); err != nil {
			return err
		}
	}

	return nil
}

// configureRouter godoc
func (app *PocketBase) configureRouter(ctx *core.ServeEvent) error {
	var controllers []AbstractController

	if err := app.Resolve(&controllers); err != nil {
		if !errors.Is(err, di.ErrTypeNotExists) {
			return err
		}
	}

	for _, controller := range controllers {
		controller.New(ctx)
	}

	return ctx.Next()

}
