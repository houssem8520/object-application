package entity

type Object struct {
	ObjectID int64  `gorm:"primary_key;column:ObjectID;type:INTEGER;"`
	Payload  string `gorm:"column:Payload;type:TEXT;"`
	Bucket   string `gorm:"column:Bucket;type:VARCHAR;size:255;"`
}
