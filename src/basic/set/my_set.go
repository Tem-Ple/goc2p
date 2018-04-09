package set

func main() {
	var countryCapitalMap map[string]int
	/* 创建集合 */
	countryCapitalMap = make(map[string]int)

	/* map 插入 key-value 对，各个国家对应的首都 */
	countryCapitalMap["France"] = 1
	countryCapitalMap["Italy"] = 2
	countryCapitalMap["Japan"] = 3
	countryCapitalMap["India"] = 4
}
