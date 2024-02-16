package database

import (
	"ablufus/internal/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database interface {
	AutoMigrate() error
	Post(u *entities.User) (*entities.User, error)
	List(ids []string, limit int, page int) ([]*entities.User, error)
	Update(id string, amount float64) error
}

type AblufusDatabase struct {
	db *gorm.DB
}

func New(dsn string) (Database, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &AblufusDatabase{db}, nil
}

func (orm *AblufusDatabase) AutoMigrate() error {
	return orm.db.AutoMigrate(&entities.User{})
}

func (orm *AblufusDatabase) Post(u *entities.User) (*entities.User, error) {
	findUser := &entities.User{ID: u.ID}
	result := orm.db.Create(&u)
	if result.Error != nil {
		result = orm.db.Find(&findUser)
		if result.Error != nil {
			return nil, result.Error
		}
		return findUser, nil
	}
	return u, nil
}

func (orm *AblufusDatabase) List(ids []string, limit int, page int) ([]*entities.User, error) {
	users := []*entities.User{}
	var result *gorm.DB
	if len(ids) > 0 {
		result = orm.db.Offset((page-1)*limit).Limit(limit).Where("id IN ?", ids).Find(&users)
	} else {
		result = orm.db.Offset((page - 1) * limit).Limit(limit).Find(&users)
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (orm *AblufusDatabase) Update(id string, amount float64) error {
	result := orm.db.Model(&entities.User{ID: id}).Update("amount", amount)
	return result.Error
}
