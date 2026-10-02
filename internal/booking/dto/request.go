package dto

type CreateBookingRequestDto struct {
	EventID  uint `json:"event_id" validate:"required"`
	Quantity int  `json:"quantity" validate:"required,gt=0"`
}
