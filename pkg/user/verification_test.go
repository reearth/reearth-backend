package user

import (
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewVerification(t *testing.T) {

	type fields struct {
		verified   bool
		code       bool
		expiration bool
	}

	tests := []struct {
		name string
		want fields
	}{
		{
			name: "init verification struct",

			want: fields{
				verified:   false,
				code:       true,
				expiration: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewVerification()
			assert.Equal(t, tt.want.verified, got.IsVerified())
			assert.Equal(t, tt.want.code, len(got.Code()) > 0)
			assert.Equal(t, tt.want.expiration, !got.Expiration().IsZero())
		})
	}
}

func TestVerification_Code(t *testing.T) {
	tests := []struct {
		name         string
		verification *Verification
		want         string
	}{
		{
			name: "should return a code string",
			verification: &Verification{
				verified:   false,
				code:       "xxx",
				expiration: time.Time{},
			},
			want: "xxx",
		},
		{
			name: "should return a empty string",
			want: "",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			assert.Equal(tt, tc.want, tc.verification.Code())
		})
	}
}

func TestVerification_Expiration(t *testing.T) {
	e := time.Now()

	tests := []struct {
		name         string
		verification *Verification
		want         time.Time
	}{
		{
			name: "should return now date",
			verification: &Verification{
				verified:   false,
				code:       "",
				expiration: e,
			},
			want: e,
		},
		{
			name:         "should return zero time",
			verification: nil,
			want:         time.Time{},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.want, tc.verification.Expiration())
		})
	}
}

func TestVerification_IsExpired(t *testing.T) {
	type fields struct {
		verified   bool
		code       string
		expiration time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Verification{
				verified:   tt.fields.verified,
				code:       tt.fields.code,
				expiration: tt.fields.expiration,
			}
			if got := v.IsExpired(); got != tt.want {
				t.Errorf("IsExpired() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerification_IsVerified(t *testing.T) {
	tests := []struct {
		name         string
		verification *Verification
		want         bool
	}{
		{
			name: "should return true",
			verification: &Verification{
				verified: true,
			},
			want: true,
		},
		{
			name:         "should return false",
			verification: nil,
			want:         false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.want, tc.verification.IsVerified())
		})
	}
}

func TestVerification_SetVerified(t *testing.T) {
	tests := []struct {
		name         string
		verification *Verification
		input        bool
		want         bool
	}{
		{
			name: "should set true",
			verification: &Verification{
				verified: false,
			},
			input: true,
			want:  true,
		},
		{
			name:         "should return false",
			verification: nil,
			want:         false,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			tc.verification.SetVerified(tc.input)
			assert.Equal(tt, tc.want, tc.verification.IsVerified())
		})
	}
}

func Test_generateCode(t *testing.T) {
	var regx = regexp.MustCompile(`[a-zA-Z0-9]{5}`)
	str := generateCode()

	tests := []struct {
		name string
		want string
	}{
		{
			name: "should generate a valid code",
			want: str,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := regx.FindString(str)
			assert.Equal(t, tt.want, got)
		})
	}
}
