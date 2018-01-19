package controllers

import (
	Response "../responses"
	Config "../config"
)

func Auth(email string, password string) (bool) {
	var _response []Response.User

	Config.Database().Where("email = ?", email).Where("password = ?", password).First(&_response)

	if len(_response) <= 0 {
		return false
	}
	return true
}

func AuthEmail(email string) (bool) {
	var _response []Response.User

	Config.Database().Where("email = ?", email).First(&_response)

	if len(_response) <= 0 {
		return false
	}
	return true
}