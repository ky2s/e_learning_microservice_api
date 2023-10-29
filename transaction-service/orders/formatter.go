package orders

type OrderDataFormat struct {
	OrderID    int    `json:"order_id"`
	UserName   string `json:"user_name"`
	UserEmail  string `json:"user_email"`
	CourseName string `json:"course_name"`
	MentorName string `json:"mentor_name"`
}

func FormatOrder(order Orders) OrderDataFormat {
	formatter := OrderDataFormat{}
	formatter.OrderID = order.ID
	formatter.UserName = order.UserName
	formatter.UserEmail = order.UserEmail
	formatter.CourseName = order.CourseName
	formatter.MentorName = order.MentorName
	return formatter
}
