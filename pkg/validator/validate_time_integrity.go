package validator

import (
	"database/sql"
	"fmt"
	"trading-data-harvester/pkg/datamodel"
)

// ValidateTimeIntegrity 验证时间完整性，验证以往数据是否存在时间缺口
func ValidateTimeIntegrity(db *sql.DB, databaseName string, tableName string, intervalMs int64) ([]datamodel.ValidatorTimeGap, error) {

	var gaps []datamodel.ValidatorTimeGap
	var lastOpenTime int64 = -1 // 初始化为-1，表示尚未读取任何 open_time

	for {
		// 构造查询SQL，使用 lastOpenTime 作为起始点分页
		var query string
		if lastOpenTime == -1 {
			// 初始查询，获取最早的 open_time
			query = fmt.Sprintf("SELECT open_time FROM %s.%s ORDER BY open_time ASC LIMIT %d", databaseName, tableName, pageSize)
		} else {
			// 后续查询，获取比 lastOpenTime 更晚的 open_time
			query = fmt.Sprintf("SELECT open_time FROM %s.%s WHERE open_time > %d ORDER BY open_time ASC LIMIT %d", databaseName, tableName, lastOpenTime, pageSize)
		}

		// 执行查询
		rows, err := db.Query(query)
		if err != nil {
			return nil, fmt.Errorf("查询 open_time 失败: %w", err)
		}

		var openTimes []int64
		for rows.Next() {
			var openTime int64
			if err := rows.Scan(&openTime); err != nil {
				rows.Close()
				return nil, fmt.Errorf("读取 open_time 失败: %w", err)
			}
			openTimes = append(openTimes, openTime) // 收集查询到的 open_time
		}
		rows.Close()

		// 数据读取完毕，退出循环
		if len(openTimes) == 0 {
			break
		}

		// 校验时间差，寻找缺口
		for _, t := range openTimes {
			if lastOpenTime != -1 {
				expected := lastOpenTime + intervalMs // 计算期望的下一个 open_time
				if t != expected {
					// 缺口出现，记录缺口信息
					gaps = append(gaps, datamodel.ValidatorTimeGap{
						MissingFrom: expected,
						MissingTo:   t - intervalMs,
					})
				}
			}
			lastOpenTime = t // 更新 lastOpenTime 为当前的 open_time
		}

		// 如果本批次数据不足一页，说明已到达数据末尾，退出循环
		if len(openTimes) < pageSize {
			break
		}
	}

	return gaps, nil // 返回所有找到的时间缺口
}
