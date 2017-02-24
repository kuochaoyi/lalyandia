package model

import (
	"errors"
	"time"
)

type AttackType int

type DefendType struct {
	UpperBody int
	LowerBody int
	Helmet    int
	Boots     int
	Gloves    int
	Underwear int
}

type MoveSpeedType struct {
	GroundMin     int
	GroundMax     int
	UnderWaterMin int
	UnderWaterMax int
	FlyingMin     int
	FlyingMax     int
	FloatingMin   int
	FloatingMax   int
}

type DamageRange struct {
	X int
	Y int
	Width int
	Height int
}

type CollisionBox struct {
	Width  int
	Height int
}

type Range int

const (
	OneToTenLevelRange Range = iota
	ElevenToTwentyLevelRange
	TwentyOneToThirtyLevelRange
	ThirtyOneToFortyLevelRange
	FortyOneToFiftyLevelRange
	FiftyOneToSixtyLevelRange
	SixtyOneToSeventyLevelRange

	UnknownLevelRange
)

type ValueWithCondition [7]float32

func (v *ValueWithCondition) GetRange(level int) Range {
	result := UnknownLevelRange
	switch level {
	case level > 0 && level <= 10:
		result = OneToTenLevelRange
	case level > 10 && level <= 20:
		result = ElevenToTwentyLevelRange
	case level > 20 && level <= 30:
		result = TwentyOneToThirtyLevelRange
	case level > 30 && level <= 40:
		result = ThirtyOneToFortyLevelRange
	case level > 40 && level <= 50:
		result = FortyOneToFiftyLevelRange
	case level > 50 && level <= 60:
		result = FiftyOneToSixtyLevelRange
	case level > 60 && level <= 70:
		result = SixtyOneToSeventyLevelRange
	}

	return result
}

func (v *ValueWithCondition) CheckAndGet(levelRange Range) (float32, error) {
	if int(levelRange) > len(*v) {
		return 0, errors.New("")
	}

	return v[levelRange], nil
}

type PlayerRace int
type PlayerSpecializationType int
type PlayerSpecialization int

/****************************************/

type ItemId int
type Type int
type Slot int
type ArmorType int
type WeaponType int
type EtcItemType int
type CrystalType int
type ItemAction int
type ConsumeType int
type MaterialType int

/****************************************/

type NpcId int
type NpcClass int
type NpcSex int

type TerritoryPoint struct {
	X    int
	Y    int
	MinZ int
	MaxZ int
}

type Position struct {
	X int
	Y int
	Z int
}

type CharacterPosition struct {
	Position Position
	Heading  int
}

type SpawnDefine struct {
	NpcName               string
	Position              CharacterPosition
	Amount                int
	RespawnDuration       time.Duration
	RespawnRandomDuration time.Duration

	//TODO Think about saving into database ?
}

type NpcMaker struct {
	Name         string
	SpawnDefines []SpawnDefine
	//TODO: Think about struct
	SpawnCondition string
	MaxAmount      int
	AI             string
}

type Territory struct {
	Points    []TerritoryPoint
	NpcMakers []NpcMaker
}

