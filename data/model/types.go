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
	X      int
	Y      int
	Width  int
	Height int
}

type CollisionBox struct {
	Width  int
	Height int
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

/****************************************/

type SkillId int
type SkillOperate int
type SkillTargetType int
type SkillAffectScope int
type SkillNextAction int

type SkillAffectLimit struct {
	Min int
	Max int
}
