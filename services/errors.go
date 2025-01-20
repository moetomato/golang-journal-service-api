package services

import "errors"

var ErrNoData = errors.New("got 0 record from db.Query")
