package model

import "time"

type UserRole string

const (
	Student UserRole = "STUDENT"
	Teacher UserRole = "TEACHER"
)

type User struct {
	ID        string    `gorm:"primary_key;column:id;not null" json:"id"`
	NIM       string    `gorm:"column:nim;not null;uniqueIndex" json:"nim"`
	Role      UserRole  `gorm:"column:role;not null"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"not null;uniqueIndex" json:"email"`
	PhotoUrl  string    `gorm:"not null" json:"photo_url"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updated_at"`

	UserCompetenceReports []UserCompetenceReports `gorm:"foreignKey:user_id;references:id"`
	DiagnosticReports     []DiagnosticReport      `gorm:"many2many:user_diagnostic_reports;foreignKey:id;joinForeignKey:user_id;references:name;joinReferences:diagnostic_report_id"`
	UserActivity          []Activity              `gorm:"many2many:user_activities;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:activities_id"`
}

func (u *User) TableName() string {
	return "users"
}
