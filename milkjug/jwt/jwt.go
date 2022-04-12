package jwt

import (
	"log"
	"milkshake/helper"
	"milkshake/redis"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// expMins = 0 use default time in conf
func GenerateTokenMapClaims(expMins int) *jwt.MapClaims {
	claims := make(jwt.MapClaims)
	if expMins == 0 {
		expMins = helper.GetJwtTimeOut()
	}
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expMins)).Unix()
	claims["iat"] = time.Now().Unix()

	return &claims
}

func GenerateTokenString(claims *jwt.MapClaims) (string, error) {
	// log.Println(claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(helper.GetJwtKey()))
}

func LazyTokenString(content map[string]interface{}, expMins int) (string, error) {
	claims := make(jwt.MapClaims)
	if expMins == 0 {
		expMins = helper.GetJwtTimeOut()
	}
	for k, v := range content {
		claims[k] = v
	}
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expMins)).Unix()
	// log.Println(claims["exp"])
	claims["iat"] = time.Now().Unix()
	return GenerateTokenString(&claims)
}

// Check whether token is valid
func CheckToken(token string) (bool, *jwt.Token) {
	t, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		// check token signing method etc
		return []byte(helper.GetJwtKey()), nil
	})
	if err != nil {
		log.Println("CheckToken Convert to jwt claims fail.", err)
		return false, nil
	}
	return true, t
}

func ReadToken(token string) (*jwt.Token, *jwt.MapClaims, bool) {
	claims := jwt.MapClaims{}
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(helper.GetJwtKey()), nil
	})
	if err != nil {
		log.Println("Read Token Convert to jwt claims fail.", err)
		return nil, nil, false
	}
	return t, &claims, true
}

func RenewToken(old string, tokenExpTime int, removeTokenTime time.Duration) (string, error) {
	new, err := GenerateTokenString(GenerateTokenMapClaims(tokenExpTime))
	// log.Println("new key is same as old key", old == new)
	// Sometime the new key is same as the old key??
	for new == old {
		new, err = GenerateTokenString(GenerateTokenMapClaims(int(tokenExpTime)))
	}
	if err == nil {
		if err = redis.SwapKey(old, new, removeTokenTime); err == nil {
			return new, nil
		}
	}
	log.Println(err)
	return "", err
}
