package main

type Data struct {
	ID    primitive.ObjectID `json:"id,omitemoty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"namne"`
	Email string             `json:"email" bson:"email"`
}

type Data2 struct {
	ID    string `json:"id,omitemoty' bson:"_id,omitempty"`
	Name  string `json:"name " bson:"name"`
	Email string `json:"email" bson:"email"`
}
