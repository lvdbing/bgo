package convert

import (
	"errors"
)

var unitNumList = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var chineseUnitNumList = []string{"零", "一", "二", "三", "四", "五", "六", "七", "八", "九"}
var chineseOldUnitNumList = []string{"零", "壹", "贰", "叁", "肆", "伍", "陆", "柒", "捌", "玖"}

// GetChineseUnitNum 个位阿拉伯数字，转换为汉字。
func GetChineseUnitNum(num int) (string, error) {
	return matchChineseUnitNum(num, chineseUnitNumList)
}

// GetChineseOldUnitNum 个位阿拉伯数字，转换为大写汉字。
func GetChineseOldUnitNum(num int) (string, error) {
	return matchChineseUnitNum(num, chineseOldUnitNumList)
}

func matchChineseUnitNum(num int, list []string) (string, error) {
	if num > 9 || num < 0 {
		return "", errors.New("数字必须为0-9之间")
	}
	for i, v := range unitNumList {
		if v == num {
			return list[i], nil
		}
	}
	return "", errors.New("没有匹配到")
}

// GetCalendarDay 日期数字转换为农历日期。
func GetLunarCalendarDay(num int) (string, error) {
	if num > 30 || num < 1 {
		return "", errors.New("日期必须为1-30之间")
	}
	switch num {
	case 10:
		return "初十", nil
	case 20:
		return "二十", nil
	case 30:
		return "三十", nil
	}
	tenBit := num % 10  // 十位
	unitBit := num / 10 // 个位
	unitStr, _ := GetChineseUnitNum(unitBit)
	switch tenBit {
	case 0:
		return "初" + unitStr, nil
	case 1:
		return "十" + unitStr, nil
	case 2:
		return "廿" + unitStr, nil
	}

	return "", errors.New("没有匹配到")
}
