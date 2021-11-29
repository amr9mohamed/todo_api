package todorepo

import (
	"errors"
	"fmt"
	"log"

	"github.com/amr9mohamed/todoAPI/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Psql struct {
	*gorm.DB
}

func NewPsqlDB(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *Psql {
	DBURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName, DbPort)
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database")
	}
	log.Println("Connected to the database")
	err = db.Debug().AutoMigrate(&domain.Todo{})
	if err != nil {
		log.Fatal("can't do the migration")
	}

	return &Psql{
		db,
	}
}

func (db *Psql) Get(id uint64) (domain.Todo, error) {
	var todo domain.Todo
	if err := db.Debug().Model(&domain.Todo{}).First(&todo, id).Error; err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}

func (db *Psql) List() ([]domain.Todo, error) {
	var todos []domain.Todo
	err := db.Debug().Model(&domain.Todo{}).Limit(250).Find(&todos).Error
	if err != nil {
		return []domain.Todo{}, err
	}
	return todos, nil
}

func (db *Psql) Delete(id uint64) error {
	if err := db.Debug().Delete(&domain.Todo{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (db *Psql) Add(newTodo domain.Todo) error {
	result := db.Debug().Model(&domain.Todo{}).Create(&newTodo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *Psql) Edit(id uint64, editedTodo domain.Todo) error {
	var todo domain.Todo
	if err := db.Model(&domain.Todo{}).First(&todo, id).Error; err != nil {
		return errors.New("record not found to be edited")
	}
	if err := db.Debug().Model(&todo).Updates(editedTodo).Error; err != nil {
		return err
	}
	return nil
}
