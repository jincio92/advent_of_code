package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("day4/day4.input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	message := []string{}
	for scanner.Scan() {
		message = append(message, scanner.Text())
	}

	total := 0
	for i := 0; i < len(message); i++ {
		for j := 0; j < len(message[i]); j++ {

			if string(message[i][j]) == "X" {
				if i > 2 {
					if string(message[i-1][j]) == "M" {
						if string(message[i-2][j]) == "A" {
							if string(message[i-3][j]) == "S" {
								total++
							}
						}
					}
					if j > 2 {
						if string(message[i-1][j-1]) == "M" {
							if string(message[i-2][j-2]) == "A" {
								if string(message[i-3][j-3]) == "S" {
									total++
								}
							}
						}
					}
					if j < len(message[i])-3 {
						if string(message[i-1][j+1]) == "M" {
							if string(message[i-2][j+2]) == "A" {
								if string(message[i-3][j+3]) == "S" {
									total++
								}
							}
						}
					}
				}
				if i < len(message)-3 {
					if string(message[i+1][j]) == "M" {
						if string(message[i+2][j]) == "A" {
							if string(message[i+3][j]) == "S" {
								total++
							}
						}
					}
					if j > 2 {
						if string(message[i+1][j-1]) == "M" {
							if string(message[i+2][j-2]) == "A" {
								if string(message[i+3][j-3]) == "S" {
									total++
								}
							}
						}
					}
					if j < len(message[i])-3 {
						if string(message[i+1][j+1]) == "M" {
							if string(message[i+2][j+2]) == "A" {
								if string(message[i+3][j+3]) == "S" {
									total++
								}
							}
						}
					}
				}
				if j > 2 {
					if string(message[i][j-1]) == "M" {
						if string(message[i][j-2]) == "A" {
							if string(message[i][j-3]) == "S" {
								total++
							}
						}
					}
				}
				if j < len(message[i])-3 {
					if string(message[i][j+1]) == "M" {
						if string(message[i][j+2]) == "A" {
							if string(message[i][j+3]) == "S" {
								total++
							}
						}
					}
				}
			}
		}
	}
	fmt.Printf("Total XMAS: %v", total)
	println()

	total = 0
	for i := 1; i < len(message)-1; i++ {
		for j := 1; j < len(message[i])-1; j++ {

			if string(message[i][j]) == "A" {
				surrounding := fmt.Sprintf("%c%c%c%c", message[i-1][j-1], message[i+1][j+1], message[i+1][j-1], message[i-1][j+1])
				if strings.Count(surrounding, "M") == 2 && strings.Count(surrounding, "S") == 2 && surrounding[0] != surrounding[1] {
					println("found! " + surrounding)
					total++
				}
			}
		}
	}
	fmt.Printf("Total X shaped MAS: %v", total)
	println()
}
