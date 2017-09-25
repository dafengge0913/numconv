package numconv

var chnNumChar = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
var chnUnitSection = []string{"", "万", "亿", "万亿"}
var chnUnitChar = []string{"", "十", "百", "千"}

func ArabicToChinese(n int) string {
	if n == 0 {
		return chnNumChar[0]
	}
	needZero := false
	result := ""
	unit := 0
	for n > 0 {
		// 以万为节
		section := n % 10000
		if needZero {
			// 上一节千位为0 需要在上一节前面加“零”
			result = chnNumChar[0] + result
		}
		secResult := sectionToChinese(section)
		if section != 0 {
			result = secResult + chnUnitSection[unit] + result // 每节 + 节权位
		}
		needZero = section < 1000 && section > 0 // 全部为0 不需要加“零”
		n /= 10000
		unit++
	}
	return result
}

func sectionToChinese(section int) string {
	zero := true
	unit := 0
	result := ""
	for section > 0 {
		v := section % 10
		if v == 0 {
			if section == 0 || !zero {
				// 包含不为零的数字时加“零” 连续多个只加一个“零”
				zero = true
				result = chnNumChar[0] + result
			}
		} else {
			zero = false
			result = chnNumChar[v] + chnUnitChar[unit] + result // 数字对应的中文数字 + 权位
		}
		unit++
		section = section / 10
	}
	return result
}
