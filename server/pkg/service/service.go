package service

import (
	"EasyLinks/server/pkg/errors"
	"EasyLinks/server/pkg/models"
	"EasyLinks/server/storage"
	"EasyLinks/server/utils"
	"log"
	"time"
)

type Service struct {
	Collection *storage.Coll
}

func NewService(coll *storage.Coll) *Service {
	return &Service{
		Collection: coll,
	}
}

func (s *Service) AddURL(url string) (*models.ShortURL, error) {
	ShortURL := &models.ShortURL{
		ID:       utils.GenerateString(url),
		URL:      url,
		ExpireAt: getExpirationTime(1),
	}
	s.Collection.InsertOne((*models.ShortURL)(ShortURL))
	switch {
	case ShortURL.Error == 0:
		ShortURL.Error = 0
		log.Println(ShortURL)
		return ShortURL, nil
	case ShortURL.Error == 6: // If the key already exists
		return ShortURL, errors.ErrorMap[6]
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

func (s *Service) GetURL(id string) (*models.ShortURL, error) {
	ShortURL := &models.ShortURL{
		ID: id,
	}
	s.Collection.FindOne((*models.ShortURL)(ShortURL))
	switch {
	case ShortURL.Error == 0:
		return ShortURL, nil
	case ShortURL.Error == 5:

		return ShortURL, errors.ErrorMap[5]
	default:
		return ShortURL, errors.ErrorMap[2]
	}
}
