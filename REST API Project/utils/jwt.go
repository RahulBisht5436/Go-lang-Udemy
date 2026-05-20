package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// jwtSecret is the symmetric key used to sign and later verify JWTs.
//
// IMPORTANT: this MUST be moved into an environment variable (e.g. JWT_SECRET
// in your .env file, read via os.Getenv) before this code ever runs in
// production. Anyone who learns this string can forge tokens that the server
// will happily accept as authentic.
var jwtSecret = []byte("supersecret-change-me")

// GenerateToken builds a signed JSON Web Token (JWT) for an authenticated user
// and returns it as a compact string the client can attach to future requests
// (typically in the "Authorization: Bearer <token>" header).
//
// A JWT is a compact, URL-safe string with three dot-separated parts:
//
//	<base64(header)>.<base64(payload)>.<base64(signature)>
//
// The header says which signing algorithm was used. The payload carries the
// "claims" (the data we want to ship to the client). The signature proves the
// server actually issued the token: as long as no one else knows jwtSecret,
// no one else can produce a signature that validates, so the server can
// trust any token whose signature checks out.
func GenerateToken(email string, userId int64) (string, error) {
	// Step 1 — Build the *unsigned* token in memory.
	// jwt.NewWithClaims picks the algorithm (which ends up in the header) and
	// attaches the claims (which become the payload). It does NOT sign yet —
	// signing happens in step 2 below.
	token := jwt.NewWithClaims(
		// HS256 = HMAC + SHA-256. It's a *symmetric* algorithm: the same
		// secret key (jwtSecret) is used both to sign tokens on the server
		// and to verify them on the server later. Good for single-service
		// setups like this one. If you ever need other services to verify
		// tokens without sharing your secret, switch to an asymmetric
		// algorithm like RS256 (RSA + SHA-256).
		jwt.SigningMethodHS256,

		// jwt.MapClaims is just a map[string]any whose entries become fields
		// in the JSON payload. Two important rules apply to anything you put
		// in here:
		//
		//   1. NEVER put secrets in claims — the payload is base64-encoded,
		//      NOT encrypted. Anyone holding the token can read every claim.
		//      So: no passwords, no API keys, no PII you wouldn't show the
		//      user themselves.
		//
		//   2. Prefer the *registered* claim names defined by RFC 7519
		//      (e.g. "exp", "iat", "sub", "iss", "aud") for standard concepts.
		//      The jwt library automatically validates these on parse — for
		//      example it will reject an expired token without you writing
		//      any check yourself.
		jwt.MapClaims{
			// Custom claims — the data we want available on every authed
			// request without doing another DB lookup.
			"email":  email,
			"userId": userId,

			// "exp" (expiration) is a registered claim. The library will
			// automatically reject any token whose exp is in the past.
			// Stored as a Unix epoch (seconds since 1970-01-01 UTC).
			// Here: 2 hours from issue time.
			"exp": time.Now().Add(time.Hour * 2).Unix(),
		},
	)

	// Step 2 — Sign the token with our secret.
	// SignedString computes the HMAC-SHA256 over <base64(header)>.<base64(payload)>
	// using jwtSecret as the key, base64-encodes the result, appends it as the
	// third segment, and returns the final dotted string. This is what gets
	// shipped to the client.
	signed, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("sign jwt: %w", err)
	}

	return signed, nil
}

func ExtractToken(token string) (string, error) {
	// Get Authorization header
	// token := context.GetHeader("Authorization")

	// Check if token exists
	if token == "" {
		return "", errors.New("Failed to Authenicate")
	}

	// Remove "Bearer " part
	tokenString := strings.TrimPrefix(token, "Bearer ")

	fmt.Println("Token:", tokenString)
	return tokenString, nil
}
func ExtractUserInfo(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		return []byte("your-secret-key"), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
