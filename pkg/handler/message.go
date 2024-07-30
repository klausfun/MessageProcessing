package handler

import (
	"MessageProcessing/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createMessage(c *gin.Context) {
	var input models.Message
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Message.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllMessageResponse struct {
	CurNum  int              `json:"number_of_current_messages"`
	CurMes  []models.Message `json:"current"`
	CompNum int              `json:"number_of_completed_messages"`
	CompMes []models.Message `json:"completed"`
}

func (h *Handler) getAllMessages(c *gin.Context) {
	curMessages, err := h.services.Message.GetCurMessages()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	compMessages, err := h.services.Message.GetCompMessages()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllMessageResponse{
		CurNum:  len(curMessages),
		CurMes:  curMessages,
		CompNum: len(compMessages),
		CompMes: compMessages,
	})
}
