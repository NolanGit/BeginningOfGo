package main

import "fmt"

func main() {
	var m map[string]string //不初始化map的时候，map为nil，不能存放健值对
	fmt.Println(m == nil)

	m2 := make(map[string]string) // 为map开辟内存？
	m2["a"] = "aa"
	t := m2["a"]
	fmt.Println(t)

	countryCapitalMap := make(map[string]string)
	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"
	for k, v := range countryCapitalMap {
		fmt.Printf("%s's capital is %s\n", k, v)
	}

	capital, ok := countryCapitalMap["India"] /*如果确定是真实的,则存在,否则不存在 */
	fmt.Println(capital)
	fmt.Println(ok)

	delete(countryCapitalMap, "India")
	capital, ok = countryCapitalMap["India"] /*如果确定是真实的,则存在,否则不存在 */
	fmt.Println(capital)
	fmt.Println(capital == "")
	fmt.Println(ok)

}
