package users

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetAll() ([]User, error) {
	var users []User
	err := r.DB.Find(&users).Error
	return users, err
}

func (r *Repository) GetByID(id uint) (User, error) {
	var user User
	err := r.DB.First(&user, id).Error
	return user, err
}

func (r *Repository) Create(user *User) error {
	return r.DB.Create(user).Error
}

func (r *Repository) Update(user *User) error {
	return r.DB.Save(user).Error
}

func (r *Repository) Delete(id uint) error {
	return r.DB.Delete(&User{}, id).Error
}
