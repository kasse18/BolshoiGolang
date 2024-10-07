package storage

import "go.uber.org/zap"

type Storage struct {
	inner  map[string]interface{}
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
		inner:  make(map[string]interface{}),
		logger: logger,
	}, nil
}

func (s Storage) Set(key string, value any) {
	s.inner[key] = value

	s.logger.Info("new key value pair set")
	s.logger.Sync()
}

func (s Storage) Get(key string) *any {
	res, ok := s.inner[key]
	if !ok {
		return nil
	}
	s.logger.Info("got value by key")
	s.logger.Sync()
	return &res
}

func (s Storage) GetKind(key string) string {
	var v interface{} = s.inner[key]

	defer s.logger.Sync()
	s.logger.Info("got kind by key")

	switch v.(type) {
	case int:
		return "D"
	case string:
		return "S"
	default:
		return "lol"
	}
}
