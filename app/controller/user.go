package controller

import (
	"github.com/gofiber/fiber/v2"
)

// User Create
// @Summary Create new user.
// @Description -
// @Tags user
// @Accept json
// @Produce json
// @param User body model.UserForCreate true "Customer data to be updated"
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.APIError "BadRequest"
// @Failure 500 {object} model.APIError "Internal server error!!"
// @Router /user [post]
func UserCreateController(c *fiber.Ctx) error {
	// TODO
	return nil
}

// User Update
// @Summary update user data.
// @Description -
// @Tags user
// @security ApiKeyAuth
// @Accept json
// @Produce json
// @param id path int true "id of user to be updated"
// @Param message body model.UserForUpdate true "User Data"
// @Success 200 {object} model.Response "OK"
// @Failure 400 {object} model.APIError "We need ID!!"
// @Failure 404 {object} model.APIError "Can not find ID"
// @Failure 500 {object} model.APIError "Internal server error!!"
// @Router /user/:id [put]
func UserUpdateController(c *fiber.Ctx) error {
	// TODO
	return nil
}

// GetUser
// @Summary Read user from the store
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path int true "User Id"
// @Success 200 {object} model.User
// @Failure 400 {object} model.APIError "We need ID!!"
// @Failure 404 {object} model.APIError "Can not find ID"
// @Router /user/{id} [get]
func UserGetControlller(c *fiber.Ctx) error {
	// TODO
	return nil
}

// ListUsers example
// @Summary List users from the store
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Users	"ok"
// @Router /user [get]
func UserListController(c *fiber.Ctx) error {
	// TODO
	return nil
}
