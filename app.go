package main

import (
	"context"
	"math/rand"
	"strconv"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type BSN struct {
	Leading string
	I4      int
	I5      int
	I6      int
	I7      int
	I8      int
}

func (a *App) GenerateBSN(withLeadingZeroes bool) string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano())) // Create a new random number generator

	bsn := BSN{
		Leading: "9999",
		I4:      9,
		I5:      rng.Intn(10),
		I6:      rng.Intn(10),
		I7:      rng.Intn(10),
		I8:      0,
	}

	if withLeadingZeroes {
		bsn.Leading = "0000"
	}

	sum := 0
	for i, r := range bsn.Leading {
		num := int(r - '0') // Convert rune to integer
		sum += num * (9 - i)
	}

	sum += 5*bsn.I4 + 4*bsn.I5 + 3*bsn.I6 + 2*bsn.I7
	bsn.I8 = sum - (sum/11)*11

	if bsn.I8 > 9 && bsn.I7 > 0 {
		bsn.I7 -= 1
		bsn.I8 = 8
	}

	if bsn.I8 > 9 && bsn.I7 <= 0 {
		bsn.I7 += 1
		bsn.I8 = 1
	}

	return bsn.Leading + strconv.Itoa(bsn.I4) + strconv.Itoa(bsn.I5) + strconv.Itoa(bsn.I6) + strconv.Itoa(bsn.I7) + strconv.Itoa(bsn.I8)
}

func (a *App) IsValidBSN(bsn string) bool {
	length := len(bsn)
	if length < 8 || length > 9 {
		return false
	}

	// Convert the BSN string to an array of integers
	numbers := make([]int, length)
	for i, char := range bsn {
		num, err := strconv.Atoi(string(char))
		if err != nil {
			return false // If a character is not a digit, return false
		}
		numbers[i] = num
	}

	lastNumber := numbers[length-1]

	// Perform the modulus-11 check
	sum := 0
	for i, num := range numbers[:length-1] { // Exclude the last digit
		sum += num * (length - i)
	}

	return (sum-lastNumber)%11 == 0
}
