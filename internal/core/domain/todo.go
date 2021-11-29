package domain

type Todo struct {
	ID          uint64 `json:"id" gorm:"primary_key;auto_increment" uri:"id" binding:"required"`
	Title       string `json:"title" gorm:"unique,not null"`
	Description string `json:"description" gorm:"not null"`
	Completed   bool   `json:"completed" gorm:"not null"`
}

func NewTodo(id uint64, title string, description string, completed bool) Todo {
	return Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}
