package repository

import (
	"context"
	"dg-test/ent"
	"dg-test/ent/attendance"
	"log"
	"time"
)

type AttendeesRepository interface {
	InsertAttendance(ctx context.Context, v *ent.Attendance) (*ent.Attendance, error)
	GetByTypeAndDate(ctx context.Context, t int, date time.Time) (*ent.Attendance, error)
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

func (s *attendeesRepository) GetByTypeAndDate(ctx context.Context, t int, date time.Time) (*ent.Attendance, error) {
	return s.c.Attendance.Query().Where(attendance.Type(t)).Where(attendance.CreatedAtLTE(date)).
		Only(ctx)
}
