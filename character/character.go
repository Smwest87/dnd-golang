package character

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"os"

	config "github.com/smwest87/dnd-golang/configuration"
	Dice "github.com/smwest87/dnd_dice"

	pq "github.com/lib/pq"
)

//Character contains character name class and level
type Character struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Class           string   `json:"class"`
	Level           int      `json:"level"`
	HitPointMaximum int      `json:"hit_point_maximum"`
	Strength        int      `json:"strength"`
	Dexterity       int      `json:"dexterity"`
	Constitution    int      `json:"constitution"`
	Wisdom          int      `json:"wisdom"`
	Intelligence    int      `json:"intelligence"`
	Charisma        int      `json:"charisma"`
	Initiative      int      `json:"intiative"`
	Modifiers       modArray `json:"modifiers"`
	// TODO AC
	// TODO FEATS
	// TODO Class Properties
}

type modArray [6]int

var password = os.Getenv("DB_PASSWORD")

func calculateMaximumHealth(class string, modifier int) (int, error) {

	roller := Dice.NewRoller()

	switch class {
	case "Barbarian":
		value, err := roller.RollDie(12, roller.RNG)
		if err != nil {
			return -1, err
		}

		if modifier <= 0 {
			return value + 1, nil
		}

		return value + modifier, nil
	case "Bard":
		value, err := roller.RollDie(8, roller.RNG)
		if err != nil {
			return -1, err
		}

		if modifier <= 0 {
			return value + 1, nil
		}

		return value + modifier, nil
	case "Cleric":
		value, err := roller.RollDie(8, roller.RNG)
		if err != nil {
			return -1, err
		}

		if modifier <= 0 {
			return value + 1, nil
		}

		return value + modifier, nil
	case "Druid":
		value, err := roller.RollDie(8, roller.RNG)
		if err != nil {
			return -1, err
		}

		if modifier <= 0 {
			return value + 1, nil
		}

		return value + modifier, nil
	case "Fighter":
		value, err := roller.RollDie(10, roller.RNG)
		if err != nil {
			return -1, err
		}

		if modifier <= 0 {
			return value + 1, nil
		}

		return value + modifier, nil
	case "Monk":
		value, err := roller.RollDie(8, roller.RNG)
		if err != nil {
			return -1, err
		}

		if modifier <= 0 {
			return value + 1, nil
		}

		return value + modifier, nil
	case "Paladin":
		value, err := roller.RollDie(10, roller.RNG)
		if err != nil {
			return -1, err
		}

		if modifier <= 0 {
			return value + 1, nil
		}

		return value + modifier, nil
	case "Ranger":
		value, err := roller.RollDie(10, roller.RNG)
		if err != nil {
			return -1, err
		}

		if modifier <= 0 {
			return value + 1, nil
		}

		return value + modifier, nil
	case "Rogue":
		value, err := roller.RollDie(8, roller.RNG)
		if err != nil {
			return -1, err
		}

		if modifier <= 0 {
			return value + 1, nil
		}

		return value + modifier, nil
	case "Sorcerer":
		value, err := roller.RollDie(6, roller.RNG)
		if err != nil {
			return -1, err
		}

		if modifier <= 0 {
			return value + 1, nil
		}

		return value + modifier, nil
	case "Warlock":
		value, err := roller.RollDie(8, roller.RNG)
		if err != nil {
			return -1, err
		}

		if modifier <= 0 {
			return value + 1, nil
		}

		return value + modifier, nil
	case "Wizard":
		value, err := roller.RollDie(6, roller.RNG)
		if err != nil {
			return -1, err
		}

		if modifier <= 0 {
			return value + 1, nil
		}

		return value + modifier, nil

	}
	err := errors.New("Class name did not match existing classes")
	return -1, err
}

func calculateAbilityModifiers(stat int) int {
	var value = math.Floor((float64(stat) - 10) / 2)
	return int(value)
}

//GenerateCharacter -- generate new character to insert into PSQL db
func GenerateCharacter(name string, class string) (*Character, error) {
	roller := Dice.NewRoller()
	character := Character{}
	character.Name = os.Args[1]
	character.Class = os.Args[2]
	character.Level = 1
	strength, err := roller.RollStat()
	if err != nil {
		return nil, err
	}
	character.Strength = strength
	dexterity, err := roller.RollStat()
	if err != nil {
		return nil, err
	}
	character.Dexterity = dexterity
	constitution, err := roller.RollStat()
	if err != nil {
		return nil, err
	}
	character.Constitution = constitution
	hitPointMaximum, err := calculateMaximumHealth(character.Class, calculateAbilityModifiers(constitution))
	if err != nil {
		return nil, err
	}
	character.HitPointMaximum = hitPointMaximum
	wisdom, err := roller.RollStat()
	if err != nil {
		return nil, err
	}
	character.Wisdom = wisdom
	intelligence, err := roller.RollStat()
	if err != nil {
		return nil, err
	}
	character.Intelligence = intelligence
	charisma, err := roller.RollStat()
	if err != nil {
		return nil, err
	}
	character.Charisma = charisma
	character.Initiative = int(math.Floor((float64(character.Dexterity) - 10) / 2))
	//TODO Armor Class and FEATS
	character.Modifiers = assignModifiers(character)
	return &character, err
}

//InsertCharacter -- insert character into PSQL db
func InsertCharacter(character Character) (sql.Result, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, password, config.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	var charInsert = "INSERT INTO dnd.dnd_characters (name,class,level,hitpointmaximum,strength,dexterity,constitution,wisdom,intelligence,charisma, initiative, modifiers) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12);"

	var result, insertErr = db.Exec(charInsert, character.Name, character.Class, character.Level, character.HitPointMaximum, character.Strength, character.Dexterity, character.Constitution, character.Wisdom, character.Intelligence, character.Charisma, character.Initiative, pq.Array(character.Modifiers))
	if insertErr != nil {
		return nil, insertErr
	}

	return result, err
}

func assignModifiers(character Character) modArray {
	character.Modifiers[0] = calculateAbilityModifiers(character.Strength)
	character.Modifiers[1] = calculateAbilityModifiers(character.Dexterity)
	character.Modifiers[2] = calculateAbilityModifiers(character.Constitution)
	character.Modifiers[3] = calculateAbilityModifiers(character.Wisdom)
	character.Modifiers[4] = calculateAbilityModifiers(character.Intelligence)
	character.Modifiers[5] = calculateAbilityModifiers(character.Charisma)
	return character.Modifiers
}
