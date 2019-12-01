package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type node struct {
	label         rune
	childCount    int
	metadataCount int
	children      []*node
	metadata      []int
}

var metadataSum int = 0

func main() {
	root := loadTreeFromFile("input.txt")
	fmt.Printf("tree loaded: %s\n", strconv.QuoteRune(root.label))
	fmt.Printf("global metadata sum = %d\n", metadataSum)

	fmt.Println()
}

func loadTreeFromFile(filepath string) *node {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	inputList := strings.Split(scanner.Text(), " ")
	index := new(int)

	return parseNode('A', inputList, index)
}

func parseNode(label rune, inputList []string, index *int) *node {
	childCount, _ := strconv.Atoi(inputList[*index])
	*index++
	metadataCount, _ := strconv.Atoi(inputList[*index])
	*index++

	currNode := &node{label, childCount, metadataCount, make([]*node, 0), make([]int, 0)}
	for i := 0; i < childCount; i++ {
		label++
		currNode.children = append(currNode.children, parseNode(label, inputList, index))
	}

	for i := 0; i < metadataCount; i++ {
		m, _ := strconv.Atoi(inputList[*index])
		*index++
		currNode.metadata = append(currNode.metadata, m)
		metadataSum += m
	}

	return currNode
}
