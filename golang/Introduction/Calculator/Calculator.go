package Calculator

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	return x / y
}

func gcf(i, j int) int {
	if i == j {
		return i
	} else if i == 0 {
		return j
	} else if j == 0 {
		return i
	} else if i < j {
		res := j % i
		return gcf(i, res)
	} else if j < i {
		res := i % j
		return gcf(j, res)
	} else {
		return 1
	}
}

func nthroot(x float64, n int) float64 {
	guess := 2.0
	change := 1.0
	closeness := math.MaxFloat64
	lastInacc := []float64{math.MaxFloat64}

	for i := 0; i < 15; i++ {
		dif := float64(n) - math.Abs(math.Log(x)/math.Log(guess))
		change = change / float64(i+1)
		fmt.Printf("Change factor: %f\n\n", change)
		count := 0
		for dif < 0 {
			count++
			if count > 100000 {
				fmt.Printf("Infinite loop prevention at increase, adjusted change\n")
				change = 2*change + 1
				break
			}
			guess *= (1 + change)
			dif = float64(n) - math.Abs(math.Log(x)/math.Log(guess))
		}
		lastInacc = append(lastInacc, dif)

		dInacc := lastInacc[len(lastInacc)-2] - dif

		fmt.Printf("Estimate: %f\n", guess)
		fmt.Printf("Estimate inaccuracy: %f\n", dif)
		fmt.Printf("Inaccuracy change: %f\n\n", dInacc)
		if math.Abs(dif) < 0.00000000001 && dInacc < 0.00000000001 {
			fmt.Printf("Appropriate accuracy reached...\n\n")
			guess = math.Round(guess * 100000000)
			guess /= 100000000
			return guess
		}

		change /= 2
		count = 0
		for dif > 0 {
			count++
			if count > 100000 {
				fmt.Printf("Infinite loop prevention at decrease\n")
				change = 0.5
				break
			}
			guess *= (1 - change)
			dif = float64(n) - math.Abs(math.Log(x)/math.Log(guess))
		}
		if dif < closeness {
			closeness = dif
		} else {
			guess *= (1 + change)
		}

		fmt.Printf("init Guess iteration %d: %f\n\n", i, guess)
	}

	for index := 0; index < 5; index++ {
		guess = float64((n-1)/n)*guess + (x / (float64(n) * math.Pow(guess, float64(n-1))))
		fmt.Printf("recusive Guess iteration %d: %f\n", index, guess)
	}
	fmt.Printf("Final guess: %f\n", guess)
	return guess
}

func exponent(x, y float64) float64 {
	var ans float64 = 1

	if y < 0 {
		x = 1 / x
		y = -y
	}

	nth := y - math.Floor(y)

	for index := 0; index < int(math.Floor(y)); index++ {
		ans *= x
	}

	if math.Abs(nth) > 0.001 {
		numerator, denominator := int(math.Floor(nth*1000000)), 1000000
		gcf := gcf(numerator, denominator)
		numerator /= gcf
		denominator /= gcf
		// fmt.Printf("fractional exponent: %d/%d\n", numerator, denominator)

		temp := nthroot(ans, denominator)
		for index := 0; index < numerator; index++ {
			temp *= x
		}

		ans *= temp
	}

	return ans
}

type operand struct {
	operation func(float64, float64) float64
	name      string
	operator  string
}

func Calculator() {
	addition := operand{
		operation: func(x, y float64) float64 {
			return add(x, y)
		},
		name:     "addition",
		operator: "+",
	}
	subtraction := operand{
		operation: func(x, y float64) float64 {
			return subtract(x, y)
		},
		name:     "subtraction",
		operator: "-",
	}
	multiplication := operand{
		operation: func(x, y float64) float64 {
			return multiply(x, y)
		},
		name:     "multiplication",
		operator: "*",
	}
	division := operand{
		func(x, y float64) float64 {
			return divide(x, y)
		},
		"division",
		"/",
	}
	exponent := operand{
		func(x, y float64) float64 {
			return math.Pow(x, y)
		},
		"exponent",
		"^",
	}
	oparr := []operand{addition, subtraction, multiplication, division, exponent}
	operations := make(map[string]operand)

	fmt.Println("Basic calculator.\nEnter commands in the form of \n<Number> <Operator> <Number>")
	fmt.Println("The list of operators are: ")
	for _, value := range oparr {
		operator := value.operator
		fmt.Print(operator + ", ")
		operations[operator] = value
	}
	fmt.Println("")
	var ans float64
	var x float64
	var y float64
	var err error
	var err2 error
	var opstr string
	var op operand

	calcOn := true
	hasAns := false

	for calcOn {
		var command string
		fmt.Scanln(&command)
		components := strings.Split(command, "")

		var xstr string
		var ystr string
		if len(components) == 0 {
			hasAns = false
			ans = 0
			continue
		}
		if hasAns {
			x = ans
			opstr = components[0]
			for _, c := range components[1:] {
				ystr += c
			}
			y, err = strconv.ParseFloat(ystr, len(ystr)-1)
		} else {
			index := 0
			for i, c := range components {
				_, exists := operations[c]
				if exists {
					opstr = c
					index = i
					break
				}
				xstr += c
			}
			x, err = strconv.ParseFloat(xstr, len(xstr)-1)
			for _, c := range components[index+1:] {
				ystr += c
			}
			y, err2 = strconv.ParseFloat(ystr, len(ystr)-1)
		}
		// if len(components) == 2 && hasAns {
		// 	x = ans
		// 	opstr = components[0]
		// 	y, err = strconv.ParseFloat(components[1], len(components[1]))
		// } else if len(components) == 3 {
		// 	x, err = strconv.ParseFloat(components[0], len(components[0]))
		// 	opstr = components[1]
		// 	y, err2 = strconv.ParseFloat(components[2], len(components[2]))
		// } else  else {
		// 	err = errors.New("bad command")
		// }
		fmt.Println(xstr + " " + opstr + " " + ystr)
		if xstr == "quit" {
			break
		}
		if err != nil || err2 != nil {
			fmt.Println("Invalid command... ")
			if hasAns {
				fmt.Print(ans)
			}
			continue
		}
		op = operations[opstr]
		ans = op.operation(x, y)
		hasAns = true
		fmt.Print(ans)
	}

}
