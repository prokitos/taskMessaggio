package models

import (
	"errors"
)

func ResponseGood() error {
	return errors.New("request accepted")
}

func ResponseErrorAtServer() error {
	return errors.New("internal error")
}

func ResponseBadRequest() error {
	return errors.New("bad request")
}
