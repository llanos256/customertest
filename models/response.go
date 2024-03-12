package models 

type Response struct {
   Load interface{} `json:"load"`
   Error string `json:"error"`
}