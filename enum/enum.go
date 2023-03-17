package enum

type Affect_Type int32

const (
	Affect_Type_INCREASE Affect_Type = iota
	Affect_Type_DECREASE
	Affect_Type_FULL
	Affect_Type_ZERO
	Affect_Type_CNT
)

var (
	Affect_TypeName = map[Affect_Type]string{
		Affect_Type_INCREASE: "INCREASE",
		Affect_Type_DECREASE: "DECREASE",
		Affect_Type_FULL:     "FULL",
		Affect_Type_ZERO:     "ZERO",
		Affect_Type_CNT:      "CNT",
	}

	Affect_TypeValue = map[string]Affect_Type{
		"INCREASE": Affect_Type_INCREASE,
		"DECREASE": Affect_Type_DECREASE,
		"FULL":     Affect_Type_FULL,
		"ZERO":     Affect_Type_ZERO,
		"CNT":      Affect_Type_CNT,
	}
)

func (e Affect_Type) String() string {
	return Affect_TypeName[e]
}

func GetAffect_Type(s string) Affect_Type {
	return Affect_TypeValue[s]
}

type Asset int32

const (
	Asset_GOLD Asset = iota
	Asset_CRYSTAL
	Asset_MINERAL
	Asset_ESSENCE
	Asset_JEWEL
	Asset_MEAT
	Asset_HERB
	Asset_STARLIGHT
	Asset_RARE_MINERAL
	Asset_RARE_ESSENCE
	Asset_RARE_JEWEL
	Asset_RARE_MEAT
	Asset_RARE_HERB
	Asset_RARE_STARLIGHT
	Asset_FATE_FRAGMENT
	Asset_CNT
)

var (
	AssetName = map[Asset]string{
		Asset_GOLD:           "GOLD",
		Asset_CRYSTAL:        "CRYSTAL",
		Asset_MINERAL:        "MINERAL",
		Asset_ESSENCE:        "ESSENCE",
		Asset_JEWEL:          "JEWEL",
		Asset_MEAT:           "MEAT",
		Asset_HERB:           "HERB",
		Asset_STARLIGHT:      "STARLIGHT",
		Asset_RARE_MINERAL:   "RARE_MINERAL",
		Asset_RARE_ESSENCE:   "RARE_ESSENCE",
		Asset_RARE_JEWEL:     "RARE_JEWEL",
		Asset_RARE_MEAT:      "RARE_MEAT",
		Asset_RARE_HERB:      "RARE_HERB",
		Asset_RARE_STARLIGHT: "RARE_STARLIGHT",
		Asset_FATE_FRAGMENT:  "FATE_FRAGMENT",
		Asset_CNT:            "CNT",
	}

	AssetValue = map[string]Asset{
		"GOLD":           Asset_GOLD,
		"CRYSTAL":        Asset_CRYSTAL,
		"MINERAL":        Asset_MINERAL,
		"ESSENCE":        Asset_ESSENCE,
		"JEWEL":          Asset_JEWEL,
		"MEAT":           Asset_MEAT,
		"HERB":           Asset_HERB,
		"STARLIGHT":      Asset_STARLIGHT,
		"RARE_MINERAL":   Asset_RARE_MINERAL,
		"RARE_ESSENCE":   Asset_RARE_ESSENCE,
		"RARE_JEWEL":     Asset_RARE_JEWEL,
		"RARE_MEAT":      Asset_RARE_MEAT,
		"RARE_HERB":      Asset_RARE_HERB,
		"RARE_STARLIGHT": Asset_RARE_STARLIGHT,
		"FATE_FRAGMENT":  Asset_FATE_FRAGMENT,
		"CNT":            Asset_CNT,
	}
)

func (e Asset) String() string {
	return AssetName[e]
}

func GetAsset(s string) Asset {
	return AssetValue[s]
}

type Asset_Grade int32

const (
	Asset_Grade_NORMAL Asset_Grade = iota
	Asset_Grade_RARE
	Asset_Grade_EPIC
	Asset_Grade_LEGEND
	Asset_Grade_MYTH
	Asset_Grade_CNT
)

var (
	Asset_GradeName = map[Asset_Grade]string{
		Asset_Grade_NORMAL: "NORMAL",
		Asset_Grade_RARE:   "RARE",
		Asset_Grade_EPIC:   "EPIC",
		Asset_Grade_LEGEND: "LEGEND",
		Asset_Grade_MYTH:   "MYTH",
		Asset_Grade_CNT:    "CNT",
	}

	Asset_GradeValue = map[string]Asset_Grade{
		"NORMAL": Asset_Grade_NORMAL,
		"RARE":   Asset_Grade_RARE,
		"EPIC":   Asset_Grade_EPIC,
		"LEGEND": Asset_Grade_LEGEND,
		"MYTH":   Asset_Grade_MYTH,
		"CNT":    Asset_Grade_CNT,
	}
)

func (e Asset_Grade) String() string {
	return Asset_GradeName[e]
}

func GetAsset_Grade(s string) Asset_Grade {
	return Asset_GradeValue[s]
}

type Battle_Result int32

const (
	Battle_Result_NONE Battle_Result = iota
	Battle_Result_LOSE
	Battle_Result_WIN
	Battle_Result_ENCOUNTER
	Battle_Result_CNT
)

var (
	Battle_ResultName = map[Battle_Result]string{
		Battle_Result_NONE:      "NONE",
		Battle_Result_LOSE:      "LOSE",
		Battle_Result_WIN:       "WIN",
		Battle_Result_ENCOUNTER: "ENCOUNTER",
		Battle_Result_CNT:       "CNT",
	}

	Battle_ResultValue = map[string]Battle_Result{
		"NONE":      Battle_Result_NONE,
		"LOSE":      Battle_Result_LOSE,
		"WIN":       Battle_Result_WIN,
		"ENCOUNTER": Battle_Result_ENCOUNTER,
		"CNT":       Battle_Result_CNT,
	}
)

func (e Battle_Result) String() string {
	return Battle_ResultName[e]
}

func GetBattle_Result(s string) Battle_Result {
	return Battle_ResultValue[s]
}

type Bonus_Trigger int32

const (
	Bonus_Trigger_KILL_NUMBER Bonus_Trigger = iota
	Bonus_Trigger_KILL_GAP
	Bonus_Trigger_CNT
)

var (
	Bonus_TriggerName = map[Bonus_Trigger]string{
		Bonus_Trigger_KILL_NUMBER: "KILL_NUMBER",
		Bonus_Trigger_KILL_GAP:    "KILL_GAP",
		Bonus_Trigger_CNT:         "CNT",
	}

	Bonus_TriggerValue = map[string]Bonus_Trigger{
		"KILL_NUMBER": Bonus_Trigger_KILL_NUMBER,
		"KILL_GAP":    Bonus_Trigger_KILL_GAP,
		"CNT":         Bonus_Trigger_CNT,
	}
)

func (e Bonus_Trigger) String() string {
	return Bonus_TriggerName[e]
}

func GetBonus_Trigger(s string) Bonus_Trigger {
	return Bonus_TriggerValue[s]
}

type Calculation int32

const (
	Calculation_PLUS Calculation = iota
	Calculation_MULT
	Calculation_P_MULT
	Calculation_CNT
)

var (
	CalculationName = map[Calculation]string{
		Calculation_PLUS:   "PLUS",
		Calculation_MULT:   "MULT",
		Calculation_P_MULT: "P_MULT",
		Calculation_CNT:    "CNT",
	}

	CalculationValue = map[string]Calculation{
		"PLUS":   Calculation_PLUS,
		"MULT":   Calculation_MULT,
		"P_MULT": Calculation_P_MULT,
		"CNT":    Calculation_CNT,
	}
)

func (e Calculation) String() string {
	return CalculationName[e]
}

func GetCalculation(s string) Calculation {
	return CalculationValue[s]
}

type Carrier_Trigger int32

const (
	Carrier_Trigger_CONNECT Carrier_Trigger = iota
	Carrier_Trigger_RELEASE
	Carrier_Trigger_CNT
)

var (
	Carrier_TriggerName = map[Carrier_Trigger]string{
		Carrier_Trigger_CONNECT: "CONNECT",
		Carrier_Trigger_RELEASE: "RELEASE",
		Carrier_Trigger_CNT:     "CNT",
	}

	Carrier_TriggerValue = map[string]Carrier_Trigger{
		"CONNECT": Carrier_Trigger_CONNECT,
		"RELEASE": Carrier_Trigger_RELEASE,
		"CNT":     Carrier_Trigger_CNT,
	}
)

func (e Carrier_Trigger) String() string {
	return Carrier_TriggerName[e]
}

func GetCarrier_Trigger(s string) Carrier_Trigger {
	return Carrier_TriggerValue[s]
}

type Carrier_Velocity int32

const (
	Carrier_Velocity_FAST Carrier_Velocity = iota
	Carrier_Velocity_AVERAGE
	Carrier_Velocity_SLOW
	Carrier_Velocity_INSTANT
	Carrier_Velocity_CONSTANT
	Carrier_Velocity_CNT
)

var (
	Carrier_VelocityName = map[Carrier_Velocity]string{
		Carrier_Velocity_FAST:     "FAST",
		Carrier_Velocity_AVERAGE:  "AVERAGE",
		Carrier_Velocity_SLOW:     "SLOW",
		Carrier_Velocity_INSTANT:  "INSTANT",
		Carrier_Velocity_CONSTANT: "CONSTANT",
		Carrier_Velocity_CNT:      "CNT",
	}

	Carrier_VelocityValue = map[string]Carrier_Velocity{
		"FAST":     Carrier_Velocity_FAST,
		"AVERAGE":  Carrier_Velocity_AVERAGE,
		"SLOW":     Carrier_Velocity_SLOW,
		"INSTANT":  Carrier_Velocity_INSTANT,
		"CONSTANT": Carrier_Velocity_CONSTANT,
		"CNT":      Carrier_Velocity_CNT,
	}
)

func (e Carrier_Velocity) String() string {
	return Carrier_VelocityName[e]
}

func GetCarrier_Velocity(s string) Carrier_Velocity {
	return Carrier_VelocityValue[s]
}

type Character_Activity int32

const (
	Character_Activity_NONE Character_Activity = iota
	Character_Activity_BATTLE
	Character_Activity_GUARD
	Character_Activity_CNT
)

var (
	Character_ActivityName = map[Character_Activity]string{
		Character_Activity_NONE:   "NONE",
		Character_Activity_BATTLE: "BATTLE",
		Character_Activity_GUARD:  "GUARD",
		Character_Activity_CNT:    "CNT",
	}

	Character_ActivityValue = map[string]Character_Activity{
		"NONE":   Character_Activity_NONE,
		"BATTLE": Character_Activity_BATTLE,
		"GUARD":  Character_Activity_GUARD,
		"CNT":    Character_Activity_CNT,
	}
)

func (e Character_Activity) String() string {
	return Character_ActivityName[e]
}

func GetCharacter_Activity(s string) Character_Activity {
	return Character_ActivityValue[s]
}

type Character_Class int32

const (
	Character_Class_WARRIOR Character_Class = iota
	Character_Class_RANGER
	Character_Class_ASSASSIN
	Character_Class_KNIGHT
	Character_Class_WIZARD
	Character_Class_PRIEST
	Character_Class_BARD
	Character_Class_BATTLE_MAGE
	Character_Class_CNT
)

var (
	Character_ClassName = map[Character_Class]string{
		Character_Class_WARRIOR:     "WARRIOR",
		Character_Class_RANGER:      "RANGER",
		Character_Class_ASSASSIN:    "ASSASSIN",
		Character_Class_KNIGHT:      "KNIGHT",
		Character_Class_WIZARD:      "WIZARD",
		Character_Class_PRIEST:      "PRIEST",
		Character_Class_BARD:        "BARD",
		Character_Class_BATTLE_MAGE: "BATTLE_MAGE",
		Character_Class_CNT:         "CNT",
	}

	Character_ClassValue = map[string]Character_Class{
		"WARRIOR":     Character_Class_WARRIOR,
		"RANGER":      Character_Class_RANGER,
		"ASSASSIN":    Character_Class_ASSASSIN,
		"KNIGHT":      Character_Class_KNIGHT,
		"WIZARD":      Character_Class_WIZARD,
		"PRIEST":      Character_Class_PRIEST,
		"BARD":        Character_Class_BARD,
		"BATTLE_MAGE": Character_Class_BATTLE_MAGE,
		"CNT":         Character_Class_CNT,
	}
)

func (e Character_Class) String() string {
	return Character_ClassName[e]
}

func GetCharacter_Class(s string) Character_Class {
	return Character_ClassValue[s]
}

type Character_Grade int32

const (
	Character_Grade_GRADE_F Character_Grade = iota
	Character_Grade_GRADE_D
	Character_Grade_GRADE_C
	Character_Grade_GRADE_B
	Character_Grade_GRADE_A
	Character_Grade_CNT
)

