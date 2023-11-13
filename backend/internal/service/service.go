package service

import "backend/internal/entity"

type PollService interface {
	Get(id int) (*entity.Poll, error)
	Create(name string) error
}

type OptionService interface {
	GetById(optionID int) (*entity.Option, error)
	Create(pollID int, options []entity.Option) error
	Vote(optionID int) error
	GetOptionsByPollId(pollId uint) ([]*entity.Option, error)
	GetTotalVotesByPollId(pollId uint) (uint, error)
}
