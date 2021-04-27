package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is zhaoyin12580@gmail.com
My email is zhaoyin12580@gmail.com
My email is zhaoyin12580@gmail.com
My email is zhaoyin12580@gmail.com
My email is zhaoyin12580@gmail.com
My email is aaa@bbb.com.cn
`
const strings = `
	<a href="http://www.zhenai.com/zhenghun/bijie" data-v-1573aa7c>毕节</a>
<a href="http://www.zhenai.com/zhenghun/bijie" data-v-1573aa7c>毕节</a>
<a href="http://www.zhenai.com/zhenghun/bijie" data-v-1573aa7c>毕节</a>
`
func main() {
	//compile, err := regexp.Compile("zhaoyin12580@gmail.com@asadsa.com")
	//compile := regexp.MustCompile(`[a-zA-Z0-9]+@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	compile := regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/bijie" data-v-1573aa7c>毕节</a>`)
	//compile := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)\.([a-zA-Z0-9]+)`)
	//findString := compile.FindString(text)
	//allString := compile.FindAllString(text, -1)// -1 表示所有
	allSubmatchString := compile.FindAllString(strings,-1) // 用()提取
	// [zhaoyin12580@gmail.com  zhaoyin12580 gmail com]
	//fmt.Println(allString)
	fmt.Println(allSubmatchString)
}
