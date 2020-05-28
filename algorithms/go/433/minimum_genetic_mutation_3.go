package main

func minMutation3(start string, end string, bank []string) int {
	if idxOf(end, bank) == -1 {
		return -1
	}
	if start == end {
		return 0
	}
	count := 0
	isUsed := make([]bool, len(bank))

	startQueue := []string{start}
	endQueue := []string{end}
	for len(startQueue) > 0 {
		count++
		l := len(startQueue)
		for _, curr := range startQueue {
			for i := 0; i < len(curr); i++ {
				for _, c := range mutationMap[curr[i]] {
					gene := curr[:i] + c + curr[i+1:]
					if idx := idxOf(gene, endQueue); idx != -1 {
						return count
					}
					if idx := idxOf(gene, bank); idx != -1 && !isUsed[idx] {
						isUsed[idx] = true
						startQueue = append(startQueue, gene)
					}
				}
			}
		}
		startQueue = startQueue[l:]
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
		}
	}
	return -1
}
