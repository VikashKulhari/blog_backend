package entities

import "github.com/lib/pq"

type Images struct {
	Images pq.StringArray `gorm:"type:text[]"`
}
