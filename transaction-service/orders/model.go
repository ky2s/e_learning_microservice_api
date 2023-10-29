package orders

import (
	"time"

	"gorm.io/datatypes"
)

type Orders struct {
	ID          int
	Status      string
	UserID      int
	UserName    string
	UserEmail   string
	CourseID    int
	CourseName  string
	CoursePrice int
	MentorID    int
	MentorName  string
	SnapUrl     string
	Amount      int
	Metadata    datatypes.JSON
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
