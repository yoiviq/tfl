package main

type afaState struct {
	id    int
	term  bool
	edges []*afaTrans
}

type afaTrans struct {
	outStates    []*afaState
	sym          byte
}

type afaBranch struct {
	rest   string
	states []*afaState
}

type bigAFA struct {
	startLeft  *afaState // q0
	startRight *afaState // 0
}

func buildAFA() *bigAFA {

	q0 := &afaState{id: 0}
	q1 := &afaState{id: 1}
	q2 := &afaState{id: 2}
	q3 := &afaState{id: 3}
	q4 := &afaState{id: 4}
	q5 := &afaState{id: 5}
	q6 := &afaState{id: 6}
	q7 := &afaState{id: 7, term: true}

	zero := &afaState{id: 8}          
	a := &afaState{id: 9}
	b := &afaState{id: 10}
	c := &afaState{id: 11}
	aa := &afaState{id: 12}
	ab := &afaState{id: 13, term: true}
	ac := &afaState{id: 14, term: true}
	ba := &afaState{id: 15}
	bb := &afaState{id: 16, term: true}
	bc := &afaState{id: 17, term: true}
	ca := &afaState{id: 18}
	cb := &afaState{id: 19}
	cc := &afaState{id: 20}

	q0.edges = []*afaTrans{
		{outStates: []*afaState{q1}, sym: 'a'},
		{outStates: []*afaState{q3}, sym: 'b'},
	}

	q1.edges = []*afaTrans{
		{outStates: []*afaState{q2}, sym: 'b'},
		{outStates: []*afaState{q0}, sym: 'c'},
	}

	q2.edges = []*afaTrans{
		{outStates: []*afaState{q0}, sym: 'c'},
	}

	q3.edges = []*afaTrans{
		{outStates: []*afaState{q4}, sym: 'a'},
		{outStates: []*afaState{q0}, sym: 'c'},
	}

	q4.edges = []*afaTrans{
		{outStates: []*afaState{q0}, sym: 'c'},
		{outStates: []*afaState{q5}, sym: 'a'},
		{outStates: []*afaState{q5}, sym: 'b'},
		{outStates: []*afaState{q5}, sym: 'c'},
	}

	q5.edges = []*afaTrans{
		{outStates: []*afaState{q6}, sym: 'a'},
		{outStates: []*afaState{q6}, sym: 'b'},
	}

	q6.edges = []*afaTrans{
		{outStates: []*afaState{q7}, sym: 'b'},
		{outStates: []*afaState{q7}, sym: 'c'},
	}

	zero.edges = []*afaTrans{
		{outStates: []*afaState{a}, sym: 'a'},
		{outStates: []*afaState{b}, sym: 'b'},
		{outStates: []*afaState{c}, sym: 'c'},
	}

	a.edges = []*afaTrans{
		{outStates: []*afaState{aa}, sym: 'a'},
		{outStates: []*afaState{ab}, sym: 'b'},
		{outStates: []*afaState{ac}, sym: 'c'},
	}

	b.edges = []*afaTrans{
		{outStates: []*afaState{ba}, sym: 'a'},
		{outStates: []*afaState{bb}, sym: 'b'},
		{outStates: []*afaState{bc}, sym: 'c'},
	}

	c.edges = []*afaTrans{
		{outStates: []*afaState{ca}, sym: 'a'},
		{outStates: []*afaState{cb}, sym: 'b'},
		{outStates: []*afaState{cc}, sym: 'c'},
	}

	aa.edges = []*afaTrans{
		{outStates: []*afaState{aa}, sym: 'a'},
		{outStates: []*afaState{ab}, sym: 'b'},
		{outStates: []*afaState{ac}, sym: 'c'},
	}

	ab.edges = []*afaTrans{
		{outStates: []*afaState{ba}, sym: 'a'},
		{outStates: []*afaState{bb}, sym: 'b'},
		{outStates: []*afaState{bc}, sym: 'c'},
	}

	ac.edges = []*afaTrans{
		{outStates: []*afaState{ca}, sym: 'a'},
		{outStates: []*afaState{cb}, sym: 'b'},
		{outStates: []*afaState{cc}, sym: 'c'},
	}

	ba.edges = []*afaTrans{
		{outStates: []*afaState{aa}, sym: 'a'},
		{outStates: []*afaState{ab}, sym: 'b'},
		{outStates: []*afaState{ac}, sym: 'c'},
	}

	bb.edges = []*afaTrans{
		{outStates: []*afaState{ba}, sym: 'a'},
		{outStates: []*afaState{bb}, sym: 'b'},
		{outStates: []*afaState{bc}, sym: 'c'},
	}

	bc.edges = []*afaTrans{
		{outStates: []*afaState{ca}, sym: 'a'},
		{outStates: []*afaState{cb}, sym: 'b'},
		{outStates: []*afaState{cc}, sym: 'c'},
	}

	ca.edges = []*afaTrans{
		{outStates: []*afaState{aa}, sym: 'a'},
		{outStates: []*afaState{ab}, sym: 'b'},
		{outStates: []*afaState{ac}, sym: 'c'},
	}

	cb.edges = []*afaTrans{
		{outStates: []*afaState{ba}, sym: 'a'},
		{outStates: []*afaState{bb}, sym: 'b'},
		{outStates: []*afaState{bc}, sym: 'c'},
	}

	cc.edges = []*afaTrans{
		{outStates: []*afaState{ca}, sym: 'a'},
		{outStates: []*afaState{cb}, sym: 'b'},
		{outStates: []*afaState{cc}, sym: 'c'},
	}

	return &bigAFA{
		startLeft:  q0,
		startRight: zero,
	}
}

func runAFAOne(start *afaState, word string) bool {
	startConf := &afaBranch{
		rest:   word,
		states: []*afaState{start},
	}

	var stack []*afaBranch
	stack = append(stack, startConf)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		sts := curr.states
		if len(sts) == 1 {
			s := sts[0]
			if s.term && len(curr.rest) == 0 {
				return true
			}
			if len(curr.rest) == 0 {
				continue
			}
			ch := curr.rest[0]
			newRest := curr.rest[1:]
			for _, tr := range s.edges {
				if tr.sym == ch {
					stack = append(stack, &afaBranch{
						rest:   newRest,
						states: tr.outStates,
					})
				}
			}
		} else if len(sts) == 2 {
			if runAFAOne(sts[0], curr.rest) && runAFAOne(sts[1], curr.rest) {
				return true
			}
		}
	}

	return false
}

func AFA(m *bigAFA, word string) bool {
	return runAFAOne(m.startLeft, word) && runAFAOne(m.startRight, word)
}
