# Huffman Coding

The Huffman Coding problem is a popular encoding algorithm used for data compression. The goal is to encode a given set of characters in such a way that the resulting binary code is as short as possible while maintaining the ability to decode it accurately. This is achieved by assigning a variable-length code to each character, with the code length inversely proportional to the character frequency.

Here is an example to illustrate the Huffman Coding problem:

Suppose we want to encode the message "ABBCCCDDDDEEEEE" using Huffman Coding. We first compute the frequency of each character:

| Character | Frequency |
| --- | --- |
| A   | 1   |
| B   | 2   |
| C   | 3   |
| D   | 4   |
| E   | 5   |

We then construct a binary tree as follows:

1.  Create a node for each character, with the weight equal to its frequency.
2.  Select the two nodes with the lowest weight, and create a new parent node with weight equal to the sum of the children weights. Assign the new node as the parent of the two selected nodes.
3.  Repeat step 2 until there is only one node left, which becomes the root of the tree.

The resulting binary tree for our example looks like this:

```
            15
           /  \
          7    \
        /  \    \
       3    4    \
      / \  / \    \
     A  B C   D    E
```

To encode each character, we traverse the binary tree from the root to the leaf corresponding to the character, assigning a "0" for each left branch and a "1" for each right branch. The resulting codes are:

| Character | Frequency | Code |
| --- | --- | --- |
| A   | 1   | 000 |
| B   | 2   | 001 |
| C   | 3   | 01  |
| D   | 4   | 10  |
| E   | 5   | 11  |

Therefore, the encoded message is "00100011101010111111", which is shorter than the original message.

In general, the candidate solutions for the Huffman Coding problem are all possible binary trees that can be constructed from the given character frequencies. The optimal solution is the one that minimizes the total length of the encoded message.

```go
type Node struct {
    Char  byte
    Freq  int
    Left  *Node
    Right *Node
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
    return pq[i].Freq < pq[j].Freq
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(*Node)
    *pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    *pq = old[0 : n-1]
    return item
}

func BuildHuffmanTree(freq map[byte]int) *Node {
    pq := make(PriorityQueue, len(freq))
    i := 0
    for char, f := range freq {
        pq[i] = &Node{Char: char, Freq: f}
        i++
    }
    heap.Init(&pq)

    for pq.Len() > 1 {
        left := heap.Pop(&pq).(*Node)
        right := heap.Pop(&pq).(*Node)
        parent := &Node{
            Freq:  left.Freq + right.Freq,
            Left:  left,
            Right: right,
        }
        heap.Push(&pq, parent)
    }

    return heap.Pop(&pq).(*Node)
}

func BuildCodeTable(root *Node) map[byte]string {
    codeTable := make(map[byte]string)
    var traverse func(*Node, string)
    traverse = func(node *Node, path string) {
        if node.Left == nil && node.Right == nil {
            codeTable[node.Char] = path
            return
        }
        traverse(node.Left, path+"0")
        traverse(node.Right, path+"1")
    }
    traverse(root, "")
    return codeTable
}

func EncodeMessage(msg string, codeTable map[byte]string) string {
    encodedMsg := ""
    for i := 0; i < len(msg); i++ {
        code, ok := codeTable[msg[i]]
        if !ok {
            panic(fmt.Sprintf("Character %c not found in code table", msg[i]))
        }
        encodedMsg += code
    }
    return encodedMsg
}

func main() {
    msg := "ABBCCCDDDDEEEEE"
    freq := make(map[byte]int)
    for i := 0; i < len(msg); i++ {
        freq[msg[i]]++
    }
    root := BuildHuffmanTree(freq)
    codeTable := BuildCodeTable(root)
    encodedMsg := EncodeMessage(msg, codeTable)
    fmt.Printf("Encoded message: %s\n", encodedMsg)
    fmt.Printf("Code table: %v\n", codeTable)
}
```
This implementation first builds a priority queue of nodes, with each node representing a character and its frequency. It then repeatedly extracts the two nodes with the lowest frequency from the queue, and creates a new parent node with weight equal to the sum of the children weights, until only one node is left, which becomes the root of the tree.

The code then traverses the tree recursively to build a table of codes for each character, and finally uses the code table to encode the message by replacing each character with its corresponding code. The resulting encoded message and code table are printed to the console.

Note that this implementation assumes that the input message

The greedy property applies in the Huffman Coding problem in the following way:

At each step of the algorithm, we merge the two smallest frequency nodes to form a new internal node. This merging operation is greedy, as we always select the two smallest frequency nodes to merge.

The optimal solution to the Huffman Coding problem is achieved by using a binary tree with the smallest total weighted path length. The greedy merging operation of the two smallest frequency nodes ensures that we always construct a binary tree that has a small total weighted path length.

Therefore, the greedy approach to the Huffman Coding problem guarantees an optimal solution.
