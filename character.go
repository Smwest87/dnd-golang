package main

import (
	Dice "dnd_dice"
	"fmt"
	"os"
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
	var level = 1
	var strength = Dice.RollStat()
	var dexterity = Dice.RollStat()
	var constitution = Dice.RollStat()
	var hitPointMaximum = calculateMaximumHealth(os.Args[2]) + constitution/2
	var wisdom = Dice.RollStat()
	var intelligence = Dice.RollStat()
	var charisma = Dice.RollStat()
	var newCharacter = Character{os.Args[1], os.Args[2], level, hitPointMaximum, strength, dexterity, constitution, wisdom, intelligence, charisma}
	return newCharacter
}

func main() {
	var NewCharacter Character = generateCharacter(os.Args[1], os.Args[2])
	fmt.Println(NewCharacter)

}
