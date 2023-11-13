package repository

import "backend/internal/entity"

type PollRepository interface {
	Get(id int) (*entity.Poll, error)
	Create(name string) error
}

type OptionRepository interface {
	GetById(optionId int) (*entity.Option, error)
	Create(pollId int, options []entity.Option) error
	Vote(optionId int) error
	GetOptionsByPollId(pollId uint) ([]*entity.Option, error)
	GetTotalVotesByPollId(pollId uint) (uint, error)
}