var (
	Character_GradeName = map[Character_Grade]string{
		Character_Grade_GRADE_F: "GRADE_F",
		Character_Grade_GRADE_D: "GRADE_D",
		Character_Grade_GRADE_C: "GRADE_C",
		Character_Grade_GRADE_B: "GRADE_B",
		Character_Grade_GRADE_A: "GRADE_A",
		Character_Grade_CNT:     "CNT",
	}

	Character_GradeValue = map[string]Character_Grade{
		"GRADE_F": Character_Grade_GRADE_F,
		"GRADE_D": Character_Grade_GRADE_D,
		"GRADE_C": Character_Grade_GRADE_C,
		"GRADE_B": Character_Grade_GRADE_B,
		"GRADE_A": Character_Grade_GRADE_A,
		"CNT":     Character_Grade_CNT,
	}
)

func (e Character_Grade) String() string {
	return Character_GradeName[e]
}

func GetCharacter_Grade(s string) Character_Grade {
	return Character_GradeValue[s]
}

type Character_Property int32

const (
	Character_Property_DARK Character_Property = iota
	Character_Property_FIRE
	Character_Property_WATER
	Character_Property_FOREST
	Character_Property_ELECTRIC
	Character_Property_LIGHT
	Character_Property_EARTH
	Character_Property_CNT
)

var (
	Character_PropertyName = map[Character_Property]string{
		Character_Property_DARK:     "DARK",
		Character_Property_FIRE:     "FIRE",
		Character_Property_WATER:    "WATER",
		Character_Property_FOREST:   "FOREST",
		Character_Property_ELECTRIC: "ELECTRIC",
		Character_Property_LIGHT:    "LIGHT",
		Character_Property_EARTH:    "EARTH",
		Character_Property_CNT:      "CNT",
	}

	Character_PropertyValue = map[string]Character_Property{
		"DARK":     Character_Property_DARK,
		"FIRE":     Character_Property_FIRE,
		"WATER":    Character_Property_WATER,
		"FOREST":   Character_Property_FOREST,
		"ELECTRIC": Character_Property_ELECTRIC,
		"LIGHT":    Character_Property_LIGHT,
		"EARTH":    Character_Property_EARTH,
		"CNT":      Character_Property_CNT,
	}
)

func (e Character_Property) String() string {
	return Character_PropertyName[e]
}

func GetCharacter_Property(s string) Character_Property {
	return Character_PropertyValue[s]
}

type Character_Quest_Trigger int32

const (
	Character_Quest_Trigger_AUTO_BATTLE Character_Quest_Trigger = iota
	Character_Quest_Trigger_TILE_GET
	Character_Quest_Trigger_EQUIPMENT_GET
	Character_Quest_Trigger_CNT
)

var (
	Character_Quest_TriggerName = map[Character_Quest_Trigger]string{
		Character_Quest_Trigger_AUTO_BATTLE:   "AUTO_BATTLE",
		Character_Quest_Trigger_TILE_GET:      "TILE_GET",
		Character_Quest_Trigger_EQUIPMENT_GET: "EQUIPMENT_GET",
		Character_Quest_Trigger_CNT:           "CNT",
	}

	Character_Quest_TriggerValue = map[string]Character_Quest_Trigger{
		"AUTO_BATTLE":   Character_Quest_Trigger_AUTO_BATTLE,
		"TILE_GET":      Character_Quest_Trigger_TILE_GET,
		"EQUIPMENT_GET": Character_Quest_Trigger_EQUIPMENT_GET,
		"CNT":           Character_Quest_Trigger_CNT,
	}
)

func (e Character_Quest_Trigger) String() string {
	return Character_Quest_TriggerName[e]
}

func GetCharacter_Quest_Trigger(s string) Character_Quest_Trigger {
	return Character_Quest_TriggerValue[s]
}

type Character_Quest_Type int32

const (
	Character_Quest_Type_COLLECT Character_Quest_Type = iota
	Character_Quest_Type_TILE_GET
	Character_Quest_Type_CNT
)

var (
	Character_Quest_TypeName = map[Character_Quest_Type]string{
		Character_Quest_Type_COLLECT:  "COLLECT",
		Character_Quest_Type_TILE_GET: "TILE_GET",
		Character_Quest_Type_CNT:      "CNT",
	}

	Character_Quest_TypeValue = map[string]Character_Quest_Type{
		"COLLECT":  Character_Quest_Type_COLLECT,
		"TILE_GET": Character_Quest_Type_TILE_GET,
		"CNT":      Character_Quest_Type_CNT,
	}
)

func (e Character_Quest_Type) String() string {
	return Character_Quest_TypeName[e]
}

func GetCharacter_Quest_Type(s string) Character_Quest_Type {
	return Character_Quest_TypeValue[s]
}

type Character_Skill int32

const (
	Character_Skill_CLASS Character_Skill = iota
	Character_Skill_LEADER
	Character_Skill_BASIC
	Character_Skill_ACTIVE
	Character_Skill_ULTIMATE
	Character_Skill_PASSIVE
	Character_Skill_CNT
)

var (
	Character_SkillName = map[Character_Skill]string{
		Character_Skill_CLASS:    "CLASS",
		Character_Skill_LEADER:   "LEADER",
		Character_Skill_BASIC:    "BASIC",
		Character_Skill_ACTIVE:   "ACTIVE",
		Character_Skill_ULTIMATE: "ULTIMATE",
		Character_Skill_PASSIVE:  "PASSIVE",
		Character_Skill_CNT:      "CNT",
	}

	Character_SkillValue = map[string]Character_Skill{
		"CLASS":    Character_Skill_CLASS,
		"LEADER":   Character_Skill_LEADER,
		"BASIC":    Character_Skill_BASIC,
		"ACTIVE":   Character_Skill_ACTIVE,
		"ULTIMATE": Character_Skill_ULTIMATE,
		"PASSIVE":  Character_Skill_PASSIVE,
		"CNT":      Character_Skill_CNT,
	}
)

func (e Character_Skill) String() string {
	return Character_SkillName[e]
}

func GetCharacter_Skill(s string) Character_Skill {
	return Character_SkillValue[s]
}

type Character_Species int32

const (
	Character_Species_ANGEL Character_Species = iota
	Character_Species_DEVIL
	Character_Species_DRAGON
	Character_Species_BEAST
	Character_Species_CNT
)

var (
	Character_SpeciesName = map[Character_Species]string{
		Character_Species_ANGEL:  "ANGEL",
		Character_Species_DEVIL:  "DEVIL",
		Character_Species_DRAGON: "DRAGON",
		Character_Species_BEAST:  "BEAST",
		Character_Species_CNT:    "CNT",
	}

	Character_SpeciesValue = map[string]Character_Species{
		"ANGEL":  Character_Species_ANGEL,
		"DEVIL":  Character_Species_DEVIL,
		"DRAGON": Character_Species_DRAGON,
		"BEAST":  Character_Species_BEAST,
		"CNT":    Character_Species_CNT,
	}
)

func (e Character_Species) String() string {
	return Character_SpeciesName[e]
}

func GetCharacter_Species(s string) Character_Species {
	return Character_SpeciesValue[s]
}

type Chat_Trigger int32

const (
	Chat_Trigger_GIFT Chat_Trigger = iota
	Chat_Trigger_STAMINA_LOW
	Chat_Trigger_STAMINA_ZERO
	Chat_Trigger_GAME_ACCESS
	Chat_Trigger_COMMON_TILE_OCCUPY
	Chat_Trigger_COMMON_BIGTILE_OCCUPY
	Chat_Trigger_SPECIFIED_TILE_OCCUPY
	Chat_Trigger_NIGHT
	Chat_Trigger_DAY
	Chat_Trigger_CHARACTER_QUEST
	Chat_Trigger_CHARACTER_GET
	Chat_Trigger_CNT
)

var (
	Chat_TriggerName = map[Chat_Trigger]string{
		Chat_Trigger_GIFT:                  "GIFT",
		Chat_Trigger_STAMINA_LOW:           "STAMINA_LOW",
		Chat_Trigger_STAMINA_ZERO:          "STAMINA_ZERO",
		Chat_Trigger_GAME_ACCESS:           "GAME_ACCESS",
		Chat_Trigger_COMMON_TILE_OCCUPY:    "COMMON_TILE_OCCUPY",
		Chat_Trigger_COMMON_BIGTILE_OCCUPY: "COMMON_BIGTILE_OCCUPY",
		Chat_Trigger_SPECIFIED_TILE_OCCUPY: "SPECIFIED_TILE_OCCUPY",
		Chat_Trigger_NIGHT:                 "NIGHT",
		Chat_Trigger_DAY:                   "DAY",
		Chat_Trigger_CHARACTER_QUEST:       "CHARACTER_QUEST",
		Chat_Trigger_CHARACTER_GET:         "CHARACTER_GET",
		Chat_Trigger_CNT:                   "CNT",
	}

	Chat_TriggerValue = map[string]Chat_Trigger{
		"GIFT":                  Chat_Trigger_GIFT,
		"STAMINA_LOW":           Chat_Trigger_STAMINA_LOW,
		"STAMINA_ZERO":          Chat_Trigger_STAMINA_ZERO,
		"GAME_ACCESS":           Chat_Trigger_GAME_ACCESS,
		"COMMON_TILE_OCCUPY":    Chat_Trigger_COMMON_TILE_OCCUPY,
		"COMMON_BIGTILE_OCCUPY": Chat_Trigger_COMMON_BIGTILE_OCCUPY,
		"SPECIFIED_TILE_OCCUPY": Chat_Trigger_SPECIFIED_TILE_OCCUPY,
		"NIGHT":                 Chat_Trigger_NIGHT,
		"DAY":                   Chat_Trigger_DAY,
		"CHARACTER_QUEST":       Chat_Trigger_CHARACTER_QUEST,
		"CHARACTER_GET":         Chat_Trigger_CHARACTER_GET,
		"CNT":                   Chat_Trigger_CNT,
	}
)

func (e Chat_Trigger) String() string {
	return Chat_TriggerName[e]
}

func GetChat_Trigger(s string) Chat_Trigger {
	return Chat_TriggerValue[s]
}

type Common_Grade int32

const (
	Common_Grade_NORMAL Common_Grade = iota
	Common_Grade_RARE
	Common_Grade_EPIC
	Common_Grade_LEGEND
	Common_Grade_MYTH
	Common_Grade_CNT
)

var (
	Common_GradeName = map[Common_Grade]string{
		Common_Grade_NORMAL: "NORMAL",
		Common_Grade_RARE:   "RARE",
		Common_Grade_EPIC:   "EPIC",
		Common_Grade_LEGEND: "LEGEND",
		Common_Grade_MYTH:   "MYTH",
		Common_Grade_CNT:    "CNT",
	}

	Common_GradeValue = map[string]Common_Grade{
		"NORMAL": Common_Grade_NORMAL,
		"RARE":   Common_Grade_RARE,
		"EPIC":   Common_Grade_EPIC,
		"LEGEND": Common_Grade_LEGEND,
		"MYTH":   Common_Grade_MYTH,
		"CNT":    Common_Grade_CNT,
	}
)

func (e Common_Grade) String() string {
	return Common_GradeName[e]
}

func GetCommon_Grade(s string) Common_Grade {
	return Common_GradeValue[s]
}

type Common_Type int32

const (
	Common_Type_ASSET Common_Type = iota
	Common_Type_CHARACTER
	Common_Type_STARGEM
	Common_Type_ITEM
	Common_Type_RECIPE
	Common_Type_STORY
	Common_Type_AFFECTION
	Common_Type_GACHA
	Common_Type_FATE_CARD
	Common_Type_CNT
)

var (
	Common_TypeName = map[Common_Type]string{
		Common_Type_ASSET:     "ASSET",
		Common_Type_CHARACTER: "CHARACTER",
		Common_Type_STARGEM:   "STARGEM",
		Common_Type_ITEM:      "ITEM",
		Common_Type_RECIPE:    "RECIPE",
		Common_Type_STORY:     "STORY",
		Common_Type_AFFECTION: "AFFECTION",
		Common_Type_GACHA:     "GACHA",
		Common_Type_FATE_CARD: "FATE_CARD",
		Common_Type_CNT:       "CNT",
	}

	Common_TypeValue = map[string]Common_Type{
		"ASSET":     Common_Type_ASSET,
		"CHARACTER": Common_Type_CHARACTER,
		"STARGEM":   Common_Type_STARGEM,
		"ITEM":      Common_Type_ITEM,
		"RECIPE":    Common_Type_RECIPE,
		"STORY":     Common_Type_STORY,
		"AFFECTION": Common_Type_AFFECTION,
		"GACHA":     Common_Type_GACHA,
		"FATE_CARD": Common_Type_FATE_CARD,
		"CNT":       Common_Type_CNT,
	}
)

func (e Common_Type) String() string {
	return Common_TypeName[e]
}

func GetCommon_Type(s string) Common_Type {
	return Common_TypeValue[s]
}

type Cost_Type int32

const (
	Cost_Type_STACK Cost_Type = iota
	Cost_Type_CNT
)

var (
	Cost_TypeName = map[Cost_Type]string{
		Cost_Type_STACK: "STACK",
		Cost_Type_CNT:   "CNT",
	}

	Cost_TypeValue = map[string]Cost_Type{
		"STACK": Cost_Type_STACK,
		"CNT":   Cost_Type_CNT,
	}
)

func (e Cost_Type) String() string {
	return Cost_TypeName[e]
}

func GetCost_Type(s string) Cost_Type {
	return Cost_TypeValue[s]
}

type Costume_Condition int32

