package domain

type Plant struct {
	Id        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	CreatedAt int64  `json:"created_at" bson:"created_at"`
}
