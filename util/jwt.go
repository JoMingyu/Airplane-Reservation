package util

import (
	"Airplane-Reservation/model"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
	mgo "gopkg.in/mgo.v2-unstable"
	"gopkg.in/mgo.v2-unstable/bson"
)

func generateToken(owner *model.User, tokenCollection *mgo.Collection, expire time.Duration) string {
	token := jwt.New(jwt.SigningMethodHS256)

	identity, _ := uuid.NewV4()
	identityStr := fmt.Sprintf("%s", identity)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = fmt.Sprintf("%s", identityStr)
	claims["exp"] = time.Now().Add(expire).Unix()

	t, _ := token.SignedString([]byte("secret"))

	tokenCollection.Remove(bson.M{"owner": *owner})
	tokenCollection.Insert(model.Token{
		Owner:    *owner,
		Identity: identityStr,
	})

	return t
}

func GenerateAccessToken(owner *model.User) string {
	return generateToken(owner, model.AccessTokenCol, time.Hour*24)
}

func GenerateRefreshToken(owner *model.User) string {
	return generateToken(owner, model.RefreshTokenCol, time.Hour*24*30)
}

func ExtractUserFromEchoContext(c echo.Context) *model.User {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	identity := claims["identity"]
	token := &model.Token{}

	if err := model.AccessTokenCol.Find(bson.M{"identity": identity}).One(token); err != nil {
		return nil
	}

	return &token.Owner
}
