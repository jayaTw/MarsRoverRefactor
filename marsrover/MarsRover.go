package marsrover

import (
	"fmt"
	"strconv"
	"strings"
)

func Run(input string) string {
	out := ""
	lines := strings.Split(input, "\n")
	numberOfRovers := (len(lines) - 1) / 2

	for i := 0; i < numberOfRovers; i++ {
		positionLineIndex := i*2 + 1
		commandLineIndex := positionLineIndex + 1

		var xwidth, ywidth int64

		split := strings.Fields(lines[positionLineIndex])
		xwidth, err1 := strconv.ParseInt(split[0], 10, 64)
		ywidth, err2 := strconv.ParseInt(split[1], 10, 64)

		if err1 != nil || err2 != nil {
			panic(fmt.Sprintf("illegal argument: could not parse position from %v", lines[positionLineIndex]))
		}

		position := []int{int(xwidth), int(ywidth)}

		direction := strings.Fields(lines[positionLineIndex])[2]
		if !contains([]string{"N", "E", "S", "W"}, direction) {
			panic(fmt.Sprintf("illegal argument: could not parse position from %v",direction))
		}

		commandArray := strings.Split(lines[commandLineIndex],"")
		for _,char:= range commandArray{
			if !contains([]string{"L", "M", "R"}, char) && char != " " {
						panic(fmt.Sprintf("illegal argument : invalid command sequence : %v",lines[commandLineIndex]))
					}
		}

		for i := 0; i < len(commandArray); i++ {
			command := string(commandArray[i])
			if command == "M" {
				newPosition := position
				if direction == "N" {
					newPosition[1] += 1
				} else if direction == "S" {
					newPosition[1] += (-1)
				} else if direction == "E" {
					newPosition[0] += (1)
				} else if direction == "W" {
					newPosition[0] += (-1)
				}
				position = newPosition
			} else if command == "R" {

				allDirections := []string{"N", "E", "S", "W"}
				indexOfDirection := getIndexOf(allDirections, direction)
				if indexOfDirection == -1 {
					panic(fmt.Sprintf("invalid argument : invalid direction : %v",direction))
				}
				sizeOfDirections := len(allDirections)
				direction = allDirections[((indexOfDirection + 1) % sizeOfDirections)]

			} else if command == "L" {
				allDirections := []string{"N", "E", "S", "W"}
				indexOfDirection := getIndexOf(allDirections, direction)
				if indexOfDirection == -1 {
					panic(fmt.Sprintf("invalid argument : invalid direction : %v",direction))
				}
				sizeOfDirections := len(allDirections)
				direction = allDirections[((indexOfDirection + 3) % sizeOfDirections)]
			}
		}

		out += fmt.Sprintf("%v %v %v\n", position[0], position[1], direction)

	}
	return out
}

func contains(list []string, charToSearch string) bool {
	for _, char := range list {
		if char == charToSearch {
			return true
		}
	}
	return false
}

func getIndexOf(list []string, char string) int {
	for index, element := range list {
		if element == char {
			return index
		}
	}
	return -1
}
