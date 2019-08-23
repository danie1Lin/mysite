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
)

func httpGetSessionData(sessionStore cache.SessionStore) httptransport.RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		cookie, err := r.Cookie(sessionStore.HeaderKey())
		sessionID := ""
		if err != nil && err != http.ErrNoCookie {
			return ctx
		} else if err == nil {
			sessionID = cookie.Value
		}

		if sessionID == "" {
			sessionID = r.Header.Get(sessionStore.HeaderKey())
			if sessionID == "" {
				return ctx
			}
		}

		data, err := sessionStore.GetData(sessionID)
		if err != nil {
			return ctx
		}
		ctx = context.WithValue(ctx, sessionStore.ContextKey(), sessionID)
		ctx = context.WithValue(ctx, "session-data", data)
		return ctx
	}
}

func httpAssignSession(sessionStore cache.SessionStore) httptransport.ServerResponseFunc {
	return func(ctx context.Context, w http.ResponseWriter) context.Context {
		var err error
		sessionID, ok := ctx.Value(sessionStore.ContextKey()).(string)
		if !ok {
			sessionID, err = sessionStore.New(map[string]interface{}{"test": "Hey"})
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

	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
		httptransport.ServerErrorLogger(errorLogger),
		httptransport.ServerBefore(httpGetSessionData(sessionStore)),
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

	m := http.NewServeMux()
	m.Handle("/login", LoginHTTPHandler)
	m.Handle("/sign_up", SignUpHTTPHandler)

	return m
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

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if response == nil {
		return nil
	}

	r := response.(endpoints.Response)
	data := map[string]string{
		"error": r.Error(),
		"code":  r.Code(),
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
	w.Write([]byte(err.Error()))
}
