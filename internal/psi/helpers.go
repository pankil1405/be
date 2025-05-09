package psi

import (
	"github.com/welltodohamm/be/types"
	"github.com/onsi/gomega"
)

// IsMatcher returns true if given input is either Omega or Gomock or a Psi matcher
func IsMatcher(a any) bool {
	switch a.(type) {
	case types.GomegaMatcher, types.GomockMatcher:
		return true
	default:
		return false
	}
}

// AsMatcher returns BeMatcher that is made from given input
func AsMatcher(m any) types.BeMatcher {
	switch t := m.(type) {
	case types.BeMatcher:
		return t
	case types.GomegaMatcher:
		return FromGomega(t)
	case types.GomockMatcher:
		return FromGomock(t)
	default:
		return FromGomega(gomega.Equal(t))
	}
}

// asGomegaMatchFunc returns given `m any` as match func (func (any) (bool, error))
func asGomegaMatchFunc(m any) func(any) (bool, error) {
	v, ok := m.(func(any) (bool, error))
	if !ok {
		return nil
	}
	return v
}
