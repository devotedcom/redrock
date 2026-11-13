package scoreboard

import "sort"

// Totals 把每位学生的多次成绩累加到一张新 map。
// records 本身不能被修改；某个学生没有成绩时，总分为 0。
func Totals(records map[string][]int) map[string]int {
	result := make(map[string]int)
	for name, scores := range records {
		sum := 0
		for _, score := range scores {
			sum += score
		}
		result[name] = sum
	}
	// TODO: 创建结果 map，遍历 records，把每条成绩加起来。
	// 不要直接改 records 里的切片。
	return result
}

// Entry 是排行榜上的一行。
type Entry struct {
	Name  string
	Score int
}

// Top 返回分数最高的前 n 名。
// 分数相同时按姓名升序；n <= 0 时返回空结果；n 大于人数时返回全部。
func Top(totals map[string]int, n int) []Entry {
	if n <= 0 {
		return []Entry{}
	}
	entries := make([]Entry, 0, len(totals))
	for name, score := range totals {
		entries = append(entries, Entry{name, int(score)})
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].Score == entries[j].Score {
			return entries[i].Name < entries[j].Name
		}
		return entries[i].Score > entries[j].Score
	})
	if n > len(entries) {
		n = len(entries)
	}
	// TODO: 把 map 转成切片，按「分数降序、姓名升序」排序，再截取前 n 个。
	return entries[:n]
}
