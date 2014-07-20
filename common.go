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
	"crypto/md5"
	"os"
	"os/user"
	"path"
	"strconv"
)

func md5ByteToStr(bytes [16]byte) string {
	str := ""

	for _, b := range bytes {
		s := strconv.FormatInt(int64(b), 16)
		if len(s) == 1 {
			str += "0" + s
		} else {
			str += s
		}
	}

	return str
}

func sumStrMd5(str string) (string, bool) {
	if len(str) < 1 {
		return "", false
	}

	md5Byte := md5.Sum([]byte(str))
	md5Str := md5ByteToStr(md5Byte)
	if len(md5Str) < 32 {
		return "", false
	}

	return md5Str, true
}

func getHomeDir() string {
	info, err := user.Current()
	if err != nil {
		return ""
	}

	return info.HomeDir
}

func getCacheDir() string {
	if homeDir := getHomeDir(); len(homeDir) > 0 {
		return path.Join(homeDir, ".cache")
	}
	return ""
}

func isFileExist(filename string) bool {
	if len(filename) < 1 {
		return false
	}

	_, err := os.Stat(filename)

	return err == nil || os.IsExist(err)
}
