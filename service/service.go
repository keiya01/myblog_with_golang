package service

import (
	"github.com/keiya01/myblog/database"
	"github.com/pkg/errors"
)

type Service struct {
	*database.Handler
}

// NewService データベースとの連携を行う
func NewService(db *database.Handler) *Service {
	return &Service{db}
}

// FindAll DBからPostに関する全てのデータを取得する。
func (s *Service) FindAll(model interface{}, order string, where ...interface{}) error {
	if db := s.Order(order).Find(model, where...); db.Error != nil {
		return errors.Wrap(db.Error, "FindAll()")
	}

	return nil
}

func (s *Service) FindOne(model interface{}, where ...interface{}) error {
	if db := s.First(model, where...); db.Error != nil {
		return errors.Wrap(db.Error, "FindOne()")
	}

	return nil
}

func (s *Service) Save(model interface{}) error {
	if db := s.Create(model); db.Error != nil {
		return errors.Wrap(db.Error, "Save()")
	}

	return nil
}
