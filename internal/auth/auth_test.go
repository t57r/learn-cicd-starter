package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		title          string
		headers        http.Header
		expectedAPIKey string
		hasErr         bool
	}{
		{
			title:          "No authorization header",
			headers:        http.Header{},
			expectedAPIKey: "",
			hasErr:         true,
		},
		{
			title:          "Auth header single key error",
			headers:        http.Header{"Authorization": []string{"single_key"}},
			expectedAPIKey: "",
			hasErr:         true,
		},
		{
			title:          "Auth header two words error",
			headers:        http.Header{"Authorization": []string{"Bearer api_key"}},
			expectedAPIKey: "",
			hasErr:         true,
		},
		{
			title:          "Valid api key",
			headers:        http.Header{"Authorization": []string{"ApiKey users_api_token"}},
			expectedAPIKey: "users_api_token",
			hasErr:         false,
		},
	}

	for _, c := range cases {
		apiKey, err := GetAPIKey(c.headers)
		if err != nil != c.hasErr {
			t.Fatalf("%s: failed hasErr expecting %t", c.title, c.hasErr)
		}
		if apiKey != c.expectedAPIKey {
			t.Fatalf("%s: failed expectedAPIKey expecting %s", c.title, c.expectedAPIKey)
		}
	}
}

// func GetAPIKey(headers http.Header) (string, error) {
// 	authHeader := headers.Get("Authorization")
// 	if authHeader == "" {
// 		return "", ErrNoAuthHeaderIncluded
// 	}
// 	splitAuth := strings.Split(authHeader, " ")
// 	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
// 		return "", errors.New("malformed authorization header")
// 	}
//
// 	return splitAuth[1], nil
// }
