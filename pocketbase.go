package pocketbase

import (
	"log"

	"github.com/defval/di"
	pb "github.com/pocketbase/pocketbase"
)

type PocketBase struct {
	*pb.PocketBase
	*di.Container
}

// New Create pocketbase instance
func New(providers ...di.Option) *PocketBase {
	container, err := di.New(providers...)
	if err != nil {
		log.Fatal(err)
	}

	wrapper := pb.New()

	return &PocketBase{
		PocketBase: wrapper,
		Container:  container,
	}
}
