package apperrors

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/moetomato/golang-journal-service-api/common"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err error) {
	var appErr *JournalAppError
	if !errors.As(err, &appErr) {
		appErr = &JournalAppError{
			ErrCode: Unknown,
			Message: "internal process failed",
			Err:     err,
		}
	}

	traceID := common.GetTraceID(req.Context())
	log.Printf("TraceID : %d, Error: %s\n", traceID, appErr)

	var statusCode int

	switch appErr.ErrCode {
	case NAData:
		statusCode = http.StatusNotFound
	case NoTargetData, ReqBodyDecodeFailed, BadParam:
		statusCode = http.StatusBadRequest
	case RequiredAuthorizationHeader, Unauthorizated:
		statusCode = http.StatusUnauthorized
	case UserUnmatched:
		statusCode = http.StatusForbidden
	default:
		statusCode = http.StatusInternalServerError
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(appErr)
}
