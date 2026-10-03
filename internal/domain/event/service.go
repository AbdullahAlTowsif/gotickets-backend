package event

import "gotickets/internal/domain/event/dto"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateEvent(req dto.CreateEventRequestDto) (*dto.EventResponseDto, error) {
	event := Event{
		Title:            req.Title,
		Description:      req.Description,
		Location:         req.Location,
		StartsAt:         req.StartsAt,
		TotalTickets:     req.TotalTickets,
		AvailableTickets: req.TotalTickets,
		Price:            req.Price,
	}
	if err := s.repo.CreateEvent(&event); err != nil {
		return nil, err
	}
	return event.ToResponse(), nil
}

func (s *service) GetAllEvents() ([]dto.EventResponseDto, error) {
	events, err := s.repo.GetAllEvents()
	if err != nil {
		return nil, err
	}

	response := make([]dto.EventResponseDto, len(events))
	for i, event := range events {
		response[i] = *event.ToResponse()
	}

	return response, nil
}

func (s *service) GetEventByID(eventId uint) (*dto.EventResponseDto, error) {
	event, err := s.repo.GetEventByID(eventId)

	if err != nil {
		return nil, err
	}

	return event.ToResponse(), nil
}

func (s *service) UpdateEvent(eventId uint, req dto.UpdateEventRequestDto) (*dto.EventResponseDto, error) {
	event, err := s.repo.GetEventByID(eventId) // getting existing event by the id first
	if err != nil {
		return nil, err
	}

	if req.Title != "" {
		event.Title = req.Title
	}

	if req.Description != "" {
		event.Description = req.Description
	}

	if req.Location != "" {
		event.Location = req.Location
	}

	if !req.StartsAt.IsZero() {
		event.StartsAt = req.StartsAt
	}

	if req.Price != 0 {
		event.Price = req.Price
	}

	if err := s.repo.UpdateEvent(event); err != nil {
		return nil, err
	}

	return event.ToResponse(), nil
}
