package main

import (
	"database/sql"
	Dice "dnd_dice"
	"fmt"
	"log"
	"math"
	"os"

	pq "github.com/lib/pq"
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
	initiative      int
	modifiers       modArray
	// TODO AC
}

type modArray [6]int

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
	var value = math.Floor((float64(stat) - 10) / 2)
	return int(value)
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
	var initiative = int(math.Floor((float64(dexterity) - 10) / 2))
	var mods modArray
	var newCharacter = Character{charName, charClass, level, hitPointMaximum, strength, dexterity, constitution, wisdom, intelligence, charisma, initiative, mods}
	newCharacter.modifiers = assignModifiers(newCharacter)
	return newCharacter
}

func insertCharacter(character Character) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	var charInsert = "INSERT INTO dnd.dnd_characters (name,class,level,hitpointmaximum,strength,dexterity,constitution,wisdom,intelligence,charisma, initiative, modifiers) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12);"

	var _, insertErr = db.Exec(charInsert, character.name, character.class, character.level, character.hitPointMaximum, character.strength, character.dexterity, character.constitution, character.wisdom, character.intelligence, character.charisma, character.initiative, pq.Array(character.modifiers))
	if insertErr != nil {
		panic(err)
	}
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
