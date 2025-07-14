package validator

import (
	"database/sql"
	"fmt"
	"trading-data-harvester/pkg/datamodel"
	utils "trading-data-harvester/utils/convert"
)

// ValidateTimeIntegrity 验证时间完整性，验证以往数据是否存在时间缺口
func ValidateTimeIntegrity(db *sql.DB, databaseName string, tableName string, intervalMs string, symbol string) ([]datamodel.ValidatorTimeGap, error) {

	var gaps []datamodel.ValidatorTimeGap

	// 间隔时间转换为毫秒
	intervalMsInt := utils.Interval2Ms(intervalMs)

	// 获取最早的时间
	var earliestOpenTime uint64
	query := fmt.Sprintf("SELECT MIN(open_time) FROM %s.%s WHERE symbol = '%s' AND interval = '%s'", databaseName, tableName, symbol, intervalMs)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&earliestOpenTime)
		if err != nil {
			return nil, err
		}
	}

	for {
		query := buildQuery(databaseName, tableName, symbol, intervalMs, earliestOpenTime, pageSize)
		rows, err := db.Query(query)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var openTimes []uint64
		for rows.Next() {
			var openTime uint64
			if err := rows.Scan(&openTime); err != nil {
				rows.Close()
				return nil, fmt.Errorf("读取 open_time 失败: %w", err)
			}
			openTimes = append(openTimes, openTime)
		}
		rows.Close()

		// 数据读取完毕，退出
		if len(openTimes) == 0 {
			break
		}

		expected := earliestOpenTime
		// 校验时间差
		for _, t := range openTimes {
			for expected <= t {
				if t != expected {
					gaps = append(gaps, datamodel.ValidatorTimeGap{
						Symbol:      symbol,
						Interval:    intervalMs,
						MissingFrom: expected,
						MissingTo:   expected + uint64(intervalMsInt),
					})
				}
				expected += uint64(intervalMsInt)
			}
		}

		// 每次循环后，earliestOpenTime 增加 interval*pageSize，防止死循环
		earliestOpenTime = expected + uint64(intervalMsInt)

		// 如果本批不足一页，也退出
		// if len(openTimes) < pageSize {
		// 	break
		// }
	}

	return gaps, nil // 返回所有找到的时间缺口
}

// 构造查询语句
func buildQuery(databaseName string, tableName string, symbol string, intervalMs string, earliestOpenTime uint64, pageSize int) string {
	query := fmt.Sprintf("SELECT open_time FROM %s.%s WHERE symbol = '%s' AND interval = '%s' AND open_time >= %d ORDER BY open_time ASC LIMIT %d", databaseName, tableName, symbol, intervalMs, earliestOpenTime, pageSize)
	return query
}
