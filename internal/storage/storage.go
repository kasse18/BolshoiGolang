package storage

import "go.uber.org/zap"

type Storage struct {
	inner  map[string]string
	logger *zap.Logger
}

func (s Storage) NewStorage() (Storage, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return Storage{}, err
	}

	defer logger.Sync()
	logger.Info("new storage created")

	return Storage{
		inner:  make(map[string]string),
		logger: logger,
	}, nil
}

func (s Storage) Set(key, value string) {
	s.inner[key] = value

	s.logger.Info("new key value pair set")
	s.logger.Sync()
}

func (s Storage) Get(key string) *string {
	res, ok := s.inner[key]
	if !ok {
		return nil
	}
	s.logger.Info("got value by key")
	s.logger.Sync()
	return &res
}
