package models

type TimeRequest struct {
	AccessToken int    `json:"access_token"`
	Time        string `json:"time"`
}
