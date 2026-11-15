package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Yahar4/internal/graph"
)

func main() {
	var g *graph.Graph

	if len(os.Args) > 1 {
		path := os.Args[1]
		var err error
		g, err = loadFromFile(path)
		if err != nil {
			fmt.Println("Ошибка:", err)
			os.Exit(1)
		}
		fmt.Println("Загружен граф из", path)
	} else {
		g = readFromStdin()
	}

	analyze(g)
}

func analyze(g *graph.Graph) {
	if g.Vertices() == 0 {
		fmt.Println("Граф пуст (нет вершин).")
		return
	}

	components := g.Components()

	if len(components) == 1 {
		fmt.Println("Граф связен")
		return
	}

	fmt.Printf("Граф несвязен. Компонент связности: %d\n", len(components))
	for i, c := range components {
		fmt.Printf("Компонента %d: %v\n", i+1, c)
	}
}

func loadFromFile(path string) (*graph.Graph, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	var n int
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		n, err = strconv.Atoi(line)
		if err != nil || n < 0 {
			return nil, fmt.Errorf("некорректное n: %q", line)
		}
		break
	}

	g := graph.New(n)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		colon := strings.IndexByte(line, ':')
		if colon < 0 {
			return nil, fmt.Errorf("нет ':' в строке %q", line)
		}
		from, err := strconv.Atoi(strings.TrimSpace(line[:colon]))
		if err != nil || from < 0 || from >= n {
			return nil, fmt.Errorf("некорректный номер вершины %q", line[:colon])
		}
		sep := func(r rune) bool { return r == ' ' || r == ',' || r == '\t' }
		for _, s := range strings.FieldsFunc(line[colon+1:], sep) {
			to, err := strconv.Atoi(s)
			if err != nil || to < 0 || to >= n {
				return nil, fmt.Errorf("некорректный сосед %q у вершины %d", s, from)
			}
			g.AddEdge(from, to)
		}
	}
	return g, sc.Err()
}

func readFromStdin() *graph.Graph {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Количество вершин: ")
	var n int
	if _, err := fmt.Fscan(r, &n); err != nil || n < 0 {
		fmt.Println("Ошибка ввода n")
		os.Exit(1)
	}

	g := graph.New(n)
	r.ReadString('\n')

	fmt.Printf("Введите %d строк(и) вида \"i: сосед1 сосед2 ...\"\n", n)
	for i := 0; i < n; {
		fmt.Printf("%d> ", i)
		line, _ := r.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		colon := strings.IndexByte(line, ':')
		if colon < 0 {
			fmt.Println("  Ошибка: нет ':'. Повторите.")
			continue
		}
		from, err := strconv.Atoi(strings.TrimSpace(line[:colon]))
		if err != nil || from < 0 || from >= n {
			fmt.Println("  Ошибка: некорректный номер вершины.")
			continue
		}
		sep := func(r rune) bool { return r == ' ' || r == ',' || r == '\t' }
		for _, s := range strings.FieldsFunc(line[colon+1:], sep) {
			to, err := strconv.Atoi(s)
			if err != nil || to < 0 || to >= n {
				fmt.Println("  Пропущен некорректный сосед:", s)
				continue
			}
			g.AddEdge(from, to)
		}
		i++
	}
	return g
}
