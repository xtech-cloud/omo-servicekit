package main

import (
	"osk/api/account"

	"github.com/gin-gonic/gin"
)

func route(_group *gin.RouterGroup) {

	account.HandleSignout("/signout", _group)
	// /account
	{
		account.HandleCurrent("/account/current", _group)
	}
}
