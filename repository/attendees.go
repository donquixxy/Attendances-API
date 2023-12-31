package repository

import (
	"context"
	"dg-test/ent"
	"dg-test/ent/attendance"
	"dg-test/exception"
	"log"
	"time"
)

type AttendeesRepository interface {
	InsertAttendance(ctx context.Context, v *ent.Attendance) (*ent.Attendance, error)
	GetByTypeAndDate(ctx context.Context, t int, date time.Time, idUser string) (*ent.Attendance, error)
	GetAllAttendances(ctx context.Context) ([]*ent.Attendance, error)
	GetAttendancesByIDUser(ctx context.Context, idUser string) ([]*ent.Attendance, error)
	UpdateAttendance(ctx context.Context, v *ent.Attendance) (*ent.Attendance, error)
	DeleteAttendance(ctx context.Context, id string) error
}

type attendeesRepository struct {
	c *ent.Client
}

func NewAttendeesRepository(
	c *ent.Client,
) AttendeesRepository {
	return &attendeesRepository{
		c: c,
	}
}

func (s *attendeesRepository) DeleteAttendance(ctx context.Context, id string) error {
	return s.c.Attendance.DeleteOneID(id).Exec(ctx)
}

func (s *attendeesRepository) UpdateAttendance(ctx context.Context, v *ent.Attendance) (*ent.Attendance, error) {
	res := s.c.Attendance.UpdateOneID(v.ID)

	if v.Type != 0 {
		res.SetType(v.Type)
	}

	if v.IDUser != "" {
		res.SetIDUser(v.IDUser)
	}

	result, err := res.SetUpdatedAt(time.Now()).Save(ctx)

	return result, err
}

func (s *attendeesRepository) GetAttendancesByIDUser(ctx context.Context, idUser string) ([]*ent.Attendance, error) {
	result, _ := s.c.Attendance.Query().Where(attendance.IDUser(idUser)).All(ctx)

	if len(result) == 0 {
		return nil, &exception.RecordNotFoundError{
			Message: "not found",
		}
	}

	return result, nil
}

func (s *attendeesRepository) GetAllAttendances(ctx context.Context) ([]*ent.Attendance, error) {
	result, _ := s.c.Attendance.Query().All(ctx)

	if len(result) == 0 {
		return nil, &exception.RecordNotFoundError{
			Message: "attendances not found",
		}
	}

	return result, nil
}

func (s *attendeesRepository) InsertAttendance(ctx context.Context, v *ent.Attendance) (*ent.Attendance, error) {
	result, err := s.c.Attendance.Create().SetID(v.ID).
		SetIDUser(v.IDUser).SetType(v.Type).SetCreatedAt(v.CreatedAt).
		SetUpdatedAt(v.UpdatedAt).Save(ctx)

	if err != nil {
		log.Printf("Failed create attendee %v", err)
		return nil, err
	}

	return result, nil
}

func (s *attendeesRepository) GetByTypeAndDate(ctx context.Context, t int, date time.Time, iduser string) (*ent.Attendance, error) {
	result, err := s.c.Attendance.Query().
		Where(attendance.IDUser(iduser)).
		Where(attendance.Type(t)).
		Where(attendance.
			CreatedAtLTE(date)).
		Only(ctx)

	if err != nil {
		return nil, &exception.RecordNotFoundError{
			Message: err.Error(),
		}
	}

	return result, nil
}