const (
	Costume_Condition_NONE Costume_Condition = iota
	Costume_Condition_BASIC
	Costume_Condition_DUEL_WIN
	Costume_Condition_DUEL_JOIN
	Costume_Condition_AFFECTION_LEVEL
	Costume_Condition_UNUSED
	Costume_Condition_CNT
)

var (
	Costume_ConditionName = map[Costume_Condition]string{
		Costume_Condition_NONE:            "NONE",
		Costume_Condition_BASIC:           "BASIC",
		Costume_Condition_DUEL_WIN:        "DUEL_WIN",
		Costume_Condition_DUEL_JOIN:       "DUEL_JOIN",
		Costume_Condition_AFFECTION_LEVEL: "AFFECTION_LEVEL",
		Costume_Condition_UNUSED:          "UNUSED",
		Costume_Condition_CNT:             "CNT",
	}

	Costume_ConditionValue = map[string]Costume_Condition{
		"NONE":            Costume_Condition_NONE,
		"BASIC":           Costume_Condition_BASIC,
		"DUEL_WIN":        Costume_Condition_DUEL_WIN,
		"DUEL_JOIN":       Costume_Condition_DUEL_JOIN,
		"AFFECTION_LEVEL": Costume_Condition_AFFECTION_LEVEL,
		"UNUSED":          Costume_Condition_UNUSED,
		"CNT":             Costume_Condition_CNT,
	}
)

func (e Costume_Condition) String() string {
	return Costume_ConditionName[e]
}

func GetCostume_Condition(s string) Costume_Condition {
	return Costume_ConditionValue[s]
}

type Craft_Category int32

const (
	Craft_Category_OBSERVATORY Craft_Category = iota
	Craft_Category_COOKING
	Craft_Category_ALCHEMY
	Craft_Category_JEWELRY
	Craft_Category_SMITHERY
	Craft_Category_BASIC
	Craft_Category_FATE_CARD
	Craft_Category_CNT
)

var (
	Craft_CategoryName = map[Craft_Category]string{
		Craft_Category_OBSERVATORY: "OBSERVATORY",
		Craft_Category_COOKING:     "COOKING",
		Craft_Category_ALCHEMY:     "ALCHEMY",
		Craft_Category_JEWELRY:     "JEWELRY",
		Craft_Category_SMITHERY:    "SMITHERY",
		Craft_Category_BASIC:       "BASIC",
		Craft_Category_FATE_CARD:   "FATE_CARD",
		Craft_Category_CNT:         "CNT",
	}

	Craft_CategoryValue = map[string]Craft_Category{
		"OBSERVATORY": Craft_Category_OBSERVATORY,
		"COOKING":     Craft_Category_COOKING,
		"ALCHEMY":     Craft_Category_ALCHEMY,
		"JEWELRY":     Craft_Category_JEWELRY,
		"SMITHERY":    Craft_Category_SMITHERY,
		"BASIC":       Craft_Category_BASIC,
		"FATE_CARD":   Craft_Category_FATE_CARD,
		"CNT":         Craft_Category_CNT,
	}
)

func (e Craft_Category) String() string {
	return Craft_CategoryName[e]
}

func GetCraft_Category(s string) Craft_Category {
	return Craft_CategoryValue[s]
}

type Crafting_Grade int32

const (
	Crafting_Grade_NORMAL Crafting_Grade = iota
	Crafting_Grade_RARE
	Crafting_Grade_EPIC
	Crafting_Grade_UNIQUE
	Crafting_Grade_MYTH
	Crafting_Grade_CNT
)

var (
	Crafting_GradeName = map[Crafting_Grade]string{
		Crafting_Grade_NORMAL: "NORMAL",
		Crafting_Grade_RARE:   "RARE",
		Crafting_Grade_EPIC:   "EPIC",
		Crafting_Grade_UNIQUE: "UNIQUE",
		Crafting_Grade_MYTH:   "MYTH",
		Crafting_Grade_CNT:    "CNT",
	}

	Crafting_GradeValue = map[string]Crafting_Grade{
		"NORMAL": Crafting_Grade_NORMAL,
		"RARE":   Crafting_Grade_RARE,
		"EPIC":   Crafting_Grade_EPIC,
		"UNIQUE": Crafting_Grade_UNIQUE,
		"MYTH":   Crafting_Grade_MYTH,
		"CNT":    Crafting_Grade_CNT,
	}
)

func (e Crafting_Grade) String() string {
	return Crafting_GradeName[e]
}

func GetCrafting_Grade(s string) Crafting_Grade {
	return Crafting_GradeValue[s]
}

type Crew int32

const (
	Crew_SEIRIOS Crew = iota
	Crew_INFINITY_EIGHT
	Crew_OFFICE_RAID_PARTY
	Crew_ALETHEIA
	Crew_KLEIN_HAUS
	Crew_LAUREL_GARDEN
	Crew_CLEAN_CLOUD
	Crew_CNT
)

var (
	CrewName = map[Crew]string{
		Crew_SEIRIOS:           "SEIRIOS",
		Crew_INFINITY_EIGHT:    "INFINITY_EIGHT",
		Crew_OFFICE_RAID_PARTY: "OFFICE_RAID_PARTY",
		Crew_ALETHEIA:          "ALETHEIA",
		Crew_KLEIN_HAUS:        "KLEIN_HAUS",
		Crew_LAUREL_GARDEN:     "LAUREL_GARDEN",
		Crew_CLEAN_CLOUD:       "CLEAN_CLOUD",
		Crew_CNT:               "CNT",
	}

	CrewValue = map[string]Crew{
		"SEIRIOS":           Crew_SEIRIOS,
		"INFINITY_EIGHT":    Crew_INFINITY_EIGHT,
		"OFFICE_RAID_PARTY": Crew_OFFICE_RAID_PARTY,
		"ALETHEIA":          Crew_ALETHEIA,
		"KLEIN_HAUS":        Crew_KLEIN_HAUS,
		"LAUREL_GARDEN":     Crew_LAUREL_GARDEN,
		"CLEAN_CLOUD":       Crew_CLEAN_CLOUD,
		"CNT":               Crew_CNT,
	}
)

func (e Crew) String() string {
	return CrewName[e]
}

func GetCrew(s string) Crew {
	return CrewValue[s]
}

type Dialogue_Format int32

const (
	Dialogue_Format_NORMAL Dialogue_Format = iota
	Dialogue_Format_CNT
)

var (
	Dialogue_FormatName = map[Dialogue_Format]string{
		Dialogue_Format_NORMAL: "NORMAL",
		Dialogue_Format_CNT:    "CNT",
	}

	Dialogue_FormatValue = map[string]Dialogue_Format{
		"NORMAL": Dialogue_Format_NORMAL,
		"CNT":    Dialogue_Format_CNT,
	}
)

func (e Dialogue_Format) String() string {
	return Dialogue_FormatName[e]
}

func GetDialogue_Format(s string) Dialogue_Format {
	return Dialogue_FormatValue[s]
}

type Duel_Trigger int32

const (
	Duel_Trigger_ENEMY_SURRENDER Duel_Trigger = iota
	Duel_Trigger_TIME_LAST
	Duel_Trigger_BOSS_STEAL
	Duel_Trigger_DOGFIGHT
	Duel_Trigger_LARGE_DONATION
	Duel_Trigger_ENEMY_OUT
	Duel_Trigger_BASE_30
	Duel_Trigger_ENEMY_ALLKILL
	Duel_Trigger_KILL_MANY
	Duel_Trigger_BOSS_KILL
	Duel_Trigger_BOSS_RESPAWN
	Duel_Trigger_BASE_ATTACKED
	Duel_Trigger_CHARACTER_DEADLY
	Duel_Trigger_ALLY_KILL
	Duel_Trigger_ENEMY_KILL
	Duel_Trigger_FIRSTBLOOD
	Duel_Trigger_EXCLUSION
	Duel_Trigger_RESOURCE_FULL
	Duel_Trigger_GAME_START
	Duel_Trigger_BASE_DESTROY
	Duel_Trigger_EXTERMINATE
	Duel_Trigger_ALL_CHARACTER_DOWN
	Duel_Trigger_KILL
	Duel_Trigger_TUTORIAL_0
	Duel_Trigger_TUTORIAL_1
	Duel_Trigger_CNT
)

var (
	Duel_TriggerName = map[Duel_Trigger]string{
		Duel_Trigger_ENEMY_SURRENDER:    "ENEMY_SURRENDER",
		Duel_Trigger_TIME_LAST:          "TIME_LAST",
		Duel_Trigger_BOSS_STEAL:         "BOSS_STEAL",
		Duel_Trigger_DOGFIGHT:           "DOGFIGHT",
		Duel_Trigger_LARGE_DONATION:     "LARGE_DONATION",
		Duel_Trigger_ENEMY_OUT:          "ENEMY_OUT",
		Duel_Trigger_BASE_30:            "BASE_30",
		Duel_Trigger_ENEMY_ALLKILL:      "ENEMY_ALLKILL",
		Duel_Trigger_KILL_MANY:          "KILL_MANY",
		Duel_Trigger_BOSS_KILL:          "BOSS_KILL",
		Duel_Trigger_BOSS_RESPAWN:       "BOSS_RESPAWN",
		Duel_Trigger_BASE_ATTACKED:      "BASE_ATTACKED",
		Duel_Trigger_CHARACTER_DEADLY:   "CHARACTER_DEADLY",
		Duel_Trigger_ALLY_KILL:          "ALLY_KILL",
		Duel_Trigger_ENEMY_KILL:         "ENEMY_KILL",
		Duel_Trigger_FIRSTBLOOD:         "FIRSTBLOOD",
		Duel_Trigger_EXCLUSION:          "EXCLUSION",
		Duel_Trigger_RESOURCE_FULL:      "RESOURCE_FULL",
		Duel_Trigger_GAME_START:         "GAME_START",
		Duel_Trigger_BASE_DESTROY:       "BASE_DESTROY",
		Duel_Trigger_EXTERMINATE:        "EXTERMINATE",
		Duel_Trigger_ALL_CHARACTER_DOWN: "ALL_CHARACTER_DOWN",
		Duel_Trigger_KILL:               "KILL",
		Duel_Trigger_TUTORIAL_0:         "TUTORIAL_0",
		Duel_Trigger_TUTORIAL_1:         "TUTORIAL_1",
		Duel_Trigger_CNT:                "CNT",
	}

	Duel_TriggerValue = map[string]Duel_Trigger{
		"ENEMY_SURRENDER":    Duel_Trigger_ENEMY_SURRENDER,
		"TIME_LAST":          Duel_Trigger_TIME_LAST,
		"BOSS_STEAL":         Duel_Trigger_BOSS_STEAL,
		"DOGFIGHT":           Duel_Trigger_DOGFIGHT,
		"LARGE_DONATION":     Duel_Trigger_LARGE_DONATION,
		"ENEMY_OUT":          Duel_Trigger_ENEMY_OUT,
		"BASE_30":            Duel_Trigger_BASE_30,
		"ENEMY_ALLKILL":      Duel_Trigger_ENEMY_ALLKILL,
		"KILL_MANY":          Duel_Trigger_KILL_MANY,
		"BOSS_KILL":          Duel_Trigger_BOSS_KILL,
		"BOSS_RESPAWN":       Duel_Trigger_BOSS_RESPAWN,
		"BASE_ATTACKED":      Duel_Trigger_BASE_ATTACKED,
		"CHARACTER_DEADLY":   Duel_Trigger_CHARACTER_DEADLY,
		"ALLY_KILL":          Duel_Trigger_ALLY_KILL,
		"ENEMY_KILL":         Duel_Trigger_ENEMY_KILL,
		"FIRSTBLOOD":         Duel_Trigger_FIRSTBLOOD,
		"EXCLUSION":          Duel_Trigger_EXCLUSION,
		"RESOURCE_FULL":      Duel_Trigger_RESOURCE_FULL,
		"GAME_START":         Duel_Trigger_GAME_START,
		"BASE_DESTROY":       Duel_Trigger_BASE_DESTROY,
		"EXTERMINATE":        Duel_Trigger_EXTERMINATE,
		"ALL_CHARACTER_DOWN": Duel_Trigger_ALL_CHARACTER_DOWN,
		"KILL":               Duel_Trigger_KILL,
		"TUTORIAL_0":         Duel_Trigger_TUTORIAL_0,
		"TUTORIAL_1":         Duel_Trigger_TUTORIAL_1,
		"CNT":                Duel_Trigger_CNT,
	}
)

func (e Duel_Trigger) String() string {
	return Duel_TriggerName[e]
}

func GetDuel_Trigger(s string) Duel_Trigger {
	return Duel_TriggerValue[s]
}

type Effect_Category int32

const (
	Effect_Category_INSTANT Effect_Category = iota
	Effect_Category_MOVE
	Effect_Category_LASTING
	Effect_Category_NEG_STATUS_EFFECT
	Effect_Category_POS_STATUS_EFFECT
	Effect_Category_STAT_INCREASE
	Effect_Category_STAT_DECREASE
	Effect_Category_CNT
)

