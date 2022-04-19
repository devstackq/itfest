package handler

import (
	"context"
	"log"
	"net/http"

	"bimbo/internal/model"

	"github.com/gin-gonic/gin"
)

// todo: - token expires & user_id in redis -> check token_uuid;

func (h *Handler) SignIn(c *gin.Context) {
	var (
		user   *model.User
		err    error
		userID int
		token  *model.TokenDetails
	)

	err = c.ShouldBindJSON(&user)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	userID, err = h.services.AuthService.SignIn(ctx, user.Email, user.Password)
	if err != nil {
		h.logger.Error(err)
		if err.Error() == "sql: no rows in result set" {
			responseWithStatus(c, http.StatusBadRequest, err.Error(), "Incorrect password or email", nil)
			return
		}
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "Internal server error", nil)
		return
	}

	// call signin/signup/refresh
	token, err = CreateToken(userID, h.cfg.App.SecretAccess, h.cfg.App.SecretRefresh)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	token.UserID = userID

	// err = h.service.CreateSession(ctx, token)
	// refresh, /signin call
	// err = h.service.UpdateSession(ctx, token)

	c.Writer.Header().Set("refresh_token", token.RefreshToken)
	c.Writer.Header().Set("access_token", token.AccessToken)
	responseWithStatus(c, http.StatusOK, "success signin", "OK", nil)
}

func (h *Handler) SignUp(c *gin.Context) {
	var (
		user   *model.User
		err    error
		lastID int
		token  *model.TokenDetails
	)

	err = c.ShouldBindJSON(&user)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusBadRequest, err.Error(), "Input error", nil)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lastID, err = h.services.AuthService.SignUp(ctx, user)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}

	token, err = CreateToken(lastID, h.cfg.App.SecretAccess, h.cfg.App.SecretRefresh)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}

	c.Writer.Header().Set("refresh_token", token.RefreshToken)
	c.Writer.Header().Set("access_token", token.AccessToken)

	responseWithStatus(c, http.StatusOK, "success signup", "OK", lastID)
}

// refresh - invalid -> redirect signin; else newToken
func (h *Handler) RefreshJwt(c *gin.Context) {
	var (
		err   error
		token *model.TokenDetails
	)

	userID, ok := c.Get("user_id")
	if !ok {
		log.Println("not exist userId value")
		responseWithStatus(c, http.StatusBadRequest, "refresh token incorrect", "no have userI_id, to refresh token", nil)
		return
	}

	token, err = CreateToken(int(userID.(float64)), h.cfg.App.SecretAccess, h.cfg.App.SecretRefresh)
	if err != nil {
		h.logger.Error(err)
		responseWithStatus(c, http.StatusInternalServerError, err.Error(), "internal server error", nil)
		return
	}
	c.Writer.Header().Set("refresh_token", token.RefreshToken)
	c.Writer.Header().Set("access_token", token.AccessToken)

	responseWithStatus(c, http.StatusCreated, "refresh token created", "Created", nil)
}

// remove client & db jwt token; by user_id
func (h *Handler) Logout(c *gin.Context) {
	c.Writer.Header().Set("refresh_token", "")
	c.Writer.Header().Set("access_token", "")
	responseWithStatus(c, http.StatusOK, "Success logout, remove tokens", "OK", nil)
}
