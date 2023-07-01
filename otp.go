package rice

import (
	"github.com/pquerna/otp/totp"
)

const ISSUER = "rice"

func GenerateOTP(account string) (string, string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      ISSUER,
		AccountName: account,
	})
	if err != nil {
		return "", "", err
	}
	return key.Secret(), key.URL(), nil
}
