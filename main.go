package main

import (
	"fmt"
	"practice/dp"
)

func main() {

	//m := make(map[int]int64, 100)
	//
	//fmt.Println(cal.Fibonacci(30, m))
	//
	//fmt.Println()
	//fmt.Println("SelectSort========")
	//
	//var a = []int{80, 3, 9, 8, 9, 77, 35, 78, 90}
	//
	//sort.SelectSort(a)
	//
	//for i := 0; i < len(a); i++ {
	//	fmt.Print(a[i]);
	//	fmt.Print(" ")
	//}
	//
	//fmt.Println()
	//fmt.Println("InsertSort========")
	//
	//var b = []int{80, 3, 9, 8, 9, 77, 35, 78, 90}
	//
	//sort.InsertSort(b)
	//
	//for i := 0; i < len(b); i++ {
	//	fmt.Print(b[i]);
	//	fmt.Print(" ")
	//}
	//
	//fmt.Println()
	//fmt.Println("BubbleSort========")
	//var c = []int{80, 3, 9, 8, 9, 77, 35, 78, 90}
	//
	//sort.BubbleSort(c)
	//for i := 0; i < len(c); i++ {
	//	fmt.Print(c[i]);
	//	fmt.Print(" ")
	//}
	//
	//fmt.Println()
	//fmt.Println("MergeSort========")
	//var mc = []int{80, 3, 9, 8, 9, 77, 35, 78, 90, 100, 101, 23, 3, 5, 6, 7, 890, 2, 32, 3431, 233}
	//
	//sort.MergeSort(mc, 0, len(mc)-1)
	//for i := 0; i < len(mc); i++ {
	//	fmt.Print(mc[i]);
	//	fmt.Print(" ")
	//}
	//
	//fmt.Println()
	//fmt.Println("QuickSort========")
	//var qc = []int{80, 3, 9, 8, 9, 77, 35, 78, 90, 100, 101, 23, 3, 5, 6, 7, 890, 2, 32, 3431, 233}
	//
	//sort.QuickSort(qc, 0, len(qc)-1)
	//for i := 0; i < len(qc); i++ {
	//	fmt.Print(qc[i]);
	//	fmt.Print(" ")
	//}

	//fmt.Println()
	//fmt.Println("Tree========")
	//treeNode := tree.InitTree();
	//fmt.Println("Tree preOrder========")
	//tree.PreOrder(&treeNode)
	//fmt.Println()
	//fmt.Println("Tree midOrder========")
	//tree.MidOrder(&treeNode)
	//fmt.Println()
	//fmt.Println("Tree postOrder========")
	//tree.PostOrder(&treeNode)
	//fmt.Println()
	//
	//t := &tree.TreeNode{}
	//t.Value = 33
	//tree.InsertNode(t, 16)
	//tree.InsertNode(t, 50)
	//tree.InsertNode(t, 13)
	//tree.InsertNode(t, 18)
	//tree.InsertNode(t, 34)
	//tree.InsertNode(t, 58)
	//tree.InsertNode(t, 15)
	//tree.InsertNode(t, 17)
	//tree.InsertNode(t, 25)
	//tree.InsertNode(t, 51)
	//tree.InsertNode(t, 66)
	//tree.InsertNode(t, 19)
	//tree.InsertNode(t, 27)
	//tree.InsertNode(t, 55)
	//fmt.Println("Tree insert========")
	//t.PreIterate()
	//fmt.Println(t)

	//s := "abbacabczbcbcbacabc"
	//p := "cab"
	//i := stringsearch.BM(s, p)
	//
	//if i > 0 {
	//	fmt.Printf("idx:%d,string:%s", i, s[i:i+len(p)])
	//}

	//fmt.Println(dp.Knapsack(dp.Pkg, dp.Value, 5, 9))
	//
	//fmt.Println(dp.ShortDir(dp.Triangle))
	//

	sum := 102
	money := []int{1, 2, 5}
	fmt.Println(dp.ChargeDP(money, sum))

	dp.ChargeRC(0, sum, money, 0)
	fmt.Println(dp.MinVal)
}
