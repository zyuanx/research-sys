package models

import "time"

type ResearchList struct {
	CreatedAt   time.Time                `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time                `json:"updatedAt" bson:"updatedAt"`
	DeletedAt   time.Time                `json:"deletedAt" bson:"deletedAt"`
	Title       string                   `json:"title"`
	Desc        string                   `json:"desc"`
	FieldsValue map[string]interface{}   `json:"fieldsValue" bson:"fieldsValue"`
	Detail      []map[string]interface{} `json:"detail"`
	Rules       map[string]interface{}   `json:"rules"`
	Confirm     string                   `json:"confirm"`
	Status      int                      `json:"status"`
}
