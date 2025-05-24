package repositories

import (
	"errors"
	"log"

	"github.com/sergeyiksanov/help-on-road/internal/dto"
	"github.com/sergeyiksanov/help-on-road/internal/models"
	"github.com/sergeyiksanov/help-on-road/internal/services"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(tx services.TransactionContext, model *models.User) (int64, error) {
	log.Print(model)
	userDto := dto.User{
		PhoneNumber:   model.PhoneNumber,
		Password:      model.Password,
		FirstName:     model.FirstName,
		LastName:      model.LastName,
		Surname:       model.Surname,
		AutoModel:     model.AutoModel,
		AutoGosNumber: model.AutoGosNumber,
		VinCode:       model.VinCode,
		IsValid:       model.IsValid,
		IsModerate:    model.IsModerate,
	}

	ctx := ur.db
	if tx != nil {
		var ok bool
		ctx, ok = tx.(*gorm.DB)
		if !ok {
			return -1, errors.New("Invalid ctx")
		}
	}

	if err := ctx.Create(&userDto).Error; err != nil {
		return -1, err
	}

	model.Id = userDto.ID
	return userDto.ID, nil
}

func (ur *UserRepository) GetById(id int64) (*models.User, error) {
	var userDto dto.User
	if err := ur.db.Where("id = ?", id).First(&userDto).Error; err != nil {
		return nil, err
	}
	log.Print(userDto)

	return &models.User{
		Id:            userDto.ID,
		PhoneNumber:   userDto.PhoneNumber,
		Password:      userDto.Password,
		FirstName:     userDto.FirstName,
		LastName:      userDto.LastName,
		Surname:       userDto.Surname,
		AutoModel:     userDto.AutoModel,
		AutoGosNumber: userDto.AutoGosNumber,
		VinCode:       userDto.VinCode,
		IsValid:       userDto.IsValid,
		IsModerate:    userDto.IsModerate,
	}, nil
}

func (ur *UserRepository) Update(tx services.TransactionContext, model *models.User) error {
	userDto := dto.User{
		FirstName:     model.FirstName,
		LastName:      model.LastName,
		Surname:       model.Surname,
		AutoModel:     model.AutoModel,
		AutoGosNumber: model.AutoGosNumber,
		VinCode:       model.VinCode,
		IsValid:       model.IsValid,
		IsModerate:    model.IsModerate,
	}

	ctx := ur.db
	if tx != nil {
		var ok bool
		ctx, ok = tx.(*gorm.DB)
		if !ok {
			return errors.New("Invalid ctx")
		}
	}

	return ctx.Save(&userDto).Error
}

func (ur *UserRepository) Delete(tx services.TransactionContext, id int64) error {
	ctx := ur.db
	if tx != nil {
		var ok bool
		ctx, ok = tx.(*gorm.DB)
		if !ok {
			return errors.New("Invalid ctx")
		}
	}

	return ctx.Delete(&dto.User{}, id).Error
}

func (ur *UserRepository) GetCountByNumber(number string) (int64, error) {
	var count int64
	err := ur.db.Model(&dto.User{}).Where("phone_number = ?", number).Count(&count).Error
	return count, err
}

func (ur *UserRepository) GetByPhoneNumber(phoneNumber string) (*models.User, error) {
	var userDto dto.User
	if err := ur.db.Where("phone_number = ?", phoneNumber).First(&userDto).Error; err != nil {
		return nil, err
	}

	return &models.User{
		Id:            userDto.ID,
		PhoneNumber:   userDto.PhoneNumber,
		Password:      userDto.Password,
		FirstName:     userDto.FirstName,
		LastName:      userDto.LastName,
		Surname:       userDto.Surname,
		AutoModel:     userDto.AutoModel,
		AutoGosNumber: userDto.AutoGosNumber,
		VinCode:       userDto.VinCode,
		IsValid:       userDto.IsValid,
		IsModerate:    userDto.IsModerate,
	}, nil
}

func (ur *UserRepository) GetNotValidUsers() ([]*models.User, error) {
	var users []dto.User
	if err := ur.db.Where("is_moderate = ?", false).Find(&users).Error; err != nil {
		return nil, err
	}

	userModels := make([]*models.User, len(users))
	for i, user := range users {
		userModels[i] = &models.User{
			Id:            user.ID,
			PhoneNumber:   user.PhoneNumber,
			Password:      user.Password,
			FirstName:     user.FirstName,
			LastName:      user.LastName,
			Surname:       user.Surname,
			AutoModel:     user.AutoModel,
			AutoGosNumber: user.AutoGosNumber,
			VinCode:       user.VinCode,
			IsValid:       user.IsValid,
			IsModerate:    user.IsModerate,
		}
	}

	return userModels, nil
}
