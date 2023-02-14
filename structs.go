package maildis

type Recipent struct {
	Email string `json:"email" bson:"email" db:"email"`
}
