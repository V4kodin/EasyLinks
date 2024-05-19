package service

import (
	"EasyLinks/server/storage"
	"EasyLinks/server/utils"
	"time"
)

type ShortURL struct {
	ID       string     `bson:"_id"`
	URL      string     `bson:"url"`
	ExpireAt *time.Time `bson:"expireAt,omitempty"`
}

type Service struct {
	Collection *storage.Coll
}

func NewService(coll *storage.Coll) *Service {
	return &Service{
		Collection: coll,
	}
}

func (s *Service) AddURL(url string) (*ShortURL, error) {
	shortURL := &ShortURL{
		ID:       utils.GenerateString(url),
		URL:      url,
		ExpireAt: getExpirationTime(1),
	}
	s.Collection.InsertOne((*storage.ShortURL)(shortURL))
	return shortURL, nil
}

func getExpirationTime(ttlDays int) *time.Time {
	if ttlDays <= 0 {
		return nil
	}
	t := time.Now().Add(time.Hour * 24 * time.Duration(ttlDays))
	return &t
}
