package minimax

const (
	maxPossibleVal = 999999
	minPossibleVal = -999999
)

type State interface {
	Eval() int
	GetChildren(isMaximizer bool) []State
}

func FindBestUsingMinimax(currState State, isMaximizer bool) State {
	_, bestState := minimax(currState, minPossibleVal, maxPossibleVal, isMaximizer, 0)
	return bestState
}

func minimax(state State, alpha, beta int, isMaximizer bool, depth uint) (int, State) {
	children := state.GetChildren(isMaximizer)
	if len(children) == 0 {
		return state.Eval() - int(depth), state
	}

	// 21 sec time for first move
	if depth+1 > 8 {
		return 0, children[0]
	}

	var bestState State
	var bestVal int
	if isMaximizer {
		maxVal := minPossibleVal
		for i, chState := range children {
			val, _ := minimax(chState, alpha, beta, false, depth+1)
			if i == 0 || val > maxVal {
				maxVal = val
				bestState = chState

				if val > alpha {
					alpha = val
					// if worst play for maximizer becomes better than max available
					// then no need to search more
					if alpha >= beta {
						break
					}
				}
			}
		}
		bestVal = maxVal
	} else {
		minVal := maxPossibleVal
		for i, chState := range children {
			val, _ := minimax(chState, alpha, beta, true, depth+1)
			if i == 0 || val < minVal {
				minVal = val
				bestState = chState

				if val < beta {
					beta = val
					if beta <= alpha {
						break
					}
				}
			}
		}
		bestVal = minVal
	}

	return bestVal, bestState
}
