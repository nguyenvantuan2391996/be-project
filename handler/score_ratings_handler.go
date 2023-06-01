package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nguyenvantuan2391996/be-project/handler/constant"
	"github.com/nguyenvantuan2391996/be-project/handler/model"
	"github.com/sirupsen/logrus"
)

// BulkCreateScoreRating
// @Summary bulk create the score rating
// @Description bulk create the score rating
// @Tags score-ratings
// @Accept json
// @Produce json
// @Param request body model.ScoreRatingRequest true "score rating request"
// @Success 200 {object} model.BulkCreateResponse
// @Failure 400 {object} model.ErrorSystem
// @Router /v1/api/score-ratings [post]
func (h *Handler) BulkCreateScoreRating(c *gin.Context) {
	logrus.Info("Start api bulk create score rating...")

	// Parse request
	var scoreRatingRequest []*model.ScoreRatingRequest
	err := c.ShouldBindJSON(&scoreRatingRequest)
	if err != nil {
		logrus.Errorf("Parse request create score rating fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.ParseRequestFail,
		})
		return
	}

	// Validate request
	var listErr []*model.ErrorValidateRequest
	for _, request := range scoreRatingRequest {
		listErr = append(listErr, h.ValidateRequest(request)...)
	}
	if len(listErr) > 0 {
		logrus.Errorf("Request is invalid: %v", listErr)
		c.JSON(http.StatusBadRequest, listErr)
		return
	}

	err = h.scoreRatingDomain.BulkCreateScoreRating(c, scoreRatingRequest)
	if err != nil {
		logrus.Errorf("Bulk create new score rating fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.BulkCreateNewScoreRatingFail,
		})
		return
	}

	logrus.Info("Bulk create new score rating success")
	c.JSON(http.StatusCreated, &model.BulkCreateResponse{
		Type:   "Bulk create",
		Status: model.SUCCESS,
	})
}

// GetScoreRatings
// @Summary gets the score rating
// @Description gets the score rating
// @Tags score-ratings
// @Accept json
// @Produce json
// @Param user_id query string false "user_id"
// @Success 200 {object} []model.ScoreRatingResponse
// @Failure 400 {object} model.ErrorSystem
// @Router /v1/api/score-ratings/:user_id [get]
func (h *Handler) GetScoreRatings(c *gin.Context) {
	logrus.Info("Start api get list score rating by user id...")

	// Get request param
	userID := c.Param("user_id")

	result, err := h.scoreRatingDomain.GetListScoreRating(c, &model.ScoreRating{
		UserId: userID,
	})
	if err != nil {
		logrus.Errorf("Get list score rating fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.GetListScoreRatingFail,
		})
		return
	}

	logrus.Info("Get list score rating success")
	res := make([]*model.ScoreRatingResponse, 0)
	for _, value := range result {
		res = append(res, &model.ScoreRatingResponse{
			ID:       value.ID,
			Metadata: value.Metadata,
		})
	}
	c.JSON(http.StatusOK, res)
}

// UpdateScoreRating
// @Summary updates the score rating
// @Description updates the score rating
// @Tags score-ratings
// @Accept json
// @Produce json
// @Param request body model.ScoreRatingRequest true "score rating request"
// @Success 200 {object} model.UpdateResponse
// @Failure 400 {object} model.ErrorSystem
// @Router /v1/api/score-ratings [put]
func (h *Handler) UpdateScoreRating(c *gin.Context) {
	logrus.Info("Start api update score rating...")

	// Parse request
	var scoreRatingRequest *model.ScoreRatingRequest
	err := c.ShouldBindJSON(&scoreRatingRequest)
	if err != nil {
		logrus.Errorf("Parse request update score rating fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.ParseRequestFail,
		})
		return
	}

	err = h.scoreRatingDomain.UpdateScoreRating(c, &model.ScoreRating{
		ID:       scoreRatingRequest.ID,
		Metadata: scoreRatingRequest.Metadata,
	})
	if err != nil {
		logrus.Errorf("Update score rating fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.UpdateScoreRatingFail,
		})
		return
	}

	logrus.Info("Update score rating success")
	c.JSON(http.StatusOK, &model.UpdateResponse{
		Id:     scoreRatingRequest.ID,
		Status: model.SUCCESS,
	})
}

// DeleteScoreRating
// @Summary deletes the score rating
// @Description deletes the score rating
// @Tags score-ratings
// @Accept json
// @Produce json
// @Param id query string false "id"
// @Success 200 {object} model.DeletedResponse
// @Failure 400 {object} model.ErrorSystem
// @Router /v1/api/score-ratings/:id [delete]
func (h *Handler) DeleteScoreRating(c *gin.Context) {
	logrus.Info("Start api delete score rating...")

	// Get request param
	id := c.Param("id")

	err := h.scoreRatingDomain.DeleteScoreRating(c, &model.ScoreRating{
		ID: id,
	})
	if err != nil {
		logrus.Errorf("Delete score rating fail: %v", err)
		c.JSON(http.StatusBadRequest, &model.ErrorSystem{
			Code:    http.StatusBadRequest,
			Message: constant.DeleteScoreRatingFail,
		})
		return
	}

	logrus.Info("Delete score rating success")
	c.JSON(http.StatusOK, &model.DeletedResponse{
		Id:     id,
		Status: model.SUCCESS,
	})
}
