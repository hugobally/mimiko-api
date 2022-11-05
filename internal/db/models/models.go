package models

import (
	"gorm.io/gorm"
)

// TODO Constraints (not null?, cascade/set null)

type User struct {
	gorm.Model
	Username *string
	Maps     []Map `gorm:"foreignKey:AuthorID"`
	Admin    bool  `gorm:"default:false"`
}

type Map struct {
	gorm.Model
	Title      string
	FlagshipID string // Spotify Track ID [Not a foreign key]
	Public     bool

	AuthorID uint

	Knots []Knot
	Links []Link
}

type Knot struct {
	gorm.Model
	TrackID     string // Spotify Track ID [Not a foreign key]
	Level       int
	Visited     *bool // TODO Remove
	MapID       uint
	ChildLinks  []Link `gorm:"foreignKey:SourceID"`
	ParentLinks []Link `gorm:"foreignKey:TargetID"`
}

type Link struct {
	gorm.Model
	ID       uint
	SourceID uint `gqlgen:"source"`
	TargetID uint `gqlgen:"target"`
	MapID    uint
}

//type LinkedAccount struct {
//	gorm.Model
//	RemoteID string
//	Kind     string
//	UserID   uint
//}
