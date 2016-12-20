package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type item struct {
	Name                 string
	Cost, Damage, Armour int
}

var weapons = []item{
	{"Dagger", 8, 4, 0},
	{"Shortsword", 10, 5, 0},
	{"Warhammer", 25, 6, 0},
	{"Longsword", 40, 7, 0},
	{"Greataxe", 74, 8, 0},
}
var armours = []item{
	{"None", 0, 0, 0},
	{"Leather", 13, 0, 1},
	{"Chainmail", 31, 0, 2},
	{"Splintmail", 53, 0, 3},
	{"Bandedmail", 75, 0, 4},
	{"Platemail", 102, 0, 5},
}
var rings = []item{
	{"None", 0, 0, 0},
	{"Damage +1", 25, 1, 0},
	{"Damage +2", 50, 2, 0},
	{"Damage +3", 100, 3, 0},
	{"Defense +1", 20, 0, 1},
	{"Defense +2", 40, 0, 2},
	{"Defense +3", 80, 0, 3},
}

type character struct {
	Hitpoints, Damage, Armor int
}

func parseInput(s string) (character, error) {
	var boss character

	cnt, err := fmt.Sscanf(s, "Hit Points: %d\nDamage: %d\nArmor: %d\n",
		&boss.Hitpoints, &boss.Damage, &boss.Armor)
	if err != nil {
		return character{}, err
	}
	if cnt != 3 {
		err := fmt.Errorf("invalid input %q", s)
		return character{}, err
	}

	return boss, nil
}

func fight(p, b character) bool {
	playerHitpoints := p.Hitpoints
	playerDamage := p.Damage - b.Armor
	if playerDamage < 1 {
		playerDamage = 1
	}
	bossHitpoints := b.Hitpoints
	bossDamage := b.Damage - p.Armor
	if bossDamage < 1 {
		bossDamage = 1
	}

	for {
		bossHitpoints -= playerDamage
		if bossHitpoints < 0 {
			bossHitpoints = 0
		}
		//fmt.Printf("The player deals %d-%d = %d damage; the boss goes down to %d hit points.\n",
		//	p.Damage, b.Armor, playerDamage, bossHitpoints)
		if bossHitpoints == 0 {
			break
		}
		playerHitpoints -= bossDamage
		if playerHitpoints < 0 {
			playerHitpoints = 0
		}
		//fmt.Printf("The boss deals %d-%d = %d damage; the player goes down to %d hit points.\n",
		//	b.Damage, p.Armor, bossDamage, playerHitpoints)
		if playerHitpoints == 0 {
			break
		}
	}

	return playerHitpoints > 0
}

func process(s string, h int) (int, error) {
	boss, err := parseInput(s)
	if err != nil {
		return 0, err
	}

	minCost := -1

	for _, weapon := range weapons {
		for _, armour := range armours {
			for _, leftRing := range rings {
				for _, rightRing := range rings {
					if leftRing == rightRing {
						if leftRing.Cost != 0 {
							continue
						}
					} else if rightRing.Cost == 0 {
						continue
					}
					cost := weapon.Cost + armour.Cost + leftRing.Cost + rightRing.Cost
					player := character{h, weapon.Damage + leftRing.Damage + rightRing.Damage, armour.Armour + leftRing.Armour + rightRing.Armour}

					w := fight(player, boss)
					if w {
						if minCost > cost || minCost == -1 {
							minCost = cost
						}
					}
					//fmt.Println(player, boss, cost, weapon.Name, armour.Name, leftRing.Name, rightRing.Name, w)
				}
			}
		}
	}

	return minCost, nil
}

func run() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%s filename\n", os.Args[0])
		return 1
	}

	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	s := string(b)

	minCost, err := process(s, 100)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	fmt.Printf("min_cost: %d\n", minCost)
	return 0
}

func main() {
	os.Exit(run())
}