var (
	Effect_CategoryName = map[Effect_Category]string{
		Effect_Category_INSTANT:           "INSTANT",
		Effect_Category_MOVE:              "MOVE",
		Effect_Category_LASTING:           "LASTING",
		Effect_Category_NEG_STATUS_EFFECT: "NEG_STATUS_EFFECT",
		Effect_Category_POS_STATUS_EFFECT: "POS_STATUS_EFFECT",
		Effect_Category_STAT_INCREASE:     "STAT_INCREASE",
		Effect_Category_STAT_DECREASE:     "STAT_DECREASE",
		Effect_Category_CNT:               "CNT",
	}

	Effect_CategoryValue = map[string]Effect_Category{
		"INSTANT":           Effect_Category_INSTANT,
		"MOVE":              Effect_Category_MOVE,
		"LASTING":           Effect_Category_LASTING,
		"NEG_STATUS_EFFECT": Effect_Category_NEG_STATUS_EFFECT,
		"POS_STATUS_EFFECT": Effect_Category_POS_STATUS_EFFECT,
		"STAT_INCREASE":     Effect_Category_STAT_INCREASE,
		"STAT_DECREASE":     Effect_Category_STAT_DECREASE,
		"CNT":               Effect_Category_CNT,
	}
)

func (e Effect_Category) String() string {
	return Effect_CategoryName[e]
}

func GetEffect_Category(s string) Effect_Category {
	return Effect_CategoryValue[s]
}

type Enable_Cond int32

const (
	Enable_Cond_NONE Enable_Cond = iota
	Enable_Cond_TIME
	Enable_Cond_CRIT
	Enable_Cond_HIT
	Enable_Cond_TARGET_HAS_STATUSEFFECT
	Enable_Cond_SELF_HAS_STATUSEFFECT
	Enable_Cond_TARGET_HAS_STACKS
	Enable_Cond_SELF_HAS_STACKS
	Enable_Cond_RANDOM_CHANCE
	Enable_Cond_STACK_COUNT
	Enable_Cond_TARGET_ALLY_HP_LESS_THAN
	Enable_Cond_TARGET_HP_LESS_THAN
	Enable_Cond_LAUNCH
	Enable_Cond_CNT
)

var (
	Enable_CondName = map[Enable_Cond]string{
		Enable_Cond_NONE:                     "NONE",
		Enable_Cond_TIME:                     "TIME",
		Enable_Cond_CRIT:                     "CRIT",
		Enable_Cond_HIT:                      "HIT",
		Enable_Cond_TARGET_HAS_STATUSEFFECT:  "TARGET_HAS_STATUSEFFECT",
		Enable_Cond_SELF_HAS_STATUSEFFECT:    "SELF_HAS_STATUSEFFECT",
		Enable_Cond_TARGET_HAS_STACKS:        "TARGET_HAS_STACKS",
		Enable_Cond_SELF_HAS_STACKS:          "SELF_HAS_STACKS",
		Enable_Cond_RANDOM_CHANCE:            "RANDOM_CHANCE",
		Enable_Cond_STACK_COUNT:              "STACK_COUNT",
		Enable_Cond_TARGET_ALLY_HP_LESS_THAN: "TARGET_ALLY_HP_LESS_THAN",
		Enable_Cond_TARGET_HP_LESS_THAN:      "TARGET_HP_LESS_THAN",
		Enable_Cond_LAUNCH:                   "LAUNCH",
		Enable_Cond_CNT:                      "CNT",
	}

	Enable_CondValue = map[string]Enable_Cond{
		"NONE":                     Enable_Cond_NONE,
		"TIME":                     Enable_Cond_TIME,
		"CRIT":                     Enable_Cond_CRIT,
		"HIT":                      Enable_Cond_HIT,
		"TARGET_HAS_STATUSEFFECT":  Enable_Cond_TARGET_HAS_STATUSEFFECT,
		"SELF_HAS_STATUSEFFECT":    Enable_Cond_SELF_HAS_STATUSEFFECT,
		"TARGET_HAS_STACKS":        Enable_Cond_TARGET_HAS_STACKS,
		"SELF_HAS_STACKS":          Enable_Cond_SELF_HAS_STACKS,
		"RANDOM_CHANCE":            Enable_Cond_RANDOM_CHANCE,
		"STACK_COUNT":              Enable_Cond_STACK_COUNT,
		"TARGET_ALLY_HP_LESS_THAN": Enable_Cond_TARGET_ALLY_HP_LESS_THAN,
		"TARGET_HP_LESS_THAN":      Enable_Cond_TARGET_HP_LESS_THAN,
		"LAUNCH":                   Enable_Cond_LAUNCH,
		"CNT":                      Enable_Cond_CNT,
	}
)

func (e Enable_Cond) String() string {
	return Enable_CondName[e]
}

func GetEnable_Cond(s string) Enable_Cond {
	return Enable_CondValue[s]
}

type Equip_State int32

const (
	Equip_State_UNEQUIP Equip_State = iota
	Equip_State_EQUIP
	Equip_State_CNT
)

var (
	Equip_StateName = map[Equip_State]string{
		Equip_State_UNEQUIP: "UNEQUIP",
		Equip_State_EQUIP:   "EQUIP",
		Equip_State_CNT:     "CNT",
	}

	Equip_StateValue = map[string]Equip_State{
		"UNEQUIP": Equip_State_UNEQUIP,
		"EQUIP":   Equip_State_EQUIP,
		"CNT":     Equip_State_CNT,
	}
)

func (e Equip_State) String() string {
	return Equip_StateName[e]
}

func GetEquip_State(s string) Equip_State {
	return Equip_StateValue[s]
}

type Equipment_Attr int32

const (
	Equipment_Attr_WEAR Equipment_Attr = iota
	Equipment_Attr_OBTAIN
	Equipment_Attr_CNT
)

var (
	Equipment_AttrName = map[Equipment_Attr]string{
		Equipment_Attr_WEAR:   "WEAR",
		Equipment_Attr_OBTAIN: "OBTAIN",
		Equipment_Attr_CNT:    "CNT",
	}

	Equipment_AttrValue = map[string]Equipment_Attr{
		"WEAR":   Equipment_Attr_WEAR,
		"OBTAIN": Equipment_Attr_OBTAIN,
		"CNT":    Equipment_Attr_CNT,
	}
)

func (e Equipment_Attr) String() string {
	return Equipment_AttrName[e]
}

func GetEquipment_Attr(s string) Equipment_Attr {
	return Equipment_AttrValue[s]
}

type Equipment_Grade int32

const (
	Equipment_Grade_NORMAL Equipment_Grade = iota
	Equipment_Grade_RARE
	Equipment_Grade_EPIC
	Equipment_Grade_LEGEND
	Equipment_Grade_MYTH
	Equipment_Grade_CNT
)

var (
	Equipment_GradeName = map[Equipment_Grade]string{
		Equipment_Grade_NORMAL: "NORMAL",
		Equipment_Grade_RARE:   "RARE",
		Equipment_Grade_EPIC:   "EPIC",
		Equipment_Grade_LEGEND: "LEGEND",
		Equipment_Grade_MYTH:   "MYTH",
		Equipment_Grade_CNT:    "CNT",
	}

	Equipment_GradeValue = map[string]Equipment_Grade{
		"NORMAL": Equipment_Grade_NORMAL,
		"RARE":   Equipment_Grade_RARE,
		"EPIC":   Equipment_Grade_EPIC,
		"LEGEND": Equipment_Grade_LEGEND,
		"MYTH":   Equipment_Grade_MYTH,
		"CNT":    Equipment_Grade_CNT,
	}
)

func (e Equipment_Grade) String() string {
	return Equipment_GradeName[e]
}

func GetEquipment_Grade(s string) Equipment_Grade {
	return Equipment_GradeValue[s]
}

type Equipment_Property int32

const (
	Equipment_Property_STARGEM_SHEEP Equipment_Property = iota
	Equipment_Property_STARGEM_BULL
	Equipment_Property_STARGEM_TWINS
	Equipment_Property_STARGEM_CRAB
	Equipment_Property_STARGEM_LION
	Equipment_Property_STARGEM_GIRL
	Equipment_Property_STARGEM_SCALES
	Equipment_Property_STARGEM_SCORPION
	Equipment_Property_STARGEM_ARCHER
	Equipment_Property_STARGEM_GOAT
	Equipment_Property_STARGEM_WATER
	Equipment_Property_STARGEM_FISH
	Equipment_Property_STARGEM_GREATBEAR
	Equipment_Property_STARGEM_SNAKE
	Equipment_Property_STARGEM_CRUX
	Equipment_Property_CNT
)

var (
	Equipment_PropertyName = map[Equipment_Property]string{
		Equipment_Property_STARGEM_SHEEP:     "STARGEM_SHEEP",
		Equipment_Property_STARGEM_BULL:      "STARGEM_BULL",
		Equipment_Property_STARGEM_TWINS:     "STARGEM_TWINS",
		Equipment_Property_STARGEM_CRAB:      "STARGEM_CRAB",
		Equipment_Property_STARGEM_LION:      "STARGEM_LION",
		Equipment_Property_STARGEM_GIRL:      "STARGEM_GIRL",
		Equipment_Property_STARGEM_SCALES:    "STARGEM_SCALES",
		Equipment_Property_STARGEM_SCORPION:  "STARGEM_SCORPION",
		Equipment_Property_STARGEM_ARCHER:    "STARGEM_ARCHER",
		Equipment_Property_STARGEM_GOAT:      "STARGEM_GOAT",
		Equipment_Property_STARGEM_WATER:     "STARGEM_WATER",
		Equipment_Property_STARGEM_FISH:      "STARGEM_FISH",
		Equipment_Property_STARGEM_GREATBEAR: "STARGEM_GREATBEAR",
		Equipment_Property_STARGEM_SNAKE:     "STARGEM_SNAKE",
		Equipment_Property_STARGEM_CRUX:      "STARGEM_CRUX",
		Equipment_Property_CNT:               "CNT",
	}

	Equipment_PropertyValue = map[string]Equipment_Property{
		"STARGEM_SHEEP":     Equipment_Property_STARGEM_SHEEP,
		"STARGEM_BULL":      Equipment_Property_STARGEM_BULL,
		"STARGEM_TWINS":     Equipment_Property_STARGEM_TWINS,
		"STARGEM_CRAB":      Equipment_Property_STARGEM_CRAB,
		"STARGEM_LION":      Equipment_Property_STARGEM_LION,
		"STARGEM_GIRL":      Equipment_Property_STARGEM_GIRL,
		"STARGEM_SCALES":    Equipment_Property_STARGEM_SCALES,
		"STARGEM_SCORPION":  Equipment_Property_STARGEM_SCORPION,
		"STARGEM_ARCHER":    Equipment_Property_STARGEM_ARCHER,
		"STARGEM_GOAT":      Equipment_Property_STARGEM_GOAT,
		"STARGEM_WATER":     Equipment_Property_STARGEM_WATER,
		"STARGEM_FISH":      Equipment_Property_STARGEM_FISH,
		"STARGEM_GREATBEAR": Equipment_Property_STARGEM_GREATBEAR,
		"STARGEM_SNAKE":     Equipment_Property_STARGEM_SNAKE,
		"STARGEM_CRUX":      Equipment_Property_STARGEM_CRUX,
		"CNT":               Equipment_Property_CNT,
	}
)

func (e Equipment_Property) String() string {
	return Equipment_PropertyName[e]
}

func GetEquipment_Property(s string) Equipment_Property {
	return Equipment_PropertyValue[s]
}

type Equipment_Slot int32

const (
	Equipment_Slot_STARGEM_SLOT_0 Equipment_Slot = iota
	Equipment_Slot_STARGEM_SLOT_1
	Equipment_Slot_STARGEM_SLOT_2
	Equipment_Slot_STARGEM_SLOT_3
	Equipment_Slot_STARGEM_SLOT_4
	Equipment_Slot_STARGEM_SLOT_5
	Equipment_Slot_CNT
)

var (
	Equipment_SlotName = map[Equipment_Slot]string{
		Equipment_Slot_STARGEM_SLOT_0: "STARGEM_SLOT_0",
		Equipment_Slot_STARGEM_SLOT_1: "STARGEM_SLOT_1",
		Equipment_Slot_STARGEM_SLOT_2: "STARGEM_SLOT_2",
		Equipment_Slot_STARGEM_SLOT_3: "STARGEM_SLOT_3",
		Equipment_Slot_STARGEM_SLOT_4: "STARGEM_SLOT_4",
		Equipment_Slot_STARGEM_SLOT_5: "STARGEM_SLOT_5",
		Equipment_Slot_CNT:            "CNT",
	}

	Equipment_SlotValue = map[string]Equipment_Slot{
		"STARGEM_SLOT_0": Equipment_Slot_STARGEM_SLOT_0,
		"STARGEM_SLOT_1": Equipment_Slot_STARGEM_SLOT_1,
		"STARGEM_SLOT_2": Equipment_Slot_STARGEM_SLOT_2,
		"STARGEM_SLOT_3": Equipment_Slot_STARGEM_SLOT_3,
		"STARGEM_SLOT_4": Equipment_Slot_STARGEM_SLOT_4,
		"STARGEM_SLOT_5": Equipment_Slot_STARGEM_SLOT_5,
		"CNT":            Equipment_Slot_CNT,
	}
)

func (e Equipment_Slot) String() string {
	return Equipment_SlotName[e]
}

func GetEquipment_Slot(s string) Equipment_Slot {
	return Equipment_SlotValue[s]
}

