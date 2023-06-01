package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyenvantuan2391996/be-project/handler/constant"
	"github.com/nguyenvantuan2391996/be-project/handler/model"
	"github.com/sirupsen/logrus"
)

// Consult
// @Summary consults to choose best option
// @Description consults to choose best option
// @Tags consult
// @Accept json
// @Produce json
// @Param user_id query string false "user_id"
// @Success 200 {object} []model.ConsultResponse
// @Failure 400 {object} model.ErrorSystem
// @Router /v1/api/consult/:user_id [post]
func (h *Handler) Consult(c *gin.Context) {
	logrus.Info("Start api consult...")

	// Get request param
	userId := c.Param("user_id")

	result, err := h.consultDomain.Consult(c, userId)
	if err != nil {
		logrus.Errorf("Consult result fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.ConsultResultFail,
		})
		return
	}

	logrus.Info("Consult result success")
	res := make([]*model.ConsultResponse, 0)
	for _, value := range result {
		res = append(res, value.ToResponse())
	}
	c.JSON(http.StatusOK, res)
}
