package ciphers

import (
	"fmt"
)

var alphabets = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func inTable(table [5][5]byte, val byte) bool {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if table[i][j] == val {
				return true
			}
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getLoc(char byte, table [5][5]byte) (int, int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if table[i][j] == char {
				return i, j
			}
		}
	}
	return -1, -1
}

func generateTable(key string) [5][5]byte {
	table := [5][5]byte{
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	k := 0
	i, j := 0, 0
	for ; i < 5; i++ {
		for ; j < 5; j++ {
			for k < len(key) {
				if !inTable(table, key[k]) {
					table[i][j] = key[k]
					break
				}
				k++
			}
			if k == len(key) {
				break
			}
		}
		if j == 5 {
			j = 0
		}
		if k == len(key) {
			break
		}
	}
	for _, char := range alphabets {
		if i >= len(table) {
			break
		}
		if !inTable(table, char) {
			table[i][j] = char
			j++
			if j == 5 {
				j = 0
				i++
			}
		}
	}
	return table
}

func PlayfairEncrypt(s []byte, key string) []byte {
	table := generateTable(key)
	for i := 0; i < 5; i++ {
		fmt.Println(string(table[i][:]))
	}

	type pair [2]byte
	var pairs []pair

	for i := 0; i < len(s); {
		var p pair
		p[0] = s[i]

		if i == len(s)-1 {
			p[1] = 'X'
		} else {
			p[1] = s[i+1]
		}

		if p[0] == p[1] {
			p[1] = 'X'
			i += 1
		} else {
			i += 2
		}
		pairs = append(pairs, p)
	}
	for _, el := range pairs {
		fmt.Println(string(el[:]))
	}

	var result []byte
	for _, p := range pairs {
		i1, j1 := getLoc(p[0], table)
		i2, j2 := getLoc(p[1], table)
		if j1 == j2 {
			result = append(result, table[(i1+1)%5][j1])
			result = append(result, table[(i2+1)%5][j2])
		} else if i1 == i2 {
			result = append(result, table[i1][(j1+1)%5])
			result = append(result, table[i2][(j2+1)%5])
		} else {
			diff := abs(j1 - j2)
			if j1 < j2 {
				result = append(result, table[i1][j1+diff])
				result = append(result, table[i2][j2-diff])
			} else {
				result = append(result, table[i1][j1-diff])
				result = append(result, table[i2][j2+diff])
			}
		}
	}

	return result
}

func PlayfairDecrypt(s []byte, key string) []byte {
	table := generateTable(key)
	for i := 0; i < 5; i++ {
		fmt.Println(string(table[i][:]))
	}

	type pair [2]byte
	var pairs []pair

	for i := 0; i < len(s); i += 2 {
		var p pair
		p[0] = s[i]
		p[1] = s[i+1]
		pairs = append(pairs, p)
	}
	for _, el := range pairs {
		fmt.Println(string(el[:]))
	}

	var result []byte
	for _, p := range pairs {
		i1, j1 := getLoc(p[0], table)
		i2, j2 := getLoc(p[1], table)
		if j1 == j2 {
			i1 = i1 - 1
			if i1 < 0 {
				i1 = 4
			}
			i2 = i2 - 1
			if i2 < 0 {
				i2 = 4
			}
			result = append(result, table[i1][j1])
			result = append(result, table[i2][j2])
		} else if i1 == i2 {
			j1 = j1 - 1
			if j1 < 0 {
				j1 = 4
			}
			j2 = j2 - 1
			if j2 < 0 {
				j2 = 4
			}
			result = append(result, table[i1][j1])
			result = append(result, table[i2][j2])
		} else {
			diff := abs(j1 - j2)
			if j1 < j2 {
				result = append(result, table[i1][j1+diff])
				result = append(result, table[i2][j2-diff])
			} else {
				result = append(result, table[i1][j1-diff])
				result = append(result, table[i2][j2+diff])
			}
		}
	}

	return result
}
