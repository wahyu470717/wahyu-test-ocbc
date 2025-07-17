package main

import (
	"fmt"
	"math"
)

// Fungsi utama untuk menghitung tahun minimum Bob menyalip Alice
func minYearsToSurpass(aliceInitial, bobBonus float64, aliceRate, bobRate float64) int {
	// Kasus 1: Bob sudah lebih besar dari awal
	if bobBonus > aliceInitial {
		return 0
	}

	// Kasus 2: Tingkat bunga Bob <= Alice dan Bob mulai lebih kecil
	// Dalam hal ini, Bob tidak akan pernah bisa menyalip Alice
	if bobRate <= aliceRate {
		return -1
	}

	// Simulasi pertumbuhan investasi tahun per tahun
	alice := aliceInitial
	bob := bobBonus
	years := 0

	// Batas maksimum untuk mencegah infinite loop
	const maxYears = 1000

	for years < maxYears {
		// Pertumbuhan di tahun berikutnya
		alice = alice * (1 + aliceRate)
		bob = bob * (1 + bobRate)
		years++

		// Cek apakah Bob sudah menyalip Alice
		if bob > alice {
			return years
		}
	}

	// Jika tidak bisa menyalip dalam batas maksimum
	return -1
}

// Fungsi alternatif menggunakan formula matematika
func minYearsToSurpassMath(aliceInitial, bobBonus float64, aliceRate, bobRate float64) int {
	// Kasus 1: Bob sudah lebih besar dari awal
	if bobBonus > aliceInitial {
		return 0
	}

	// Kasus 2: Tingkat bunga Bob <= Alice
	if bobRate <= aliceRate {
		return -1
	}

	// Formula: bobBonus * (1 + bobRate)^n > aliceInitial * (1 + aliceRate)^n
	// Rearrange: (bobBonus / aliceInitial) > ((1 + aliceRate) / (1 + bobRate))^n
	// Take log: log(bobBonus / aliceInitial) > n * log((1 + aliceRate) / (1 + bobRate))
	// Solve for n: n > log(bobBonus / aliceInitial) / log((1 + aliceRate) / (1 + bobRate))

	ratio := bobBonus / aliceInitial
	growthRatio := (1 + aliceRate) / (1 + bobRate)

	if ratio >= 1 {
		return 0
	}

	// Hitung n menggunakan logaritma
	n := math.Log(ratio) / math.Log(growthRatio)

	// Karena kita butuh Bob > Alice (strictly greater), ambil ceiling
	years := int(math.Ceil(-n))

	// Verifikasi hasil dengan simulasi
	alice := aliceInitial
	bob := bobBonus

	for i := 0; i < years; i++ {
		alice = alice * (1 + aliceRate)
		bob = bob * (1 + bobRate)
	}

	if bob > alice {
		return years
	}

	// Jika belum cukup, tambah 1 tahun
	return years + 1
}

// Fungsi untuk menampilkan proses simulasi tahun per tahun
func showInvestmentProgress(aliceInitial, bobBonus float64, aliceRate, bobRate float64, maxYears int) {
	fmt.Printf("=== SIMULASI PERTUMBUHAN INVESTASI ===\n")
	fmt.Printf("Alice Initial: $%.2f, Rate: %.2f%%\n", aliceInitial, aliceRate*100)
	fmt.Printf("Bob Bonus: $%.2f, Rate: %.2f%%\n", bobBonus, bobRate*100)
	fmt.Printf("=====================================\n")

	alice := aliceInitial
	bob := bobBonus

	fmt.Printf("Year 0: Alice = $%.2f, Bob = $%.2f\n", alice, bob)

	for year := 1; year <= maxYears; year++ {
		alice = alice * (1 + aliceRate)
		bob = bob * (1 + bobRate)

		fmt.Printf("Year %d: Alice = $%.2f, Bob = $%.2f", year, alice, bob)

		if bob > alice {
			fmt.Printf(" *** BOB MENYALIP ALICE! ***\n")
			break
		} else {
			fmt.Printf(" (Selisih: $%.2f)\n", alice-bob)
		}
	}
}

