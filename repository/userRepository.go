package repository

import (
	"golang_crud/model"
	"golang_crud/utils/database"
	"log"
	"strconv"
)

// InsertUser insert a user to database
func InsertUser(user *model.User) *model.User {
	fields := ""
	paramsHandler := ""
	var params []interface{}
	// Judge Username
	if user.Username != "" {
		fields = "`username`"
		paramsHandler = "?"
		params = append(params, user.Username)
	}
	// Judge RealName
	if user.RealName != "" {
		if fields != "" {
			fields = fields + ", `real_name`"
			paramsHandler = paramsHandler + ", ?"
		} else {
			fields = "`real_name`"
			paramsHandler = "?"
		}
		params = append(params, user.RealName)
	}
	// Judge PhoneNumber
	if user.PhoneNumber != "" {
		if fields != "" {
			fields = fields + ", `phone_number`"
			paramsHandler = paramsHandler + ", ?"
		} else {
			fields = "`phone_number`"
			paramsHandler = "?"
		}
		params = append(params, user.PhoneNumber)
	}
	// Judge Email
	if user.Email != "" {
		if fields != "" {
			fields = fields + ", `email`"
			paramsHandler = paramsHandler + ", ?"
		} else {
			fields = "`email`"
			paramsHandler = "?"
		}
		params = append(params, user.Email)
	}
	// Judge Password
	if user.Password != "" {
		if fields != "" {
			fields = fields + ", `password`"
			paramsHandler = paramsHandler + ", ?"
		} else {
			fields = "`password`"
			paramsHandler = "?"
		}
		params = append(params, user.Password)
	}
	// Judge Password
	if user.Birthday != "" {
		if fields != "" {
			fields = fields + ", `birthday`"
			paramsHandler = paramsHandler + ", ?"
		} else {
			fields = "`birthday`"
			paramsHandler = "?"
		}
		params = append(params, user.Birthday)
	}
	// Judge Gender
	if user.Gender != 0 {
		if fields != "" {
			fields = fields + ", `gender`"
			paramsHandler = paramsHandler + ", ?"
		} else {
			fields = "`gender`"
			paramsHandler = "?"
		}
		params = append(params, user.Gender)
	}
	var sql = "INSERT INTO `user` (" + fields + ") VALUES (" + paramsHandler + ")"
	// Record a log to console
	log.Printf("Insert database: %v\n", sql)
	log.Printf("Insert params: %v\n", params)
	result, err := database.DBConnection.Exec(sql, params...)
	if err != nil {
		log.Printf("Insert database happend an error.\n sql: %v\nparams: %v\nerror: %v\n", sql, params, err)
		return nil
	}
	userId, err := result.LastInsertId()
	if err != nil {
		log.Printf("Get insert database last id happend an error: %v\n", err)
		return nil
	}
	return SearchUserByUserId(uint64(userId))
}

// DeleteUserByUserIds delete users by user ids
func DeleteUserByUserIds(userIds *[]uint64) int64 {
	sql := "UPDATE `user` SET `del_status` = 1 WHERE `del_status` = 0 AND `user_id` IN (" + strconv.FormatUint((*userIds)[0], 10)
	for i := 1; i < len(*userIds); i++ {
		sql = sql + ", " + strconv.FormatUint((*userIds)[i], 10)
	}
	sql = sql + ")"
	result, err := database.DBConnection.Exec(sql)
	log.Printf("Delete database: %v\n", sql)
	if err != nil {
		log.Printf("Delete database happend an error.\n sql: %v\nparams: %v\nerror: %v\n", sql, userIds, err)
		return 0
	} else {
		rows, _ := result.RowsAffected()
		return rows
	}
}

// SearchUserByUserId search user by userId
func SearchUserByUserId(userId uint64) *model.User {
	// Search database sql
	var sql = "SELECT `user_id`, `username`, `real_name`, `phone_number`, `email`, `password`, `birthday`, `gender` FROM `user` WHERE `del_status` = 0 AND `user_id` = ?"
	// Record a log to console
	log.Printf("Query database: %v\n", sql)
	log.Printf("Query params: %v\n", userId)
	row := database.DBConnection.QueryRow(sql, userId)
	var user model.User
	err := row.Scan(&user.UserId, &user.Username, &user.RealName, &user.PhoneNumber, &user.Email, &user.Password, &user.Birthday, &user.Gender)
	if err != nil {
		log.Printf("Scan query result happend an error: %v\n", err)
		return nil
	}
	return &user
}

