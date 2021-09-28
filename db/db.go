package db

import "github.com/ammorteza/location_history_server/entity"

type Storage struct {
	st map[string][]entity.Location
}

var storage *Storage

func New() *Storage {
	if storage == nil {
		storage = &Storage{
			st: make(map[string][]entity.Location),
		}
	}

	return storage
}

func (s *Storage) Insert(orderID string, l entity.Location) {
	/* 	_, ok := s.st[orderID]
	   	if ok { */
	s.st[orderID] = append(s.st[orderID], l)
	/* 	}

	   	s.st */

}

func (s *Storage) Fetch(orderID string, max int) []entity.Location {
	temp := make([]entity.Location, 0)
	if max == 0 {
		temp = s.st[orderID]
		return temp
	}

	count := 0

	for _, history := range s.st[orderID] {
		temp = append(temp, history)
		count++
		if count == max {
			break
		}
	}

	return temp
}

func (s *Storage) Delete(orderID string) {
	_, ok := s.st[orderID]
	if ok {
		s.st[orderID] = make([]entity.Location, 0)
	}
}
