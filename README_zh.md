# config
[English Document](./README.md)
## 一、介绍
这是一个用golang语言实现的读取配置文件的包。
读取文本文件，生成map或者struct。
## 二、使用方法
- 下载包
```
go get github.com/tpgzcyyao/config
```
- 导入包
```
import "github.com/typzcyyao/config"
```
- 读取配置文件生成map，方法为LoadFile
```
res, err := (new(config.Config)).LoadFile("/export/config/test.conf")
if err != nil {
	return err
}
fmt.Println(fmt.Sprintf("%v", res))
```
- 读取配置文件生成struct，方法为LoadConfig，ResConfig为自定义的struct
```
resConfig := new(ResConfig)
err := (new(config.Config)).LoadConfig("/export/config/test.conf", resConfig)
if err != nil {
	return err
}
fmt.Println(fmt.Sprintf("%v", resConfig))
```
## 三、说明
### 读取文件返回map
```
func (c *Config) LoadFile(path string) (map[string]map[string]string, error)
```
- path为配置文件的绝对路径
- 返回map的value全部为string类型
### 读取文件返回结构体
```
func (c *Config) LoadConfig(path string, v interface{}) error
```
- path为配置文件的绝对路径
- v接收传入的结构体，方法执行完之后，配置会加载到结构体的变量中
### 配置文件示例
```
[string_section] # this is string config
string_first = this is string
string_second = this is second = 2 # this is comment in line
# this is comment line

[int_section]
int_first = -2147483648
int_second = -9223372036854775808
int_third = 2147483648
int_forth = 9223372036854775808

[float_section]
float_first = 1.111111111111111111111111111111
float_second = 2.1111111111111111111111111111111

[boolean_section]
boolean_first = false
boolean_second = true
```
### 自定义结构体示例
```
type ResConfig struct {
        StringSection  Strings
        IntSection     Ints
        FloatSection   Floats
        BooleanSection Booleans
}
type Strings struct {
        StringFirst  string
        StringSecond string
}
type Ints struct {
        IntFirst  int
        IntSecond int64
        IntThird  uint
        IntForth  uint64
}
type Floats struct {
        FloatFirst  float32
        FloatSecond float64
}
type Booleans struct {
        BooleanFirst  bool
        BooleanSecond bool
}
```
### 自定义结构体支持的基本类型
- string
- bool
- int
- int64
- uint
- uint64
- float32
- float64
### 配置文件中的key和结构体字段的关系
- 配置文件中key为下划线式（this\_is\_section），对应的结构体中的字段为骆驼式（ThisIsSection）
- [section]对应结构体本身的字段
- key = value对应子结构体的字段
## 四、示例
- 执行
```
cd $GOPATH/src/github.com/tpgzcyyao/config/test
go test
```
- 会得到示例结果
```
The config map is: 
map[boolean_section:map[boolean_first:false boolean_second:true] float_section:map[float_first:1.111111111111111111111111111111 float_second:2.1111111111111111111111111111111] int_section:map[int_first:-2147483648 int_forth:9223372036854775808 int_second:-9223372036854775808 int_third:2147483648] string_section:map[string_first:this is string string_second:this is second = 2]]

The config map in json format is: 
{
        "boolean_section": {
                "boolean_first": "false",
                "boolean_second": "true"
        },
        "float_section": {
                "float_first": "1.111111111111111111111111111111",
                "float_second": "2.1111111111111111111111111111111"
        },
        "int_section": {
                "int_first": "-2147483648",
                "int_forth": "9223372036854775808",
                "int_second": "-9223372036854775808",
                "int_third": "2147483648"
        },
        "string_section": {
                "string_first": "this is string",
                "string_second": "this is second = 2"
        }
}

The config struct is: 
&{{this is string this is second = 2} {-2147483648 -9223372036854775808 2147483648 9223372036854775808} {1.1111112 2.111111111111111} {false true}}
```
- 示例配置文件位置
```
$GOPATH/src/github.com/tpgzcyyao/config/test/test.conf
```
