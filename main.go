package main

import (
	"fmt"
	"os"
)

// Regular Coffee
const (
	WaterNeededPerCup  = 200
	MilkNeededPerCup   = 50
	CoffeeNeededPerCup = 15
)

// Espresso
const (
	WaterPerEspresso  = 250
	MilkPerEspresso   = 0
	CoffeePerEspresso = 16
	CostPerEspresso   = 4
)

// Latte
const (
	WaterPerLatte  = 350
	MilkPerLatte   = 75
	CoffeePerLatte = 20
	CostPerLatte   = 7
)

// Cappuccino
const (
	WaterPerCappuccino  = 200
	MilkPerCappuccino   = 100
	CoffeePerCappuccino = 12
	CostPerCappuccino   = 6
)

func MinIntSlice(array []int) (minNum int) {
	if len(array) > 0 {
		minNum = array[0]
	}
	for i, e := range array {
		if i == 0 || e < minNum {
			minNum = e
		}
	}
	return
}

func processCalculation(getWater, getMilk, getCoffee, getCups int) (int, int) {
	calcWaterRatio := getWater / WaterNeededPerCup
	calcMilkRatio := getMilk / MilkNeededPerCup
	calcCoffeeRatio := getCoffee / CoffeeNeededPerCup

	ratioSlice := []int{calcWaterRatio, calcMilkRatio, calcCoffeeRatio}

	minCups := MinIntSlice(ratioSlice)
	maxCups := 0

	if minCups > getCups {
		maxCups = minCups - getCups
	}

	return minCups, maxCups
}


func processBuyFillTake(
	storedMoney,
	storedWater,
	storedMilk,
	storedCoffee,
	storedCups int,
) {

	var remainingMoney, remainingWater, remainingMilk, remainingCoffee, remainingCups int

	remainingMoney += storedMoney
	remainingWater += storedWater
	remainingMilk += storedMilk
	remainingCoffee += storedCoffee
	remainingCups += storedCups

	for {
		userAction := getUserAction() // buy | fill | take

		switch userAction {
		case "buy":
			remainingMoney, remainingWater, remainingMilk, remainingCoffee, remainingCups = processBuy(
				remainingMoney, remainingWater, remainingMilk, remainingCoffee,
				remainingCups,
			)
			continue

		case "fill":
			remainingWater, remainingMilk, remainingCoffee, remainingCups = processFill(
				remainingWater,
				remainingMilk,
				remainingCoffee,
				remainingCups,
			)
			continue

		case "take":
			remainingMoney = processTake(
				storedMoney,
			)
			continue

		case "remaining":
			processRemaining(
				remainingMoney, remainingWater, remainingMilk, remainingCoffee,
				remainingCups,
			)
			continue

		case "exit":
			os.Exit(0)

		default:

		}
	}

}

