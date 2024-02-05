package main

func shortestSequence(rolls []int, k int) int {

}

func next(rolls []int, k int, currSeq []int) int {
	if !possible(currSeq, rolls) {
		return len(currSeq)
	}
}

func possible(currSeq, rolls []int) bool {
	if len(currSeq) > len(rolls) {
		return false
	}

	var i1, i2 int

	for i1 < len(currSeq) && i2 < len(rolls) {
		if currSeq[i1] == rolls[i2] {

		}
	}
}
