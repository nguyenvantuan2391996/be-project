package handler

import (
	"net/http"

	"github.com/nguyenvantuan2391996/be-project/handler/constant"
	"github.com/nguyenvantuan2391996/be-project/handler/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// CreateUser
// @Summary creates the user account
// @Description creates the user account
// @Tags users
// @Accept json
// @Produce json
// @Param request body model.UserRequest true "user request"
// @Success 200 {object} model.UserResponse
// @Failure 400 {object} model.ErrorSystem
// @Router /v1/api/users [post]
func (h *Handler) CreateUser(c *gin.Context) {
	logrus.Info("Start api create user...")

	// Parse request
	var userRequest *model.UserRequest
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		logrus.Errorf("Parse request create user fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.ParseRequestFail,
		})
		return
	}

	// Validate request
	listErr := h.ValidateRequest(userRequest)
	if len(listErr) > 0 {
		logrus.Errorf("Request is invalid: %v", listErr)
		c.JSON(http.StatusBadRequest, listErr)
		return
	}

	userCreated, err := h.userDomain.CreateUser(c, userRequest.Name)
	if err != nil {
		logrus.Errorf("Create new user fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.CreateNewUserFail,
		})
		return
	}

	logrus.Info("Create new user success")
	c.JSON(http.StatusCreated, &model.UserResponse{
		ID:   userCreated.ID,
		Name: userCreated.Name,
	})
}
