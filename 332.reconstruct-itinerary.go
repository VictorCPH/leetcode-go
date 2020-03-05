package main

import (
	"fmt"
	"sort"
)

func main() {
	tickets := [][]string{
		[]string{"MUC", "LHR"},
		[]string{"JFK", "MUC"},
		[]string{"SFO", "SJC"},
		[]string{"LHR", "SFO"},
	}
	fmt.Println(findItinerary(tickets))

	tickets1 := [][]string{
		[]string{"JFK", "SFO"},
		[]string{"JFK", "ATL"},
		[]string{"SFO", "ATL"},
		[]string{"ATL", "JFK"},
		[]string{"ATL", "SFO"},
	}
	fmt.Println(findItinerary(tickets1))

	tickets2 := [][]string{
		[]string{"JFK", "KUL"},
		[]string{"JFK", "NRT"},
		[]string{"NRT", "JFK"},
	}
	fmt.Println(findItinerary(tickets2))
}

func findItinerary(tickets [][]string) []string {
	m := make(map[string][]string, len(tickets)+1)
	var ans []string

	for _, t := range tickets {
		m[t[0]] = append(m[t[0]], t[1])
	}

	for k := range m {
		sort.Strings(m[k])
	}

	DFS("JFK", m, &ans)
	return ans
}
func DFS(start string, m map[string][]string, ans *[]string) {
	for len(m[start]) > 0 {
		cur := m[start][0]
		m[start] = m[start][1:]
		DFS(cur, m, ans)
	}

	*ans = append([]string{start}, *ans...)
}
