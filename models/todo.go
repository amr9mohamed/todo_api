package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Description string    `json:"description" gorm:"not null"`
	Completed   bool      `json:"completed" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

//func Paginate(pageNumber int, pageSize int) func(db *gorm.DB) *gorm.DB {
//	return func(db *gorm.DB) *gorm.DB {
//		page := pageNumber
//		if page == 0 {
//			page = 1
//		}
//
//		pageSize := pageSize
//		switch {
//		case pageSize > 100:
//			pageSize = 100
//		case pageSize <= 0:
//			pageSize = 10
//		}
//
//		offset := (page - 1) * pageSize
//		return db.Offset(int(offset)).Limit(int(pageSize))
//	}
//}

func (t *Todo) GetTodos(db *gorm.DB) (*[]Todo, error) {
	var todos []Todo
	err := db.Debug().Model(&Todo{}).Limit(250).Find(&todos).Error
	if err != nil {
		return &[]Todo{}, err
	}
	return &todos, err
}

func (t *Todo) PostTodo(db *gorm.DB, newTodo Todo) error {
	result := db.Debug().Model(&Todo{}).Create(&newTodo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *Todo) DeleteTodo(db *gorm.DB, id string) error {
	if err := db.Debug().Delete(&Todo{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (t *Todo) GetTodo(db *gorm.DB, id string) (Todo, error) {
	var todo Todo
	if err := db.Debug().Model(&Todo{}).First(&todo, id).Error; err != nil {
		return Todo{}, err
	}
	return todo, nil
}

func (t *Todo) EditTodo(db *gorm.DB, id string, editedTodo Todo) error {
	var todo Todo
	if err := db.Model(&Todo{}).First(&todo, id).Error; err != nil {
		return errors.New("record not found to be edited")
	}
	if err := db.Debug().Model(&todo).Updates(editedTodo).Error; err != nil {
		return err
	}
	return nil
}
