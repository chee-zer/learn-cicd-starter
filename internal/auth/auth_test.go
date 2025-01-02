package auth

import (
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
		wantErr bool
	}{
		{
			name: "No auth header",
			args: args{
				headers: http.Header{
				"Host":[]string{"example.com"},
				},
			},
			wantErr: true,
		},
		{
			name: "malformed auth header",
			args: args{
				headers: http.Header{
				"Host":[]string{"example.com"},
    			"Authorization":[]string{"ApiKeymock-token-12345"},
				},
			},
			wantErr: true,
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.args.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
