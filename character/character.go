package character

import (
	"database/sql"
	"errors"
	"fmt"
	"math"
	"os"

	Dice "github.com/smwest87/dnd_dice"

	pq "github.com/lib/pq"
)

//Character contains character name class and level
type Character struct {
	name            string
	class           string
	level           int
	hitPointMaximum int
	strength        int
	dexterity       int
	constitution    int
	wisdom          int
	intelligence    int
	charisma        int
	initiative      int
	modifiers       modArray
	// TODO AC
	// TODO FEATS
	// TODO Class Properties
}

type modArray [6]int

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "postgres"
)

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
	character.name = os.Args[1]
	character.class = os.Args[2]
	character.level = 1
	strength, err := roller.RollStat()
	if err != nil {
		return nil, err
	}
	character.strength = strength
	dexterity, err := roller.RollStat()
	if err != nil {
		return nil, err
	}
	character.dexterity = dexterity
	constitution, err := roller.RollStat()
	if err != nil {
		return nil, err
	}
	character.constitution = constitution
	hitPointMaximum, err := calculateMaximumHealth(character.class, calculateAbilityModifiers(constitution))
	if err != nil {
		return nil, err
	}
	character.hitPointMaximum = hitPointMaximum
	wisdom, err := roller.RollStat()
	if err != nil {
		return nil, err
	}
	character.wisdom = wisdom
	intelligence, err := roller.RollStat()
	if err != nil {
		return nil, err
	}
	character.intelligence = intelligence
	charisma, err := roller.RollStat()
	if err != nil {
		return nil, err
	}
	character.charisma = charisma
	character.initiative = int(math.Floor((float64(character.dexterity) - 10) / 2))
	//TODO Armor Class and FEATS
	character.modifiers = assignModifiers(character)
	return &character, err
}

//InsertCharacter -- insert character into PSQL db
func InsertCharacter(character Character) (sql.Result, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	var charInsert = "INSERT INTO dnd.dnd_characters (name,class,level,hitpointmaximum,strength,dexterity,constitution,wisdom,intelligence,charisma, initiative, modifiers) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12);"

	var result, insertErr = db.Exec(charInsert, character.name, character.class, character.level, character.hitPointMaximum, character.strength, character.dexterity, character.constitution, character.wisdom, character.intelligence, character.charisma, character.initiative, pq.Array(character.modifiers))
	if insertErr != nil {
		return nil, insertErr
	}

	return result, err
}

func assignModifiers(character Character) modArray {
	character.modifiers[0] = calculateAbilityModifiers(character.strength)
	character.modifiers[1] = calculateAbilityModifiers(character.dexterity)
	character.modifiers[2] = calculateAbilityModifiers(character.constitution)
	character.modifiers[3] = calculateAbilityModifiers(character.wisdom)
	character.modifiers[4] = calculateAbilityModifiers(character.intelligence)
	character.modifiers[5] = calculateAbilityModifiers(character.charisma)
	return character.modifiers
}
