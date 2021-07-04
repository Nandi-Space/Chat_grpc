package auth

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Claims represents the authorization claims transmitted via a JWT
type Claims struct {
	jwt.StandardClaims
}

// NewClaims constructs a Claim value for the identified user.
/* The claims expire within a specifie duration of the provided time.
Additional fields of the Claims can be set after calling NewClaims is desired
*/
func NewClaims(userID string, now time.Time, expires time.Duration) Claims {
	c := Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userID,
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(expires).Unix(),
		},
	}

	return c
}
