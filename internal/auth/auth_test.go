package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type args struct {
		headers http.Header
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr error
	}{
		{
			name: "No auth header",
			args: args{
				headers: http.Header{
				"Host":[]string{"example.com"},
				},
			},
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed auth header",
			args: args{
				headers: http.Header{
				"Host":[]string{"example.com"},
    			"Authorization":[]string{"ApiKeymock-token-12345"},
				},
			},
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name: "should work",
			args: args{
				headers: http.Header{
				"Host":[]string{"example.com"},
    			"Authorization":[]string{"ApiKey mock-token-12345"},
				},
			},
			want: "mock-token-12345",
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.args.headers)
			if (err != nil) && (err.Error() != tt.wantErr.Error()) {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
