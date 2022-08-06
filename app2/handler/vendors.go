package handler

import (
	"app2/domain"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	RetrieveVendorUseCase interface {
		Execute(context.Context, string) (*domain.Vendor, error)
	}
	vendorsHandler struct {
		retrieveVendorUseCase RetrieveVendorUseCase
	}
)

func NewVendorsHandler(retrieveVendorUseCase RetrieveVendorUseCase) *vendorsHandler {
	return &vendorsHandler{retrieveVendorUseCase: retrieveVendorUseCase}
}

func (v vendorsHandler) HandleRetrieveVendor(ctx *gin.Context) {
	request := &struct {
		ID string `uri:"vendor-id"`
	}{}

	err := ctx.ShouldBindUri(request)
	if err != nil {
		return
	}

	vendor, err := v.retrieveVendorUseCase.Execute(ctx.Request.Context(), request.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, struct {
			message string
		}{message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, vendor)

}
