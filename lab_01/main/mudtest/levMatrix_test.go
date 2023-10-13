package test

import (
	"testing"
	"../algorithms"
	//"fmt"
)

//простой тест
func BenchmarkSampleMatrixLevenshtein1(b *testing.B) {
    r1 := []rune("лабрадор")
	r2 := []rune("гибралтар")

	//fmt.Println("s1 = лабрадор, s2 = гибралтар")
	dist := algorithms.MatrixLevenshtein(r1, r2)
	
    if dist != 5 {
        b.Error("Expected 5, got ", dist)
    }
}

func BenchmarkSampleMatrixDamerauLevenshtein1(b *testing.B) {
    r1 := []rune("лабрадор")
	r2 := []rune("гибралтар")

	dist := algorithms.MatrixDamerauLevenshtein(r1, r2)

	if dist != 5 {
		b.Error("Expected 5, got ", dist)
	}
}

func BenchmarkSampleRecursiveDL1(b *testing.B) {
    r1 := []rune("лабрадор")
	r2 := []rune("гибралтар")

	dist := algorithms.RecursiveDL(r1, r2)

	if dist != 5 {
		b.Error("Expected 5, got ", dist)
	}
}

func BenchmarkSampleRecursiveDLCash1(b *testing.B) {
    r1 := []rune("лабрадор")
	r2 := []rune("гибралтар")

	dist := algorithms.RecursiveDLCash(r1, r2)

	if dist != 5 {
		b.Error("Expected 5, got ", dist)
	}
}

//ошибка 10(10, 10)
func BenchmarkSampleMatrixLevenshtein2(b *testing.B) {
    r1 := []rune("fhdjtiamhs")
	r2 := []rune("1234567890")

	dist := algorithms.MatrixLevenshtein(r1, r2)
	
    if dist != 10 {
        b.Error("Expected 5, got ", dist)
    }
}


func BenchmarkSampleMatrixDamerauLevenshtein2(b *testing.B) {
    r1 := []rune("fhdjtiamhs")
	r2 := []rune("1234567890")

	dist := algorithms.MatrixDamerauLevenshtein(r1, r2)

	if dist != 10 {
		b.Error("Expected 5, got ", dist)
	}
}

func BenchmarkSampleRecursiveDL2(b *testing.B) {
    r1 := []rune("fhdjtiamhs")
	r2 := []rune("1234567890")

	dist := algorithms.RecursiveDL(r1, r2)

	if dist != 10 {
		b.Error("Expected 5, got ", dist)
	}
}

func BenchmarkSampleRecursiveDLCash2(b *testing.B) {
    r1 := []rune("fhdjtiamhs")
	r2 := []rune("1234567890")

	dist := algorithms.RecursiveDLCash(r1, r2)

	if dist != 10 {
		b.Error("Expected 5, got ", dist)
	}
}

//ошибка 20(20, 20)
func BenchmarkSampleMatrixLevenshtein3(b *testing.B) {
    r1 := []rune("fhdjtiamhsfhdjtiamhs")
	r2 := []rune("12345678901234567890")

	dist := algorithms.MatrixLevenshtein(r1, r2)
	
    if dist != 20 {
        b.Error("Expected 5, got ", dist)
    }
}


func BenchmarkSampleMatrixDamerauLevenshtein3(b *testing.B) {
    r1 := []rune("flsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsnflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsnflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsn")
	r2 := []rune("123")

	dist := algorithms.MatrixDamerauLevenshtein(r1, r2)

	if dist != 300 {
		b.Error("Expected 5, got ", dist)
	}
}

func BenchmarkSampleRecursiveDL3(b *testing.B) {
    r1 := []rune("flsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsnflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsnflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsn")
	r2 := []rune("123")

	dist := algorithms.RecursiveDL(r1, r2)

	if dist != 300 {
		b.Error("Expected 5, got ", dist)
	}
}

func BenchmarkSampleRecursiveDLCash3(b *testing.B) {
    r1 := []rune("flsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsnflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsnflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsngjsifogmflsn")
	r2 := []rune("123")

	dist := algorithms.RecursiveDLCash(r1, r2)

	if dist != 300 {
		b.Error("Expected 5, got ", dist)
	}
}

//ошибка 30(30, 30)
func BenchmarkSampleMatrixLevenshtein4(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhs")
	r2 := []rune("123456789012345678901234567890")

	dist := algorithms.MatrixLevenshtein(r1, r2)
	
    if dist != 30 {
        b.Error("Expected 5, got ", dist)
    }
}


func BenchmarkSampleMatrixDamerauLevenshtein4(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhs")
	r2 := []rune("123456789012345678901234567890")

	dist := algorithms.MatrixDamerauLevenshtein(r1, r2)

	if dist != 30 {
		b.Error("Expected 5, got ", dist)
	}
}

// func BenchmarkSampleRecursiveDL4(b *testing.B) {
//     r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhs")
// 	r2 := []rune("123456789012345678901234567890")
// 	dist := algorithms.RecursiveDL(r1, r2)

// 	if dist != 30 {
// 		b.Error("Expected 5, got ", dist)
// 	}
// }

func BenchmarkSampleRecursiveDLCash4(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhs")
	r2 := []rune("123456789012345678901234567890")
	dist := algorithms.RecursiveDLCash(r1, r2)

	if dist != 30 {
		b.Error("Expected 5, got ", dist)
	}
}

//ошибка 40(40, 40)
func BenchmarkSampleMatrixLevenshtein5(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhs")
	r2 := []rune("1234567890123456789012345678901234567890")

	dist := algorithms.MatrixLevenshtein(r1, r2)
	
    if dist != 40 {
        b.Error("Expected 5, got ", dist)
    }
}

