package pas2

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Parse(spec string) ([]int, error) {
	if strings.TrimSpace(spec) == "" {
		return []int{}, nil
	}
	seen := map[int]bool{}
	parse := func(text string) (int, error) {
		text = strings.TrimSpace(text)
		if text == "" {
			return 0, fmt.Errorf("empty port")
		}
		for _, ch := range text {
			if ch < '0' || ch > '9' {
				return 0, fmt.Errorf("invalid port")
			}
		}
		value, e := strconv.Atoi(text)
		if e != nil || value < 1 || value > 65535 {
			return 0, fmt.Errorf("port outside 1..65535")
		}
		return value, nil
	}
	for _, segment := range strings.Split(spec, ",") {
		parts := strings.Split(segment, "-")
		if len(parts) > 2 {
			return nil, fmt.Errorf("invalid range")
		}
		start, e := parse(parts[0])
		if e != nil {
			return nil, e
		}
		end := start
		if len(parts) == 2 {
			end, e = parse(parts[1])
			if e != nil {
				return nil, e
			}
		}
		if end < start {
			return nil, fmt.Errorf("reversed range")
		}
		for port := start; port <= end; port++ {
			seen[port] = true
		}
	}
	result := make([]int, 0, len(seen))
	for port := range seen {
		result = append(result, port)
	}
	sort.Ints(result)
	return result, nil
}
