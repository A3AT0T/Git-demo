package main

import (
	"fmt"
	"os"
)

var input string

func main() {
L1:
	fmt.Println(`Оберіть потрібний варіант: 
	1 -- Налаштувати свіч Raisecom 2624
	2 -- Налаштувати свіч Raisecom 2648
	3 -- Налаштувати свіч Raisecom 2128
	3 -- Налаштувати свіч Foxgate s1
	v -- Свортити діапазон вланів на Extrime
	`)

	fmt.Scanf("%v", &input)
	switch input {
	case "1":
		fmt.Print("Raisecom 2648")
	case "2":
		fmt.Print("Raisecom 2624")
	case "3":
		fmt.Print("Raisecom 2128")
	case "4":
		fmt.Print("Foxgate s1")
	case "v":
		fmt.Print("Введіть діапазон: ")
		var first_element int
		var last_elemetn int
		fmt.Fscan(os.Stdin, &first_element, &last_elemetn)
		//fmt.Scanf("%d %d",&first_element, &last_elemetn) <<< ------------ ХЗ
		for first_element <= last_elemetn {
			fmt.Printf(
				`create vlan v%d tag %d
disable igmp proxy-query  vlan v%d
disable igmp v%d
disable igmp snooping vlan v%d
configure igmp router-alert transmit off vlan v%d `, first_element, first_element, first_element, first_element, first_element, first_element)
			first_element++
		}
	case "q":
		break
	default:
		fmt.Println("Не вірний варіант. Для виходу натисни 'q' або повтори вібір: \n ")
		goto L1

	}
}
