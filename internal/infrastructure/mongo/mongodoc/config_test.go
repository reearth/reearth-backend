package mongodoc

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		c config.Config
	}
	tests := []struct {
		name string
		args args
		want ConfigDocument
	}{
		{
			name: "new config",
			args: args{
				c: config.Config{
					Migration: 12345,
					Auth: &config.Auth{
						Cert: "cert",
						Key:  "key",
					},
				},
			},
			want: ConfigDocument{
				Migration: 12345,
				Auth: &Auth{
					Cert: "cert",
					Key:  "key",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, NewConfig(tt.args.c))
		})
	}
}

func TestConfigDocument_Model(t *testing.T) {
	tests := []struct {
		name   string
		target *ConfigDocument
		want   *config.Config
	}{
		{
			name: "config model",
			target: &ConfigDocument{
				Migration: 12345,
				Auth: &Auth{
					Cert: "cert",
					Key:  "key",
				},
			},
			want: &config.Config{
				Migration: 12345,
				Auth: &config.Auth{
					Cert: "cert",
					Key:  "key",
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := &ConfigDocument{
				Migration: tt.target.Migration,
				Auth:      tt.target.Auth,
			}
			assert.Equal(t, tt.want, c.Model())
		})
	}
}
