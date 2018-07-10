package controllers

import (
	"Airplane-Reservation/model"
	"Airplane-Reservation/util"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2-unstable/bson"
)

func checkIDIsDuplicated(id string) (bool, error) {
	count, err := model.UserCol.Find(bson.M{"_id": id}).Count()
	if err != nil {
		return true, err
	}

	if count != 0 {
		return true, nil
	}

	return false, nil
}

func IDDuplicationCheck(c echo.Context) error {
	id := c.Param("id")

	if isIDDuplicated, err := checkIDIsDuplicated(id); err != nil {
		return c.NoContent(http.StatusBadRequest)
	} else if isIDDuplicated {
		return c.NoContent(http.StatusConflict)
	} else {
		return c.NoContent(http.StatusOK)
	}

}

func Signup(c echo.Context) error {
	type binder struct {
		ID   string `json:"id"`
		Pw   string `json:"pw"`
		Name string `json:"name"`
	}

	payload := &binder{}

	if err := c.Bind(payload); err != nil {
		return err
	}

	if isIDDuplicated, err := checkIDIsDuplicated(payload.ID); err != nil {
		return c.NoContent(http.StatusBadRequest)
	} else if isIDDuplicated {
		return c.NoContent(http.StatusConflict)
	} else {
		model.UserCol.Insert(model.User{
			ID:   payload.ID,
			Pw:   payload.Pw,
			Name: payload.Name,
		})

		return c.NoContent(http.StatusOK)
	}
}

