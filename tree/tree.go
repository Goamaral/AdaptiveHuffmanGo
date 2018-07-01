package tree

import "fmt"

// Node - node list
type Node struct {
	isActive bool
	child    *Node
	brother  *Node
	letter   rune
	isWord   bool
}

func CreateNode() *Node {
	root := new(Node)
	root.isActive = false
	root.child = nil
	root.brother = nil
	root.letter = 0
	root.isWord = false

	return root
}

func (root *Node) AddWord(word string) {
	node := root
	save := root
	for _, letter := range word {
		node = node.addLetter(letter)
		save = node
		node = node.child
	}

	save.isWord = true
}

func (node *Node) ActivateNode(letter rune) {
	node.isActive = true
	node.letter = letter
	node.brother = CreateNode()
	node.child = CreateNode()
}

func (node *Node) addLetter(letter rune) *Node {
	if !node.isActive {
		node.ActivateNode(letter)
		return node
	}

	for node.isActive {
		if node.letter == letter {
			return node
		}

		node = node.brother
	}

	node.ActivateNode(letter)

	return node
}

func (node *Node) printTree() {
	node.printTreeAux("")
}

func (node *Node) printTreeAux(word string) {
	for node.isActive {
		newWord := word + string(node.letter)

		if node.isWord {
			fmt.Println(newWord)
		} else {
			if node.child != nil {
				node.child.printTreeAux(newWord)
			}
		}

		node = node.brother
	}
}

func (node *Node) GetSuggestions(word string) []string {
	curr := node
	for _, letter := range word {
		for curr.isActive && curr.letter != letter {
			curr = curr.brother
		}

		if !curr.isActive {
			return nil
		}

		curr = curr.child
	}

	var result []string

	return curr.generateSuggestion(word, "", result)
}

func (node *Node) searchLetter(letter rune) *Node {
	if !node.isActive {
		return node
	} else {
		if node.letter == letter {
			return node
		} else {
			return node.brother.searchLetter(letter)
		}
	}
}

func (node *Node) generateSuggestion(prefix string, suffix string, result []string) []string {
	if !node.isActive {
		return result
	}

	newSuffix := suffix + string(node.letter)

	if node.isWord {
		result = append(result, prefix+newSuffix)
	}

	result = node.child.generateSuggestion(prefix, newSuffix, result)

	return node.brother.generateSuggestion(prefix, suffix, result)
}
