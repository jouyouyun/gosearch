/**
 * Copyright (c) 2011 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package gosearch

import (
	"fmt"
	"github.com/jouyouyun/gopinyin"
	"path"
)

func NewSearchWithStrList(list []string) (string, bool) {
	datas := []dataInfo{}
	strs := ""

	for _, v := range list {
		strs += v + "+"
		pyList := gopinyin.HansToPinyin(v)
		if len(pyList) == 1 && pyList[0] == v {
			info := dataInfo{v, v}
			datas = append(datas, info)
			continue
		}
		for _, k := range pyList {
			info := dataInfo{k, v}
			datas = append(datas, info)
		}
		info := dataInfo{v, v}
		datas = append(datas, info)
	}

	md5Str, ok := sumStrMd5(strs)
	if !ok {
		fmt.Println("Sum MD5 Failed")
		return "", false
	}

	cachePath, ok1 := getCachePath()
	if !ok1 {
		fmt.Println("Get Cache Path Failed")
		return "", false
	}

	filename := path.Join(cachePath, md5Str)
	writeDatasToFile(&datas, filename)

	return md5Str, true
}

func SearchString(str, md5 string) []string {
	list := []string{}
	if len(str) < 1 || len(md5) < 1 {
		return list
	}

	list = searchString(str, md5)
	tmp := []string{}
	for _, v := range list {
		if !strIsInList(v, tmp) {
			tmp = append(tmp, v)
		}
	}

	return tmp
}

func SearchStartWithString(str, md5 string) []string {
	list := []string{}
	if len(str) < 1 || len(md5) < 1 {
		return list
	}

	list = searchStartWithString(str, md5)
	tmp := []string{}
	for _, v := range list {
		if !strIsInList(v, tmp) {
			tmp = append(tmp, v)
		}
	}

	return tmp
}

func strIsInList(str string, list []string) bool {
	for _, l := range list {
		if str == l {
			return true
		}
	}

	return false
}
