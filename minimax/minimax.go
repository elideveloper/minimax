package minimax

const (
	maxPossibleVal = 999999.0
	minPossibleVal = -999999.0
)

type State interface {
	Eval() float64
	GetChildren(isMaximizer bool) []State
}

func FindBestUsingMinimax(currState State, isMaximizer bool) State {
	_, bestState := minimax(currState, minPossibleVal, maxPossibleVal, isMaximizer, 1)
	return bestState
}

func minimax(state State, alpha, beta float64, isMaximizer bool, depth uint) (float64, State) {
	children := state.GetChildren(isMaximizer)
	if len(children) == 0 {
		return state.Eval() / float64(depth), state
	}

	var bestState State
	var bestVal float64
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

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
