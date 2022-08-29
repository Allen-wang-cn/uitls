/**
 * @Description: strings operate tools
 * @Author: Allen
 * @since: 2022-08-23 11:54:58
 * @lastTime: 2022-08-23 14:09:36
 * @LastAuthor: Allen
 * @FilePath: \uitls\string.go
 */
package utils

/**
 * @Description: 获取字符串长度 如果是空串 返回 0  如果非空 转为rune切片 再获取长度
 * @since: 2022-08-23 14:08:29
 * @param {string} param
 * @return {*}
 */
func StrLen(param string) int64 {
	if param == "" {
		return 0
	}
	ru := []rune(param)
	return int64(len(ru))
}

func StrToSlip(param string) int64 {

	return 0
}
