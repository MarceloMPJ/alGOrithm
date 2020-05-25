package main

import "fmt"

const alphabetSize = 26

type trieNode struct {
	children    [alphabetSize]*trieNode
	isEndOfWord bool
	point       int
}

func getNode(point int) *trieNode {
	pNode := trieNode{isEndOfWord: false, point: point}
	for i := 0; i < alphabetSize; i++ {
		pNode.children[i] = nil
	}
	return &pNode
}

func insert(root *trieNode, key []rune, point int) {
	pCrawl := root

	for i := 0; i < len(key); i++ {
		index := int(key[i] - 'a')
		if pCrawl.children[index] == nil {
			pCrawl.children[index] = getNode(point)
		}

		if pCrawl.children[index].point < point {
			pCrawl.children[index].point = point
		}

		pCrawl = pCrawl.children[index]
	}

	pCrawl.isEndOfWord = true
}

func search(root *trieNode, key []rune) int {
	pCrawl := root

	for i := 0; i < len(key); i++ {
		index := int(key[i] - 'a')
		if pCrawl.children[index] == nil {
			return -1
		}
		pCrawl = pCrawl.children[index]
	}

	return pCrawl.point
}

func main() {
	root := getNode(-1)

	var n, q int
	fmt.Scanf("%d %d", &n, &q)

	for i := 0; i < n; i++ {
		var str string
		var point int
		fmt.Scanf("%s %d", &str, &point)
		insert(root, []rune(str), point)
	}

	for i := 0; i < q; i++ {
		var str string
		fmt.Scanf("%s", &str)
		fmt.Println(search(root, []rune(str)))
	}
}
