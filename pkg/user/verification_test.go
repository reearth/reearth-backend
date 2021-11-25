package user

import (
	"reflect"
	"testing"
	"time"
)

func TestNewVerification(t *testing.T) {
	tests := []struct {
		name string
		want *Verification
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVerification(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVerification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerification_Code(t *testing.T) {
	type fields struct {
		verified   bool
		code       string
		expiration time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   string
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
			if got := v.Code(); got != tt.want {
				t.Errorf("Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerification_Expiration(t *testing.T) {
	type fields struct {
		verified   bool
		code       string
		expiration time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
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
			if got := v.Expiration(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expiration() = %v, want %v", got, tt.want)
			}
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
			if got := v.IsVerified(); got != tt.want {
				t.Errorf("IsVerified() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerification_SetVerified(t *testing.T) {
	type fields struct {
		verified   bool
		code       string
		expiration time.Time
	}
	type args struct {
		b bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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
		})
	}
}

func Test_generateCode(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateCode(); got != tt.want {
				t.Errorf("generateCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
