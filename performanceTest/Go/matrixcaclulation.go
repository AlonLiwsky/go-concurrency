package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"
)

func multiply(A [][]int64, B [][]int64) [][]int64 {
	sizeA := len(A)
	sizeB := len(B)
	n := make([][]int64, sizeA)
	for i := range n {
		n[i] = make([]int64, sizeB)
	}
	for i := 0; i < sizeA; i++ {
		for k := 0; k < sizeB; k++ {
			temp := A[i][k]
			for j := 0; j < sizeB; j++ {
				n[i][j] += temp * B[k][j]
			}
		}
	}
	return n
}

func splitMatrix(nrOfThreads int, matrix [][]int64) (matrixes [][][]int64) {
	splitter := len(matrix) / nrOfThreads
	for i := 0; i < nrOfThreads; i++ {
		matrixes = append(matrixes,
			matrix[splitter*i:(splitter*(i+1))])
	}
	return
}

//Multiplies it's section of matrix A with matrix B and save the result in it's position of the result
func multiplyStuff(finalMatrix *[][][]int64, matrix1 [][]int64, matrix2 [][]int64, i int) {
	(*finalMatrix)[i] = multiply(matrix1, matrix2)
}

func readFile(filePath string) (matrix1 [][]int64, matrix2 [][]int64) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var temp []int64
	matrixNr := 1
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rows := strings.Fields(scanner.Text())
		if len(rows) != 0 {
			cells := strings.Split(rows[0], ",")
			for _, cell := range cells {
				i, err := strconv.ParseInt(cell, 10, 64)
				if err != nil {
					log.Fatal(err)
				}
				temp = append(temp, i)
			}
			if matrixNr == 1 {
				matrix1 = append(matrix1, temp)
			} else {
				matrix2 = append(matrix2, temp)
			}
			temp = nil
		} else {
			matrixNr = 2
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func createFile() (matrix1 [][]int, matrix2 [][]int) {
	largo := 100
	ancho := 100

	matrix1 = make([][]int, largo)
	matrix2 = make([][]int, largo)

	for i := 0; i < largo; i++ {
		matrix1 = make([][]int, ancho)
		matrix2 = make([][]int, ancho)
	}

	for i := 0; i < largo; i++ {
		for j := 0; j < ancho; j++ {
			matrix1[i] = append(matrix1[i], rand.Int())
			matrix2[i] = append(matrix2[i], rand.Int())
		}
	}

	return matrix1, matrix2
}

func main() {
	numCPUs:=runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)
	nrOfThreads :=
	debug.SetGCPercent(-1)
	if nrOfThreads <= 0 {
		runtime.GOMAXPROCS(1)
	} else if nrOfThreads >= 16 {
		runtime.GOMAXPROCS(8)
	} else {
		runtime.GOMAXPROCS(nrOfThreads)
	}


	finishedMatrix := make([][][]int64, nrOfThreads)

	//Load matrices from file
	matrix1, matrix2 := readFile("matricesToCalculate.csv")
	if len(matrix1) != len(matrix2) || (nrOfThreads != 0 &&
		len(matrix1)%nrOfThreads != 0) {
		log.Fatal("USAGE: " + os.Args[0] + " <file><nrOfThreads > ")
	}

	var start int64
	if nrOfThreads == 0 {
		//Run it sequentially
		start = time.Now().UnixNano()
		multiply(matrix1, matrix2)
	} else {
		//Split matrix for each thread
		matrixes := splitMatrix(nrOfThreads, matrix1)
		start = time.Now().UnixNano()
		var wg sync.WaitGroup
		for i := 0; i < nrOfThreads; i++ {
			wg.Add(1)

			//Start each thread
			go func(index int) {
				defer wg.Done()
				multiplyStuff(&finishedMatrix, matrixes[index],
					matrix2, index)
			}(i)
		}
		wg.Wait()
	}
	end := time.Now().UnixNano()
	fmt.Printf("Execution took %d ns\n", end-start)
	runtime.GC()
}
