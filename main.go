package main

import (
	"math/rand"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

var (
	planes = []string{
		"bluebell",
		"devilfire",
		"fleetfoot",
		"hammerhead",
		"jalopy",
		"lumberjack",
	}
	firstname = []string{
		"alex",
		"brad",
		"chuck",
		"dave",
		"edgar",
		"fred",
		"greg",
		"harold",
		"ivan",
		"jack",
		"keith",
		"larry",
	}
	lastname = []string{
		"abrams",
		"beam",
		"clarkson",
		"drather",
		"evans",
		"farr",
		"gray",
		"hardy",
		"ives",
		"jones",
		"kenton",
		"lord"}
	nickname = []string{
		"ace",
		"bogey",
		"cap",
		"duck",
		"eagle",
		"fuzzy",
		"goose",
		"hawk",
		"itchy",
		"joker",
		"kite",
		"lefty",
	}
)

type Person struct {
	firstname map[string]bool
	lastname  map[string]bool
	nickname  map[string]bool
}

type Storage struct {
	plane     map[string]bool
	pilot     Person
	navigator Person
}

type PersonSolution struct {
	firstname string
	lastname  string
	nickname  string
}

type Solution struct {
	plane     string
	pilot     PersonSolution
	navigator PersonSolution
}

func main() {
	solution := solve()

	spew.Dump(solution)
}

func solve() []Solution {
	var storage Storage
	var solution Solution
	storage.pilot.firstname = make(map[string]bool)
	storage.pilot.lastname = make(map[string]bool)
	storage.pilot.nickname = make(map[string]bool)
	storage.navigator.firstname = make(map[string]bool)
	storage.navigator.lastname = make(map[string]bool)
	storage.navigator.nickname = make(map[string]bool)
	storage.plane = make(map[string]bool)
	var arraySolution []Solution

	// Shuffle all the orders so they are unique searches
	shuffle(planes)
	shuffle(firstname)
	shuffle(lastname)
	shuffle(nickname)

	for _, p := range planes {
		for _, fname := range firstname {
			// jalopy, alex, false, false --
			// jalopy, bart, false, true
			if checkValidPilotFirstName(p, fname, storage.pilot.firstname[fname], storage.plane[p]) {
				storage.pilot.firstname[fname] = true
				storage.plane[p] = true
				solution.plane = p
				solution.pilot.firstname = fname
				for _, pilotlname := range lastname {
					storage.plane = make(map[string]bool) //erase where names can be in planes
					if checkValidPilotLastname(p, pilotlname, storage.pilot.lastname[pilotlname], storage.plane[p], solution) {
						storage.pilot.lastname[pilotlname] = true
						storage.plane[p] = true
						solution.pilot.lastname = pilotlname
						for _, pilotNname := range nickname {
							storage.plane = make(map[string]bool)
							if checkValidPilotNickName(p, pilotNname, storage.pilot.nickname[pilotNname], storage.plane[p], solution) {
								storage.pilot.nickname[pilotNname] = true
								storage.plane[p] = true
								solution.pilot.nickname = pilotNname

								break
							}
						}
						break
					}

				}

			} else if checkValidNavigatorFirstName(p, fname, storage.navigator.firstname[fname], storage.plane[p], storage.pilot.firstname[fname]) {
				// else move on to the next pilot or check for navigators
				storage.navigator.firstname[fname] = true
				storage.plane[p] = true
				solution.plane = p
				solution.navigator.firstname = fname

			}

		}
		arraySolution = append(arraySolution, solution)
	}
	return arraySolution
}

func shuffle(value []string) {
	rand.Shuffle(len(value), func(i, j int) {
		value[i], value[j] = value[j], value[i]
	})
}

func checkValidPilotNickName(plane string, pilotNname string, pilotBeenUsed bool, planeTaken bool, solution Solution) bool {
	if strings.HasPrefix(plane, string(pilotNname[0])) || strings.HasPrefix(solution.pilot.firstname, string(pilotNname[0])) || strings.HasPrefix(solution.pilot.lastname, string(pilotNname[0])) || pilotBeenUsed || planeTaken {
		return false
	}
	return true
}

func checkValidPilotLastname(plane string, pilotLName string, pilotBeenUsed bool, planeTaken bool, solution Solution) bool {
	if strings.HasPrefix(plane, string(pilotLName[0])) || strings.HasPrefix(solution.pilot.firstname, string(pilotLName[0])) || pilotBeenUsed || planeTaken {
		return false
	}
	return true
}
func checkValidPilotFirstName(plane string, pilotsName string, pilotBeenUsed bool, planeTaken bool) bool {
	if strings.HasPrefix(plane, string(pilotsName[0])) || pilotBeenUsed || planeTaken || isNavigator(pilotsName) {
		// if plane and pilot have same first character, return false
		// if pilot has been used already, don't reuse
		// if plane already has a pilot, don't assign another one
		// if persons name is a navigator, return false
		return false
	}
	return true
}

func checkValidNavigatorFirstName(plane, navFname string, navBeenUsed, planeTaken, pilotBeenUsed bool) bool {
	if strings.HasPrefix(plane, string(navFname[0])) || navBeenUsed || pilotBeenUsed || planeTaken || isPilot(navFname) {
		// if plane and person have same first character, return false
		// if person has been used already, don't reuse
		// if plane already has a navigator, don't assign another one
		// if persons name is a pilot, return false
		return false
	}
	return true
}

// Given from clues in the puzzle we can deduce 4 navigators
func isNavigator(pilotsName string) bool {
	switch pilotsName {
	case "jack":
		return true
	case "harold":
		return true
	case "chuck":
		return true
	case "brad":
		return true
	default:
		return false
	}

}

func isPilot(personsName string) bool {
	switch personsName {
	default:
		return false
	}
}
