package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyenvantuan2391996/be-project/handler/constant"
	"github.com/nguyenvantuan2391996/be-project/handler/model"
	"github.com/sirupsen/logrus"
)

// CreateStandard
// @Summary creates the standard attribute
// @Description creates the standard attribute
// @Tags standards
// @Accept json
// @Produce json
// @Param request body model.StandardRequest true "standard request"
// @Success 200 {object} model.StandardResponse
// @Failure 400 {object} model.ErrorSystem
// @Router /v1/api/standards [post]
func (h *Handler) CreateStandard(c *gin.Context) {
	logrus.Info("Start api create standard...")

	// Parse request
	var standardRequest *model.StandardRequest
	err := c.ShouldBindJSON(&standardRequest)
	if err != nil {
		logrus.Errorf("Parse request create standard fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.ParseRequestFail,
		})
		return
	}

	// Validate request
	listErr := h.ValidateRequest(standardRequest)
	if len(listErr) > 0 {
		logrus.Errorf("Request is invalid: %v", listErr)
		c.JSON(http.StatusBadRequest, listErr)
		return
	}

	standardCreated, err := h.standardDomain.CreateStandard(c, standardRequest)
	if err != nil {
		logrus.Errorf("Create new standard fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.CreateNewStandardFail,
		})
		return
	}

	logrus.Info("Create new standard success")
	c.JSON(http.StatusCreated, &model.StandardResponse{
		ID:           standardCreated.ID,
		UserID:       standardCreated.UserId,
		StandardName: standardCreated.StandardName,
		Weight:       standardCreated.Weight,
	})
}

// GetStandards
// @Summary gets the standard attribute
// @Description gets the standard attribute
// @Tags standards
// @Accept json
// @Produce json
// @Param user_id query string false "user_id"
// @Success 200 {object} []model.StandardResponse
// @Failure 400 {object} model.ErrorSystem
// @Router /v1/api/standards/:user_id [get]
func (h *Handler) GetStandards(c *gin.Context) {
	logrus.Info("Start api get list standard...")

	// Get request param
	userID := c.Param("user_id")

	standards, err := h.standardDomain.GetStandards(c, &model.Standard{
		UserID: userID,
	})
	if err != nil {
		logrus.Errorf("Get list standard fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.GetStandardsFail,
		})
		return
	}

	logrus.Info("Get list standard success")
	res := make([]*model.StandardResponse, 0)
	for _, s := range standards {
		res = append(res, s.ToResponse())
	}
	c.JSON(http.StatusOK, res)
}

// DeleteStandard
// @Summary deletes the standard attribute
// @Description deletes the standard attribute
// @Tags standards
// @Accept json
// @Produce json
// @Param id query string false "id"
// @Success 200 {object} []model.DeletedResponse
// @Failure 400 {object} model.ErrorSystem
// @Router /v1/api/standards/:id [delete]
func (h *Handler) DeleteStandard(c *gin.Context) {
	logrus.Info("Start api delete standard...")

	// Get request param
	id := c.Param("id")

	err := h.standardDomain.DeleteStandard(c, &model.Standard{
		ID: id,
	})
	if err != nil {
		logrus.Errorf("Delete standard fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.DeleteStandardFail,
		})
		return
	}

	logrus.Info("Delete standard success")
	c.JSON(http.StatusOK, &model.DeletedResponse{
		Id:     id,
		Status: model.SUCCESS,
	})
}
