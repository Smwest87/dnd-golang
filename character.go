package main

import (
	"database/sql"
	Dice "dnd_dice"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Character struct {
	//contains character name class and level
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
	// TODO AC
}

const (
	host   = "localhost"
	port   = 8000
	user   = "postgres"
	dbname = "postgres"
)

func main() {
	var NewCharacter Character = generateCharacter(os.Args[1], os.Args[2])
	insertCharacter(NewCharacter)
}

func calculateMaximumHealth(class string, modifier int) int {
	value := 0
	switch class {
	case "Barbarian":
		value = Dice.RollDie(12)
	case "Bard":
		value = Dice.RollDie(8)
	case "Cleric":
		value = Dice.RollDie(8)
	case "Druid":
		value = Dice.RollDie(8)
	case "Fighter":
		value = Dice.RollDie(10)
	case "Monk":
		value = Dice.RollDie(8)
	case "Paladin":
		value = Dice.RollDie(10)
	case "Ranger":
		value = Dice.RollDie(10)
	case "Rogue":
		value = Dice.RollDie(8)
	case "Sorcerer":
		value = Dice.RollDie(6)
	case "Warlock":
		value = Dice.RollDie(8)
	case "Wizard":
		value = Dice.RollDie(6)
	}

	if modifier <= 0 {
		return value + 1
	}
	return value + modifier

}

func calculateAbilityModifiers(stat int) int {
	value := 0
	switch stat {
	case 1:
		value = -5
	case 2:
		value = -4
	case 3:
		value = -4
	case 4:
		value = -3
	case 5:
		value = -3
	case 6:
		value = -2
	case 7:
		value = -2
	case 8:
		value = -1
	case 9:
		value = -1
	case 10:
		value = 0
	case 11:
		value = 0
	case 12:
		value = 1
	case 13:
		value = 1
	case 14:
		value = 2
	case 15:
		value = 2
	case 16:
		value = 3
	case 17:
		value = 3
	case 18:
		value = 4
	case 19:
		value = 4
	case 20:
		value = 5
	case 21:
		value = 5
	case 22:
		value = 6
	case 23:
		value = 6
	case 24:
		value = 7
	case 25:
		value = 7
	case 26:
		value = 8
	case 27:
		value = 8
	case 28:
		value = 9
	case 29:
		value = 9
	case 30:
		value = 10

	}
	return value
}

func generateCharacter(name string, class string) Character {
	var charName = os.Args[1]
	var charClass = os.Args[2]
	var level = 1
	var strength = Dice.RollStat()
	var dexterity = Dice.RollStat()
	var constitution = Dice.RollStat()
	var hitPointMaximum = calculateMaximumHealth(charClass, calculateAbilityModifiers(constitution))
	var wisdom = Dice.RollStat()
	var intelligence = Dice.RollStat()
	var charisma = Dice.RollStat()
	var newCharacter = Character{charName, charClass, level, hitPointMaximum, strength, dexterity, constitution, wisdom, intelligence, charisma}
	return newCharacter
}

func insertCharacter(character Character) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	var charInsert = "INSERT INTO public.dnd_characters (name,class,level,hitpointmaximum,strength,dexterity,constitution,wisdom,intelligence,charisma) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);"

	var _, insertErr = db.Exec(charInsert, character.name, character.class, character.level, character.hitPointMaximum, character.strength, character.dexterity, character.constitution, character.wisdom, character.intelligence, character.charisma)
	if insertErr != nil {
		panic(err)
	}
}
