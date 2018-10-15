package io

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type Statistic struct {
	Id              bson.ObjectId `json:"id" bson:"_id"`
	Id_restaurant   bson.ObjectId `json:"id_restaurant" bson:"id_restaurant"`
	Date            string        `json:"date" bson:"date"`
	Sold_lunches    int           `json:"sold_lunches" bson:"sold_lunches"`
	Canceled_shifts int           `json:"canceled_shifts" bson:"canceled_shifts"`
	Av_time         float32       `json:"av_time" bson:"av_time"`
	Av_punctuation  float32       `json:"av_punctuation" bson:"av_punctuation"`
	Bonus_sold      int           `json:"bonus_sold" bson:"bonus_sold"`
	Student_sold    int           `json:"student_sold" bson:"student_sold"`
	External_sold   int           `json:"external_sold" bson:"external_sold"`
}

func (t Statistic) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
