package rice

import (
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenerateToken(t *testing.T) {
	_ = os.Setenv("JWT_SECRET_KEY", "7abbb1ad90ff9881d999")
	type args struct {
		claim MyCustomClaims
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "valid token",
			args: args{
				claim: MyCustomClaims{
					UserID:   "123",
					Username: "testuser",
					Email:    "",
					Role:     "admin",
					RegisteredClaims: jwt.RegisteredClaims{
						Issuer:    "testissuer",
						Subject:   "testsubject",
						Audience:  []string{"testaudience"},
						ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
						IssuedAt:  jwt.NewNumericDate(time.Now()),
						ID:        "testid",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateToken(tt.args.claim)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == "" {
				t.Errorf("GenerateToken() got = %v, want %v", got, tt.want)
				return
			}

			claim, err := ParseToken(got)
			if err != nil {
				t.Errorf("ParseToken() error = %v", err)
				return
			}
			if claim.UserID != tt.args.claim.UserID {
				t.Errorf("ParseToken() got = %v, want %v", claim.UserID, tt.args.claim.UserID)
			}
			if claim.Username != tt.args.claim.Username {
				t.Errorf("ParseToken() got = %v, want %v", claim.Username, tt.args.claim.Username)
			}
			if claim.Email != tt.args.claim.Email {
				t.Errorf("ParseToken() got = %v, want %v", claim.Email, tt.args.claim.Email)
			}
			if claim.Role != tt.args.claim.Role {
				t.Errorf("ParseToken() got = %v, want %v", claim.Role, tt.args.claim.Role)
			}
		})
	}
}
