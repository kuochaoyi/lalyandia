package model

const PlayerMaxLevel = 70

const (
	Human PlayerRace = iota
	Elf
	DarkElf
	Orc
	Dwarf
)

const (
	Fighter PlayerSpecializationType = iota
	Mystic
	Priest
)

const (
	HumanFighter PlayerSpecialization = iota
	Warrior
	Gladiator
	Warloard
	Knight
	Paladin
	DarkAvanger
	Rogue
	ThreasureHunter
	Hawkeye

	HumanMystic
	HumanWizard
	Sorcerer
	Necromancer
	Warlock
	Cleric
	Bishop
	Prophet

	ElvenFighter
	ElvenKnight
	TempleKnight
	SwordSinger
	ElvenScout
	PlainsWalker
	SilverRanger

	ElvenMystic
	ElvenWizard
	SpellSinger
	ElementalSummoner
	ElvenOracle
	ElvenElder

	DarkFighter
	PalusKnight
	ShillienKnight
	Bladedancer
	Assasin
	AbyssWalker
	PhantomRanger

	DarkMystic
	DarkWizard
	Shellhowler
	PhantomSummoner
	ShillienOracle
	ShillienElder

	OrcFighter
	OrcRaider
	Destroyer
	Monk
	Tyrant

	OrcMystic
	OrcShaman
	Overlord
	Warcryer

	DwarvenFigter
	Scavenger
	BountyHunter
	Artisan
	Warsmith

	PlayerSpecializationCount int = iota
)

/****************************************/

const (
	Weapon Type = iota
	Armor
	EtcItem
	Asset
	Accessary
	QuestItem
)

const (
	NoneSlot Slot = iota
	UnderWear
	RightEar
	LeftEar
	RightFinger
	LeftFinger
	Head
	RightHand
	Lefthand
	Gloves
	Chest
	Legs
	Feet
	Back
	LeftRightHand
)

const (
	NoneArmor ArmorType = iota
	Light
	Heavy
	Magic
)

const (
	NoneWeapon WeaponType = iota
	Sword
	Blunt
	Daggers
	Bow
	Pole
	Dual
	Etc
	Fist
)

const (
	NoneEtcItem EtcItemType = iota
	Arrow
	Potion
	Scroll
	Recipe
	Material
)

const (
	NoneCrystal CrystalType = iota
	D
	C
	B
	A
)

const (
	NoneAction ItemAction = iota
	Equip
	Peel
	SoulShot
	SpiritShot
	SkillMaintain
)

const (
	NormalConsumeType ConsumeType = iota
	StackableConsumeType
	AssetConsumeType
)

const (
	Steel MaterialType = iota
	FineSteel
	Wood
	Bone
	Bronze
	Leather
	Cloth
	Gold
	Mithril
	Liquid
	Oriharukon
	Damascus
	Adamantaite
	BloodSteel
	Paper
	Silver
	Chrysolite
	Crystal
	Horn
	ScaleofDragon
	Cotton
	Dyestuff
	Cobweb
)

/****************************************/

const (
	warrior NpcClass = iota
	Citizen
	Object
	Merchant
	WarehouseKeeper
	Teleporter
	GuildMaster
	GuildCoach
	Guard
	Blacksmith
	Holything
	Summon
	Chamberlain
)

const (
	NoneSex NpcSex = iota
	MaleSex
	FemaleSex
	EtcSex
)