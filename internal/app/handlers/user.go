package handlers

import (
	"fmt"
	"net/http"

	"github.com/t67y110v/web/internal/app/handlers/requests"
	model "github.com/t67y110v/web/internal/app/model/user"
	"github.com/t67y110v/web/internal/app/utils"

	"github.com/gofiber/fiber/v2"
)

// @Summary User Login
// @Description authentification user in the system
// @Tags         User
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.Login true  "login"
// @Success 200 {object} responses.Login
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/user/login [post]
func (h *Handlers) UserLogin() fiber.Handler {

	return func(c *fiber.Ctx) error {
		req := requests.Login{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		u, err := h.pgStore.Repository().FindByEmail(req.Email)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if !u.ComparePassword(req.Password) {
			return c.JSON(fiber.Map{
				"message": "wrong password",
			})
		}

		if u.Id == 0 {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		t, err := utils.CreateToken(u.Id, u.CenterID, u.Role, u.Name)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"JWT": t,
		})
	}

}

// @Summary User Registration
// @Description registration of user
// @Tags         User
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.Registration  true "registration"
// @Success 200 {object} responses.Registration
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/user/register [post]
func (h *Handlers) UserRegister() fiber.Handler {

	return func(c *fiber.Ctx) error {

		req := requests.Registration{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		u := &model.User{
			Email:          req.Email,
			Password:       req.Password,
			Name:           req.Name,
			Role:           req.Role,
			CenterID:       req.CenterID,
			OrganizationID: req.OrganizationID,
		}
		if err := h.pgStore.Repository().Create(u); err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		u.Sanitize()
		go h.journal.SaveAction(c.Context(), fmt.Sprintf("Регистрация пользователя %s", req.Name), c.Cookies("token_name"), c.Cookies("token_role"), "create", u)
		return c.JSON(u)
	}

}

// @Summary User Update
// @Description update of user
// @Tags         User
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.Update true "update"
// @Success 200 {object} responses.Registration
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/user/update [patch]
func (h *Handlers) UserUpdate() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.Update{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		u, err := h.pgStore.Repository().FindByID(req.ID)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		if err := h.pgStore.Repository().UpdateUser(req.ID, req.Role, req.CenterID, req.Email, req.Name, req.Paswword); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		go h.journal.SaveAction(c.Context(), fmt.Sprintf("Обновление пользователя %s|id:%d", req.Name, req.ID), c.Cookies("token_name"), c.Cookies("token_role"), "update", req, u)
		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}

// @Summary User Delete
// @Description delete of user
// @Tags         User
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.Delete true "delete"
// @Success 200 {object} responses.Registration
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/user/delete [delete]
func (h *Handlers) UserDelete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := requests.Delete{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		u, err := h.pgStore.Repository().FindByID(req.ID)
		if err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		u.ClearPswd()
		if err := h.pgStore.Repository().DeleteUser(req.ID); err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		go h.journal.SaveAction(c.Context(), fmt.Sprintf("Удаление пользователя %s", u.Email), c.Cookies("token_name"), c.Cookies("token_role"), "delete", u)
		return c.JSON(fiber.Map{
			"message": "success",
		})
	}
}

// @Summary Get Users
// @Description getting all  of users
// @Tags         User
//
//	@Accept       json
//
// @Produce json
// @Success 200 {object} responses.Registration
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /api/user/all [get]
func (h *Handlers) GetUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		u, err := h.pgStore.Repository().GetUsers()
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return c.JSON(u)
	}
}

func (h *Handlers) NewUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := h.pgStore.Repository().Create(&model.User{
			Id:       2,
			Email:    "admin@test.ru",
			Password: "string",
			Name:     "Admin",
			Role:     1,
			CenterID: 1,
		}); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		if err := h.pgStore.Repository().AddNewCenter("FirstCenter", 1); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		return nil
	}

}
