package main
import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}
/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	} else if l1 == nil && l2 == nil {
		return nil
	} else { 
		var dgt1, dgt2 int
		k := 1
		for l1 != nil {
			dgt1 += l1.Val*k
			k *= 10
			l1 = l1.Next
		}
		k = 1
		for l2 != nil {
			dgt2 += l2.Val*k
			k *= 10
			l2 = l2.Next
		}
		sum := dgt1 + dgt2
		fmt.Println(sum)
		if sum != 0 {
			for sum != 0 {
				res = listPushBack(res, sum%10)
				sum = sum/10
			}
		} else {
			res = listPushBack(res, sum)
		}
	}
	return res
}
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res *ListNode
	var arr1, arr2 []int
	if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 == nil {
		return l1
	} else if l1 == nil && l2 == nil {
		return nil
	} else { 
		for l1 != nil {
			arr1 = append(arr1, l1.Val)
			l1 = l1.Next
		}
		for l2 != nil {
			arr2 = append(arr2, l2.Val)
			l2 = l2.Next
		}
		l1 := len(arr1)
		l2 := len(arr2)
		fmt.Println(arr1[l1-1],arr2[l2-1])
		for l1 != 0 && l2 != 0 {
			if l1 == 0 && l2 != 0{
				res = listPushBack(res, arr2[l2-1])
			} else if l2 == 0 && l1 != 0 {
				res = listPushBack(res, arr2[l2-1])
			} else {
				res = listPushBack(res, arr1[l1-1]+arr2[l2-1])
			}
			l1--
			l2--
		} 
		//fmt.Println(arr1,arr2)
	}
	return res
}

func PrintList(l *ListNode) {
	it := l
	for it != nil {
		fmt.Print(it.Val, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *ListNode, data int) *ListNode {
	n := &ListNode{Val: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {

	var link *ListNode
	link = listPushBack(link, 1)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 0)
	link = listPushBack(link, 1)
	PrintList(link)

	var link2 *ListNode
	link2 = listPushBack(link2, 5)
	link2 = listPushBack(link2, 6)
	link2 = listPushBack(link2, 4)
	PrintList(link2)
	PrintList(addTwoNumbers(link,link2))
}
