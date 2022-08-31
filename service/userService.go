package service

import (
	"golang_crud/model"
	"golang_crud/repository"
)

// InsertUser insert a user to database
func InsertUser(user *model.User) *model.User {
	return repository.InsertUser(user)
}

// DeleteUserByUserIds delete users by user ids
func DeleteUserByUserIds(userIds *[]uint64) int64 {
	return repository.DeleteUserByUserIds(userIds)
}

// SearchUserByUserId search user by userId
func SearchUserByUserId(userId uint64) *model.User {
	return repository.SearchUserByUserId(userId)
}

// SearchUserList search user in database, this function could be brought a param
func SearchUserList(user *model.User) *[]*model.User {
	return repository.SearchUserList(user)
}

// UpdateUser update user
func UpdateUser(user *model.User) *model.User {
	return repository.UpdateUser(user)
}
