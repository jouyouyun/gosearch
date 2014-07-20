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
	"testing"
)

func TestSearch(t *testing.T) {
	strList := []string{
		"Firefox 浏览器",
		"Google Chrome 浏览器",
		"雷鸟",
		"文本编辑器",
		"Sublime Text",
		"Retext",
	}

	md5Str, ok := NewSearchWithStrList(strList)
	if !ok {
		t.Error("New search failed")
		return
	}

	retList := SearchString("liulanqi", md5Str)
	fmt.Println("Search 'liulanqi', ret:", retList)

	retList = SearchString("text", md5Str)
	fmt.Println("Search 'text', ret:", retList)

	retList = SearchStartWithString("s", md5Str)
	fmt.Println("Search 't', ret:", retList)
}
