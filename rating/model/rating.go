package model

const RecordTypeMovie = RecordType("Movie")

type RecordId string
type RecordType string
type UserId string
type RatingValue int

type Rating struct {
	RecordId   string      `json:"recordId"`
	RecordType string      `json:"recordType"`
	UserId     UserId      `json:"userId"`
	Value      RatingValue `json:"value"`
}
