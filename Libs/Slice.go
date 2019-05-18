package Libs

// 判断第一个参数是否在第二个数组中， 可以理解为PHP中的in_array
func In_slice(val interface{}, slice []interface{}) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

