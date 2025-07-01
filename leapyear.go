package main

func leapyear(y int) bool {
	// まず4で割り切れるか
	if y%4 == 0 {
		// 次に100で割り切れるか
		if y%100 == 0 {
			// 100で割り切れる場合、さらに400で割り切れるかをチェック
			if y%400 == 0 {
				return true // 400で割り切れる場合はうるう年
			}
			return false // 100で割り切れるが400で割り切れない場合はうるう年ではない
		}
		return true // 4で割り切れるが100で割り切れない場合はうるう年
	}
	return false // 4で割り切れない場合はうるう年ではない
}
