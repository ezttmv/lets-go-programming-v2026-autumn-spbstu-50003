package main

import "fmt"
func main() {
	var(
		highBord int
		lowBord int
		n int
		k int
		border string
		t int
	)

	_, errN := fmt.Scanln(&n)
	if errN != nil || n < 1 || n > 1000{
		fmt.Println("Invalid number of depatments(1;1000)")
		return
	}
	for i := 1; i <= n; i++ {
		_, errK := fmt.Scan(&k)
		if errK != nil || k < 1 || k > 1000{
			fmt.Println("Invalid number of employees(1;1000)")
			return
		}
		highBord = 30
		lowBord = 15
		for j := 1; j <= k; j++ {
			fmt.Scan(&border, &t)
			if lowBord != -1{
				switch border {
				case ">=":
					if lowBord <= t && highBord >= t{
						lowBord = t
					}else if highBord < t{
						lowBord = -1
					}
				case "<=" :
					if  lowBord <= t && highBord >= t{
						highBord = t
					}else if lowBord > t {
						lowBord = -1
					}
				default:
					fmt.Println("invalid")
					j--
					continue
				}
			}
			fmt.Println(lowBord)
		}
	}



}