// Fungsi untuk test berbagai skenario
func testScenarios() {
	fmt.Println("========== TEST BERBAGAI SKENARIO ==========")

	scenarios := []struct {
		name         string
		aliceInitial float64
		bobBonus     float64
		aliceRate    float64
		bobRate      float64
	}{
		{"Contoh Soal", 1000, 800, 0.05, 0.08},
		{"Bob Rate Lebih Rendah", 1000, 800, 0.08, 0.05},
		{"Bob Sudah Menang", 1000, 1200, 0.05, 0.08},
		{"Rate Sama", 1000, 800, 0.05, 0.05},
		{"Perbedaan Kecil", 1000, 950, 0.05, 0.06},
		{"Perbedaan Besar", 1000, 500, 0.03, 0.15},
	}

	for _, scenario := range scenarios {
		fmt.Printf("\n--- %s ---\n", scenario.name)
		fmt.Printf("Alice: $%.0f @ %.1f%%, Bob: $%.0f @ %.1f%%\n",
			scenario.aliceInitial, scenario.aliceRate*100,
			scenario.bobBonus, scenario.bobRate*100)

		result1 := minYearsToSurpass(scenario.aliceInitial, scenario.bobBonus,
			scenario.aliceRate, scenario.bobRate)
		result2 := minYearsToSurpassMath(scenario.aliceInitial, scenario.bobBonus,
			scenario.aliceRate, scenario.bobRate)

		fmt.Printf("Hasil Simulasi: %d tahun\n", result1)
		fmt.Printf("Hasil Matematika: %d tahun\n", result2)

		if result1 == result2 {
			fmt.Printf("✓ Hasil konsisten\n")
		} else {
			fmt.Printf("⚠ Hasil berbeda!\n")
		}
	}
}

// Fungsi untuk analisis compound interest
func analyzeCompoundInterest(principal, rate float64, years int) {
	fmt.Printf("\n=== ANALISIS COMPOUND INTEREST ===\n")
	fmt.Printf("Principal: $%.2f\n", principal)
	fmt.Printf("Rate: %.2f%%\n", rate*100)
	fmt.Printf("==================================\n")

	current := principal
	for i := 1; i <= years; i++ {
		current = current * (1 + rate)
		fmt.Printf("Year %d: $%.2f\n", i, current)
	}

	totalGain := current - principal
	fmt.Printf("Total Gain: $%.2f (%.2f%%)\n", totalGain, (totalGain/principal)*100)
}

func main() {
	fmt.Println("========== INVESTMENT COMPETITION SOLVER ==========")

	// Test case dari soal
	aliceInitial := 1000.0
	bobBonus := 800.0
	aliceRate := 0.05
	bobRate := 0.08

	fmt.Printf("=== CONTOH SOAL ===\n")
	result := minYearsToSurpass(aliceInitial, bobBonus, aliceRate, bobRate)
	fmt.Printf("Minimum tahun untuk Bob menyalip Alice: %d tahun\n", result)

	// Tampilkan progress detail
	if result > 0 && result <= 20 {
		fmt.Printf("\n")
		showInvestmentProgress(aliceInitial, bobBonus, aliceRate, bobRate, result+2)
	}

	// Test berbagai skenario
	testScenarios()

	// Analisis compound interest
	fmt.Printf("\n")
	analyzeCompoundInterest(1000, 0.05, 5)
	analyzeCompoundInterest(800, 0.08, 5)

	// Test edge cases
	fmt.Printf("\n========== EDGE CASES ==========\n")

	// Test Bob sudah menang dari awal
	fmt.Printf("Bob sudah menang: %d\n", minYearsToSurpass(1000, 1200, 0.05, 0.08))

	// Test rate sama
	fmt.Printf("Rate sama: %d\n", minYearsToSurpass(1000, 800, 0.05, 0.05))

	// Test Bob rate lebih rendah
	fmt.Printf("Bob rate lebih rendah: %d\n", minYearsToSurpass(1000, 800, 0.08, 0.05))

	// Test dengan nilai yang sangat kecil
	fmt.Printf("Perbedaan rate kecil: %d\n", minYearsToSurpass(1000, 999, 0.05, 0.051))

	fmt.Println("\n========== PROGRAM SELESAI ==========")
}
