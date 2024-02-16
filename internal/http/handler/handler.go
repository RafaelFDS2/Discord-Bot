package handler

import (
	"ablufus/exceptions"
	"ablufus/internal/entities"
	"ablufus/internal/service"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	s service.Service
}

func New(s service.Service) *Handler {
	return &Handler{s}
}

func (h *Handler) Post(c echo.Context) error {
	var u entities.UserRequest
	if err := c.Bind(&u); err != nil {
		res := exceptions.HandleException(exceptions.New(exceptions.ErrBadData, err))
		return c.JSON(res.Code, res)
	}
	res, err := h.s.Post(u)
	if err != nil {
		res := exceptions.HandleException(err)
		return c.JSON(res.Code, res)
	}
	return c.JSON(http.StatusCreated, res)
}

func (h *Handler) List(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 10
	}
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}
	ids := strings.Split(c.QueryParam("ids"), ",")
	res, err := h.s.List(ids, limit, page)
	if err != nil {
		res := exceptions.HandleException(err)
		return c.JSON(res.Code, res)
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) Update(c echo.Context) error {
	id := c.Param("user_id")
	var u entities.UserPatchRequest
	if err := c.Bind(&u); err != nil {
		fmt.Print(err.Error())
		res := exceptions.HandleException(exceptions.New(exceptions.ErrBadData, err))
		return c.JSON(res.Code, res)
	}
	err := entities.Validate(u)
	if err != nil {
		res := exceptions.HandleException(err)
		return c.JSON(res.Code, res)
	}

	err = h.s.Update(id, u.Amount)
	if err != nil {
		res := exceptions.HandleException(err)
		return c.JSON(res.Code, res)
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (h *Handler) Delete(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
