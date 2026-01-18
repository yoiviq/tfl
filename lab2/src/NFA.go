package main

type NfaState struct {
	id    int
	final bool
	arcs  []*NfaArc
}

type NfaArc struct {
	to   *NfaState
	sym  rune
}

type NfaMachine struct {
	states []*NfaState
	start  *NfaState
}

func buildNFA() *NfaMachine {

	q0 := &NfaState{id: 0, final: false}
	q1 := &NfaState{id: 1, final: false}
	q2 := &NfaState{id: 2, final: false}
	q3 := &NfaState{id: 3, final: false}
	q4 := &NfaState{id: 4, final: false}
	q5 := &NfaState{id: 5, final: false}
	q6 := &NfaState{id: 6, final: false}
	q7 := &NfaState{id: 7, final: true} 

	// q0 -> q1 [a]
	q0.arcs = append(q0.arcs, &NfaArc{to: q1, sym: 'a'})
	// q0 -> q3 [b]
	q0.arcs = append(q0.arcs, &NfaArc{to: q3, sym: 'b'})

	// q1 -> q2 [b]
	q1.arcs = append(q1.arcs, &NfaArc{to: q2, sym: 'b'})
	// q1 -> q0 [c]
	q1.arcs = append(q1.arcs, &NfaArc{to: q0, sym: 'c'})

	// q2 -> q0 [c]
	q2.arcs = append(q2.arcs, &NfaArc{to: q0, sym: 'c'})

	// q3 -> q4 [a]
	q3.arcs = append(q3.arcs, &NfaArc{to: q4, sym: 'a'})
	// q3 -> q0 [c]
	q3.arcs = append(q3.arcs, &NfaArc{to: q0, sym: 'c'})

	// q4 -> q0 [c]
	q4.arcs = append(q4.arcs, &NfaArc{to: q0, sym: 'c'})

	// q4 -> q5 [a,b,c]
	for _, ch := range []rune{'a', 'b', 'c'} {
		q4.arcs = append(q4.arcs, &NfaArc{to: q5, sym: ch})
	}

	// q5 -> q6 [a,b]
	for _, ch := range []rune{'a', 'b'} {
		q5.arcs = append(q5.arcs, &NfaArc{to: q6, sym: ch})
	}

	// q6 -> q7 [b,c]
	for _, ch := range []rune{'b', 'c'} {
		q6.arcs = append(q6.arcs, &NfaArc{to: q7, sym: ch})
	}

	return &NfaMachine{
		states: []*NfaState{q0, q1, q2, q3, q4, q5, q6, q7},
		start:  q0,
	}
}

type nfaConfig struct {
	rest  string
	state *NfaState
}

func NFA(m *NfaMachine, word string) bool {
	stack := []*nfaConfig{{rest: word, state: m.start}}

	for len(stack) > 0 {

		conf := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(conf.rest) == 0 {
			if conf.state.final {
				return true
			}
			continue
		}

		curr := conf.rest[0]
		nextRest := conf.rest[1:]

		for _, arc := range conf.state.arcs {
			if rune(curr) == arc.sym {
				stack = append(stack, &nfaConfig{rest: nextRest, state: arc.to})
			}
		}
	}

	return false
}
