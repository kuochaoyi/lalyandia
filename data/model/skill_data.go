package model

type Skill struct {
	Id    SkillId
	Name  string
	Level int

	IsMagic          bool
	Operate          SkillOperate
	Target           SkillTargetType
	AffectScope      SkillAffectScope
	AffectLimit      SkillAffectLimit
	MagicLevel       int
	MpConsumeForCast int
	MpConsumeForUse  int
	CastRange        int
	EffectiveRange   int
	HitTime          float32
	CoolTime         float32
	HitCancelTime    float32
	ReuseDelay       float32
	EffectPoints     int
	NextAction       SkillNextAction

	Effect     string
	SelfEffect string

	//TODO
	//attribute
	//trait
	//multi_class
	//ride_state
}
