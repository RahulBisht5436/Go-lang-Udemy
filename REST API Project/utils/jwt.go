package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// jwtSecret returns the symmetric key used to sign and verify JWTs.
//
// It is read lazily (not at package init) because main.go calls
// godotenv.Load() AFTER this package is imported. If we initialised the
// secret with `var jwtSecret = []byte(os.Getenv("JWT_SECRET"))`, the
// variable would be captured BEFORE the .env file was parsed and would
// silently be empty — every token would then be signed with an empty key.
func jwtSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

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
	signed, err := token.SignedString(jwtSecret())
	if err != nil {
		return "", fmt.Errorf("sign jwt: %w", err)
	}

	return signed, nil
}

// ExtractToken pulls the raw JWT string out of the value of an Authorization
// header. It accepts either "Bearer <token>" (the convention) or just the
// raw token, and returns an error if the header is empty.
func ExtractToken(token string) (string, error) {
	if token == "" {
		return "", errors.New("missing Authorization header")
	}
	return strings.TrimPrefix(token, "Bearer "), nil
}

// ExtractUserInfo parses a JWT, verifies its signature against jwtSecret(),
// confirms it's still valid (signature OK + not expired), and returns the
// custom claims we baked into it in GenerateToken: the user's email and
// numeric id. Any failure — bad signature, wrong algorithm, expired token,
// missing/typo'd claims — is surfaced as an error so the caller can return
// HTTP 401.
func ExtractUserInfo(tokenString string) (email string, userId int64, err error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Defence against the "alg=none" / algorithm-confusion family of
		// attacks: pin the expected signing family explicitly.
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return jwtSecret(), nil
	})
	if err != nil {
		return "", 0, err
	}
	if !token.Valid {
		return "", 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", 0, errors.New("invalid claims type")
	}

	email, ok = claims["email"].(string)
	if !ok {
		return "", 0, errors.New("email claim missing or not a string")
	}

	// IMPORTANT: numbers in MapClaims come out as float64 because they
	// were JSON-decoded. You stored userId as int64, but you must read
	// it back as float64 first and convert.
	uidFloat, ok := claims["userId"].(float64)
	if !ok {
		return "", 0, errors.New("userId claim missing or not a number")
	}
	userId = int64(uidFloat)

	return email, userId, nil
}
