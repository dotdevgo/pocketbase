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
	container, err := di.New(providers...)
	if err != nil {
		log.Fatal(err)
	}

	var config pb.Config = pb.Config{}
	if err := container.Resolve(&config); err != nil {
		if !errors.Is(err, di.ErrTypeNotExists) {
			log.Fatal(err)
		}
	}

	app := &PocketBase{
		Container:  container,
		PocketBase: pb.NewWithConfig(config),
	}

	app.Provide(func() *PocketBase {
		return app
	})

	return app
}

// NewWithConfig godoc
func NewWithConfig(config pb.Config, providers ...di.Option) *PocketBase {
	providers = append(providers, di.Provide(func() pb.Config {
		return config
	}))

	return New(providers...)
}

// Start godoc
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
	app.Logger().Info("PocketBase: router")

	var controllers []Controller

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

// providers = append(providers, NewExtension(func() *controllerExtension {
// 	return &controllerExtension{}
// }))
