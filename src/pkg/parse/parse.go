package parse

import "strconv"

func StringToInt(data string) (int, error) {
	num, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(num), nil
}
