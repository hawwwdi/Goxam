package exam

import "fmt"

// noq : number of questions
// noa : number of answers
// nom : number of incorrect answers
// forfeit : forfeit coefficient
func CalcScore(noq, noa, noi, forfeit int) string {
	var score float32
	if forfeit != 0 {
		score = (float32(noa) - (1.0/float32(forfeit))*float32(noi)) / float32(noq)
	} else {
		score = float32(noa) / float32(noq)
	}
	score *= 100
	return fmt.Sprint(score)
}
