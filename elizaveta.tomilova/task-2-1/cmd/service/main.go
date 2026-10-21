package main

import "fmt"

func chadgeBorder(lowBord *int, highBord *int, border string, temp int) bool {
	if *lowBord != -1 {
		switch border {
		case ">=":
			if *lowBord <= temp && *highBord >= temp {
				*lowBord = temp
			} else if *highBord < temp {
				*lowBord = -1
			}
		case "<=":
			if *lowBord <= temp && *highBord >= temp {
				*highBord = temp
			} else if *lowBord > temp {
				*lowBord = -1
			}
		default:
			fmt.Println("invalid")

			return false
		}
	}

	return true
}

func main() {
	var (
		highBord    int
		lowBord     int
		departments int
		employees   int
		border      string
		temp        int
	)

	_, errN := fmt.Scanln(&departments)
	if errN != nil || departments < 1 || departments > 1000 {
		fmt.Println("Invalid number of depatments(1;1000)")

		return
	}

	for i := 1; i <= departments; i++ {
		_, errK := fmt.Scan(&employees)
		if errK != nil || employees < 1 || employees > 1000 {
			fmt.Println("Invalid number of employees(1;1000)")

			return
		}

		highBord = 30
		lowBord = 15

		for emp := 1; emp <= employees; emp++ {
			_, errB := fmt.Scan(&border, &temp)
			if errB != nil {
				return
			}

			if !chadgeBorder(&lowBord, &highBord, border, temp) {
				emp--

				continue
			}

			fmt.Println(lowBord)
		}
	}
}
