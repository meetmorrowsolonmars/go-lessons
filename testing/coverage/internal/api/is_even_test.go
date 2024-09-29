package api_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/meetmorrowsolonmars/go-lessons/testing/coverage/internal/api"
)

func TestIsEvenHandleFunc(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(api.IsEvenHandleFunc))
	defer server.Close()

	type args struct {
		req *http.Request
	}

	tests := []struct {
		name string
		argv func() args
		want api.IsEvenResponse
	}{
		{
			name: "odd number",
			argv: func() args {
				body := strings.NewReader("{ \"number\": 11 }")

				req, err := http.NewRequest(
					http.MethodPost,
					server.URL,
					io.NopCloser(body),
				)
				require.NoError(t, err)

				return args{
					req: req,
				}
			},
			want: api.IsEvenResponse{
				IsEven: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			argv := tt.argv()

			res, err := server.Client().Do(argv.req)
			require.NoError(t, err)

			body := api.IsEvenResponse{}

			err = json.NewDecoder(res.Body).Decode(&body)
			require.NoError(t, err)

			assert.Equal(t, http.StatusOK, res.StatusCode)
			assert.Equal(t, tt.want, body)
		})
	}
}
