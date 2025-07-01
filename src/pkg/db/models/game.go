package models

import "time"

type Player struct {
	Id                    int32 `gorm:"primaryKey"`
	AccountId             int32
	CreatedAt             time.Time
	UpdatedAt             time.Time
	Empire                int
	PlayerClass           int
	SkillGroup            int
	PlayTime              int
	Level                 int
	Experience            int
	Gold                  int
	St                    int
	Ht                    int
	Dx                    int
	Iq                    int
	PositionX             int
	PositionY             int
	Health                int
	Mana                  int
	Stamina               int
	BodyPart              int
	HairPart              int
	Name                  string
	GivenStatusPoints     int
	AvailableStatusPoints int
	Slot                  int

	Items []Item
}

type Item struct {
	Id              int32 `gorm:"primaryKey"`
	OwnerId         int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Window          int
	Position        int
	Count           int
	ProtoId         int
	Socket0         int
	Socket1         int
	Socket2         int
	AttributeType0  int
	AttributeValue0 int
	AttributeType1  int
	AttributeValue1 int
	AttributeType2  int
	AttributeValue2 int
	AttributeType3  int
	AttributeValue3 int
	AttributeType4  int
	AttributeValue4 int
	AttributeType5  int
	AttributeValue5 int
	AttributeType6  int
	AttributeValue6 int
}

type Players []Player
type Items []Item
