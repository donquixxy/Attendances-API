package service

import (
	"context"
	"dg-test/domain/request"
	"dg-test/ent"
	"dg-test/exception"
	"dg-test/repository"
	"dg-test/utils"
	"fmt"
	"time"
)

type AttendeesService interface {
	InsertAttendance(ctx context.Context, r *request.CreateRequest) (*ent.Attendance, error)
}

type attendeesService struct {
	attendeesRepository repository.AttendeesRepository
}

func NewAttendeesService(
	attendeesRepository repository.AttendeesRepository,
) AttendeesService {
	return &attendeesService{
		attendeesRepository: attendeesRepository,
	}
}

func (s *attendeesService) InsertAttendance(ctx context.Context, r *request.CreateRequest) (*ent.Attendance, error) {

	// Type 1 -> Checkin
	// Type 2 -> Checkout

	att := &ent.Attendance{
		ID:        utils.GenerateUUID(),
		IDUser:    r.IDUser,
		Type:      r.Type,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Check if user going to entry checkin / checkout more than once
	// as per today
	var str string
	result, _ := s.attendeesRepository.GetByTypeAndDate(ctx, att.Type, att.CreatedAt)

	if att.Type == 1 {
		str = "checkin"
	} else {
		str = "checkout"
	}

	if result != nil {
		return nil, &exception.BadRequestError{
			Message: fmt.Sprintf("employee already %s today.", str),
		}
	}

	return s.attendeesRepository.InsertAttendance(ctx, att)
}
