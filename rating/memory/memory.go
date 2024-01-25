package memory

import (
	"context"
	er "movieExample/metadata/internal/repository"
	"movieExample/rating/model"
)

type Repository struct {
	data map[model.RecordType]map[model.RecordId][]model.Rating
}

func New() *Repository {
	return &Repository{map[model.RecordType]map[model.RecordId][]model.Rating{}}
}

func (r *Repository) Get(_ context.Context, recordId model.RecordId, recordType model.RecordType) ([]model.Rating, error) {
	if _, ok := r.data[recordType]; !ok {
		return nil, er.ErrNotFound
	}
	if ratings, ok := r.data[recordType][recordId]; !ok || len(ratings) == 0 {
		return nil, er.ErrNotFound
	}

	return r.data[recordType][recordId], nil
}
