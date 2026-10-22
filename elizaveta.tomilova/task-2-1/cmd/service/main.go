package main

import "fmt"

type tempRange struct {
	low  int
	high int
}

func (r *tempRange) chadgeBorder(border string, temp int) int {
	if r.low != -1 {
		switch border {
		case ">=":
			if r.low <= temp && r.high >= temp {
				r.low = temp
			} else if r.high < temp {
				r.low = -1
			}
		case "<=":
			if r.low <= temp && r.high >= temp {
				r.high = temp
			} else if r.low > temp {
				r.low = -1
			}
		default:
			fmt.Println("invalid")

			return -1
		}
	}

	return r.low
}

func main() {
	var (
		border string
		temp   int
	)
	var departments int
	_, err := fmt.Scanln(&departments)
	if err != nil || departments < 1 || departments > 1000 {
		fmt.Println("Invalid number of depatments(1;1000)")

		return
	}
	var employees int
	for i := 1; i <= departments; i++ {
		_, err := fmt.Scan(&employees)
		if err != nil || employees < 1 || employees > 1000 {
			fmt.Println("Invalid number of employees(1;1000)")

			return
		}
		const (
			highBord = 30
			lowBord  = 15
		)
		tempRng := tempRange{low: lowBord, high: highBord}
		for emp := 1; emp <= employees; emp++ {
			_, err := fmt.Scan(&border, &temp)
			if err != nil {
				fmt.Println("Invalid border/temp")
				return
			}

			fmt.Println(tempRng.chadgeBorder(border, temp))

		}
	}
}
