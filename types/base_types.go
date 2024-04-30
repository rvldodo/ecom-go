package types

import "time"

type Timestamp struct {
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;type:timestamp;not null;autoCreateTime"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp ON UPDATE CURRENT_TIMESTAMP;null;autoUpdateTime"`
}
