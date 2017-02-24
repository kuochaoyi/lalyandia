package model

type ItemData struct {
	Id          ItemId
	Type        Type
	Slot        Slot
	ArmorType   ArmorType
	WeaponType  WeaponType
	EtcItemType EtcItemType

	Blessed          bool
	EnchantEnabled   bool
	Enchanted        bool
	Tradable         bool
	Dropable         bool
	Destructable     bool
	CanPenetrate     bool
	PrivateStore     bool
	CanEquipByPlayer bool

	DefaultAction   ItemAction
	DefaultPrice    int
	Price           int
	Weight          int
	CrystalType     CrystalType
	CrystalCount    int
	SoulshotCount   int
	SpiritshotCount int
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

	Skills []string

	//TODO
	UseCondition   string
	EquipCondition string

	//TODO
	//dual_fhit_rate
	//category
	//html
	//base_attribute_attack
	//base_attribute_defend
}
