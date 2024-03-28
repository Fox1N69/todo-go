package handlers

import (
	"blog/internal/services"
	"blog/pkg/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PostHandlerI interface {
	GetAllPosts(c echo.Context) error
	CreatePost(c echo.Context) error
	UpdatePost(c echo.Context) error
	DeletePost(c echo.Context) error
}
type PostHandler struct {
	service *services.PostServeci
}

func NewPostHandler(postService *services.PostServeci) *PostHandler {
	return &PostHandler{service: postService}
}

func stringToUint(str string) (uint, error) {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, errors.New("Filed to contvert string to uint")
	}

	return uint(num), nil
}

func (h *PostHandler) GetAllPosts(c echo.Context) error {
	post, err := h.service.GetAllPosts()
	if err != nil {
		return err
	}

	return c.JSON(200, post)
}

func (h *PostHandler) CreatePost(c echo.Context) error {
	post := new(models.Post)

	//binding data request to the post model
	if err := c.Bind(post); err != nil {
		return c.JSON(400, "Invalid requst body")
	}

	//create post with service
	err := h.service.CreatePost(post)
	if err != nil {
		return c.JSON(400, "Failed create post")
	}

	return c.JSON(200, map[string]interface{}{
		"message": "Create successfully",
		"post":    post,
	})
}

func (h *PostHandler) UpdatePost(c echo.Context) error {
	idParam := c.Param("id")

	//convert param id to uint
	id, err := stringToUint(idParam)
	if err != nil {
		return c.JSON(400, "Invalid ID")
	}

	//binding request data to the post model
	var post models.Post
	if err := c.Bind(&post); err != nil {
		return c.JSON(400, "Invalid request body")
	}

	// update post with use service
	if err := h.service.UpdatePost(id, &post); err != nil {
		return c.JSON(http.StatusBadRequest, "Failed to update post")
	}

	return c.JSON(200, "Post updated successfully")
}

func (h *PostHandler) DeletePost(c echo.Context) error {
	return nil
}
