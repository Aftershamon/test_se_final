package entity

import "time"

type Video struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `valid:"required~Name not blank"`
	Reason string `valid:"maxstringlength(5)~Max 5"`
	Email  string `valid:"email~ไม่ถูกจ้า"`
	// Start  time.Time `valid:"IsAfterAndPresent~ห้ามเป็นอดีต"`
	// End    time.Time `valid:"IsAfterStartOneDay~สิ้นสุด"`
}