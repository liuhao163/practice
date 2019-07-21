package main

import "fmt"

func KMP(str string, ptr string) int {
	next := getNext(ptr)
	i := 0
	j := 0
	for i < len(str) && j < len(ptr) {
		if j < 0 || str[i] == ptr[j] {
			j++
			i++
		} else {
			j = next[j]
		}
	}

	if j == len(ptr) {
		return i - j
	}

	return -1
}

//失效函数,求解ptr的next数组，我们可以看做是ptr字符串和自己的最长前缀的匹配
//重点：
// 1、公式1：如果p[i]的最长匹配前缀子串是j,如果p[i+1]==p[j+1]，那么p[i+1]的最长匹配子串是j+1，next[i+1]=j+1
// 2、公式2：上述情况下如果p[i+1]!=p[j+1],我们认为j+1是坏字符串，应该将最长前缀字符串(在这里是模式串)挪动next[j+1]个距离假设是y，在继续查找，如果这时候p[i+1]==p[y],那么p[i+1]的最长匹配子串就是next[i+1]=y
func getNext(ptr string) []int {
	next := make([]int, len(ptr))
	next[0] = -1 //如果前缀只有一个字符是没有好前缀的

	i := 0
	j := -1

	//遍历模式串ptr
	for i < len(ptr)-1 {
		//j==-1时候，匹配失效，复位j=0说明没有匹配得上的最长前缀子串；
		// 如果ptr[i]==ptr[j],说明next[i]=j，然后俩个都++如果继续相等next[i+1]=j+1，
		if j == -1 || ptr[i] == ptr[j] {
			j++
			i++
			next[i] = j
		} else {
			//如果ptr[i] != ptr[j]并且j!=-1，相当于模式串的最长后缀和模式串的最长前缀无法匹配
			//这时候要移动j,移动的方案是假设j是坏字符，那么查找ptr[0,j]的最长前缀子串一定在next数组中（上面的分支已经匹配过了）
			//所以移动j，距离是next[j]（刚才匹配的最长前缀的长度）
			j = next[j] //****不会空指针的原因在上面的分支，上一次循环next[i++]=j++，i本身就>j，所以next[j]之前已经计算过了。一定是有值的
		}
	}
	fmt.Print(next)
	return next
}