type Equipment_Type int32

const (
	Equipment_Type_STARGEM Equipment_Type = iota
	Equipment_Type_CNT
)

var (
	Equipment_TypeName = map[Equipment_Type]string{
		Equipment_Type_STARGEM: "STARGEM",
		Equipment_Type_CNT:     "CNT",
	}

	Equipment_TypeValue = map[string]Equipment_Type{
		"STARGEM": Equipment_Type_STARGEM,
		"CNT":     Equipment_Type_CNT,
	}
)

func (e Equipment_Type) String() string {
	return Equipment_TypeName[e]
}

func GetEquipment_Type(s string) Equipment_Type {
	return Equipment_TypeValue[s]
}

type Event_Advance int32

const ()

var (
	Event_AdvanceName = map[Event_Advance]string{}

	Event_AdvanceValue = map[string]Event_Advance{}
)

func (e Event_Advance) String() string {
	return Event_AdvanceName[e]
}

func GetEvent_Advance(s string) Event_Advance {
	return Event_AdvanceValue[s]
}

type Event_Class int32

const ()

var (
	Event_ClassName = map[Event_Class]string{}

	Event_ClassValue = map[string]Event_Class{}
)

func (e Event_Class) String() string {
	return Event_ClassName[e]
}

func GetEvent_Class(s string) Event_Class {
	return Event_ClassValue[s]
}

type Fate_Card_Effect int32

const (
	Fate_Card_Effect_ATTACK Fate_Card_Effect = iota
	Fate_Card_Effect_DEFENSE
	Fate_Card_Effect_UTILITY
	Fate_Card_Effect_DEBUFF
	Fate_Card_Effect_SUPPORT
	Fate_Card_Effect_NONE
	Fate_Card_Effect_CURSE
	Fate_Card_Effect_SOUL_TRADE
	Fate_Card_Effect_ATTENTION
	Fate_Card_Effect_STUN
	Fate_Card_Effect_CNT
)

var (
	Fate_Card_EffectName = map[Fate_Card_Effect]string{
		Fate_Card_Effect_ATTACK:     "ATTACK",
		Fate_Card_Effect_DEFENSE:    "DEFENSE",
		Fate_Card_Effect_UTILITY:    "UTILITY",
		Fate_Card_Effect_DEBUFF:     "DEBUFF",
		Fate_Card_Effect_SUPPORT:    "SUPPORT",
		Fate_Card_Effect_NONE:       "NONE",
		Fate_Card_Effect_CURSE:      "CURSE",
		Fate_Card_Effect_SOUL_TRADE: "SOUL_TRADE",
		Fate_Card_Effect_ATTENTION:  "ATTENTION",
		Fate_Card_Effect_STUN:       "STUN",
		Fate_Card_Effect_CNT:        "CNT",
	}

	Fate_Card_EffectValue = map[string]Fate_Card_Effect{
		"ATTACK":     Fate_Card_Effect_ATTACK,
		"DEFENSE":    Fate_Card_Effect_DEFENSE,
		"UTILITY":    Fate_Card_Effect_UTILITY,
		"DEBUFF":     Fate_Card_Effect_DEBUFF,
		"SUPPORT":    Fate_Card_Effect_SUPPORT,
		"NONE":       Fate_Card_Effect_NONE,
		"CURSE":      Fate_Card_Effect_CURSE,
		"SOUL_TRADE": Fate_Card_Effect_SOUL_TRADE,
		"ATTENTION":  Fate_Card_Effect_ATTENTION,
		"STUN":       Fate_Card_Effect_STUN,
		"CNT":        Fate_Card_Effect_CNT,
	}
)

func (e Fate_Card_Effect) String() string {
	return Fate_Card_EffectName[e]
}

func GetFate_Card_Effect(s string) Fate_Card_Effect {
	return Fate_Card_EffectValue[s]
}

type Fate_Card_Grade int32

const (
	Fate_Card_Grade_GRADE_C Fate_Card_Grade = iota
	Fate_Card_Grade_GRADE_B
	Fate_Card_Grade_GRADE_A
	Fate_Card_Grade_CNT
)

var (
	Fate_Card_GradeName = map[Fate_Card_Grade]string{
		Fate_Card_Grade_GRADE_C: "GRADE_C",
		Fate_Card_Grade_GRADE_B: "GRADE_B",
		Fate_Card_Grade_GRADE_A: "GRADE_A",
		Fate_Card_Grade_CNT:     "CNT",
	}

	Fate_Card_GradeValue = map[string]Fate_Card_Grade{
		"GRADE_C": Fate_Card_Grade_GRADE_C,
		"GRADE_B": Fate_Card_Grade_GRADE_B,
		"GRADE_A": Fate_Card_Grade_GRADE_A,
		"CNT":     Fate_Card_Grade_CNT,
	}
)

func (e Fate_Card_Grade) String() string {
	return Fate_Card_GradeName[e]
}

func GetFate_Card_Grade(s string) Fate_Card_Grade {
	return Fate_Card_GradeValue[s]
}

type Gacha_Pool_Grade int32

const (
	Gacha_Pool_Grade_D Gacha_Pool_Grade = iota
	Gacha_Pool_Grade_C
	Gacha_Pool_Grade_B
	Gacha_Pool_Grade_A
	Gacha_Pool_Grade_CNT
)

var (
	Gacha_Pool_GradeName = map[Gacha_Pool_Grade]string{
		Gacha_Pool_Grade_D:   "D",
		Gacha_Pool_Grade_C:   "C",
		Gacha_Pool_Grade_B:   "B",
		Gacha_Pool_Grade_A:   "A",
		Gacha_Pool_Grade_CNT: "CNT",
	}

	Gacha_Pool_GradeValue = map[string]Gacha_Pool_Grade{
		"D":   Gacha_Pool_Grade_D,
		"C":   Gacha_Pool_Grade_C,
		"B":   Gacha_Pool_Grade_B,
		"A":   Gacha_Pool_Grade_A,
		"CNT": Gacha_Pool_Grade_CNT,
	}
)

func (e Gacha_Pool_Grade) String() string {
	return Gacha_Pool_GradeName[e]
}

func GetGacha_Pool_Grade(s string) Gacha_Pool_Grade {
	return Gacha_Pool_GradeValue[s]
}

type Gacha_Slot_Type int32

const (
	Gacha_Slot_Type_REGULAR Gacha_Slot_Type = iota
	Gacha_Slot_Type_PICKUP
	Gacha_Slot_Type_CNT
)

var (
	Gacha_Slot_TypeName = map[Gacha_Slot_Type]string{
		Gacha_Slot_Type_REGULAR: "REGULAR",
		Gacha_Slot_Type_PICKUP:  "PICKUP",
		Gacha_Slot_Type_CNT:     "CNT",
	}

	Gacha_Slot_TypeValue = map[string]Gacha_Slot_Type{
		"REGULAR": Gacha_Slot_Type_REGULAR,
		"PICKUP":  Gacha_Slot_Type_PICKUP,
		"CNT":     Gacha_Slot_Type_CNT,
	}
)

func (e Gacha_Slot_Type) String() string {
	return Gacha_Slot_TypeName[e]
}

func GetGacha_Slot_Type(s string) Gacha_Slot_Type {
	return Gacha_Slot_TypeValue[s]
}

type Graph int32

const (
	Graph_NONE Graph = iota
	Graph_BELL_CURVE
	Graph_MAX
	Graph_MIN
	Graph_CNT
)

var (
	GraphName = map[Graph]string{
		Graph_NONE:       "NONE",
		Graph_BELL_CURVE: "BELL_CURVE",
		Graph_MAX:        "MAX",
		Graph_MIN:        "MIN",
		Graph_CNT:        "CNT",
	}

	GraphValue = map[string]Graph{
		"NONE":       Graph_NONE,
		"BELL_CURVE": Graph_BELL_CURVE,
		"MAX":        Graph_MAX,
		"MIN":        Graph_MIN,
		"CNT":        Graph_CNT,
	}
)

func (e Graph) String() string {
	return GraphName[e]
}

func GetGraph(s string) Graph {
	return GraphValue[s]
}

type Guide_Type int32

const (
	Guide_Type_CHARACTER Guide_Type = iota
	Guide_Type_DUEL
	Guide_Type_SIGNATURE_WEAPON
	Guide_Type_TIER
	Guide_Type_CNT
)

var (
	Guide_TypeName = map[Guide_Type]string{
		Guide_Type_CHARACTER:        "CHARACTER",
		Guide_Type_DUEL:             "DUEL",
		Guide_Type_SIGNATURE_WEAPON: "SIGNATURE_WEAPON",
		Guide_Type_TIER:             "TIER",
		Guide_Type_CNT:              "CNT",
	}

	Guide_TypeValue = map[string]Guide_Type{
		"CHARACTER":        Guide_Type_CHARACTER,
		"DUEL":             Guide_Type_DUEL,
		"SIGNATURE_WEAPON": Guide_Type_SIGNATURE_WEAPON,
		"TIER":             Guide_Type_TIER,
		"CNT":              Guide_Type_CNT,
	}
)

func (e Guide_Type) String() string {
	return Guide_TypeName[e]
}

func GetGuide_Type(s string) Guide_Type {
	return Guide_TypeValue[s]
}

type Hashtag int32

const (
	Hashtag_SKYSCRAPER Hashtag = iota
	Hashtag_SIGHT
	Hashtag_SLUM
	Hashtag_SILENT
	Hashtag_NOISY
	Hashtag_FOOD
	Hashtag_MYSTERY
	Hashtag_RUMOR
	Hashtag_PEACEFUL
	Hashtag_CNT
)

var (
	HashtagName = map[Hashtag]string{
		Hashtag_SKYSCRAPER: "SKYSCRAPER",
		Hashtag_SIGHT:      "SIGHT",
		Hashtag_SLUM:       "SLUM",
		Hashtag_SILENT:     "SILENT",
		Hashtag_NOISY:      "NOISY",
		Hashtag_FOOD:       "FOOD",
		Hashtag_MYSTERY:    "MYSTERY",
		Hashtag_RUMOR:      "RUMOR",
		Hashtag_PEACEFUL:   "PEACEFUL",
		Hashtag_CNT:        "CNT",
	}

	HashtagValue = map[string]Hashtag{
		"SKYSCRAPER": Hashtag_SKYSCRAPER,
		"SIGHT":      Hashtag_SIGHT,
		"SLUM":       Hashtag_SLUM,
		"SILENT":     Hashtag_SILENT,
		"NOISY":      Hashtag_NOISY,
		"FOOD":       Hashtag_FOOD,
		"MYSTERY":    Hashtag_MYSTERY,
		"RUMOR":      Hashtag_RUMOR,
		"PEACEFUL":   Hashtag_PEACEFUL,
		"CNT":        Hashtag_CNT,
	}
)

func (e Hashtag) String() string {
	return HashtagName[e]
}

func GetHashtag(s string) Hashtag {
	return HashtagValue[s]
}

type Hashtag_Variable int32

const (
	Hashtag_Variable_RAIN Hashtag_Variable = iota
	Hashtag_Variable_STORM
	Hashtag_Variable_SNOW
	Hashtag_Variable_FESTIVAL
	Hashtag_Variable_BREEZE
	Hashtag_Variable_FIRE
	Hashtag_Variable_STAR_SHOWER
	Hashtag_Variable_CNT
)

var (
	Hashtag_VariableName = map[Hashtag_Variable]string{
		Hashtag_Variable_RAIN:        "RAIN",
		Hashtag_Variable_STORM:       "STORM",
		Hashtag_Variable_SNOW:        "SNOW",
		Hashtag_Variable_FESTIVAL:    "FESTIVAL",
		Hashtag_Variable_BREEZE:      "BREEZE",
		Hashtag_Variable_FIRE:        "FIRE",
		Hashtag_Variable_STAR_SHOWER: "STAR_SHOWER",
		Hashtag_Variable_CNT:         "CNT",
	}

	Hashtag_VariableValue = map[string]Hashtag_Variable{
		"RAIN":        Hashtag_Variable_RAIN,
		"STORM":       Hashtag_Variable_STORM,
		"SNOW":        Hashtag_Variable_SNOW,
		"FESTIVAL":    Hashtag_Variable_FESTIVAL,
		"BREEZE":      Hashtag_Variable_BREEZE,
		"FIRE":        Hashtag_Variable_FIRE,
		"STAR_SHOWER": Hashtag_Variable_STAR_SHOWER,
		"CNT":         Hashtag_Variable_CNT,
	}
)

func (e Hashtag_Variable) String() string {
	return Hashtag_VariableName[e]
}

func GetHashtag_Variable(s string) Hashtag_Variable {
	return Hashtag_VariableValue[s]
}

type Item int32

const (
	Item_NONBATTLE Item = iota
	Item_BATTLE
	Item_CNT
)

var (
	ItemName = map[Item]string{
		Item_NONBATTLE: "NONBATTLE",
		Item_BATTLE:    "BATTLE",
		Item_CNT:       "CNT",
	}

	ItemValue = map[string]Item{
		"NONBATTLE": Item_NONBATTLE,
		"BATTLE":    Item_BATTLE,
		"CNT":       Item_CNT,
	}
)

func (e Item) String() string {
	return ItemName[e]
}

func GetItem(s string) Item {
	return ItemValue[s]
}

