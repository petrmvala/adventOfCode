package main

import (
	"testing"
)

const conns = `kh-tc
qp-kh
de-cg
ka-co
yn-aq
qp-ub
cg-tb
vc-aq
tb-ka
wh-tc
yn-cg
kh-ub
ta-co
de-co
tc-td
tb-wq
wh-td
ta-ka
td-qp
aq-cg
wq-ub
ub-vc
de-ta
wq-aq
wq-vc
wh-yn
ka-de
kh-ta
co-tc
wh-qp
tb-vc
td-yn
`

const sets = `co,de,ta
co,ka,ta
de,ka,ta
qp,td,wh
tb,vc,wq
tc,td,wh
td,wh,yn
`

func TestGetSets(t *testing.T) {
	s := GetSets(conns)
	got := setsToString(t, s)
	want := sets
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func setsToString(t testing.TB, sets []string) string {
	t.Helper()
	var s string
	for _, set := range sets {
		s += set + "\n"
	}
	return s
}
