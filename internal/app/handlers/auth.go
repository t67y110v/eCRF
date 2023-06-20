package handlers

import (
	"net/http"
	"strconv"

	"time"

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
// @Router /user/login [post]
func (h *Handlers) Login() fiber.Handler {

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
				"message": err.Error(),
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
// @Router /user/register [post]
func (h *Handlers) Register() fiber.Handler {

	return func(c *fiber.Ctx) error {

		req := requests.Registration{}
		if err := c.BodyParser(&req); err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
			Name:     req.Name,
			Role:     req.Role,
			CenterID: req.CenterID,
		}
		if err := h.pgStore.Repository().Create(u); err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		u.Sanitize()

		return c.JSON(u)
	}

}

func (h *Handlers) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {

		role, err := strconv.Atoi(c.FormValue("role"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		centerID, err := strconv.Atoi(c.FormValue("center_id"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		userID, err := strconv.Atoi(c.FormValue("user_id"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}

		if err := h.pgStore.Repository().UpdateUser(userID, role, centerID, c.FormValue("email"), c.FormValue("name"), c.FormValue("password")); err != nil {
			return utils.ErrorPage(c, err)
		}
		return c.Redirect("/admin/panel")
	}
}

func (h *Handlers) Logout() fiber.Handler {

	return func(c *fiber.Ctx) error {

		cookie := fiber.Cookie{
			Name:     "jwt",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
		}

		c.Cookie(&cookie)

		return c.JSON(fiber.Map{
			"message": "success",
		})
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
			return utils.ErrorPage(c, err)
		}
		if err := h.pgStore.Repository().AddNewCenter("FirstCenter"); err != nil {
			return utils.ErrorPage(c, err)
		}
		return nil
	}

}
