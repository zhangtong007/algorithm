// n*n阶二维数组/切片的水平翻转、垂直翻转、对角线翻转、顺时针90度、逆时针90度实现
package sliceflip

import "fmt"

func Test() {
	arr := [][]interface{}{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	printArr(arr)
	// diagonalFilp(arr)
	// horizontalFlip(arr)
	// verticalFlip(arr)
	// rotateClockwise(arr)
	rotateCounterclockwise(arr)
	printArr(arr)
}

// 顺时针旋转90度
func rotateClockwise(arr [][]interface{}) {
	// 对角线旋转
	diagonalFilp(arr)
	// 水平翻转
	horizontalFlip(arr)
}

// 逆时针旋转90度
func rotateCounterclockwise(arr [][]interface{}) {
	// 对角线旋转
	diagonalFilp(arr)
	// 垂直翻转
	verticalFlip(arr)
}

// 水平翻转
func horizontalFlip(arr [][]interface{}) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			arr[i][j], arr[i][n-j-1] = arr[i][n-j-1], arr[i][j]
		}
	}
}

// 垂直翻转
func verticalFlip(arr [][]interface{}) {
	n := len(arr)
	for j := 0; j < n; j++ {
		for i := 0; i < n/2; i++ {
			arr[i][j], arr[n-i-1][j] = arr[n-i-1][j], arr[i][j]
		}
	}
}

// 对角线翻转
func diagonalFilp(arr [][]interface{}) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			arr[i][j], arr[j][i] = arr[j][i], arr[i][j]
		}
	}
}

// 元素打印
func printArr(arrs [][]interface{}) {
	fmt.Println("----------------------------------")
	for _, arr := range arrs {
		for _, v := range arr {
			fmt.Printf("%v\t", v)
		}
		fmt.Println()
	}
	fmt.Println("----------------------------------")
}
