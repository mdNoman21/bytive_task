package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectInfo struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	ProjectID     string             `json:"project_id"`
	Date          CustomDate         `json:"date" bson:"date" form:"date" validate:"required"`
	Start_Time    CustomTime         `bson:"start_time" json:"start_time" form:"start_time" validate:"required"`
	End_Time      CustomTime         `bson:"end_time" json:"end_time" form:"end_time" validate:"required"`
	Description   string             `json:"description" bson:"description" form:"description" validate:"required"`
	Project       string             `json:"project" bson:"project" form:"project" validate:"required"`
	Who           string             `json:"user_name" bson:"user_name"`
	TaskList      []TaskType         `bson:"tasklist"`
	PaymentStatus Status             `bson:"payment_status"`
	BillableTime  string             `bson:"billable_time"`
	Hours         float64            `bson:"hours"`
	Time_Spent    TimeSpent          `bson:"time_spent" form:"time_spent"`
}

type CustomDate struct {
	time.Time
}

const customDateFormat = "2006-01-02"

func (cd *CustomDate) UnmarshalText(data []byte) error {
	t, err := time.Parse(customDateFormat, string(data))
	if err != nil {
		return err
	}
	cd.Time = t
	return nil
}

func (cd CustomDate) MarshalText() ([]byte, error) {
	return []byte(cd.Time.Format(customDateFormat)), nil
}

type CustomTime struct {
	time.Time
}

const customTimeFormat = "3:04pm"

func (ct *CustomTime) UnmarshalText(data []byte) error {
	t, err := time.Parse(customTimeFormat, string(data))
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

func (ct CustomTime) MarshalText() ([]byte, error) {
	return []byte(ct.Time.Format(customTimeFormat)), nil
}

type Status struct {
	Billable bool
	Billed   bool
}

type TaskType struct {
	Type string `json:"type" bson:"type"`
}

type TimeSpent struct {
	Hours   int `form:"hours"`
	Minutes int `form:"minutes"`
}
