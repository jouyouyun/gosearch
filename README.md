Search key by regexp
====================

`NewSearchWithStrList(list []string) (md5Str string, ok bool)`

传入字符串数组，返回 `md5` 和是否成功, 其中 `md5` 作为以后搜索的句柄.

`SearchString(key, md5 string) (retList []string)`

搜索并返回 `key` 的结果。

`SearchStartWithString(str, md5 string) (retList []string)`

搜索并返回以 `key` 开头的结果。


示例请查看 `search_test.go`
---------------------------
