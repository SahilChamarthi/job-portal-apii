package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"project/internal/auth"
	"project/internal/middlewear"
	"project/internal/model"
	"project/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

type handler struct {
	a *auth.Auth
	// 	us services.UsersService
	// 	cs services.CompanyService
	r services.AllinServices
}

func NewHandler(a *auth.Auth, r services.AllinServices) (*handler, error) {
	if r == nil {
		return nil, errors.New("service implementation not given")
	}

	return &handler{a: a, r: r}, nil

}
func (h *handler) userSignin(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in userSignin handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	var userCreate model.UserSignup
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&userCreate)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(&userCreate)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid input"})
		return
	}
	us, err := h.r.UserSignup(userCreate)
	if err != nil {
		log.Error().Err(err).Str("Trace Id", traceId).Msg("user signup problem")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "user signup failed"})
		return
	}
	c.JSON(http.StatusOK, us)

}

func (h *handler) userLoginin(c *gin.Context) {
	ctx := c.Request.Context()

	traceId, ok := ctx.Value(middlewear.TraceIdKey).(string)
	if !ok {
		log.Error().Str("traceId", traceId).Msg("trace id not found in userSignin handler")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}
	var userLogin model.UserLogin
	body := c.Request.Body
	err := json.NewDecoder(body).Decode(&userLogin)
	if err != nil {
		log.Error().Err(err).Msg("error in decoding")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return
	}

	validate := validator.New()
	err = validate.Struct(&userLogin)
	if err != nil {
		log.Error().Err(err).Msg("error in validating ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid input"})
		return
	}

	regClaims, err := h.r.UserLogin(userLogin)
	if err != nil {
		log.Error().Err(err).Msg("error in Loginin ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": "invalid input"})
		return
	}

	token, err := h.a.GenerateToken(regClaims)
	if err != nil {
		log.Error().Err(err).Msg("error in Gneerating toek ")
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"msg": http.StatusText(http.StatusInternalServerError)})
		return

	}

	c.JSON(http.StatusOK, token)

}
