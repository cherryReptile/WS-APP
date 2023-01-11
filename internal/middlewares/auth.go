package middlewares

import (
	"context"
	"errors"
	"github.com/cherryReptile/WS-AUTH/api"
	"github.com/gofiber/fiber/v2"
	"strings"
)

//func CheckAuthHeader() fiber.Handler {
//	return func(c *fiber.Ctx) error {
//		h := c.GetReqHeaders()
//		_, ok := h["Authorization"]
//		if !ok {
//			return errors.New("auth header is required")
//		}
//		return c.Next()
//	}
//}

func CheckAuth(as api.CheckAuthServiceClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		t, err := getAuthHeader(c)
		if err != nil {
			return err
		}
		_, err = as.CheckAuth(context.Background(), &api.TokenRequest{Token: t})
		if err != nil {
			return err
		}
		return c.Next()
	}
}

func getAuthHeader(c *fiber.Ctx) (string, error) {
	h := c.GetReqHeaders()
	t, ok := h["Authorization"]
	if !ok {
		return "", errors.New("auth header is required")
	}

	v := strings.Split(t, " ")
	if len(v) < 2 || v[1] == "" || len(v) > 2 {
		return "", errors.New("incorrect auth header")
	}

	return v[1], nil
}
