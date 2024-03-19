package leetcode2129

// 快慢指针
// l, r
// l指向一个单词的开头，r指向空格
/*
策略：
if r-l > 2, bytes[l] -> 大写， 其他 小写
else 全部小写
l -> r
*/
func capitalizeTitle(title string) string {
	bytesTitle := []byte(title)
	l, r := 0, 0
	n := len(title)
	for l < n {
		// 挪动r，找到下一个空格或者结束
		for r < n && bytesTitle[r] != ' ' {
			r++
		}
		// 判断长度
		if r-l > 2 {
			// l变大写
			bytesTitle[l] = toUpper(bytesTitle[l])
		} else {
			// 保持小写
			bytesTitle[l] = toLower(bytesTitle[l])
		}
		l++
		for l < r {
			// 其他全部小写
			bytesTitle[l] = toLower(bytesTitle[l])
			l++
		}
		l++
		r++
	}
	return string(bytesTitle)
}

func toUpper(v byte) byte {
	// 字母只含有英文，小于'a' 则认为大写
	if v < 'a' {
		return v
	}
	return v - 32
}

func toLower(v byte) byte {
	if v >= 'a' {
		return v
	}
	return v + 32
}
