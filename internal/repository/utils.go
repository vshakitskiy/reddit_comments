package repository

import (
	"errors"
)

type Identifiable interface {
	GetID() string
}

func onIdOrder[T Identifiable](
	ids []string,
	data []T,
) ([]T, error) {
	mapData := make(map[string]T, len(ids))
	for _, value := range data {
		mapData[value.GetID()] = value
	}

	ordered := make([]T, len(ids))
	for i, id := range ids {
		if value, ok := mapData[id]; ok {
			ordered[i] = value
		} else {
			return nil, errors.New("unable to get data")
		}
	}

	return ordered, nil
}
