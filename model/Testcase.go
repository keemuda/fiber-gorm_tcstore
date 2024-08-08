package model

import (
	"time"
)

type Testcase struct {
	TestCaseID      uint    `gorm:"primaryKey;autoIncrement"`
	StoryID         string  `gorm:"not null;size:45"`
	Version         *string `gorm:"size:45"`
	ApplicationName string  `gorm:"not null;size:255"`
	Description     *string `gorm:"size:255"`
	Date            *time.Time
	FileName        string `gorm:"not null;size:255"`
}

//I'm not sure. Is this proper declare struct for gorm?
