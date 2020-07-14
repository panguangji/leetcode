/*
*创建者：pangj
*创建时间：2020/7/14
*描述：
 */
package main

import "fmt"

func main()  {
	strs := []string{"a","a","a"}
	res := LongestCommonPrefix(strs)
	if res == "" {
		fmt.Println("输入不存在公共前缀")
	}else {
		fmt.Println(res)
	}
}

func LongestCommonPrefix(strs []string) string {

	minLength := 0
	if len(strs) == 0  {
		return ""
	}else if len(strs) == 1 {
		return strs[0]
	}
	minLength = len(strs[0])
	for _, str := range strs {
		if len(str) < minLength{
			minLength = len(str)
		}
	}

	var index int
	for i:=0; i< minLength ;i++  {
		isEnd := false
		for j := 0; j< len(strs)-1 ; j++  {
			if strs[j][i] != strs[j+1][i]{
				isEnd = true
				break
			}
			//全部检查通过，i: 索引等于长度减1， j:索引等于长度减1，然后是前一个值和后一个值相比较，如果比较到了最后个值，则代表所有都一致
			if i == minLength -1  && j == len(strs)-2{
				i++
				fmt.Println("i = ",i )
				isEnd = true
				break
			}
		}
		if isEnd{
			index = i
			break
		}
	}
	if index  == 0 {
		return ""
	}else {
		return strs[0][0:index]
	}
}