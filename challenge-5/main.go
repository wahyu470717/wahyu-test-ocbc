package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Struktur untuk menyimpan state dalam memoization
type State struct {
	arr string // representasi array sebagai string
	k   int    // jumlah bom yang tersisa
}

// Global map untuk memoization
var memo map[State]int64

// Fungsi utama untuk menghitung skor demolisi maksimal
func getDemolitionScore(arr []int, k int) int64 {
	// Inisialisasi memoization
	memo = make(map[State]int64)

	// Panggil fungsi rekursif
	return solve(arr, k)
}

// Fungsi rekursif untuk menyelesaikan masalah
func solve(arr []int, k int) int64 {
	// Base case: jika tidak ada bom lagi atau array kosong
	if k == 0 || len(arr) == 0 {
		return 0
	}

	// Buat state untuk memoization
	state := State{
		arr: arrayToString(arr),
		k:   k,
	}

	// Cek apakah sudah ada di memo
	if val, exists := memo[state]; exists {
		return val
	}

	var maxScore int64 = 0

	// Coba semua posisi untuk demolisi
	for i := 0; i < len(arr); i++ {
		// Hitung skor dari demolisi posisi i
		currentScore := int64(arr[i])

		// Buat partisi kiri dan kanan
		leftPartition := make([]int, i)
		copy(leftPartition, arr[:i])

		rightPartition := make([]int, len(arr)-i-1)
		copy(rightPartition, arr[i+1:])

		// Tentukan partisi mana yang akan dibuang
		var remainingPartition []int

		if len(leftPartition) < len(rightPartition) {
			// Buang partisi kiri, gunakan partisi kanan
			remainingPartition = rightPartition
		} else if len(leftPartition) > len(rightPartition) {
			// Buang partisi kanan, gunakan partisi kiri
			remainingPartition = leftPartition
		} else {
			// Ukuran sama, buang partisi kiri
			remainingPartition = rightPartition
		}

		// Lemahkan partisi yang tersisa
		weakenedPartition := make([]int, len(remainingPartition))
		for j := 0; j < len(remainingPartition); j++ {
			weakenedPartition[j] = max(0, remainingPartition[j]-arr[i])
		}

		// Rekursif untuk langkah berikutnya
		futureScore := solve(weakenedPartition, k-1)

		// Update skor maksimal
		totalScore := currentScore + futureScore
		if totalScore > maxScore {
			maxScore = totalScore
		}
	}

	// Simpan hasil ke memo
	memo[state] = maxScore
	return maxScore
}

// Fungsi helper untuk mengkonversi array ke string (untuk memoization)
func arrayToString(arr []int) string {
	var parts []string
	for _, v := range arr {
		parts = append(parts, strconv.Itoa(v))
	}
	return strings.Join(parts, ",")
}

// Fungsi helper untuk mendapatkan nilai maksimal
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Fungsi untuk menampilkan langkah-langkah demolisi (untuk debugging)
func showDemolitionSteps(arr []int, k int) {
	fmt.Printf("Array awal: %v, k = %d\n", arr, k)

	original := make([]int, len(arr))
	copy(original, arr)

	current := arr
	totalScore := int64(0)

	for move := 1; move <= k && len(current) > 0; move++ {
		bestScore := int64(-1)
		bestIndex := -1
		var bestNext []int

		// Cari langkah terbaik
		for i := 0; i < len(current); i++ {
			// Simulasi demolisi di posisi i
			leftPartition := make([]int, i)
			copy(leftPartition, current[:i])

			rightPartition := make([]int, len(current)-i-1)
			copy(rightPartition, current[i+1:])

			var remainingPartition []int
			if len(leftPartition) < len(rightPartition) {
				remainingPartition = rightPartition
			} else if len(leftPartition) > len(rightPartition) {
				remainingPartition = leftPartition
			} else {
				remainingPartition = rightPartition
			}

			weakenedPartition := make([]int, len(remainingPartition))
			for j := 0; j < len(remainingPartition); j++ {
				weakenedPartition[j] = max(0, remainingPartition[j]-current[i])
			}

			// Hitung skor future
			futureScore := solve(weakenedPartition, k-move)
			totalCurrentScore := int64(current[i]) + futureScore

			if totalCurrentScore > bestScore {
				bestScore = totalCurrentScore
				bestIndex = i
				bestNext = weakenedPartition
			}
		}

		if bestIndex != -1 {
			fmt.Printf("Langkah %d: Pilih posisi %d (nilai %d)\n", move, bestIndex, current[bestIndex])
			totalScore += int64(current[bestIndex])
			fmt.Printf("  Skor sekarang: %d\n", totalScore)

			if len(bestNext) > 0 {
				fmt.Printf("  Array baru: %v\n", bestNext)
			} else {
				fmt.Printf("  Array kosong\n")
			}

			current = bestNext
		}
	}

	fmt.Printf("Skor akhir: %d\n\n", totalScore)
}

func main() {
	// Test case dari soal
	arr1 := []int{10, 2, 8, 5}
	k1 := 2

	fmt.Println("=== Test Case 1 ===")
	result1 := getDemolitionScore(arr1, k1)
	fmt.Printf("Hasil: %d\n", result1)
	showDemolitionSteps(arr1, k1)

	// Test case tambahan
	arr2 := []int{5, 3, 7, 1, 9}
	k2 := 3

	fmt.Println("=== Test Case 2 ===")
	result2 := getDemolitionScore(arr2, k2)
	fmt.Printf("Hasil: %d\n", result2)
	showDemolitionSteps(arr2, k2)

	// Test case kecil
	arr3 := []int{1, 2, 3}
	k3 := 2

	fmt.Println("=== Test Case 3 ===")
	result3 := getDemolitionScore(arr3, k3)
	fmt.Printf("Hasil: %d\n", result3)
	showDemolitionSteps(arr3, k3)
}
