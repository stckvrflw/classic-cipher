package ciphers

func generateMatrixKey(key string, n int) [][]byte {
	k := 0
	var matrixKey [][]byte
	for i := 0; i < n; i++ {
		var row []byte
		for j := 0; j < n; j++ {
			row = append(row, key[k]%65)
			k++
		}
		matrixKey = append(matrixKey, row)
	}
	return matrixKey
}

func pad(s []byte, n int) []byte {
	return nil
}

func MulVectorMatrix(vector []byte, matrix [][]byte) []int {
	resultVector := make([]int, len(vector))
	for i := range matrix {
		sum := 0
		for j := range vector {
			sum += int(matrix[i][j]) * int(vector[j])
		}
		resultVector[i] = sum
	}
	return resultVector
}

func HillEncrypt(s []byte, key string, partition int) []byte {
	matrixKey := generateMatrixKey(key, partition)
	leftover := len(s) % partition
	if leftover != 0 {
		s = pad(s, leftover)
	}
	var result []byte
	for i := 0; i < len(s); i += partition {
		var vector []byte
		// vector := s[i : i+partition]
		vector = append(vector, s[i:i+partition]...)
		for j := 0; j < len(vector); j++ {
			// newVal := vector[j] - 65
			vector[j] %= 65

		}
		encryptVector := MulVectorMatrix([]byte(vector), matrixKey)
		for j := 0; j < len(vector); j++ {
			// result = append(result, byte(encryptVector[j]%26))
			result = append(result, byte(encryptVector[j]%26+'A'))

		}
	}

	return result
}
