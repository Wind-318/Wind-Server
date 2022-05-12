package anime

import (
	"spider/callClient/callAlgorithm"
	"spider/callClient/callDatabase"
	"spider/callClient/callUser"
	"spider/callClient/data"
	"spider/config"
	"spider/registerCenter"
	"strconv"
)

// 搜索功能
func Search(request registerCenter.AnimeData) ([]string, []data.SelectAnimeDataInterface) {
	ids, infos := []string{}, []data.SelectAnimeDataInterface{}
	// 检查权限，非登录状态直接返回空
	if ok, err := callUser.CallUserIsLogin(request.Cookie); err != nil || !ok {
		return ids, infos
	}
	// 用户输入字段
	text := request.Text

	// 若搜索内容全为空格，直接返回空
	isEmpty := true
	for i := range text {
		if text[i] != ' ' {
			// 有不为空的，跳出
			isEmpty = false
			break
		}
	}
	if isEmpty {
		return ids, infos
	}

	names, err := callDatabase.CallMySQLSelectAnimeName(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider")
	if err != nil {
		return ids, infos
	}

	// 轮询数据库查询
	for index := len(names) - 1; index >= 0; index-- {
		// 进行匹配
		if ok, err := callAlgorithm.CallAlgorithmMatch(names[index], text); err == nil && ok {
			// 选出数据
			tempInfo := data.SelectAnimeDataInterface{}
			tempInfo, err := callDatabase.CallMySQLSelectAnimeInfoByName(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", names[index])
			// 错误处理
			if err != nil {
				return ids, infos
			}
			// 选取 source
			tempInfo.Source, err = callDatabase.CallMySQLSelectAnimeSourceByName(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", names[index])
			// 错误处理
			if err != nil {
				return ids, infos
			}
			// 选取 url
			tempInfo.Urls, err = callDatabase.CallMySQLSelectAnimeUrlByName(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", names[index])
			// 错误处理
			if err != nil {
				return ids, infos
			}

			ids = append(ids, strconv.Itoa(index))
			infos = append(infos, tempInfo)
		}
	}

	// 返回数据
	return ids, infos
}

// 选出新番
func SearchNewAnime(request registerCenter.AnimeData) ([]string, []data.SelectAnimeDataInterface) {
	ids, infos := []string{}, []data.SelectAnimeDataInterface{}
	// 非登录状态直接返回空数据
	if ok, err := callUser.CallUserIsLogin(request.Cookie); err != nil || !ok {
		return ids, infos
	}

	// isNew 为 1 则为当季动漫
	names, err := callDatabase.CallMySQLSelectNewAnime(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider")
	if err != nil {
		return ids, infos
	}
	// 逻辑同上
	for index := len(names) - 1; index >= 0; index-- {
		tempInfo := data.SelectAnimeDataInterface{}
		tempInfo, err := callDatabase.CallMySQLSelectAnimeInfoByName(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", names[index])

		// 错误处理
		if err != nil {
			return ids, infos
		}
		// 选取 source
		tempInfo.Source, err = callDatabase.CallMySQLSelectAnimeSourceByName(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", names[index])

		// 错误处理
		if err != nil {
			return ids, infos
		}
		// 选取 url
		tempInfo.Urls, err = callDatabase.CallMySQLSelectAnimeUrlByName(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", names[index])

		// 错误处理
		if err != nil {
			return ids, infos
		}

		ids = append(ids, strconv.Itoa(index))
		infos = append(infos, tempInfo)
	}

	// 返回数据
	return ids, infos
}

// 选出指定年份番剧
func SearchByYear(request registerCenter.AnimeData) ([]string, []data.SelectAnimeDataInterface) {
	ids, infos := []string{}, []data.SelectAnimeDataInterface{}
	// 非登录直接返回空
	if ok, err := callUser.CallUserIsLogin(request.Cookie); err != nil || !ok {
		return ids, infos
	}
	// 选取的年份
	year := request.Year

	// 依据年份搜索
	y, _ := strconv.Atoi(year)
	names, err := callDatabase.CallMySQLSelectAnimeByYear(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", y)
	if err != nil {
		return ids, infos
	}
	// 逻辑同上
	for index := len(names) - 1; index >= 0; index-- {
		tempInfo := data.SelectAnimeDataInterface{}
		tempInfo, err := callDatabase.CallMySQLSelectAnimeInfoByName(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", names[index])
		// 错误处理
		if err != nil {
			return ids, infos
		}
		// 选取 source
		tempInfo.Source, err = callDatabase.CallMySQLSelectAnimeSourceByName(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", names[index])
		// 错误处理
		if err != nil {
			return ids, infos
		}
		// 选取 url
		tempInfo.Urls, err = callDatabase.CallMySQLSelectAnimeUrlByName(config.MySQLAccount, config.MySQLPassword, config.MySQLIP, config.MySQLPort, "spider", names[index])
		// 错误处理
		if err != nil {
			return ids, infos
		}

		ids = append(ids, strconv.Itoa(index))
		infos = append(infos, tempInfo)
	}

	return ids, infos
}
