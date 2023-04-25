package handlers

import (
	"strconv"

	"time"

	model "github.com/t67y110v/web/internal/app/model/user"
	"github.com/t67y110v/web/internal/app/utils"

	"github.com/gofiber/fiber/v2"
)

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
		role, err := strconv.Atoi(c.FormValue("role"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}
		centerID, err := strconv.Atoi(c.FormValue("center_id"))
		if err != nil {
			return utils.ErrorPage(c, err)
		}

		u := &model.User{
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
			Name:     c.FormValue("name"),
			Role:     role,
			CenterID: centerID,
		}

		if err := h.pgStore.Repository().Create(u); err != nil {
			return utils.ErrorPage(c, err)
		}

		u.Sanitize()

		return c.Redirect("/admin/panel")

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

		email := c.FormValue("email")
		password := c.FormValue("password")
		u, err := h.pgStore.Repository().FindByEmail(email)
		if err != nil {
			return utils.LoginError(c)
		}

		if !u.ComparePassword(password) {
			return utils.PasswordError(c)
		}

		if u.Id == 0 {
			return utils.LoginError(c)
		}

		t, err := utils.CreateToken(u.Id, u.CenterID, u.Role, u.Name)
		if err != nil {
			return utils.LoginError(c) // 502 internal
		}
		cookie := &fiber.Cookie{
			Name:  "JWT",
			Value: t,
		}

		c.Cookie(cookie)
		return c.Redirect("/main/filter=0")
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
