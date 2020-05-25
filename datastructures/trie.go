package datastructures

const alphabetSize = 26

type trieNode struct {
	children    [alphabetSize]*trieNode
	isEndOfWord bool
}

func getNode() *trieNode {
	pNode := trieNode{isEndOfWord: false}
	for i := 0; i < alphabetSize; i++ {
		pNode.children[i] = nil
	}
	return &pNode
}

// If not present, inserts key into trie
func insert(root *trieNode, key []rune) {
	pCrawl := root

	for i := 0; i < len(key); i++ {
		index := int(key[i] - 'a')
		if pCrawl.children[index] == nil {
			pCrawl.children[index] = getNode()
		}

		pCrawl = pCrawl.children[index]
	}

	pCrawl.isEndOfWord = true
}

// Returns true if key presents in trie, else false
func search(root *trieNode, key []rune) bool {
	pCrawl := root

	for i := 0; i < len(key); i++ {
		index := int(key[i] - 'a')
		if pCrawl.children[index] == nil {
			return false
		}
		pCrawl = pCrawl.children[index]
	}

	return pCrawl != nil && pCrawl.isEndOfWord
}