func BenchmarkSampleMatrixDamerauLevenshtein5(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhs")
	r2 := []rune("1234567890123456789012345678901234567890")

	dist := algorithms.MatrixDamerauLevenshtein(r1, r2)

	if dist != 40 {
		b.Error("Expected 5, got ", dist)
	}
}

// func BenchmarkSampleRecursiveDL5(b *testing.B) {
//     r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhs")
// 	r2 := []rune("1234567890123456789012345678901234567890")
// 	dist := algorithms.RecursiveDL(r1, r2)

// 	if dist != 40 {
// 		b.Error("Expected 5, got ", dist)
// 	}
// }

func BenchmarkSampleRecursiveDLCash5(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhs")
	r2 := []rune("1234567890123456789012345678901234567890")
	dist := algorithms.RecursiveDLCash(r1, r2)

	if dist != 40 {
		b.Error("Expected 5, got ", dist)
	}
}

//ошибка 50(50, 50)
func BenchmarkSampleMatrixLevenshtein6(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamh")
	r2 := []rune("12345678901234567890123456789012345678901234567890")

	dist := algorithms.MatrixLevenshtein(r1, r2)
	
    if dist != 50 {
        b.Error("Expected 5, got ", dist)
    }
}

func BenchmarkSampleMatrixDamerauLevenshtein6(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamh")
	r2 := []rune("12345678901234567890123456789012345678901234567890")

	dist := algorithms.MatrixDamerauLevenshtein(r1, r2)

	if dist != 50 {
		b.Error("Expected 5, got ", dist)
	}
}

// func BenchmarkSampleRecursiveDL6(b *testing.B) {
// 	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamh")
// 	r2 := []rune("12345678901234567890123456789012345678901234567890")
// 	dist := algorithms.RecursiveDL(r1, r2)

// 	if dist != 50 {
// 		b.Error("Expected 5, got ", dist)
// 	}
// }

func BenchmarkSampleRecursiveDLCash6(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamh")
	r2 := []rune("12345678901234567890123456789012345678901234567890")
	dist := algorithms.RecursiveDLCash(r1, r2)

	if dist != 50 {
		b.Error("Expected 5, got ", dist)
	}
}

//ошибка 100(100, 100)
func BenchmarkSampleMatrixLevenshtein7(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamh")
	r2 := []rune("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")

	dist := algorithms.MatrixLevenshtein(r1, r2)
	
    if dist != 100 {
        b.Error("Expected 5, got ", dist)
    }
}

func BenchmarkSampleMatrixDamerauLevenshtein7(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamh")
	r2 := []rune("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")

	dist := algorithms.MatrixDamerauLevenshtein(r1, r2)

	if dist != 100 {
		b.Error("Expected 5, got ", dist)
	}
}

func BenchmarkSampleRecursiveDLCash7(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamh")
	r2 := []rune("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
	dist := algorithms.RecursiveDLCash(r1, r2)

	if dist != 100 {
		b.Error("Expected 5, got ", dist)
	}
}


//ошибка 200(200, 200)
func BenchmarkSampleMatrixLevenshtein8(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh")
	r2 := []rune("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")

	dist := algorithms.MatrixLevenshtein(r1, r2)
	
    if dist != 200 {
        b.Error("Expected 5, got ", dist)
    }
}

func BenchmarkSampleMatrixDamerauLevenshtein8(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh")
	r2 := []rune("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")

	dist := algorithms.MatrixDamerauLevenshtein(r1, r2)

	if dist != 200 {
		b.Error("Expected 5, got ", dist)
	}
}

func BenchmarkSampleRecursiveDLCash8(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh")
	r2 := []rune("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")

	dist := algorithms.RecursiveDLCash(r1, r2)

	if dist != 200 {
		b.Error("Expected 5, got ", dist)
	}
}

//ошибка 400(400, 400)
func BenchmarkSampleMatrixLevenshtein9(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh" + 
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" + 
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh")

	r2 := []rune("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")

	dist := algorithms.MatrixLevenshtein(r1, r2)
	
    if dist != 400 {
        b.Error("Expected 5, got ", dist)
    }
}

func BenchmarkSampleMatrixDamerauLevenshtein9(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh" + 
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" + 
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh")

	r2 := []rune("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")

	dist := algorithms.MatrixDamerauLevenshtein(r1, r2)

	if dist != 400 {
		b.Error("Expected 5, got ", dist)
	}
}

func BenchmarkSampleRecursiveDLCash9(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh" + 
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" + 
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh")

	r2 := []rune("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")

	dist := algorithms.RecursiveDLCash(r1, r2)

	if dist != 400 {
		b.Error("Expected 5, got ", dist)
	}
}

//ошибка 800(800, 800)
func BenchmarkSampleMatrixLevenshtein10(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh")

	r2 := []rune("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")

	dist := algorithms.MatrixLevenshtein(r1, r2)
	
    if dist != 800 {
        b.Error("Expected 5, got ", dist)
    }
}

func BenchmarkSampleMatrixDamerauLevenshtein10(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh")

	r2 := []rune("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")

	dist := algorithms.MatrixDamerauLevenshtein(r1, r2)

	if dist != 800 {
		b.Error("Expected 5, got ", dist)
	}
}

func BenchmarkSampleRecursiveDLCash10(b *testing.B) {
	r1 := []rune("fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiam" +
	"fhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhfhdjtiamhsfhdjtiamhsfhdjtiamhsfhdjtiamhssfhdjtiamhh")

	r2 := []rune("1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890" +
	"1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")

	dist := algorithms.RecursiveDLCash(r1, r2)

	if dist != 800 {
		b.Error("Expected 5, got ", dist)
	}
}
