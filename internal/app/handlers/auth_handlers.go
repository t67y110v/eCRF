package handlers

import (
	"bytes"

	"strconv"

	"encoding/json"

	"time"

	"github.com/t67y110v/web/internal/app/handlers/requests"
	model "github.com/t67y110v/web/internal/app/model/user"
	"github.com/t67y110v/web/internal/app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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
		//id, name, centerId, role, nil
		_, _, _, _, err := utils.CheckToken(c.Cookies("JWT"))
		if err != nil {
			return utils.LoginError(c)
		}
		role, err := strconv.Atoi(c.FormValue("role"))
		if err != nil {
			return err
		}
		centerID, err := strconv.Atoi(c.FormValue("center_id"))
		if err != nil {
			return err
		}

		u := &model.User{
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
			Name:     c.FormValue("name"),
			Role:     role,
			CenterID: centerID,
		}

		if err := h.pgStore.Repository().Create(u); err != nil {
			return err
		}

		u.Sanitize()

		return c.Redirect("/admin/panel")

	}

}

func (h *Handlers) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {

		_, _, _, _, err := utils.CheckToken(c.Cookies("JWT"))
		if err != nil {
			return utils.LoginError(c)
		}
		role, err := strconv.Atoi(c.FormValue("role"))
		if err != nil {
			return err
		}
		centerID, err := strconv.Atoi(c.FormValue("center_id"))
		if err != nil {
			return err
		}
		userID, err := strconv.Atoi(c.FormValue("user_id"))
		if err != nil {
			return err
		}

		if err := h.pgStore.Repository().UpdateUser(userID, role, centerID, c.FormValue("email"), c.FormValue("name"), c.FormValue("password")); err != nil {
			return err
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

// @Summary Check session
// @Description Validation user token
// @Tags         User
//
//	@Accept       json
//
// @Produce json
// @Param  data body requests.CheckToken true  "Check token"
// @Success 200 {object} responses.Login
// @Failure 400 {object} responses.Error
// @Failure 500 {object} responses.Error
// @Router /user/check [post]
func (h *Handlers) CheckJWT() fiber.Handler {

	return func(c *fiber.Ctx) error {

		req := &requests.CheckToken{}

		reader := bytes.NewReader(c.Body())

		if err := json.NewDecoder(reader).Decode(req); err != nil {
			h.logger.Warningf("handle checkJWT,  error :%e", err)
		}

		cookie := req.Cookie

		tokenString := cookie

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if claims["id"] == nil {
			h.logger.Warningf("handle checkJWT,  error :%e", err)
			return c.JSON(fiber.Map{
				"message": "token id is nil",
			})
		}

		id := float64(claims["id"].(float64))

		u, err := h.pgStore.Repository().FindByID(strconv.Itoa(int(id)))
		if err != nil {
			return err
		}

		return c.JSON(u)

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
