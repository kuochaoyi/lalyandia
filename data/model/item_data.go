package model

type ItemData struct {
	Id          ItemId
	Type        Type
	Slot        Slot
	ArmorType   ArmorType
	WeaponType  WeaponType
	EtcItemType EtcItemType

	Blessed      bool
	Enchanted    bool
	Tradable     bool
	Dropable     bool
	Destructable bool
	CanPenetrate bool

	DefaultAction   ItemAction
	DefaultPrice    int
	Price           int
	Weight          int
	CrystalType     CrystalType
	CrystalCount    int
	SoulshotCount   int
	ConsumeType     ConsumeType
	InitialCount    int
	MaximumCount    int
	ImmediateEffect int
	MaterialType    MaterialType

	Durability int
	Damaged    int

	PhysicalDamage  int
	PhysicalDefense int
	RandomDamage    int
	Critical        int
	AttackSpeed     int

	AttackRange    int
	EffectiveRange int
	DamageRange    DamageRange

	AvoidModify int
	HitModify   float32

	MagicalDamage  int
	MagicalDefense int

	ShieldDefense     int
	ShieldDefenseRate int

	MpConsume  int
	MpBonus    int
	ReuseDelay int

	Skill               string
	CriticalAttackSkill string

	//TODO
	//dual_fhit_rate
	//category
	//html
}
