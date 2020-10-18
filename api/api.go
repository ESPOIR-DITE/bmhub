package api

import (
	"bmhub/config"
	"gopkg.in/resty.v1"
)

const Base_URL = ""

func Rest() *resty.Request {

	return resty.R().SetAuthToken("").
		SetHeader("Accept", "application/json").
		SetHeader("email", "email").
		SetHeader("site", "site").
		SetHeader("Content-Type", "application/json")
}

var JSON = config.ConfigWithCustomTimeFormat
