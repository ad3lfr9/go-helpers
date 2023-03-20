package randomHelper

import (
	"fmt"
	"sort"

	"bitbucket.org/kidozteam/adstats/pkg/platform/fastrand64"
)

// modified version of https://github.com/milosgajdos/go-estimate/blob/511e405c965a7d58f6ce8d5e5fbf4f25618babad/rnd/rnd.go

// RouletteDraw draws a random integer from [0,len(p)) from a probability mass function (PMF) defined by weights in p.
// RouletteDraw implements the Roulette Wheel Draw a.k.a. Fitness Proportionate Selection:
// - https://en.wikipedia.org/wiki/Fitness_proportionate_selection
// - http://www.keithschwarz.com/darts-dice-coins/
func RouletteDrawInt(p []int) (int, error) {
	if len(p) == 0 {
		return 0, fmt.Errorf("Invalid probability weights: %v", p)
	}
	// create the discrete CDF
	cdf := make([]int, len(p))
	cdf[0] = p[0]
	for i, v := range p[1:] {
		cdf[i+1] = cdf[i] + v
	}
	// multiply the sample with the largest CDF value;
	val := int(fastrand64.Uint32n((cdf[len(cdf)-1])))
	// Binary search to return the smallest index i such that cdf[i] > val
	return sort.Search(len(cdf), func(i int) bool { return cdf[i] > val }), nil
}
