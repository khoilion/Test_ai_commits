package models

type Todo struct {
	ID        string `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

type CreateTodoInput struct {
	Task string `json:"task" binding:"required"` // binding required nếu gửi yêu cầu trống thì báo lỗi
}
