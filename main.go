package main

import (
	"bufio"
	"os"
	"strconv"
)

var sum int

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	sc.Scan()
	n := nextInt(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		a := nextInt(sc.Text())
		sum = 0
		recur(a, 0)
		wr.WriteString(strconv.Itoa(sum) + "\n")
	}
}

func recur(a, x int) {
	if x > a {
		return
	}
	if x == a {
		sum++
		return
	}
	for i := 1; i <= 3; i++ {
		recur(a, x+i)
	}
}

func nextInt(str string) int {
	ret, _ := strconv.Atoi(str)
	return ret
}
