## 时间函数
### date
```php
date("Y-m-d H:i:s");
```
```golang
time.Now().Format("2006-01-02 15:04:05")
```
### time
```php
time();
```
```golang
time.Now().Unix()
```
## 字符串函数
### substr
```php
// 纯数字或者英文
$str = "";
substr($str, startOffset, length);
// 带中文
str := "我是一名Golang开发工程师!"
mt_substr(str, startOffset, length);
```
```golang
// 纯数字或者英文
str = "";
str[start:start+length]
// 带中文
str := "我是一名Golang开发工程师!"
string([]rune(str)[:3]
```
### strlen
```php
$str = ""l;
// 带中文
mt_strlen($str);
// 纯数字或者英文
strlen($str)
```
```golang
// 带中文
str := ""
utf8.RuneCountInString(str)
// 纯数字或者英文
len(str)
```
### json_encode
```php
// 要encode的数组
$array = [];
json_encode($array);
```
```golang
func jsonEncode(data interface{}) (string, error) {
  // 返回的是[]byte
	jsons, err := json.Marshal(data)
  // 使用string将[]byte转成字符串
	return string(jsons), err
}
```
### json_decode
```php
// json字符串
$jsonStr = '{"name": "张三"}';
// 返回一个数组格式
json_decode($jsonStr, true);
```
```golang
func jsonDecode(data string) (map[string]interface{}, error) {
  // 定义一个map接收转换后的数据
	var dat map[string]interface{}
  // 将json字符串转为[]bye
	err := json.Unmarshal([]byte(data), &dat)
	return dat, err
}
```
### explode
```php
$str = "被分割的字符"
$delimiter = "分割符"
if (mb_strlen($str) < mb_strlen($delimiter)) {
  return $str;
} else {
  return explode($delimiter, $str);
}
```
```golang
str := "被分割的字符"
delimiter := "分割符"
if len(delimiter) > len(str) {
  return strings.Split(delimiter, str)
} else {
  return strings.Split(str, delimiter)
}
```
### implode
```php
$array = [];
implode(delimiter, $array);
```
```golang
str := []string
strings.Join(str, delimiter)
```
### md5
```golang
md5($str);
```
```golang
h := md5.New()
io.WriteString(h, str)
fmt.Sprintf("%x", h.Sum(nil))
```
## 数组
### sort
```php
$array = [];
sort($array);
```
```golang
// int slice
intSlice := []int{4, 2, 3}
sort.Ints(intSlice)
fmt.Println(intSlice)

// string slice
stringSlice := []string{"e", "c", "b", "a"}
sort.Strings(stringSlice)
fmt.Println(stringSlice)
```
### ksort
```php
$array = [];
ksort($array);
```
```golang
```