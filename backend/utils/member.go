package utils

import (
	"cashier-machine/model"
	repository "cashier-machine/repository/config"
	"cashier-machine/repository/modelfunc"
	"log"
)

// InsertMemberData inserts a new member into the database
func InsertMemberData(member model.Member) (model.Member, error) {
	newMember, err := modelfunc.InsertMember(repository.Mysql.DB, member)
	if err != nil {
		log.Println("Error inserting member data:", err)
		return model.Member{}, err
	}
	return newMember, nil
}

// UpdateMemberData updates an existing member's details
func UpdateMemberData(member model.Member) (model.Member, error) {
	updatedMember, err := modelfunc.UpdateMember(repository.Mysql.DB, member)
	if err != nil {
		log.Println("Error updating member data:", err)
		return model.Member{}, err
	}
	return updatedMember, nil
}

// SoftDeleteMemberData performs a soft delete on a member (sets deleted_at field)
func SoftDeleteMemberData(memberID string) error {
	err := modelfunc.SoftDeleteMember(repository.Mysql.DB, memberID)
	if err != nil {
		log.Println("Error soft deleting member:", err)
		return err
	}
	return nil
}
