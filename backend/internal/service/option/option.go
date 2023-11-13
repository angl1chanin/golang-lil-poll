package option

import (
	"backend/internal/entity"
	"backend/internal/repository"
	"backend/internal/service"
)

type optionService struct {
	r service.OptionService
}

var _ service.OptionService = (*optionService)(nil)

func NewOptionService(r repository.OptionRepository) *optionService {
	return &optionService{
		r: r,
	}
}

func (s *optionService) GetById(optionID int) (*entity.Option, error) {
	return nil, nil
}

func (s *optionService) Create(pollID int, options []entity.Option) error {
	return nil
}

func (s *optionService) Vote(optionID int) error {
	return s.r.Vote(optionID)
}

func (s *optionService) GetOptionsByPollId(pollId uint) ([]*entity.Option, error) {
	return s.r.GetOptionsByPollId(pollId)
}

func (s *optionService) GetTotalVotesByPollId(pollId uint) (uint, error) {
	return s.r.GetTotalVotesByPollId(pollId)
}
