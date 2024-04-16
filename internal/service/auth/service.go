package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	GenerateToken(userID int) (string, error)
	// ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewAuthService() *jwtService {
	return &jwtService{}
}

// var SECRET_KEY = os.Getenv("SECRET_KEY")
var SECRET_KEY = "SECRET_KEY_OM"

func (s *jwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	// claim["exp"] = time.Now().Add(time.Duration(1) * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
