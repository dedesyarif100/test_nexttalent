// PSBB ( Pembatasan Sosial Berskala Besar ) (60 Poin) (NDL011)
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var totalFamily int
	var memberFamily string
	var minimumPeople = 4
	var totalPeople = 0
	var totalBusway = 0
	var arrayFamily []int
	var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

	fmt.Print("Input the number of families : ")
	fmt.Scanln(&totalFamily)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Input the number of members in the family (separated by a space) : ")
	scanner.Scan()

	memberFamily = scanner.Text()
	words := strings.Split(memberFamily, " ")
	if len(words) != totalFamily {
		fmt.Println("\nInput must be equal with count of family")
		return
	}

	for key, val := range words {
		if IsLetter(val) {
			fmt.Println("\nValue must number, but in order ", key+1, " value is a : ", val)
			return
		}
		num, _ := strconv.Atoi(val)
		arrayFamily = append(arrayFamily, num)
	}

	for _, val := range arrayFamily {
		totalPeople = totalPeople + val
	}

	for i := totalPeople; i >= 1; i = i - minimumPeople {
		totalBusway++
	}
	fmt.Println("\nMinimum bus required is : ", totalBusway)
}
