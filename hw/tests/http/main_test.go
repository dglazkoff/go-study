//package main
//
//import (
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/require"
//	"io"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestUserViewHandler(t *testing.T) {
//
//	tests := []struct {
//		name     string
//		code     int
//		userId   string
//		response string
//	}{
//		{
//			name:     "simple test",
//			code:     http.StatusOK,
//			userId:   "u1",
//			response: `{"ID":"u1","FirstName":"Misha","LastName":"Popov"}`,
//		},
//	}
//
//	handler := UserViewHandler(map[string]User{
//		"u1": {
//			ID:        "u1",
//			FirstName: "Misha",
//			LastName:  "Popov",
//		},
//		"u2": {
//			ID:        "u2",
//			FirstName: "Sasha",
//			LastName:  "Popov",
//		},
//	})
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			req := httptest.NewRequest(http.MethodGet, "/users"+"?user_id="+tt.userId, nil)
//			w := httptest.NewRecorder()
//
//			handler(w, req)
//			res := w.Result()
//
//			assert.Equal(t, tt.code, res.StatusCode)
//			defer res.Body.Close()
//			resBody, err := io.ReadAll(res.Body)
//
//			require.NoError(t, err)
//			assert.JSONEq(t, tt.response, string(resBody))
//			assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
//		})
//	}
//}

package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserViewHandler(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
		user        User
	}
	tests := []struct {
		name    string
		request string
		users   map[string]User
		want    want
	}{
		{
			name: "simple test #1",
			users: map[string]User{
				"id1": {
					ID:        "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			want: want{
				contentType: "application/json",
				statusCode:  200,
				user: User{ID: "id1",
					FirstName: "Misha",
					LastName:  "Popov",
				},
			},
			request: "/users?user_id=id1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, tt.request, nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(UserViewHandler(tt.users))
			h(w, request)

			result := w.Result()

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

			userResult, err := ioutil.ReadAll(result.Body)
			require.NoError(t, err)
			err = result.Body.Close()
			require.NoError(t, err)

			var user User
			err = json.Unmarshal(userResult, &user)
			require.NoError(t, err)

			assert.Equal(t, tt.want.user, user)
		})
	}
}