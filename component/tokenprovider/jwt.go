package tokenprovider

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims 
}
func NewUserClaims(userID uuid.UUID,duration time.Duration)(*Claims,error){
	tokenId,err:=uuid.NewRandom();
	if err!=nil{
		return nil,fmt.Errorf("error generating token id: %w",err)
	}
	return &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID: tokenId.String(),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	},nil
}
// GenerateToken generates a new JWT token with claims and expiration time
func GenerateToken(secretKey string, userID uuid.UUID, expiration time.Duration) (string, error) {
	claims,err := NewUserClaims(userID,expiration)
	if err!=nil{
		return "",err
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenStr,err:=token.SignedString([]byte(""))
	if err!=nil{
		return "",fmt.Errorf("error signing token: %w", err)
	}
	return tokenStr,nil
}

// ParseToken parses and validates the JWT token and returns the claims
func ParseToken(tokenString string, secretKey string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		_,ok:=token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrTokenInvalidClaims
}