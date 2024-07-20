package entities

import "github.com/lib/pq"

type Tags struct {
	Tags pq.StringArray `gorm:"type:text[]"`
}
