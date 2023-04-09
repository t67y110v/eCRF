package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (h *Handlers) ProtocolPage() fiber.Handler {

	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("JWT")
		tokenString := cookie

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if claims["id"] == nil {
			h.logger.Warningf("handle checkJWT,  error :%e", err)
			return c.Redirect("/auth")
		}

		id := float64(claims["id"].(float64))

		u, err := h.pgStore.UserRepository().FindByID(strconv.Itoa(int(id)))
		if err != nil {
			return err
		}

		protocol_id := c.Params("id")
		p_id, err := strconv.Atoi(protocol_id)

		if err != nil {
			return c.Redirect("/auth")
		}

		p, err := h.pgStore.UserRepository().GetProtocolById(p_id)
		if err != nil {
			return c.Redirect("/auth")
		}
		return c.Render("protocol_index", fiber.Map{
			"Name":         u.Name,
			"Role":         u.Role,
			"CLinicCenter": u.CenterID,
			"Protocol":     p,
		})

	}

}

func (h *Handlers) ProtocolEdit() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("JWT")

		tokenString := cookie

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if claims["id"] == nil {
			h.logger.Warningf("handle checkJWT,  error :%e", err)
			return c.Redirect("/auth")
		}

		id := float64(claims["id"].(float64))

		u, err := h.pgStore.UserRepository().FindByID(strconv.Itoa(int(id)))
		if err != nil {
			return err
		}

		if err != nil {
			return c.Redirect("/auth")
		}

		protocol_id := c.Params("id")
		p_id, err := strconv.Atoi(protocol_id)
		if err != nil {
			return c.Redirect(fmt.Sprintf("/protocol/edit/%s", protocol_id))
		}
		p, err := h.pgStore.UserRepository().GetProtocolById(p_id)
		if err != nil {
			return c.Redirect("/auth")
		}

		return c.Render("protocol_edit", fiber.Map{
			"Name":         u.Name,
			"Role":         u.Role,
			"CLinicCenter": u.CenterID,
			"Protocol":     p,
		})
	}
}

func (h *Handlers) ProtocolSave() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("JWT")
		name := c.FormValue("name")
		status := c.FormValue("status")
		center_id := c.FormValue("center")
		protocol_id := c.FormValue("id")
		tokenString := cookie
		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if claims["id"] == nil {
			h.logger.Warningf("handle checkJWT,  error :%e", err)
			return c.Redirect("/auth")
		}

		id := float64(claims["id"].(float64))

		_, err = h.pgStore.UserRepository().FindByID(strconv.Itoa(int(id)))
		if err != nil {
			return err
		}

		p_id, err := strconv.Atoi(protocol_id)
		if err != nil {
			return c.Redirect(fmt.Sprintf("/protocol/edit/%s", protocol_id))
		}
		s, err := strconv.Atoi(status)
		if err != nil {
			return c.Redirect(fmt.Sprintf("/protocol/edit/%s", protocol_id))
		}
		c_id, err := strconv.Atoi(center_id)
		if err != nil {
			return c.Redirect(fmt.Sprintf("/protocol/edit/%s", protocol_id))
		}
		err = h.pgStore.UserRepository().UpdateProtocolById(p_id, s, name, c_id)
		if err != nil {
			return c.Redirect(fmt.Sprintf("/protocol/edit/%s", protocol_id))
		}

		return c.Redirect("/main/filter=0")
	}
}
func (h *Handlers) AuthPage() fiber.Handler {

	return func(c *fiber.Ctx) error {
		c.ClearCookie("JWT")
		er := c.Cookies("error")
		fmt.Print(er)
		if er == "password" {
			return c.Render("auth_page", fiber.Map{
				"WrongPassword": "True",
			})
		} else if er == "login" {
			return c.Render("auth_page", fiber.Map{
				"WrongLogin": "True",
			})
		}

		return c.Render("auth_page", fiber.Map{})

	}

}

func (h *Handlers) MainPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		filter := c.Params("filter")
		fmt.Println(filter)

		p, err := h.pgStore.UserRepository().GetProtocolsByFilter(filter)
		if err != nil {
			return err
		}

		c.ClearCookie("error")
		cookie := c.Cookies("JWT")
		if cookie == "" {
			return c.Redirect("/auth")
		}
		tokenString := cookie

		claims := jwt.MapClaims{}

		_, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if claims["id"] == nil {
			h.logger.Warningf("handle checkJWT,  error :%e", err)
			return c.Redirect("/auth")
		}

		id := float64(claims["id"].(float64))

		u, err := h.pgStore.UserRepository().FindByID(strconv.Itoa(int(id)))
		if err != nil {
			c.Status(fiber.StatusNotFound)
			er := &fiber.Cookie{
				Name:  "error",
				Value: "login",
			}
			c.Cookie(er)
			return c.RedirectBack("/auth")
		}

		return c.Render("main_page", fiber.Map{
			"Name":         u.Name,
			"Role":         u.Role,
			"CLinicCenter": u.CenterID,
			"Protocols":    p,
		})
	}
}

func (h *Handlers) NewProtocol() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("JWT")

		tokenString := cookie

		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if claims["id"] == nil {
			h.logger.Warningf("handle checkJWT,  error :%e", err)
			return c.Redirect("/auth")
		}

		id := float64(claims["id"].(float64))

		u, err := h.pgStore.UserRepository().FindByID(strconv.Itoa(int(id)))
		if err != nil {
			return err
		}

		if err != nil {
			return c.Redirect("/auth")
		}

		return c.Render("protocol_new", fiber.Map{
			"Name":         u.Name,
			"Role":         u.Role,
			"CLinicCenter": u.CenterID,
		})
	}
}

func (h *Handlers) AddProtocol() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("JWT")
		name := c.FormValue("name")
		status := c.FormValue("status")
		center_id := c.FormValue("center")
		tokenString := cookie
		claims := jwt.MapClaims{}

		_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		if claims["id"] == nil {
			h.logger.Warningf("handle checkJWT,  error :%e", err)
			return c.Redirect("/auth")
		}

		id := float64(claims["id"].(float64))

		_, err = h.pgStore.UserRepository().FindByID(strconv.Itoa(int(id)))
		if err != nil {
			return err
		}

		if err != nil {
			return c.Redirect("/main/filter=0")
		}
		s, err := strconv.Atoi(status)
		if err != nil {
			return c.Redirect("/main/filter=0")
		}
		c_id, err := strconv.Atoi(center_id)
		if err != nil {
			return c.Redirect("/main/filter=0")
		}
		err = h.pgStore.UserRepository().AddProtocol(name, s, c_id)
		if err != nil {
			return c.Redirect("/main/filter=0")
		}

		return c.Redirect("/main/filter=0")
	}
}
