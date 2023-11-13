package usecase

import (
	"backend/internal/entity"
	"backend/internal/service"
)

type UseCase interface {
	GetPollById(id int) (*entity.Poll, error)
	GetOptionById(id int) (*entity.Option, error)
	CreatePoll(name string, pollID int, options []*entity.Option) error
	VoteForOption(optionID int) error
	DeletePoll(id int) error
}

type useCase struct {
	pollService   service.PollService
	optionService service.OptionService
}

var _ UseCase = (*useCase)(nil)

func NewUseCase(pollService service.PollService, optionService service.OptionService) *useCase {
	return &useCase{
		pollService:   pollService,
		optionService: optionService,
	}
}

func (uc *useCase) GetPollById(id int) (*entity.Poll, error) {
	poll, err := uc.pollService.Get(id)
	if err != nil {
		return nil, err
	}

	poll.Options, err = uc.optionService.GetOptionsByPollId(poll.ID)
	if err != nil {
		return nil, err
	}

	poll.TotalVotes, err = uc.optionService.GetTotalVotesByPollId(poll.ID)
	if err != nil {
		return nil, err
	}

	return poll, nil
}

func (uc *useCase) GetOptionById(id int) (*entity.Option, error) {
	return nil, nil
}

func (uc *useCase) CreatePoll(name string, pollID int, options []*entity.Option) error {
	return nil
}

func (uc *useCase) VoteForOption(optionID int) error {
	return nil
}

func (uc *useCase) DeletePoll(id int) error {
	return nil
}
