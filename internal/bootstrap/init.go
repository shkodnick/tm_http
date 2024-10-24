package bootstrap

import (
	"github.com/pkg/errors"

	"app/internal/config"
	"app/internal/core/profile"
	"app/internal/store"
)

func CreateUserService() (*profile.ProfileService, error) {
	cfg, err := config.NewConfig()
	if err != nil {
		return nil, errors.Wrap(err, "CreateConfig")
	}

	dbStore, err := store.New(cfg.Database)
	if err != nil {
		return nil, err
	}

	return profile.New(dbStore)
}