type Item_Grade int32

const (
	Item_Grade_NORMAL Item_Grade = iota
	Item_Grade_RARE
	Item_Grade_EPIC
	Item_Grade_LEGEND
	Item_Grade_MYTH
	Item_Grade_CNT
)

var (
	Item_GradeName = map[Item_Grade]string{
		Item_Grade_NORMAL: "NORMAL",
		Item_Grade_RARE:   "RARE",
		Item_Grade_EPIC:   "EPIC",
		Item_Grade_LEGEND: "LEGEND",
		Item_Grade_MYTH:   "MYTH",
		Item_Grade_CNT:    "CNT",
	}

	Item_GradeValue = map[string]Item_Grade{
		"NORMAL": Item_Grade_NORMAL,
		"RARE":   Item_Grade_RARE,
		"EPIC":   Item_Grade_EPIC,
		"LEGEND": Item_Grade_LEGEND,
		"MYTH":   Item_Grade_MYTH,
		"CNT":    Item_Grade_CNT,
	}
)

func (e Item_Grade) String() string {
	return Item_GradeName[e]
}

func GetItem_Grade(s string) Item_Grade {
	return Item_GradeValue[s]
}

type Item_Sub int32

const (
	Item_Sub_CHARACTER_EXP Item_Sub = iota
	Item_Sub_EQUIPMENT_EXP
	Item_Sub_AFFECTION_EXP
	Item_Sub_CHARACTER_STAMINA
	Item_Sub_BATTLE
	Item_Sub_CNT
)

var (
	Item_SubName = map[Item_Sub]string{
		Item_Sub_CHARACTER_EXP:     "CHARACTER_EXP",
		Item_Sub_EQUIPMENT_EXP:     "EQUIPMENT_EXP",
		Item_Sub_AFFECTION_EXP:     "AFFECTION_EXP",
		Item_Sub_CHARACTER_STAMINA: "CHARACTER_STAMINA",
		Item_Sub_BATTLE:            "BATTLE",
		Item_Sub_CNT:               "CNT",
	}

	Item_SubValue = map[string]Item_Sub{
		"CHARACTER_EXP":     Item_Sub_CHARACTER_EXP,
		"EQUIPMENT_EXP":     Item_Sub_EQUIPMENT_EXP,
		"AFFECTION_EXP":     Item_Sub_AFFECTION_EXP,
		"CHARACTER_STAMINA": Item_Sub_CHARACTER_STAMINA,
		"BATTLE":            Item_Sub_BATTLE,
		"CNT":               Item_Sub_CNT,
	}
)

func (e Item_Sub) String() string {
	return Item_SubName[e]
}

func GetItem_Sub(s string) Item_Sub {
	return Item_SubValue[s]
}

type Language int32

const (
	Language_LANGUAGE_KR Language = iota
	Language_LANGUAGE_EN
	Language_CNT
)

var (
	LanguageName = map[Language]string{
		Language_LANGUAGE_KR: "LANGUAGE_KR",
		Language_LANGUAGE_EN: "LANGUAGE_EN",
		Language_CNT:         "CNT",
	}

	LanguageValue = map[string]Language{
		"LANGUAGE_KR": Language_LANGUAGE_KR,
		"LANGUAGE_EN": Language_LANGUAGE_EN,
		"CNT":         Language_CNT,
	}
)

func (e Language) String() string {
	return LanguageName[e]
}

func GetLanguage(s string) Language {
	return LanguageValue[s]
}

type Mail_Status int32

const (
	Mail_Status_NEW Mail_Status = iota
	Mail_Status_READ
	Mail_Status_CONFIRM
	Mail_Status_CNT
)

var (
	Mail_StatusName = map[Mail_Status]string{
		Mail_Status_NEW:     "NEW",
		Mail_Status_READ:    "READ",
		Mail_Status_CONFIRM: "CONFIRM",
		Mail_Status_CNT:     "CNT",
	}

	Mail_StatusValue = map[string]Mail_Status{
		"NEW":     Mail_Status_NEW,
		"READ":    Mail_Status_READ,
		"CONFIRM": Mail_Status_CONFIRM,
		"CNT":     Mail_Status_CNT,
	}
)

func (e Mail_Status) String() string {
	return Mail_StatusName[e]
}

func GetMail_Status(s string) Mail_Status {
	return Mail_StatusValue[s]
}

type Monster_Class int32

const (
	Monster_Class_BOSS Monster_Class = iota
	Monster_Class_ELITE
	Monster_Class_MINOR
	Monster_Class_TOWER
	Monster_Class_CNT
)

var (
	Monster_ClassName = map[Monster_Class]string{
		Monster_Class_BOSS:  "BOSS",
		Monster_Class_ELITE: "ELITE",
		Monster_Class_MINOR: "MINOR",
		Monster_Class_TOWER: "TOWER",
		Monster_Class_CNT:   "CNT",
	}

	Monster_ClassValue = map[string]Monster_Class{
		"BOSS":  Monster_Class_BOSS,
		"ELITE": Monster_Class_ELITE,
		"MINOR": Monster_Class_MINOR,
		"TOWER": Monster_Class_TOWER,
		"CNT":   Monster_Class_CNT,
	}
)

func (e Monster_Class) String() string {
	return Monster_ClassName[e]
}

func GetMonster_Class(s string) Monster_Class {
	return Monster_ClassValue[s]
}

type Node_Branch int32

const (
	Node_Branch_NONE Node_Branch = iota
	Node_Branch_RANDOM
	Node_Branch_GIFT
	Node_Branch_QUEST_RESPONSE
	Node_Branch_CNT
)

var (
	Node_BranchName = map[Node_Branch]string{
		Node_Branch_NONE:           "NONE",
		Node_Branch_RANDOM:         "RANDOM",
		Node_Branch_GIFT:           "GIFT",
		Node_Branch_QUEST_RESPONSE: "QUEST_RESPONSE",
		Node_Branch_CNT:            "CNT",
	}

	Node_BranchValue = map[string]Node_Branch{
		"NONE":           Node_Branch_NONE,
		"RANDOM":         Node_Branch_RANDOM,
		"GIFT":           Node_Branch_GIFT,
		"QUEST_RESPONSE": Node_Branch_QUEST_RESPONSE,
		"CNT":            Node_Branch_CNT,
	}
)

func (e Node_Branch) String() string {
	return Node_BranchName[e]
}

func GetNode_Branch(s string) Node_Branch {
	return Node_BranchValue[s]
}

type Oracle_Broadcast_Type int32

const (
	Oracle_Broadcast_Type_EVENT Oracle_Broadcast_Type = iota
	Oracle_Broadcast_Type_REGULAR
	Oracle_Broadcast_Type_IRREGULAR
	Oracle_Broadcast_Type_NOGET
	Oracle_Broadcast_Type_CNT
)

var (
	Oracle_Broadcast_TypeName = map[Oracle_Broadcast_Type]string{
		Oracle_Broadcast_Type_EVENT:     "EVENT",
		Oracle_Broadcast_Type_REGULAR:   "REGULAR",
		Oracle_Broadcast_Type_IRREGULAR: "IRREGULAR",
		Oracle_Broadcast_Type_NOGET:     "NOGET",
		Oracle_Broadcast_Type_CNT:       "CNT",
	}

	Oracle_Broadcast_TypeValue = map[string]Oracle_Broadcast_Type{
		"EVENT":     Oracle_Broadcast_Type_EVENT,
		"REGULAR":   Oracle_Broadcast_Type_REGULAR,
		"IRREGULAR": Oracle_Broadcast_Type_IRREGULAR,
		"NOGET":     Oracle_Broadcast_Type_NOGET,
		"CNT":       Oracle_Broadcast_Type_CNT,
	}
)

func (e Oracle_Broadcast_Type) String() string {
	return Oracle_Broadcast_TypeName[e]
}

func GetOracle_Broadcast_Type(s string) Oracle_Broadcast_Type {
	return Oracle_Broadcast_TypeValue[s]
}

type Oracle_Chat_Format int32

const (
	Oracle_Chat_Format_COMMON Oracle_Chat_Format = iota
	Oracle_Chat_Format_SPECIAL
	Oracle_Chat_Format_GOD
	Oracle_Chat_Format_USER
	Oracle_Chat_Format_CNT
)

var (
	Oracle_Chat_FormatName = map[Oracle_Chat_Format]string{
		Oracle_Chat_Format_COMMON:  "COMMON",
		Oracle_Chat_Format_SPECIAL: "SPECIAL",
		Oracle_Chat_Format_GOD:     "GOD",
		Oracle_Chat_Format_USER:    "USER",
		Oracle_Chat_Format_CNT:     "CNT",
	}

	Oracle_Chat_FormatValue = map[string]Oracle_Chat_Format{
		"COMMON":  Oracle_Chat_Format_COMMON,
		"SPECIAL": Oracle_Chat_Format_SPECIAL,
		"GOD":     Oracle_Chat_Format_GOD,
		"USER":    Oracle_Chat_Format_USER,
		"CNT":     Oracle_Chat_Format_CNT,
	}
)

func (e Oracle_Chat_Format) String() string {
	return Oracle_Chat_FormatName[e]
}

func GetOracle_Chat_Format(s string) Oracle_Chat_Format {
	return Oracle_Chat_FormatValue[s]
}

type Order_Type int32

const (
	Order_Type_FARTHEST Order_Type = iota
	Order_Type_LOW_HP
	Order_Type_CNT
)

var (
	Order_TypeName = map[Order_Type]string{
		Order_Type_FARTHEST: "FARTHEST",
		Order_Type_LOW_HP:   "LOW_HP",
		Order_Type_CNT:      "CNT",
	}

	Order_TypeValue = map[string]Order_Type{
		"FARTHEST": Order_Type_FARTHEST,
		"LOW_HP":   Order_Type_LOW_HP,
		"CNT":      Order_Type_CNT,
	}
)

func (e Order_Type) String() string {
	return Order_TypeName[e]
}

func GetOrder_Type(s string) Order_Type {
	return Order_TypeValue[s]
}

type Recipe_Category int32

const (
	Recipe_Category_OFFENSIVE Recipe_Category = iota
	Recipe_Category_DEFFENSIVE
	Recipe_Category_ESSENTIAL
	Recipe_Category_UNIQUE
	Recipe_Category_STAMINA
	Recipe_Category_AFFECTION
	Recipe_Category_EXP
	Recipe_Category_STARGEM_EXP
	Recipe_Category_UTILITY
	Recipe_Category_CONDENSATION
	Recipe_Category_DECONDENSATION
	Recipe_Category_ALTERATION
	Recipe_Category_WARRIOR
	Recipe_Category_WIZARD
	Recipe_Category_RANGER
	Recipe_Category_ASSASSIN
	Recipe_Category_KNIGHT
	Recipe_Category_BARD
	Recipe_Category_PRIEST
	Recipe_Category_BATTLE_MAGE
	Recipe_Category_FAVORITES
	Recipe_Category_CNT
)

var (
	Recipe_CategoryName = map[Recipe_Category]string{
		Recipe_Category_OFFENSIVE:      "OFFENSIVE",
		Recipe_Category_DEFFENSIVE:     "DEFFENSIVE",
		Recipe_Category_ESSENTIAL:      "ESSENTIAL",
		Recipe_Category_UNIQUE:         "UNIQUE",
		Recipe_Category_STAMINA:        "STAMINA",
		Recipe_Category_AFFECTION:      "AFFECTION",
		Recipe_Category_EXP:            "EXP",
		Recipe_Category_STARGEM_EXP:    "STARGEM_EXP",
		Recipe_Category_UTILITY:        "UTILITY",
		Recipe_Category_CONDENSATION:   "CONDENSATION",
		Recipe_Category_DECONDENSATION: "DECONDENSATION",
		Recipe_Category_ALTERATION:     "ALTERATION",
		Recipe_Category_WARRIOR:        "WARRIOR",
		Recipe_Category_WIZARD:         "WIZARD",
		Recipe_Category_RANGER:         "RANGER",
		Recipe_Category_ASSASSIN:       "ASSASSIN",
		Recipe_Category_KNIGHT:         "KNIGHT",
		Recipe_Category_BARD:           "BARD",
		Recipe_Category_PRIEST:         "PRIEST",
		Recipe_Category_BATTLE_MAGE:    "BATTLE_MAGE",
		Recipe_Category_FAVORITES:      "FAVORITES",
		Recipe_Category_CNT:            "CNT",
	}

	Recipe_CategoryValue = map[string]Recipe_Category{
		"OFFENSIVE":      Recipe_Category_OFFENSIVE,
		"DEFFENSIVE":     Recipe_Category_DEFFENSIVE,
		"ESSENTIAL":      Recipe_Category_ESSENTIAL,
		"UNIQUE":         Recipe_Category_UNIQUE,
		"STAMINA":        Recipe_Category_STAMINA,
		"AFFECTION":      Recipe_Category_AFFECTION,
		"EXP":            Recipe_Category_EXP,
		"STARGEM_EXP":    Recipe_Category_STARGEM_EXP,
		"UTILITY":        Recipe_Category_UTILITY,
		"CONDENSATION":   Recipe_Category_CONDENSATION,
		"DECONDENSATION": Recipe_Category_DECONDENSATION,
		"ALTERATION":     Recipe_Category_ALTERATION,
		"WARRIOR":        Recipe_Category_WARRIOR,
		"WIZARD":         Recipe_Category_WIZARD,
		"RANGER":         Recipe_Category_RANGER,
		"ASSASSIN":       Recipe_Category_ASSASSIN,
		"KNIGHT":         Recipe_Category_KNIGHT,
		"BARD":           Recipe_Category_BARD,
		"PRIEST":         Recipe_Category_PRIEST,
		"BATTLE_MAGE":    Recipe_Category_BATTLE_MAGE,
		"FAVORITES":      Recipe_Category_FAVORITES,
		"CNT":            Recipe_Category_CNT,
	}
)

