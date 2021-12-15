package models

type APIResponse struct {
	Data  []Deployment `json:"data" bson:"data"`
	Error interface{} `json:"error" bson:"error"`
}