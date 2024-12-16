package ceknumbercontroller

import (
	"main/helper"
	ceknumberrepository "main/repository/cekNumber_repository"
	"main/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PhoneHandler struct
type PhoneHandler struct {
	service *service.AllService
}

// NewPhoneHandler creates a new handler
func NewPhoneHandler(service *service.AllService) PhoneHandler {
	return PhoneHandler{service: service}
}

// CekNumber checks if the phone number is "cantik" and processes it
func (ctrl *PhoneHandler) CekNumber(ctx *gin.Context) {
	number := ctx.Query("number")
	_, err := strconv.Atoi(number)
	if err != nil {
		helper.Responses(ctx, http.StatusBadRequest, "Number yang dimasukan tidak valid", nil)
		ctx.Abort()
		return
	}

	msg := ctrl.service.CekNum.FilterCantikNumbers(number)
	helper.Responses(ctx, http.StatusOK, msg, nil)
}

// GetAllNumbers retrieves all stored phone numbers
func (ctrl *PhoneHandler) GetAllNumbers(ctx *gin.Context) {
	data := ctrl.service.CekNum.GetAllNumbers()
	if len(data) == 0 {
		helper.Responses(ctx, http.StatusInternalServerError, "Database kosong", nil)
		ctx.Abort()
		return
	}

	helper.Responses(ctx, http.StatusOK, "Berhasil mengambil semua nomor cantik", data)
}

// FindNumber searches for a specific phone number
func (ctrl *PhoneHandler) FindNumber(ctx *gin.Context) {
	number := ctx.Query("number")
	_, err := strconv.Atoi(number)
	if err != nil {
		helper.Responses(ctx, http.StatusBadRequest, "Number yang dimasukan tidak valid", nil)
		ctx.Abort()
		return
	}

	phone, msg := ctrl.service.CekNum.FindNumber(number)
	if phone == nil {
		helper.Responses(ctx, http.StatusNotFound, msg, nil)
		return
	}

	helper.Responses(ctx, http.StatusOK, msg, phone)
}

// UpdateNumber updates a specific phone number by ID
func (ctrl *PhoneHandler) UpdateNumber(ctx *gin.Context) {
	number := ctx.Query("number")
	_, err := strconv.Atoi(number)
	if err != nil {
		helper.Responses(ctx, http.StatusBadRequest, "Number yang dimasukan tidak valid", nil)
		ctx.Abort()
		return
	}

	newNumber := ctx.Query("new_number")
	_, err = strconv.Atoi(newNumber)
	if err != nil {
		helper.Responses(ctx, http.StatusBadRequest, "Number yang dimasukan tidak valid", nil)
		ctx.Abort()
		return
	}

	if !ceknumberrepository.IsCantik(newNumber) {
		helper.Responses(ctx, http.StatusBadRequest, "Nomor Baru Tidak Cantik Update digagalkan", nil)
		ctx.Abort()
		return
	}

	msg := ctrl.service.CekNum.UpdateNumber(number, newNumber)
	helper.Responses(ctx, http.StatusOK, msg, nil)
}

// DeleteNumber deletes a specific phone number by ID
func (ctrl *PhoneHandler) DeleteNumber(ctx *gin.Context) {
	oldNumber := ctx.Query("number")
	_, err := strconv.Atoi(oldNumber)
	if err != nil {
		helper.Responses(ctx, http.StatusBadRequest, "Number yang dimasukan tidak valid", nil)
		ctx.Abort()
		return
	}

	msg := ctrl.service.CekNum.DeleteNumber(oldNumber)
	helper.Responses(ctx, http.StatusOK, msg, nil)
}
