package service

import (
	"EasyLinks/server/pkg/errors"
	"EasyLinks/server/storage"
	"EasyLinks/server/utils"
	"log"
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
	ShortURL := &ShortURL{
		ID:       utils.GenerateString(url),
		URL:      url,
		ExpireAt: getExpirationTime(1),
	}
	_, err := s.Collection.InsertOne((*storage.ShortURL)(ShortURL))
	switch {
	case err == 0:
		log.Println(ShortURL)
		return ShortURL, nil
	case err == 6: // If the key already exists
		return nil, errors.ErrorMap[6]
	default:
		return nil, errors.ErrorMap[2]
	}
}

func getExpirationTime(ttlDays int) *time.Time {
	if ttlDays <= 0 {
		return nil
	}
	t := time.Now().Add(time.Hour * 24 * time.Duration(ttlDays))
	return &t
}

func (s *Service) GetURL(id string) (*ShortURL, error) {

	url, err := s.Collection.FindOne(id)
	switch {
	case err == 0:
		return (*ShortURL)(url), nil
	case err == 5:
		return nil, errors.ErrorMap[5]
	default:
		return nil, errors.ErrorMap[2]
	}
}