func processBuy(storedMoney, storedWater, storedMilk, storedCoffee, storedCups int) (
	int, int, int, int, int,
) {
	var wantToBuy int

	var remainingMoney, remainingWater, remainingMilk, remainingCoffee, remainingCups int

	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino: ")
	_, _ = fmt.Scan(&wantToBuy)

	switch wantToBuy {
	case 1: // Espresso

		if storedWater >= WaterPerEspresso && storedCoffee >= CoffeePerEspresso {
			remainingMoney = storedMoney + CostPerEspresso
			remainingWater = storedWater - WaterPerEspresso
			remainingMilk = storedMilk - MilkPerEspresso
			remainingCoffee = storedCoffee - CoffeePerEspresso
			remainingCups = storedCups - 1

			fmt.Println("I have enough resources, making you a coffee!")

		} else {
			remainingMoney = storedMoney
			remainingWater = storedWater
			remainingMilk = storedMilk
			remainingCoffee = storedCoffee
			remainingCups = storedCups
			if storedWater < WaterPerEspresso {
				fmt.Println("Sorry, not enough water!")
			} else if storedCoffee < CoffeePerEspresso {
				fmt.Println("Sorry, not enough coffee!")
			}

		}

	case 2: // Latte

		if storedWater >= WaterPerLatte && storedMilk >= MilkPerLatte && storedCoffee >= CoffeePerLatte {
			remainingMoney = storedMoney + CostPerLatte
			remainingWater = storedWater - WaterPerLatte
			remainingMilk = storedMilk - MilkPerLatte
			remainingCoffee = storedCoffee - CoffeePerLatte
			remainingCups = storedCups - 1
			fmt.Println("I have enough resources, making you a coffee!")

		} else {
			remainingMoney = storedMoney
			remainingWater = storedWater
			remainingMilk = storedMilk
			remainingCoffee = storedCoffee
			remainingCups = storedCups
			if storedWater < WaterPerLatte {
				fmt.Println("Sorry, not enough water!")
			} else if storedMilk < MilkPerLatte {
				fmt.Println("Sorry, not enough milk!")
			} else if storedCoffee < CoffeePerLatte {
				fmt.Println("Sorry, not enough coffee!")
			}

		}

	case 3: // Cappuccino

		if storedWater >= WaterPerCappuccino && storedMilk >= MilkPerCappuccino && storedCoffee >= CoffeePerCappuccino {
			remainingMoney = storedMoney + CostPerCappuccino
			remainingWater = storedWater - WaterPerCappuccino
			remainingMilk = storedMilk - MilkPerCappuccino
			remainingCoffee = storedCoffee - CoffeePerCappuccino
			remainingCups = storedCups - 1
			fmt.Println("I have enough resources, making you a coffee!")

		} else {
			remainingMoney = storedMoney
			remainingWater = storedWater
			remainingMilk = storedMilk
			remainingCoffee = storedCoffee
			remainingCups = storedCups
			if storedWater < WaterPerCappuccino {
				fmt.Println("Sorry, not enough water!")
			} else if storedMilk < MilkPerCappuccino {
				fmt.Println("Sorry, not enough milk!")
			} else if storedCoffee < CoffeePerCappuccino {
				fmt.Println("Sorry, not enough coffee!")
			}

		}

	default:
		remainingMoney = storedMoney
		remainingWater = storedWater
		remainingMilk = storedMilk
		remainingCoffee = storedCoffee
		remainingCups = storedCups

	}

	return remainingMoney, remainingWater, remainingMilk, remainingCoffee, remainingCups

}

func processFill(
	remainingWater,
	remainingMilk,
	remainingCoffee,
	remainingCups int,
) (
	int, int, int, int,
) {
	var addedWater, addedMilk, addedCoffee, addedCups int

	fmt.Println("Write how many ml of water you want to add:")
	_, _ = fmt.Scan(&addedWater)

	fmt.Println("Write how many ml of milk you want to add:")
	_, _ = fmt.Scan(&addedMilk)

	fmt.Println("Write how many grams of coffee beans you want to add:")
	_, _ = fmt.Scan(&addedCoffee)

	fmt.Println("Write how many disposable cups you want to add:")
	_, _ = fmt.Scan(&addedCups)

	remainingWater = remainingWater + addedWater
	remainingMilk = remainingMilk + addedMilk
	remainingCoffee = remainingCoffee + addedCoffee
	remainingCups = remainingCups + addedCups

	return remainingWater, remainingMilk, remainingCoffee, remainingCups
}

func processTake(storedMoney int) int {
	remainingMoney := storedMoney - storedMoney
	fmt.Printf("I gave you $%d\n", remainingMoney)
	return remainingMoney

}

func processRemaining(
	remainingMoney, remainingWater, remainingMilk, remainingCoffee,
	remainingCups int,
) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", remainingWater)
	fmt.Printf("%d ml of milk\n", remainingMilk)
	fmt.Printf("%d g of coffee beans\n", remainingCoffee)
	fmt.Printf("%d  disposable cups\n", remainingCups)
	fmt.Printf("$%d of money\n", remainingMoney)
}

func getUserAction() string {
	var userAction string

	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	_, _ = fmt.Scan(&userAction)

	return userAction
}

func main() {
	//	Declaring Variables
	storedMoney := 550
	storedWater := 400
	storedMilk := 540
	storedCoffee := 120
	storedCups := 9

	processBuyFillTake(
		storedMoney,
		storedWater,
		storedMilk,
		storedCoffee,
		storedCups,
	)
}
