package orders

import (
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

	//get user
	course, err := helper.GetCourseByID(input.CourseID)
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

	newOrder, err := s.repository.Create(order)
	if err != nil {
		return newOrder, err
	}

	// paymentOrder := payments.Order{
	// 	ID:     newOrder.ID,
	// 	Amount: input.Amount,
	// }

	// paymentUrl, err := s.paymentService.GetPaymentUrl(paymentOrder, payments.User{Email: user.Data.Email, Name: user.Data.Name})

	// newOrder.SnapUrl = paymentUrl

	// newOrder, err = s.repository.Update(newOrder)

	// if err != nil {
	// 	return newOrder, err
	// }

	return newOrder, nil

}

func (s *service) _()
