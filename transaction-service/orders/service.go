package orders

import (
	"errors"
	"transaction-service/helper"
	"transaction-service/payments"
)

type Service interface {
	Create(input CreateOrderInput) (Orders, error)
	// Index(input CreateOrderInput) (Orders, error)
	// Show(input CreateOrderInput) (Orders, error)
}

type service struct {
	repository     Repository
	paymentService payments.Service
}

func NewService(repository Repository) *service {
	return &service{
		repository: repository,
	}
}

func (s *service) Create(input CreateOrderInput) (Orders, error) {
	order := Orders{}

	//get user
	user, err := helper.GetUserByID(input.UserID)
	if err != nil {
		return Orders{}, err
	}

	if user.Data.ID <= 0 {
		return Orders{}, errors.New("user data is empty")
	}

	//get user
	course, err := helper.GetCourseByID(input.CourseID)
	if err != nil {
		return Orders{}, err
	}
	if course.Data.ID <= 0 {
		return Orders{}, errors.New("course data is empty")
	}

	//get mentor
	mentor, err := helper.GetMentorByID(course.Data.MentorID)
	if err != nil {
		return Orders{}, err
	}

	order.Status = input.Status
	order.UserID = user.Data.ID
	order.UserName = user.Data.Name
	order.UserEmail = user.Data.Email
	order.CourseID = course.Data.ID
	order.CourseName = course.Data.Name
	order.CoursePrice = course.Data.Price
	order.MentorID = mentor.Data.ID
	order.MentorName = mentor.Data.Name
	order.Amount = input.Amount

	newOrder, err := s.repository.Create(order)
	if err != nil {
		return newOrder, err
	}

	// PAYMENT HERE

	return newOrder, nil

}

func (s *service) _()
