package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {

	f, _ := os.Open("input.txt")
	numbers := []string{}
	var gamma_rate int

	for {
		var temp string
		_, err := fmt.Fscanf(f, "%s\n", &temp)
		if err != nil {
			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err) //error in the file
			os.Exit(1)
		}
		numbers = append(numbers, temp)
	}

	for i := 0; i < len(numbers[0]); i++ {
		gamma_rate += calculate_gamma(numbers, i) * int((math.Pow(10, float64(len(numbers[0])-1-i)))) //just number * 10/100/1000
	}

	fmt.Println("gamma rate is ", gamma_rate)
	fmt.Println("or ", toDecimal(gamma_rate), "in decimal")

	epsilon_rate := reverse_int(gamma_rate) //epsilon is just inversed binary gamma
	fmt.Println("gamma rate is ", epsilon_rate)

	epsilon_rate_int, _ := strconv.Atoi(epsilon_rate) //cast to int
	fmt.Println("or ", toDecimal(epsilon_rate_int), "in decimal")

	power_consumption := toDecimal(gamma_rate) * toDecimal(epsilon_rate_int)
	fmt.Println("power consumption is ", power_consumption)

	//----- PART TWO -----//
	fmt.Println("\n----- PART TWO -----")
	oxygen_rate := oxygen_rating(numbers, 0)
	oxygen_rate_int, _ := strconv.Atoi(oxygen_rate)
	fmt.Println("Oxygen rating : ", oxygen_rate)
	fmt.Println("or ", toDecimal(oxygen_rate_int), "in decimal")

	CO2_rate := CO2_rating(numbers, 0)
	CO2_rate_int, _ := strconv.Atoi(CO2_rate)
	fmt.Println("CO2 rating : ", CO2_rate)
	fmt.Println("or ", toDecimal(CO2_rate_int), "in decimal")

	life_support := toDecimal(oxygen_rate_int) * toDecimal(CO2_rate_int)
	fmt.Println("life support consumption is ", life_support)

}

func calculate_gamma(numbers []string, i int) int {

	var counter int // 0 = +1, 1 = -1

	for j := 0; j < len(numbers); j++ {
		//fmt.Println("index ", i, j, string(numbers[j][i]))
		num, err := strconv.Atoi(string(numbers[j][i]))
		if err != nil {
			panic(err)
		}
		if num == 0 {
			counter += 1
		} else {
			counter -= 1
		}
	}

	if counter > 0 { //there are more 0 than 1
		return 0
	} else {
		return 1
	}
}

func reverse_int(value int) string {

	intString := strconv.Itoa(value)

	invString := ""

	for x := 0; x < len(intString); x++ {
		if intString[x] == '1' {
			invString += "0"
		} else {
			invString += "1"

		}
	}

	return invString
}

func toDecimal(num int) int {
	var dec_num float64
	inc := 0
	for num > 0 {
		remainder := num % 10
		dec_num += float64(remainder) * math.Pow(2, float64(inc))
		inc += 1
		num /= 10
	}

	return int(dec_num)
}

//----- PART TWO -----//

func oxygen_rating(numbers []string, i int) string {
	counter := 0
	var choosen_num int
	var choosen_number = []string{}

	if len(numbers) == 1 {
		return numbers[0]
	}

	for j := 0; j < len(numbers); j++ {
		//fmt.Println("index ", i, j, string(numbers[j][i]))
		num, err := strconv.Atoi(string(numbers[j][i]))
		if err != nil {
			panic(err)
		}
		if num == 0 {
			counter += 1
		} else {
			counter -= 1
		}
	}

	if counter > 0 { //there are more 0 than 1
		choosen_num = 0
	} else if counter < 0 {
		choosen_num = 1
	} else if counter == 0 {
		choosen_num = 1
	}

	for j := 0; j < len(numbers); j++ {
		num, _ := strconv.Atoi(string(numbers[j][i]))
		if num == choosen_num {
			choosen_number = append(choosen_number, numbers[j])
		}
	}

	//fmt.Println(choosen_number) //print selection

	i += 1

	return oxygen_rating(choosen_number, i)
}

func CO2_rating(numbers []string, i int) string {
	counter := 0
	var choosen_num int
	var choosen_number = []string{}

	if len(numbers) == 1 {
		return numbers[0]
	}

	for j := 0; j < len(numbers); j++ {
		//fmt.Println("index ", i, j, string(numbers[j][i]))
		num, err := strconv.Atoi(string(numbers[j][i]))
		if err != nil {
			panic(err)
		}
		if num == 0 {
			counter += 1
		} else {
			counter -= 1
		}
	}

	if counter > 0 { //there are more 0 than 1
		choosen_num = 1
	} else if counter < 0 {
		choosen_num = 0
	} else if counter == 0 {
		choosen_num = 0
	}

	for j := 0; j < len(numbers); j++ {
		num, _ := strconv.Atoi(string(numbers[j][i]))
		if num == choosen_num {
			choosen_number = append(choosen_number, numbers[j])
		}
	}

	//fmt.Println(choosen_number) //print selection

	i += 1

	return CO2_rating(choosen_number, i)
}
