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
	InsertAttendance(ctx context.Context, r *request.CreateRequest) (*ent.Attendance, error) // User id is by input
	StoreAttendance(ctx context.Context, idUser string, typ int) (*ent.Attendance, error)    // User id claims by token
	GetAllAttendances(ctx context.Context) ([]*ent.Attendance, error)
	GetAttendancesByIDUser(ctx context.Context, idUser string) ([]*ent.Attendance, error)
	UpdateAttendance(ctx context.Context, r *request.UpdateAttendanceRequest, id string) (*ent.Attendance, error)
	DeleteAttendance(ctx context.Context, id string) error
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

func (s *attendeesService) DeleteAttendance(ctx context.Context, id string) error {
	return s.attendeesRepository.DeleteAttendance(ctx, id)
}

func (s *attendeesService) UpdateAttendance(ctx context.Context, r *request.UpdateAttendanceRequest, id string) (*ent.Attendance, error) {

	att := &ent.Attendance{
		ID:        id,
		IDUser:    r.IDUser,
		Type:      r.Type,
		UpdatedAt: time.Now(),
	}

	return s.attendeesRepository.UpdateAttendance(ctx, att)

}

func (s *attendeesService) GetAttendancesByIDUser(ctx context.Context, idUser string) ([]*ent.Attendance, error) {
	return s.attendeesRepository.GetAttendancesByIDUser(ctx, idUser)
}

func (s *attendeesService) GetAllAttendances(ctx context.Context) ([]*ent.Attendance, error) {
	return s.attendeesRepository.GetAllAttendances(ctx)
}

func (s *attendeesService) StoreAttendance(ctx context.Context, idUser string, typ int) (*ent.Attendance, error) {
	// Type 1 -> Checkin
	// Type 2 -> Checkout

	att := &ent.Attendance{
		ID:        utils.GenerateUUID(),
		IDUser:    idUser,
		Type:      typ,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.IsValidAttendees(ctx, att)

	if err != nil {
		return nil, &exception.BadRequestError{
			Message: err.Error(),
		}
	}

	return s.attendeesRepository.InsertAttendance(ctx, att)
}

func (s *attendeesService) IsValidAttendees(ctx context.Context, att *ent.Attendance) error {

	// Check if user going to entry checkin / checkout more than once
	// as per today
	var str string
	result, _ := s.attendeesRepository.GetByTypeAndDate(ctx, att.Type, att.CreatedAt, att.IDUser)

	if att.Type == 1 {
		str = "checkin"
	} else {
		str = "checkout"
	}
	if result != nil {
		return &exception.BadRequestError{
			Message: fmt.Sprintf("employee already %s today.", str),
		}
	}

	return nil
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

	err := s.IsValidAttendees(ctx, att)

	if err != nil {
		return nil, &exception.BadRequestError{
			Message: err.Error(),
		}
	}

	return s.attendeesRepository.InsertAttendance(ctx, att)
}
