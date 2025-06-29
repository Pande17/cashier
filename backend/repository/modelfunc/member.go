package modelfunc

import (
	"cashier-machine/model"

	"gorm.io/gorm"
)

// InsertMember inserts a new Member into the database
func InsertMember(db *gorm.DB, member model.Member) (model.Member, error) {
	err := db.Create(&member).Error
	if err != nil {
		return model.Member{}, err
	}
	return member, nil
}

// UpdateMember updates the details of an existing member
func UpdateMember(db *gorm.DB, member model.Member) (model.Member, error) {
	err := db.Model(&model.Member{}).Where("id = ?", member.ID).Updates(member).Error
	if err != nil {
		return model.Member{}, err
	}
	return member, nil
}

// SoftDeleteMember sets the deleted_at field to the current time (soft delete)
func SoftDeleteMember(db *gorm.DB, memberID string) error {
	err := db.Model(&model.Member{}).Where("id = ?", memberID).Update("deleted_at", gorm.Expr("NOW()")).Error
	if err != nil {
		return err
	}
	return nil
}

// GetMemberByID retrieves a member by their ID
func GetMemberByID(db *gorm.DB, memberID string) (model.Member, error) {
	var member model.Member
	err := db.First(&member, "id = ?", memberID).Error
	if err != nil {
		return model.Member{}, err
	}
	return member, nil
}

// GetAllMembers retrieves all members from the database
func GetAllMembers(db *gorm.DB) ([]model.Member, error) {
	var members []model.Member
	err := db.Find(&members).Error
	if err != nil {
		return nil, err
	}
	return members, nil
}
