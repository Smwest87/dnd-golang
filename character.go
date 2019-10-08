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

func calculateMaximumHealth(class string) int {
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
	return value
}

func generateCharacter(name string, class string) Character {
	var charName = os.Args[1]
	var charClass = os.Args[2]
	var level = 1
	var strength = Dice.RollStat()
	var dexterity = Dice.RollStat()
	var constitution = Dice.RollStat()
	var hitPointMaximum = calculateMaximumHealth(charClass) + constitution/2
	var wisdom = Dice.RollStat()
	var intelligence = Dice.RollStat()
	var charisma = Dice.RollStat()
	var newCharacter = Character{charName, charClass, level, hitPointMaximum, strength, dexterity, constitution, wisdom, intelligence, charisma}
	return newCharacter
}

func main() {
	var NewCharacter Character = generateCharacter(os.Args[1], os.Args[2])
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	//connStr := "postgres://postgres:@localhost/8000/postgres"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	var charInsert = "INSERT INTO public.dnd_characters (name,class,level,hitpointmaximum,strength,dexterity,constitution,wisdom,intelligence,charisma) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10);"

	var _, insertErr = db.Exec(charInsert, NewCharacter.name, NewCharacter.class, NewCharacter.level, NewCharacter.hitPointMaximum, NewCharacter.strength, NewCharacter.dexterity, NewCharacter.constitution, NewCharacter.wisdom, NewCharacter.intelligence, NewCharacter.charisma)
	if insertErr != nil {
		panic(err)
	}

}
