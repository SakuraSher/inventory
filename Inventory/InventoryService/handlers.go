package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type InventoryHandlers struct {
	service *InventoryService
}

func NewInventoryHandlers(service *InventoryService) *InventoryHandlers {
	return &InventoryHandlers{service: service}
}

func (h *InventoryHandlers) GetItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "InvalidItemID"})
	}
	item, err := h.service.GetItem(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "InvalidId"})

	}
	return (c.JSON(http.StatusOK, item))

}

func (h *InventoryHandlers) CreateItem(c echo.Context) error {
	item := new(Item)
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, item)

}

func (h *InventoryHandlers) UpdateItemQuantity(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	var req struct {
		Delta int `json:"delta"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.service.UpdateItemQuantity(uint(id), req.Delta); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusOK)
}

func (h *InventoryHandlers) DeleteItem(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	if err := h.service.DeleteItem(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.NoContent(http.StatusOK)

}

func (h *InventoryHandlers) ListItems(c echo.Context) error {
	items, err := h.service.ListItems()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, items)

}
