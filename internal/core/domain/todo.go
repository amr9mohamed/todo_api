package domain

type Todo struct {
	ID          uint64 `json:"id" gorm:"primary_key;auto_increment"`
	Description string `json:"description" gorm:"not null"`
	Completed   bool   `json:"completed" gorm:"not null"`
}

func NewTodo(id uint64, description string, completed bool) Todo {
	return Todo{
		ID:          id,
		Description: description,
		Completed:   completed,
	}
}
