package http

import (
	"net/http"
	"github.com/PhipattanachaiDev/golang_api-migration/internal/domain"
	"github.com/PhipattanachaiDev/golang_api-migration/internal/ports"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service ports.UserService
}

func NewHandler(service ports.UserService) *Handler {
	return &Handler{service: service}
}

// @Summary Create user
// @Produce json
// @Param user body domain.User true "User"
// @Success 201 {object} domain.User
// @Router /users [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.service.CreateUser(&user)
	c.JSON(http.StatusCreated, user)
}

// @Summary Get user
// @Param id path string true "User ID"
// @Produce json
// @Success 200 {object} domain.User
// @Router /users/{id} [get]
func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary List users
// @Produce json
// @Success 200 {array} domain.User
// @Router /users [get]
func (h *Handler) GetUsers(c *gin.Context) {
	users, _ := h.service.GetUsers()
	c.JSON(http.StatusOK, users)
}

// @Summary Update user
// @Param id path string true "User ID"
// @Param user body domain.User true "User"
// @Produce json
// @Router /users/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	h.service.UpdateUser(&user)
	c.JSON(http.StatusOK, user)
}

// @Summary Delete user
// @Param id path string true "User ID"
// @Produce json
// @Router /users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	h.service.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}