# Algorithms and Data Structures in Go

A comprehensive collection of algorithms and data structures implemented in Go (Golang). This repository serves as both a learning resource and a reference implementation for common computer science concepts.

## Contents

### Data Structures
- **Cache**
  - LRU (Least Recently Used)
  - LFU (Least Frequently Used)
- **Trees**
  - AVL Tree
  - Binary Search Tree
  - B-Tree
  - Red-Black Tree
  - Trie
  - Segment Tree
  - Fenwick Tree
- **Lists and Queues**
  - Linked Lists (Singly, Doubly, Circular)
  - Queue
  - Stack
  - Deque
  - Dynamic Array
- **Hash Structures**
  - HashMap
  - Set

### Algorithms

#### Sorting
- Bubble Sort
- Quick Sort
- Merge Sort
- Heap Sort
- Insertion Sort
- Selection Sort
- Shell Sort
- Counting Sort
- Radix Sort
- Bucket Sort
- Tim Sort
- Pancake Sort
- Patience Sort
- Circle Sort
- Cocktail Sort
- Comb Sort
- Cycle Sort
- Exchange Sort
- Pigeonhole Sort
- Stooge Sort

#### Searching
- Binary Search
- Linear Search
- Jump Search
- Interpolation Search
- Ternary Search
- Select K-th Element

#### Graph Algorithms
- Depth-First Search (DFS)
- Breadth-First Search (BFS)
- Dijkstra's Algorithm
- Bellman-Ford Algorithm
- Floyd-Warshall Algorithm
- Kruskal's Algorithm
- Prim's Algorithm
- Topological Sort
- Kosaraju's Algorithm
- Edmonds-Karp Algorithm
- Graph Coloring Algorithms
- Articulation Points
- Union Find

#### Dynamic Programming
- Longest Common Subsequence
- Longest Increasing Subsequence
- Longest Palindromic Subsequence
- Knapsack Problem
- Matrix Chain Multiplication
- Rod Cutting
- Coin Change
- Edit Distance
- Word Break
- Burst Balloons

#### String Algorithms
- Knuth-Morris-Pratt (KMP)
- Aho-Corasick
- Boyer-Moore
- Rabin-Karp
- Levenshtein Distance
- Manacher's Algorithm
- String Matching
- Palindrome Check
- Pangram Check

#### Cryptography
- RSA
- Caesar Cipher
- ROT13
- XOR Cipher
- Diffie-Hellman Key Exchange
- Digital Signature Algorithm (DSA)
- Polybius Cipher
- Rail Fence Cipher

#### Mathematical Algorithms
- Greatest Common Divisor (GCD)
- Least Common Multiple (LCM)
- Prime Numbers
- Fibonacci Numbers
- Pascal's Triangle
- Matrix Operations
- Catalan Numbers
- Euler Totient
- Fast Exponentiation
- Binary Operations

#### Compression
- Huffman Coding
- Run-Length Encoding (RLE)

## Features
- Clean, well-documented implementations
- Comprehensive test coverage
- Standard Go coding conventions
- Generic implementations where applicable
- Performance-optimized algorithms
- Educational comments and documentation

## Requirements
- Go 1.18 or higher

## Installation
```bash
git clone https://github.com/deeper-x/data_struct_algs.git
cd data_struct_algs
go mod tidy
```

## Usage
Each algorithm and data structure is contained in its own package. Import and use them as needed:

```go
import (
    "github.com/deeper-x/data_struct_algs/sort"
    "github.com/deeper-x/data_struct_algs/search"
)

// Example usage of quicksort
numbers := []int{64, 34, 25, 12, 22, 11, 90}
sorted := sort.QuickSort(numbers)

// Example usage of binary search
index := search.Binary(sorted, 22)
```

## Testing
Run all tests:
```bash
go test ./...
```

Run specific package tests:
```bash
go test ./sort/...
go test ./search/...
```

## Contributing
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## Project Structure
```
├── cache/          # Caching implementations
├── cipher/         # Cryptography algorithms
├── compression/    # Data compression algorithms
├── conversion/     # Data conversion utilities
├── dynamic/        # Dynamic programming solutions
├── graph/          # Graph algorithms
├── math/          # Mathematical algorithms
├── search/        # Search algorithms
├── sort/          # Sorting algorithms
├── strings/       # String manipulation algorithms
└── structure/     # Data structure implementations
```

## License
This project is licensed under the MIT License - see the LICENSE file for details.

## Contact
For questions and feedback, please open an issue in the GitHub repository.