package main

type afaState struct {
	id    int
	term  bool
	edges []*afaTrans
}

type afaTrans struct {
	outStates []*afaState
	sym       byte
	isEps     bool
}

type bigAFA struct {
	start *afaState 
}

func buildAFA() *bigAFA {

	p0 := &afaState{id: 0}
	p1 := &afaState{id: 1}
	p2 := &afaState{id: 2}
	p3 := &afaState{id: 3}
	p4 := &afaState{id: 4}
	p5 := &afaState{id: 5}
	p6 := &afaState{id: 6}
	p7 := &afaState{id: 7, term: true}
	p8 := &afaState{id: 8}
	p9 := &afaState{id: 9}
	p10 := &afaState{id: 10, term: true}
	p11 := &afaState{id: 11, term: true}
	p12 := &afaState{id: 12}
	p13 := &afaState{id: 13, term: true}

	p0.edges = []*afaTrans{
		{outStates: []*afaState{p1}, sym: 'a'},
		{outStates: []*afaState{p3}, sym: 'b'},
	}
	p1.edges = []*afaTrans{
		{outStates: []*afaState{p2}, sym: 'b'},
		{outStates: []*afaState{p0}, sym: 'c'},
	}
	p2.edges = []*afaTrans{
		{outStates: []*afaState{p0}, sym: 'c'},
	}
	p3.edges = []*afaTrans{
		{outStates: []*afaState{p4}, sym: 'a'},
		{outStates: []*afaState{p0}, sym: 'c'},
	}
	p4.edges = []*afaTrans{
		{outStates: []*afaState{p5}, sym: 'a'},
		{outStates: []*afaState{p5}, sym: 'b'},
		{outStates: []*afaState{p8}, sym: 'c'},
	}
	p5.edges = []*afaTrans{
		{outStates: []*afaState{p6}, sym: 'a'},
		{outStates: []*afaState{p6}, sym: 'b'},
		{outStates: []*afaState{p6}, sym: 'c'},
	}
	p6.edges = []*afaTrans{
		{outStates: []*afaState{p7}, sym: 'a'},
		{outStates: []*afaState{p7}, sym: 'b'},
		{outStates: []*afaState{p7}, sym: 'c'},
	}
	p8.edges = []*afaTrans{
		{outStates: []*afaState{p9}, sym: 'a'},
		{outStates: []*afaState{p12}, sym: 'b'},
		{outStates: []*afaState{p6}, sym: 'c'},
	}
	p9.edges = []*afaTrans{
		{outStates: []*afaState{p7}, sym: 'a'},
		{outStates: []*afaState{p11}, sym: 'b'},
		{outStates: []*afaState{p10}, sym: 'c'},
	}
	p10.edges = []*afaTrans{
		{outStates: []*afaState{p1}, sym: 'a'},
		{outStates: []*afaState{p3}, sym: 'b'},
	}
	p11.edges = []*afaTrans{
		{outStates: []*afaState{p0}, sym: 'c'},
	}
	p12.edges = []*afaTrans{
		{outStates: []*afaState{p13}, sym: 'a'},
		{outStates: []*afaState{p7}, sym: 'b'},
		{outStates: []*afaState{p10}, sym: 'c'},
	}
	p13.edges = []*afaTrans{
		{outStates: []*afaState{p5}, sym: 'a'},
		{outStates: []*afaState{p5}, sym: 'b'},
		{outStates: []*afaState{p8}, sym: 'c'},
	}

	zero := &afaState{id: 14}
	ab := &afaState{id: 15}             
	_b := &afaState{id: 16, term: true} 
	_c := &afaState{id: 17, term: true}    

	zero.edges = []*afaTrans{
		{outStates: []*afaState{ab}, sym: 'a'},
		{outStates: []*afaState{ab}, sym: 'b'},
		{outStates: []*afaState{zero}, sym: 'c'},
	}
	ab.edges = []*afaTrans{
		{outStates: []*afaState{ab}, sym: 'a'},
		{outStates: []*afaState{_b}, sym: 'b'},
		{outStates: []*afaState{_c}, sym: 'c'},
	}
	_b.edges = []*afaTrans{
		{outStates: []*afaState{ab}, sym: 'a'},
		{outStates: []*afaState{_b}, sym: 'b'},
		{outStates: []*afaState{_c}, sym: 'c'},
	}
	_c.edges = []*afaTrans{
		{outStates: []*afaState{ab}, sym: 'a'},
		{outStates: []*afaState{ab}, sym: 'b'},
		{outStates: []*afaState{zero}, sym: 'c'},
	}

	amp := &afaState{id: 100}
	amp.edges = []*afaTrans{
		{outStates: []*afaState{p0, zero}, isEps: true},
	}

	return &bigAFA{start: amp}
}

type memoKey struct {
	id  int
	pos int
}

func acceptAll(states []*afaState, w []byte, pos int, memo map[memoKey]byte) bool {
	for _, st := range states {
		if !acceptState(st, w, pos, memo) {
			return false
		}
	}
	return true
}

func acceptState(s *afaState, w []byte, pos int, memo map[memoKey]byte) bool {
	key := memoKey{id: s.id, pos: pos}

	if v, ok := memo[key]; ok {
		switch v {
		case 2:
			return false
		case 3:
			return true
		case 1:

			return false
		}
	}

	memo[key] = 1

	if pos == len(w) && s.term {
		memo[key] = 3
		return true
	}

	for _, tr := range s.edges {
		if tr.isEps {
			if acceptAll(tr.outStates, w, pos, memo) {
				memo[key] = 3
				return true
			}
			continue
		}
		if pos < len(w) && w[pos] == tr.sym {
			if acceptAll(tr.outStates, w, pos+1, memo) {
				memo[key] = 3
				return true
			}
		}
	}

	memo[key] = 2
	return false
}

func AFA(m *bigAFA, word string) bool {
	memo := make(map[memoKey]byte, 128)
	return acceptState(m.start, []byte(word), 0, memo)
}
