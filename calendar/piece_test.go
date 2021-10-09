package calendar

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func TestNewPiece(t *testing.T) {
	g := NewGomegaWithT(t)

	//p, ok := NewPiece(P_L, 3, 5, O_LT, true)
	//g.Expect(ok).To(BeTrue())
	//
	//fmt.Printf("%s\n", p.String())

	p, ok := NewPiece(P_L, 0, 2, O_RT, false)
	g.Expect(ok).To(BeTrue())

	fmt.Printf("%s\n", p.String())
}
