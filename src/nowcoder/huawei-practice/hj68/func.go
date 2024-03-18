package main

import (
	"fmt"
	"sort"
)

type Student struct {
	score int
	name  string
}

func main() {
	var n, flag int
	fmt.Scan(&n)
	fmt.Scan(&flag)
	var (
		name  string
		score int
	)
	students := make([]Student, 0, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&name)
		fmt.Scan(&score)
		students = append(students, Student{
			name:  name,
			score: score,
		})
	}
	sort.SliceStable(students, func(i, j int) bool {
		if flag == 0 {
			return students[i].score > students[j].score
		}
		return students[i].score < students[j].score
	})
	// 打印
	for _, student := range students {
		fmt.Printf("%s %d\n", student.name, student.score)
	}
}
