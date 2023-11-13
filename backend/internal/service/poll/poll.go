package poll

import (
	"backend/internal/entity"
	"backend/internal/repository"
	"backend/internal/service"
)

type pollService struct {
	r repository.PollRepository
}

var _ service.PollService = (*pollService)(nil)

func NewPollService(r repository.PollRepository) *pollService {
	return &pollService{
		r: r,
	}
}

func (s *pollService) Get(id int) (*entity.Poll, error) {
	return s.r.Get(id)
}

func (s *pollService) Create(name string) error {
	return nil
}
