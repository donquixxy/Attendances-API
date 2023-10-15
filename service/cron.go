package service

import (
	"context"
	"dg-test/ent"
	"dg-test/repository"
	"log"
	"sync"
)

type CronJobService interface {
	SendCheckinReminder(ctx context.Context) error
	SendCheckoutReminder(ctx context.Context) error
}

type cronJobService struct {
	mailRepository repository.MailRepostory
	userRepository repository.UserRepository
}

func NewCronJOBService(
	mailRepository repository.MailRepostory,
	userRepository repository.UserRepository,
) CronJobService {
	return &cronJobService{
		userRepository: userRepository,
		mailRepository: mailRepository,
	}
}

func (s *cronJobService) SendCheckinReminder(ctx context.Context) error {
	// Get all list user email
	users, err := s.userRepository.GetAllUser(ctx)

	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, i := range users {
		wg.Add(1)

		go func(u *ent.User) {
			defer wg.Done()

			err := s.mailRepository.SendEmailCheckin(u.Email)

			if err != nil {
				log.Printf("failed to send email to %s with error %v", u.Email, err)
			}
		}(i)
	}
	wg.Wait()

	return nil
}

func (s *cronJobService) SendCheckoutReminder(ctx context.Context) error {
	// Get all list user email
	users, err := s.userRepository.GetAllUser(ctx)

	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for _, i := range users {
		wg.Add(1)

		go func(u *ent.User) {
			defer wg.Done()

			err := s.mailRepository.SendEmailCheckout(u.Email)

			if err != nil {
				log.Printf("failed to send checkout email to %s with error %v", u.Email, err)
			}
		}(i)
	}
	wg.Wait()

	return nil
}

// func (s *cronJobService) doTheJob(users []*ent.user) {
// 	var wg *sync.WaitGroup
// 	for _, i := range users {
// 		wg.Add(1)

// 		go func() {
// 			err := s.mailRepository.SendEmailCheckin(i.Email)

// 			if err != nil {
// 				log.Printf("failed send email to %s", i.email)
// 			}

// 		}(users)
// 	}
// }