// SearchUserList search user in database, this function could be brought a param
func SearchUserList(user *model.User) *[]*model.User {
	// Search database sql
	var sql = "SELECT `user_id`, `username`, `real_name`, `phone_number`, `email`, `password`, `birthday`, `gender` FROM `user` WHERE `del_status` = 0"
	// Search Params
	var params []interface{}
	// Judge UserId
	if user.UserId != 0 {
		sql = sql + " AND `user_id` = ?"
		params = append(params, user.UserId)
	}
	// Judge Username
	if user.Username != "" {
		sql = sql + " AND `username` LIKE CONCAT('%', ?, '%')"
		params = append(params, user.Username)
	}
	// Judge RealName
	if user.RealName != "" {
		sql = sql + " AND `real_name` LIKE CONCAT('%', ?, '%')"
		params = append(params, user.RealName)
	}
	// Judge PhoneNumber
	if user.PhoneNumber != "" {
		sql = sql + " AND `phone_number` LIKE CONCAT('%', ?, '%')"
		params = append(params, user.PhoneNumber)
	}
	// Judge Email
	if user.Email != "" {
		sql = sql + " AND `email` LIKE CONCAT('%', ?, '%')"
		params = append(params, user.Email)
	}
	// Judge Password
	if user.Password != "" {
		sql = sql + " AND `password` = ?"
		params = append(params, user.Password)
	}
	// Judge Password
	if user.Birthday != "" {
		sql = sql + " AND `birthday` = ?"
		params = append(params, user.Birthday)
	}
	// Judge Gender
	if user.Gender != 0 {
		sql = sql + " AND `gender` = ?"
		params = append(params, user.Gender)
	}
	// Record a log to console
	log.Printf("Query database: %v\n", sql)
	log.Printf("Query params: %v\n", params)
	// Do query database
	rows, err := database.DBConnection.Query(sql, params...)
	// Handle error
	if err != nil {
		log.Printf("Query database happend an error.\n sql: %v\nparams: %v\nerror: %v\n", sql, params, err)
		return nil
	}
	// Handle all user result
	var users []*model.User
	for rows.Next() {
		// Scan user entity
		var user model.User
		err = rows.Scan(&user.UserId, &user.Username, &user.RealName, &user.PhoneNumber, &user.Email, &user.Password, &user.Birthday, &user.Gender)
		if err != nil {
			log.Printf("Scan query result happend an error: %v\n", err)
			continue
		}
		// Append user
		users = append(users, &user)
	}
	return &users
}

// UpdateUser update user
func UpdateUser(user *model.User) *model.User {
	// Get need to update entry userId
	userId := user.UserId
	// Handle user exists
	resultUser := SearchUserByUserId(userId)
	if resultUser == nil {
		return nil
	}
	// Append params
	fields := ""
	var params []interface{}
	// Judge Username
	if user.Username != "" {
		fields = "`username` = ?"
		params = append(params, user.Username)
	}
	// Judge RealName
	if user.RealName != "" {
		if fields != "" {
			fields = fields + ", `real_name` = ?"
		} else {
			fields = "`real_name` = ?"
		}
		params = append(params, user.RealName)
	}
	// Judge PhoneNumber
	if user.PhoneNumber != "" {
		if fields != "" {
			fields = fields + ", `phone_number` = ?"
		} else {
			fields = "`phone_number` = ?"
		}
		params = append(params, user.PhoneNumber)
	}
	// Judge Email
	if user.Email != "" {
		if fields != "" {
			fields = fields + ", `email` = ?"
		} else {
			fields = "`email` = ?"
		}
		params = append(params, user.Email)
	}
	// Judge Password
	if user.Password != "" {
		if fields != "" {
			fields = fields + ", `password` = ?"
		} else {
			fields = "`password` = ?"
		}
		params = append(params, user.Password)
	}
	// Judge Password
	if user.Birthday != "" {
		if fields != "" {
			fields = fields + ", `birthday` = ?"
		} else {
			fields = "`birthday` = ?"
		}
		params = append(params, user.Birthday)
	}
	// Judge Gender
	if user.Gender != 0 {
		if fields != "" {
			fields = fields + ", `gender` = ?"
		} else {
			fields = "`gender` = ?"
		}
		params = append(params, user.Gender)
	}
	if fields == "" {
		return nil
	}
	params = append(params, userId)
	sql := "UPDATE `user` SET " + fields + " WHERE `del_status` = 0 AND `user_id` = ?"
	// Record a log to console
	log.Printf("Update database: %v\n", sql)
	log.Printf("Insert params: %v\n", params)
	result, err := database.DBConnection.Exec(sql, params...)
	if err != nil {
		log.Printf("Update database happend an error.\n sql: %v\nparams: %v\nerror: %v\n", sql, params, err)
		return nil
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Printf("Get Update database rows happend an error: %v\n", err)
		return nil
	} else if rows < 1 {
		return nil
	}
	return SearchUserByUserId(userId)
}
