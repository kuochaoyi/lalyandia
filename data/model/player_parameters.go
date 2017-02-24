package model

type PlayerParameters struct {
	BasePhysicalAttack int
	BaseCritical       int
	BaseAttackType     WeaponType
	BaseAttackSpeed    int
	BaseDefendType     DefendType
	BaseMagicAttack    int
	BaseMagicDefend    DefendType
	BaseCanPenetrate   int
	BaseAttackRange    int
	BaseDamageRange    DamageRange
	BaseRandDamage     int

	WeightLevelBonus [PlayerMaxLevel]int

	StrBonus [99]int
	IntBonus [99]int
	ConBonus [99]int
	MenBonus [99]int
	DexBonus [99]int
	WitBonus [99]int

	OrgHpRegen [99]float32
	OrgMpRegen [99]float32
	OrgCpRegen [99]float32

	MoveSpeed          MoveSpeedType
	MoveModeSpeedBonus float32

	HpTable [PlayerMaxLevel]float32
	MpTable [PlayerMaxLevel]float32
	CpTable [PlayerMaxLevel]float32

	HitConditionBonus float32

	BreathBonus    int
	SafeFallHeight int
	CollisionBox   CollisionBox

	KarmaIncrease []float32

	//TODO
	//karma_increase_constant
}

type PlayerSpecializationParameters [PlayerSpecializationCount]PlayerParameters
