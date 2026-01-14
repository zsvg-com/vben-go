package tesk

// 无参数
func NoParamsMethod() {
	println("无参数方法")
}

// 单个参数
func ParamsMethod(data string) {
	println(data)
}

// 多个参数 -一般参数为固定，模式
func MultipleParamsMethod(param1 string, param2 bool, param3 string, param4 string) {
	println(param1)
	println(param2)
	println(param3)
	println(param4)
}