func (e Recipe_Category) String() string {
	return Recipe_CategoryName[e]
}

func GetRecipe_Category(s string) Recipe_Category {
	return Recipe_CategoryValue[s]
}

type Remove_Cond int32

const (
	Remove_Cond_GET_HIT Remove_Cond = iota
	Remove_Cond_CNT
)

var (
	Remove_CondName = map[Remove_Cond]string{
		Remove_Cond_GET_HIT: "GET_HIT",
		Remove_Cond_CNT:     "CNT",
	}

	Remove_CondValue = map[string]Remove_Cond{
		"GET_HIT": Remove_Cond_GET_HIT,
		"CNT":     Remove_Cond_CNT,
	}
)

func (e Remove_Cond) String() string {
	return Remove_CondName[e]
}

func GetRemove_Cond(s string) Remove_Cond {
	return Remove_CondValue[s]
}

type Request_Type int32

const (
	Request_Type_ASSET Request_Type = iota
	Request_Type_STARGEM
	Request_Type_ITEM
	Request_Type_TILE
	Request_Type_CNT
)

var (
	Request_TypeName = map[Request_Type]string{
		Request_Type_ASSET:   "ASSET",
		Request_Type_STARGEM: "STARGEM",
		Request_Type_ITEM:    "ITEM",
		Request_Type_TILE:    "TILE",
		Request_Type_CNT:     "CNT",
	}

	Request_TypeValue = map[string]Request_Type{
		"ASSET":   Request_Type_ASSET,
		"STARGEM": Request_Type_STARGEM,
		"ITEM":    Request_Type_ITEM,
		"TILE":    Request_Type_TILE,
		"CNT":     Request_Type_CNT,
	}
)

func (e Request_Type) String() string {
	return Request_TypeName[e]
}

func GetRequest_Type(s string) Request_Type {
	return Request_TypeValue[s]
}

type Resource_Type int32

const (
	Resource_Type_COMMON_GRADE_BG Resource_Type = iota
	Resource_Type_COMMON_GRADE_TEXT
	Resource_Type_CNT
)

var (
	Resource_TypeName = map[Resource_Type]string{
		Resource_Type_COMMON_GRADE_BG:   "COMMON_GRADE_BG",
		Resource_Type_COMMON_GRADE_TEXT: "COMMON_GRADE_TEXT",
		Resource_Type_CNT:               "CNT",
	}

	Resource_TypeValue = map[string]Resource_Type{
		"COMMON_GRADE_BG":   Resource_Type_COMMON_GRADE_BG,
		"COMMON_GRADE_TEXT": Resource_Type_COMMON_GRADE_TEXT,
		"CNT":               Resource_Type_CNT,
	}
)

func (e Resource_Type) String() string {
	return Resource_TypeName[e]
}

func GetResource_Type(s string) Resource_Type {
	return Resource_TypeValue[s]
}

type Result_Time int32

const (
	Result_Time_STARTCASTING Result_Time = iota
	Result_Time_ENDCASTING
	Result_Time_CNT
)

var (
	Result_TimeName = map[Result_Time]string{
		Result_Time_STARTCASTING: "STARTCASTING",
		Result_Time_ENDCASTING:   "ENDCASTING",
		Result_Time_CNT:          "CNT",
	}

	Result_TimeValue = map[string]Result_Time{
		"STARTCASTING": Result_Time_STARTCASTING,
		"ENDCASTING":   Result_Time_ENDCASTING,
		"CNT":          Result_Time_CNT,
	}
)

func (e Result_Time) String() string {
	return Result_TimeName[e]
}

func GetResult_Time(s string) Result_Time {
	return Result_TimeValue[s]
}

type Shop_Type int32

const (
	Shop_Type_GACHA Shop_Type = iota
	Shop_Type_CNT
)

var (
	Shop_TypeName = map[Shop_Type]string{
		Shop_Type_GACHA: "GACHA",
		Shop_Type_CNT:   "CNT",
	}

	Shop_TypeValue = map[string]Shop_Type{
		"GACHA": Shop_Type_GACHA,
		"CNT":   Shop_Type_CNT,
	}
)

func (e Shop_Type) String() string {
	return Shop_TypeName[e]
}

func GetShop_Type(s string) Shop_Type {
	return Shop_TypeValue[s]
}

type Skill_Subtype int32

const (
	Skill_Subtype_MELEE Skill_Subtype = iota
	Skill_Subtype_PROJECTILE
	Skill_Subtype_EFFECT
	Skill_Subtype_RADIAL
	Skill_Subtype_CNT
)

var (
	Skill_SubtypeName = map[Skill_Subtype]string{
		Skill_Subtype_MELEE:      "MELEE",
		Skill_Subtype_PROJECTILE: "PROJECTILE",
		Skill_Subtype_EFFECT:     "EFFECT",
		Skill_Subtype_RADIAL:     "RADIAL",
		Skill_Subtype_CNT:        "CNT",
	}

	Skill_SubtypeValue = map[string]Skill_Subtype{
		"MELEE":      Skill_Subtype_MELEE,
		"PROJECTILE": Skill_Subtype_PROJECTILE,
		"EFFECT":     Skill_Subtype_EFFECT,
		"RADIAL":     Skill_Subtype_RADIAL,
		"CNT":        Skill_Subtype_CNT,
	}
)

func (e Skill_Subtype) String() string {
	return Skill_SubtypeName[e]
}

func GetSkill_Subtype(s string) Skill_Subtype {
	return Skill_SubtypeValue[s]
}

type Skill_Trigger int32

const (
	Skill_Trigger_NEXT_NODE Skill_Trigger = iota
	Skill_Trigger_BASIC_ATTACK
	Skill_Trigger_GET_HIT
	Skill_Trigger_RANDOM_CHANCE
	Skill_Trigger_ALLY_LEFT
	Skill_Trigger_LEADER
	Skill_Trigger_CNT
)

var (
	Skill_TriggerName = map[Skill_Trigger]string{
		Skill_Trigger_NEXT_NODE:     "NEXT_NODE",
		Skill_Trigger_BASIC_ATTACK:  "BASIC_ATTACK",
		Skill_Trigger_GET_HIT:       "GET_HIT",
		Skill_Trigger_RANDOM_CHANCE: "RANDOM_CHANCE",
		Skill_Trigger_ALLY_LEFT:     "ALLY_LEFT",
		Skill_Trigger_LEADER:        "LEADER",
		Skill_Trigger_CNT:           "CNT",
	}

	Skill_TriggerValue = map[string]Skill_Trigger{
		"NEXT_NODE":     Skill_Trigger_NEXT_NODE,
		"BASIC_ATTACK":  Skill_Trigger_BASIC_ATTACK,
		"GET_HIT":       Skill_Trigger_GET_HIT,
		"RANDOM_CHANCE": Skill_Trigger_RANDOM_CHANCE,
		"ALLY_LEFT":     Skill_Trigger_ALLY_LEFT,
		"LEADER":        Skill_Trigger_LEADER,
		"CNT":           Skill_Trigger_CNT,
	}
)

func (e Skill_Trigger) String() string {
	return Skill_TriggerName[e]
}

func GetSkill_Trigger(s string) Skill_Trigger {
	return Skill_TriggerValue[s]
}

type Skill_Type int32

const (
	Skill_Type_ULTIMATE Skill_Type = iota
	Skill_Type_ACTIVE
	Skill_Type_PASSIVE
	Skill_Type_MOVEMENT
	Skill_Type_BASIC
	Skill_Type_LEADER
	Skill_Type_CONDITIONAL
	Skill_Type_DAMAGE
	Skill_Type_IMMEDIATE
	Skill_Type_HEAL
	Skill_Type_CNT
)

var (
	Skill_TypeName = map[Skill_Type]string{
		Skill_Type_ULTIMATE:    "ULTIMATE",
		Skill_Type_ACTIVE:      "ACTIVE",
		Skill_Type_PASSIVE:     "PASSIVE",
		Skill_Type_MOVEMENT:    "MOVEMENT",
		Skill_Type_BASIC:       "BASIC",
		Skill_Type_LEADER:      "LEADER",
		Skill_Type_CONDITIONAL: "CONDITIONAL",
		Skill_Type_DAMAGE:      "DAMAGE",
		Skill_Type_IMMEDIATE:   "IMMEDIATE",
		Skill_Type_HEAL:        "HEAL",
		Skill_Type_CNT:         "CNT",
	}

	Skill_TypeValue = map[string]Skill_Type{
		"ULTIMATE":    Skill_Type_ULTIMATE,
		"ACTIVE":      Skill_Type_ACTIVE,
		"PASSIVE":     Skill_Type_PASSIVE,
		"MOVEMENT":    Skill_Type_MOVEMENT,
		"BASIC":       Skill_Type_BASIC,
		"LEADER":      Skill_Type_LEADER,
		"CONDITIONAL": Skill_Type_CONDITIONAL,
		"DAMAGE":      Skill_Type_DAMAGE,
		"IMMEDIATE":   Skill_Type_IMMEDIATE,
		"HEAL":        Skill_Type_HEAL,
		"CNT":         Skill_Type_CNT,
	}
)

func (e Skill_Type) String() string {
	return Skill_TypeName[e]
}

func GetSkill_Type(s string) Skill_Type {
	return Skill_TypeValue[s]
}

type Stat int32

const (
	Stat_ATK Stat = iota
	Stat_DEF
	Stat_HP
	Stat_CRITD
	Stat_CRIT
	Stat_TD
	Stat_REC
	Stat_CDR
	Stat_SHI
	Stat_NEU
	Stat_EVA
	Stat_ACC
	Stat_SERES
	Stat_ATKSPD
	Stat_DM
	Stat_DMGRES
	Stat_CRUX
	Stat_DRAIN
	Stat_NEUREC
	Stat_PERSISTANCE
	Stat_CRITR
	Stat_CRITDR
	Stat_BLOCKE
	Stat_BLOCK
	Stat_SPD
	Stat_RANGE
	Stat_CNT
)

var (
	StatName = map[Stat]string{
		Stat_ATK:         "ATK",
		Stat_DEF:         "DEF",
		Stat_HP:          "HP",
		Stat_CRITD:       "CRITD",
		Stat_CRIT:        "CRIT",
		Stat_TD:          "TD",
		Stat_REC:         "REC",
		Stat_CDR:         "CDR",
		Stat_SHI:         "SHI",
		Stat_NEU:         "NEU",
		Stat_EVA:         "EVA",
		Stat_ACC:         "ACC",
		Stat_SERES:       "SERES",
		Stat_ATKSPD:      "ATKSPD",
		Stat_DM:          "DM",
		Stat_DMGRES:      "DMGRES",
		Stat_CRUX:        "CRUX",
		Stat_DRAIN:       "DRAIN",
		Stat_NEUREC:      "NEUREC",
		Stat_PERSISTANCE: "PERSISTANCE",
		Stat_CRITR:       "CRITR",
		Stat_CRITDR:      "CRITDR",
		Stat_BLOCKE:      "BLOCKE",
		Stat_BLOCK:       "BLOCK",
		Stat_SPD:         "SPD",
		Stat_RANGE:       "RANGE",
		Stat_CNT:         "CNT",
	}

	StatValue = map[string]Stat{
		"ATK":         Stat_ATK,
		"DEF":         Stat_DEF,
		"HP":          Stat_HP,
		"CRITD":       Stat_CRITD,
		"CRIT":        Stat_CRIT,
		"TD":          Stat_TD,
		"REC":         Stat_REC,
		"CDR":         Stat_CDR,
		"SHI":         Stat_SHI,
		"NEU":         Stat_NEU,
		"EVA":         Stat_EVA,
		"ACC":         Stat_ACC,
		"SERES":       Stat_SERES,
		"ATKSPD":      Stat_ATKSPD,
		"DM":          Stat_DM,
		"DMGRES":      Stat_DMGRES,
		"CRUX":        Stat_CRUX,
		"DRAIN":       Stat_DRAIN,
		"NEUREC":      Stat_NEUREC,
		"PERSISTANCE": Stat_PERSISTANCE,
		"CRITR":       Stat_CRITR,
		"CRITDR":      Stat_CRITDR,
		"BLOCKE":      Stat_BLOCKE,
		"BLOCK":       Stat_BLOCK,
		"SPD":         Stat_SPD,
		"RANGE":       Stat_RANGE,
		"CNT":         Stat_CNT,
	}
)

func (e Stat) String() string {
	return StatName[e]
}

func GetStat(s string) Stat {
	return StatValue[s]
}

type Stat_Subtype int32

const (
	Stat_Subtype_CURRENT Stat_Subtype = iota
	Stat_Subtype_MAX
	Stat_Subtype_PROCESSED
	Stat_Subtype_CNT
)

