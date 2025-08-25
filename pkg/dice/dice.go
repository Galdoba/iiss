package dice

import (
	"math/rand"
	"time"
)

// Modifier определяет функцию-модификатор для результатов броска
type Modifier func(*DicePool)

// DicePool представляет пул кубиков с текущими результатами
type DicePool struct {
	seed    int64
	rng     *rand.Rand
	results []int
	sumMod  int
}

// NewDicePool создает новый пул кубиков со случайным seed (на основе времени)
func NewDicePool() *DicePool {
	seed := time.Now().UnixNano()
	return &DicePool{
		seed:    seed,
		rng:     rand.New(rand.NewSource(seed)),
		results: make([]int, 0),
	}
}

func (dp *DicePool) Roll(n int, opts ...RollOption) *DicePool {
	dp.results = make([]int, n)
	for i := range n {
		dp.results[i] = dp.rng.Intn(6) + 1
	}
	for _, modify := range opts {
		modify(dp)
	}
	return dp
}

type RollOption func(*DicePool)

func Mods(mods ...int) RollOption {
	return func(dp *DicePool) {
		dp.results = append(dp.results, mods...)
	}
}

func TreatAs(a, b int) RollOption {
	return func(dp *DicePool) {
		for i := range dp.results {
			if dp.results[i] == a {
				dp.results[i] = b
			}
		}
	}
}

/////////////////////

func (dp *DicePool) Sum() int {
	s := 0
	for _, r := range dp.results {
		s += r
	}
	return s
}
