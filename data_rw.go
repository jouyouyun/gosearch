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
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
)

func writeDatasToFile(datas interface{}, filename string) {
	if datas == nil {
		fmt.Print("writeDatasToFile args error")
		return
	}

	var w bytes.Buffer
	enc := gob.NewEncoder(&w)
	if err := enc.Encode(datas); err != nil {
		fmt.Print("Gob Encode Datas Failed:", err)
		return
	}

	fp, err := os.Create(filename)
	if err != nil {
		fmt.Print("Open '%s' failed:", err)
		return
	}
	defer fp.Close()

	fp.WriteString(w.String())
	fp.Sync()
}

func readDatasFromFile(datas interface{}, filename string) bool {
	if !isFileExist(filename) || datas == nil {
		fmt.Print("readDatasFromFile args error")
		return false
	}

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print("ReadFile '%s' failed:", err)
		return false
	}

	r := bytes.NewBuffer(contents)
	dec := gob.NewDecoder(r)
	if err = dec.Decode(datas); err != nil {
		fmt.Print("Decode Datas Failed:", err)
		return false
	}

	return true
}
