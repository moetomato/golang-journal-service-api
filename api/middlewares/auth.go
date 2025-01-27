package middlewares

import (
	"context"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/moetomato/golang-journal-service-api/apperrors"
	"github.com/moetomato/golang-journal-service-api/common"
	"google.golang.org/api/idtoken"
)

var googleClientID = os.Getenv("GOOGLE_CLIENT_ID")

func AuthMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorization := req.Header.Get("Authorization")

		authHeaders := strings.Split(authorization, " ")
		if len(authHeaders) != 2 {
			apperrors.RequiredAuthorizationHeader.Wrap(errors.New("invalid header"), "no authorization header found")
		}

		bearer, token := authHeaders[0], authHeaders[1]

		if bearer != "Bearer" || token == "" {
			err := apperrors.RequiredAuthorizationHeader.Wrap(errors.New("authentication failed"), "invalid authorization header")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		tokenValidator, err := idtoken.NewValidator(context.Background())
		if err != nil {
			apperrors.MakeValidatorFailed.Wrap(err, "could not create token validator")
		}
		payload, err := tokenValidator.Validate(context.Background(), token, googleClientID)

		if err != nil {
			apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}

		name, ok := payload.Claims["name"]
		if !ok {
			err = apperrors.Unauthorizated.Wrap(err, "invalid id token")
			apperrors.ErrorHandler(w, req, err)
			return
		}
		req = common.SetUserName(req, name.(string))
		h.ServeHTTP(w, req)
	})
}
