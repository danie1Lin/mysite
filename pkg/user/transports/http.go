package transports

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mysite/pkg/user/endpoints"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"

	"mysite/pkg/cache"

	jsonpatch "github.com/evanphx/json-patch"

	e "mysite/pkg/errors"
)

func httpGetSessionData(sessionStore cache.SessionStore, logger log.Logger) httptransport.RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		cookie, err := r.Cookie(sessionStore.HeaderKey())
		logger = log.With(logger, "before server", "httpGetSessionData")
		sessionID := ""
		if err != nil && err != http.ErrNoCookie {
			return ctx
		} else if err == nil {
			sessionID = cookie.Value
		}
		newSession := false
		if sessionID == "" {
			sessionID = r.Header.Get(sessionStore.HeaderKey())
			if sessionID == "" {
				newSession = true
				sessionID, err = sessionStore.New(map[string]interface{}{"user_id": ""})
				if err != nil {
					logger.Log("err", err)
					return nil
				}
			}
		}

		ctx = context.WithValue(ctx, sessionStore.ContextKey(), sessionID)

		if !newSession {
			data, err := sessionStore.GetData(sessionID)
			if err != nil {
				logger.Log("err", err)
				return ctx
			}
			ctx = context.WithValue(ctx, "session-data", data)
		}

		//ctx = context.WithValue(ctx, "session-data", data)
		return ctx
	}
}

func httpAssignSession(sessionStore cache.SessionStore) httptransport.ServerResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter) context.Context {
		var err error
		sessionID, ok := ctx.Value(sessionStore.ContextKey()).(string)
		if !ok {
			sessionID, err = sessionStore.New(map[string]interface{}{"user": nil})
			if err != nil {

				return ctx
			}
		}
		w.Header().Set(sessionStore.HeaderKey(), sessionID)
		cookie := &http.Cookie{
			Name:  sessionStore.HeaderKey(),
			Value: sessionID,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
		return ctx
	}
}

func NewHTTPHandler(ctx context.Context, endpoints endpoints.EndPoints, logger log.Logger, sessionStore cache.SessionStore) http.Handler {
	errorLogger := level.Error(logger)
	errorLogger = log.With(errorLogger, "logger", "error logger")
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
		httptransport.ServerErrorLogger(errorLogger),
		httptransport.ServerBefore(httpGetSessionData(sessionStore, errorLogger)),
		httptransport.ServerAfter(
			httpAssignSession(sessionStore),
		),
	}

	LoginHTTPHandler := httptransport.NewServer(
		endpoints.LoginEndPoint,
		decodeLoginRequest,
		encodeResponse,
		options...,
	)

	SignUpHTTPHandler := httptransport.NewServer(
		endpoints.SignUpEndPoint,
		decodeSignUpRequest,
		encodeResponse,
		options...,
	)

	GetProfileHandler := httptransport.NewServer(
		endpoints.GetProfileEndPoint,
		decodeGetProfileRequest,
		encodeResponse,
		options...,
	)

	m := http.NewServeMux()
	m.Handle("/login", LoginHTTPHandler)
	m.Handle("/sign_up", SignUpHTTPHandler)
	m.Handle("/profile", GetProfileHandler)
	return m
}

func decodeGetProfileRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeLoginRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeSignUpRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

type GeneralResponseData struct {
	Error   string `json:"error,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if response == nil {
		return nil
	}

	r := response.(endpoints.Response)
	data := GeneralResponseData{
		r.Error(),
		r.Code(),
		r.Message(),
	}

	if r.StatusCode() != 0 {
		w.WriteHeader(r.StatusCode())
	}

	errorData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	res, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return err
	}
	patch, err := jsonpatch.CreateMergePatch([]byte("{}"), errorData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	result, err := jsonpatch.MergePatch(res, patch)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = io.WriteString(w, string(result))
	return err
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	switch v := err.(type) {
	case *e.Err:
		fmt.Println(v.Details())
	}
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
