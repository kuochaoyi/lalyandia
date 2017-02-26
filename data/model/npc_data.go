package model

type NpcData struct {
	Id    NpcId
	Class NpcClass

	Name                string
	Sex                 NpcSex
	Level               int
	BaseAttackType      AttackType
	BaseCanPenetrate    bool
	BaseAttackRange     int
	BaseDamageRange     DamageRange
	BaseRandomDamage    int
	BaseDefence         float32
	BaseMagicAttack     float32
	BaseMagicDefence    float32
	BasePhysicalAttack  float32
	PhysicalHitModify   float32
	PhysicalAvoidModify int
	BaseAttackSpeed     int
	BaseReuseDelay      int
	BaseCritical        int
	MagicUseSpeedModify int

	Undying       bool
	CanBeAttacked bool

	Strength     int
	Intelligence int
	Dexterity    int
	Wisdom       int
	Constitution int
	Mental       int

	OriginalHp      float32
	OriginalHpRegen float32
	HpIncrease      int

	OriginalMp      float32
	OriginalMpRegen float32
	MpIncrease      int

	SlotChest     string
	SlotRightHand string
	SlotLeftHand  string

	Experience        int
	AcquireExpRate    float32
	AcquireSkillPoint int

	SkillList []string

	AggressionRange int
	ClanHelpRange   int
	//TODO create type
	Clan           []string
	IgnoreClanList []string

	SafeHeight      int
	CollisionRadius int
	CollisionHeight int

	CorpseTime         int
	ItemMakeList       []string
	CorpseMakeList     []string
	AdditionalMakeList []string

	AI string

	//TODO
	//magic_list
	//temper
	//category
	//org_jump
	//max_item
	//ground_low
	//ground_high
	//underwater_low
	//underwater_high
	//fly_low
	//fly_high
	//float_low
	//float_high
	//org_earshot
	//sight
	//status
	//align
	//guild
	//pledge
	//pledge_leader
	//alliance
	//alliance_leader
	//ruler
	//ruling_domain
}
