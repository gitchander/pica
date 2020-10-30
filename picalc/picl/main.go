package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"time"

	"github.com/gitchander/pica/picalc/pigla"
)

func main() {

	var (
		destFileName string
		digits       int
		showSteps    bool
		formatNumber string
	)

	flag.StringVar(&destFileName, "dest", "", "Destination file name")
	flag.IntVar(&digits, "digits", 50, "Number of decimal digits")
	flag.BoolVar(&showSteps, "steps", false, "Show steps")
	flag.StringVar(&formatNumber, "format", "dec", "Format of a result number")

	flag.Parse()

	var (
		nod10 = float64(digits)                // Number of decimal digits
		nod2  = pigla.ConvertNOD(10, nod10, 2) // Number of binary digits
	)

	prec := uint(math.Floor(nod2))
	fmt.Println("prec:", prec)

	//steps := int(math.Log2(float64(prec)))
	steps := int(math.Log2(nod10)) + 1
	fmt.Println("steps:", steps)

	var sf pigla.StepFunc
	if showSteps {
		sf = func(step int, pi *big.Float) {
			if digits < 10000 {
				fmt.Printf("%d: %v\n", step, pi)
			}
		}
	}

	start := time.Now()
	pi := pigla.CalcBigPi(prec, steps, sf)
	fmt.Println("calculate duration:", time.Since(start))

	if digits < 10000 {
		fmt.Println("Pi =", pi)
	}

	if destFileName != "" {

		start := time.Now()

		// text := pi.Text('g', -1) // base 10
		// text := pi.Text('x', -1) // base 16

		text := make([]byte, 0, 1024)

		switch formatNumber {
		case "dec": // decimal - base 10
			text = pi.Append(text, 'g', -1)
			fmt.Println("convert to decimal duration:", time.Since(start))

		case "hex": // hexadecimal - base 16
			text = pi.Append(text, 'x', -1)
			fmt.Println("convert to hexadecimal duration:", time.Since(start))

		default:
			log.Fatalf("invalid number format %q", formatNumber)
		}

		err := ioutil.WriteFile(destFileName, text, 0666)
		checkError(err)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func exampleBigFloatFormatParse(x *big.Float) {

	prec := x.Prec()

	//var formatByte byte = 'g' // base 10
	var formatByte byte = 'x' // base 16

	textX := x.Text(formatByte, -1)

	y, _, err := big.ParseFloat(textX, 0, prec, big.ToNearestEven)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("equal:", x.Cmp(y))

	textY := y.Text(formatByte, -1)

	fmt.Println(textX)
	fmt.Println(textY)
}