var (
	Stat_SubtypeName = map[Stat_Subtype]string{
		Stat_Subtype_CURRENT:   "CURRENT",
		Stat_Subtype_MAX:       "MAX",
		Stat_Subtype_PROCESSED: "PROCESSED",
		Stat_Subtype_CNT:       "CNT",
	}

	Stat_SubtypeValue = map[string]Stat_Subtype{
		"CURRENT":   Stat_Subtype_CURRENT,
		"MAX":       Stat_Subtype_MAX,
		"PROCESSED": Stat_Subtype_PROCESSED,
		"CNT":       Stat_Subtype_CNT,
	}
)

func (e Stat_Subtype) String() string {
	return Stat_SubtypeName[e]
}

func GetStat_Subtype(s string) Stat_Subtype {
	return Stat_SubtypeValue[s]
}

type Story_Type int32

const (
	Story_Type_SCRIPT Story_Type = iota
	Story_Type_ORACLE
	Story_Type_CNT
)

var (
	Story_TypeName = map[Story_Type]string{
		Story_Type_SCRIPT: "SCRIPT",
		Story_Type_ORACLE: "ORACLE",
		Story_Type_CNT:    "CNT",
	}

	Story_TypeValue = map[string]Story_Type{
		"SCRIPT": Story_Type_SCRIPT,
		"ORACLE": Story_Type_ORACLE,
		"CNT":    Story_Type_CNT,
	}
)

func (e Story_Type) String() string {
	return Story_TypeName[e]
}

func GetStory_Type(s string) Story_Type {
	return Story_TypeValue[s]
}

type Streamer int32

const (
	Streamer_RANDOM_CHARACTER Streamer = iota
	Streamer_BASE_DESTROYER
	Streamer_LAST_ENEMY_KILLER
	Streamer_BOSS_KILLER
	Streamer_KILLER
	Streamer_TUTORIAL_0
	Streamer_TUTORIAL_1
	Streamer_CNT
)

var (
	StreamerName = map[Streamer]string{
		Streamer_RANDOM_CHARACTER:  "RANDOM_CHARACTER",
		Streamer_BASE_DESTROYER:    "BASE_DESTROYER",
		Streamer_LAST_ENEMY_KILLER: "LAST_ENEMY_KILLER",
		Streamer_BOSS_KILLER:       "BOSS_KILLER",
		Streamer_KILLER:            "KILLER",
		Streamer_TUTORIAL_0:        "TUTORIAL_0",
		Streamer_TUTORIAL_1:        "TUTORIAL_1",
		Streamer_CNT:               "CNT",
	}

	StreamerValue = map[string]Streamer{
		"RANDOM_CHARACTER":  Streamer_RANDOM_CHARACTER,
		"BASE_DESTROYER":    Streamer_BASE_DESTROYER,
		"LAST_ENEMY_KILLER": Streamer_LAST_ENEMY_KILLER,
		"BOSS_KILLER":       Streamer_BOSS_KILLER,
		"KILLER":            Streamer_KILLER,
		"TUTORIAL_0":        Streamer_TUTORIAL_0,
		"TUTORIAL_1":        Streamer_TUTORIAL_1,
		"CNT":               Streamer_CNT,
	}
)

func (e Streamer) String() string {
	return StreamerName[e]
}

func GetStreamer(s string) Streamer {
	return StreamerValue[s]
}

type Target int32

const (
	Target_SELF Target = iota
	Target_ALLY
	Target_ENEMY
	Target_NOT_SELF_ALLY
	Target_TARGET
	Target_CNT
)

var (
	TargetName = map[Target]string{
		Target_SELF:          "SELF",
		Target_ALLY:          "ALLY",
		Target_ENEMY:         "ENEMY",
		Target_NOT_SELF_ALLY: "NOT_SELF_ALLY",
		Target_TARGET:        "TARGET",
		Target_CNT:           "CNT",
	}

	TargetValue = map[string]Target{
		"SELF":          Target_SELF,
		"ALLY":          Target_ALLY,
		"ENEMY":         Target_ENEMY,
		"NOT_SELF_ALLY": Target_NOT_SELF_ALLY,
		"TARGET":        Target_TARGET,
		"CNT":           Target_CNT,
	}
)

func (e Target) String() string {
	return TargetName[e]
}

func GetTarget(s string) Target {
	return TargetValue[s]
}

type Target_Condition int32

const (
	Target_Condition_STUN Target_Condition = iota
	Target_Condition_FROZEN
	Target_Condition_STATUSEFFECT_3_OVER
	Target_Condition_CNT
)

var (
	Target_ConditionName = map[Target_Condition]string{
		Target_Condition_STUN:                "STUN",
		Target_Condition_FROZEN:              "FROZEN",
		Target_Condition_STATUSEFFECT_3_OVER: "STATUSEFFECT_3_OVER",
		Target_Condition_CNT:                 "CNT",
	}

	Target_ConditionValue = map[string]Target_Condition{
		"STUN":                Target_Condition_STUN,
		"FROZEN":              Target_Condition_FROZEN,
		"STATUSEFFECT_3_OVER": Target_Condition_STATUSEFFECT_3_OVER,
		"CNT":                 Target_Condition_CNT,
	}
)

func (e Target_Condition) String() string {
	return Target_ConditionName[e]
}

func GetTarget_Condition(s string) Target_Condition {
	return Target_ConditionValue[s]
}

type Tile_Type int32

const (
	Tile_Type_SINGLE Tile_Type = iota
	Tile_Type_CITY
	Tile_Type_RUIN
	Tile_Type_BASE
	Tile_Type_CNT
)

var (
	Tile_TypeName = map[Tile_Type]string{
		Tile_Type_SINGLE: "SINGLE",
		Tile_Type_CITY:   "CITY",
		Tile_Type_RUIN:   "RUIN",
		Tile_Type_BASE:   "BASE",
		Tile_Type_CNT:    "CNT",
	}

	Tile_TypeValue = map[string]Tile_Type{
		"SINGLE": Tile_Type_SINGLE,
		"CITY":   Tile_Type_CITY,
		"RUIN":   Tile_Type_RUIN,
		"BASE":   Tile_Type_BASE,
		"CNT":    Tile_Type_CNT,
	}
)

func (e Tile_Type) String() string {
	return Tile_TypeName[e]
}

func GetTile_Type(s string) Tile_Type {
	return Tile_TypeValue[s]
}

type Toast_Message_Type int32

const (
	Toast_Message_Type_MY_TILE_ATTACKED Toast_Message_Type = iota
	Toast_Message_Type_MY_TILE_OCCUPIED
	Toast_Message_Type_OTHER_BIGTILE_OCCUPIED
	Toast_Message_Type_CNT
)

var (
	Toast_Message_TypeName = map[Toast_Message_Type]string{
		Toast_Message_Type_MY_TILE_ATTACKED:       "MY_TILE_ATTACKED",
		Toast_Message_Type_MY_TILE_OCCUPIED:       "MY_TILE_OCCUPIED",
		Toast_Message_Type_OTHER_BIGTILE_OCCUPIED: "OTHER_BIGTILE_OCCUPIED",
		Toast_Message_Type_CNT:                    "CNT",
	}

	Toast_Message_TypeValue = map[string]Toast_Message_Type{
		"MY_TILE_ATTACKED":       Toast_Message_Type_MY_TILE_ATTACKED,
		"MY_TILE_OCCUPIED":       Toast_Message_Type_MY_TILE_OCCUPIED,
		"OTHER_BIGTILE_OCCUPIED": Toast_Message_Type_OTHER_BIGTILE_OCCUPIED,
		"CNT":                    Toast_Message_Type_CNT,
	}
)

func (e Toast_Message_Type) String() string {
	return Toast_Message_TypeName[e]
}

func GetToast_Message_Type(s string) Toast_Message_Type {
	return Toast_Message_TypeValue[s]
}

type Tutorial_Event int32

const (
	Tutorial_Event_TUTORIAL_INTRO Tutorial_Event = iota
	Tutorial_Event_CNT
)

var (
	Tutorial_EventName = map[Tutorial_Event]string{
		Tutorial_Event_TUTORIAL_INTRO: "TUTORIAL_INTRO",
		Tutorial_Event_CNT:            "CNT",
	}

	Tutorial_EventValue = map[string]Tutorial_Event{
		"TUTORIAL_INTRO": Tutorial_Event_TUTORIAL_INTRO,
		"CNT":            Tutorial_Event_CNT,
	}
)

func (e Tutorial_Event) String() string {
	return Tutorial_EventName[e]
}

func GetTutorial_Event(s string) Tutorial_Event {
	return Tutorial_EventValue[s]
}

type Tutorial_Resource_Type int32

const (
	Tutorial_Resource_Type_TIMELINE Tutorial_Resource_Type = iota
	Tutorial_Resource_Type_NAVI
	Tutorial_Resource_Type_CNT
)

var (
	Tutorial_Resource_TypeName = map[Tutorial_Resource_Type]string{
		Tutorial_Resource_Type_TIMELINE: "TIMELINE",
		Tutorial_Resource_Type_NAVI:     "NAVI",
		Tutorial_Resource_Type_CNT:      "CNT",
	}

	Tutorial_Resource_TypeValue = map[string]Tutorial_Resource_Type{
		"TIMELINE": Tutorial_Resource_Type_TIMELINE,
		"NAVI":     Tutorial_Resource_Type_NAVI,
		"CNT":      Tutorial_Resource_Type_CNT,
	}
)

func (e Tutorial_Resource_Type) String() string {
	return Tutorial_Resource_TypeName[e]
}

func GetTutorial_Resource_Type(s string) Tutorial_Resource_Type {
	return Tutorial_Resource_TypeValue[s]
}

type Tutorial_Trigger int32

const (
	Tutorial_Trigger_GAME_START Tutorial_Trigger = iota
	Tutorial_Trigger_CITY_MEET
	Tutorial_Trigger_FIRST_TILE
	Tutorial_Trigger_SECOND_TILE
	Tutorial_Trigger_SECOND_CAPTURE
	Tutorial_Trigger_FIRST_ENCOUNTER
	Tutorial_Trigger_PREBATTLE
	Tutorial_Trigger_STRIKER
	Tutorial_Trigger_FINAL
	Tutorial_Trigger_SECOND_BATTLE_WIN
	Tutorial_Trigger_CNT
)

var (
	Tutorial_TriggerName = map[Tutorial_Trigger]string{
		Tutorial_Trigger_GAME_START:        "GAME_START",
		Tutorial_Trigger_CITY_MEET:         "CITY_MEET",
		Tutorial_Trigger_FIRST_TILE:        "FIRST_TILE",
		Tutorial_Trigger_SECOND_TILE:       "SECOND_TILE",
		Tutorial_Trigger_SECOND_CAPTURE:    "SECOND_CAPTURE",
		Tutorial_Trigger_FIRST_ENCOUNTER:   "FIRST_ENCOUNTER",
		Tutorial_Trigger_PREBATTLE:         "PREBATTLE",
		Tutorial_Trigger_STRIKER:           "STRIKER",
		Tutorial_Trigger_FINAL:             "FINAL",
		Tutorial_Trigger_SECOND_BATTLE_WIN: "SECOND_BATTLE_WIN",
		Tutorial_Trigger_CNT:               "CNT",
	}

	Tutorial_TriggerValue = map[string]Tutorial_Trigger{
		"GAME_START":        Tutorial_Trigger_GAME_START,
		"CITY_MEET":         Tutorial_Trigger_CITY_MEET,
		"FIRST_TILE":        Tutorial_Trigger_FIRST_TILE,
		"SECOND_TILE":       Tutorial_Trigger_SECOND_TILE,
		"SECOND_CAPTURE":    Tutorial_Trigger_SECOND_CAPTURE,
		"FIRST_ENCOUNTER":   Tutorial_Trigger_FIRST_ENCOUNTER,
		"PREBATTLE":         Tutorial_Trigger_PREBATTLE,
		"STRIKER":           Tutorial_Trigger_STRIKER,
		"FINAL":             Tutorial_Trigger_FINAL,
		"SECOND_BATTLE_WIN": Tutorial_Trigger_SECOND_BATTLE_WIN,
		"CNT":               Tutorial_Trigger_CNT,
	}
)

func (e Tutorial_Trigger) String() string {
	return Tutorial_TriggerName[e]
}

func GetTutorial_Trigger(s string) Tutorial_Trigger {
	return Tutorial_TriggerValue[s]
}

type Win_Condition int32

const (
	Win_Condition_STANDARD Win_Condition = iota
	Win_Condition_EXTERMINATE
	Win_Condition_CNT
)

var (
	Win_ConditionName = map[Win_Condition]string{
		Win_Condition_STANDARD:    "STANDARD",
		Win_Condition_EXTERMINATE: "EXTERMINATE",
		Win_Condition_CNT:         "CNT",
	}

	Win_ConditionValue = map[string]Win_Condition{
		"STANDARD":    Win_Condition_STANDARD,
		"EXTERMINATE": Win_Condition_EXTERMINATE,
		"CNT":         Win_Condition_CNT,
	}
)

func (e Win_Condition) String() string {
	return Win_ConditionName[e]
}

func GetWin_Condition(s string) Win_Condition {
	return Win_ConditionValue[s]
}
