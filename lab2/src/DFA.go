package main

func DFA(word string) bool {
	
	const (
		q0 = 0
		q1 = 1
		q2 = 2
		q3 = 3
		q4 = 4
		q5 = 5
		q6 = 6
		q7 = 7
		q8 = 8
		q9 = 9
		q10 = 10
		q11 = 11
		q12 = 12
		qDead = 13
	)

	// det[state][symbolIndex] = nextState
	det := [][]int{
		// a   b    c
		/*q0*/ {q1, q2, qDead},
		/*q1*/ {qDead, q3, q0},
		/*q2*/ {q4, qDead, q0},
		/*q3*/ {qDead, qDead, q0},
		/*q4*/ {q5, q5, q6},
		/*q5*/ {q7, q7, qDead},
		/*q6*/ {q8, q9, qDead},
		/*q7*/ {qDead, q10, q10},
		/*q8*/ {qDead, q11, q12},
		/*q9*/ {q4, q10, q12},
		/*q10*/ {qDead, qDead, qDead},
		/*q11*/ {qDead, qDead, q0},
		/*q12*/ {q1, q2, qDead},
		/*dead*/ {qDead, qDead, qDead},
	}

	finalStates := map[int]bool{
		q10: true,
		q11: true,
		q12: true,
	}

	state := q0
	for _, r := range word {
		var idx int
		switch r {
		case 'a':
			idx = 0
		case 'b':
			idx = 1
		case 'c':
			idx = 2
		default:
			return false
		}
		state = det[state][idx]
	}

	return finalStates[state]
}
