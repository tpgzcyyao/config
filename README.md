# config
[中文文档](./README_zh.md)
## I. Introduction
This is a golang package using for reading config file and producing map or struct。
## II. Instructions
- download the package
```
go get github.com/tpgzcyyao/config
```
- import the package
```
import "github.com/typzcyyao/config"
```
- The function LoadFile is used for reading config file and producing map.
```
res, err := (new(config.Config)).LoadFile("/export/config/test.conf")
if err != nil {
	return err
}
fmt.Println(fmt.Sprintf("%v", res))
```
- The function LoadConfig is used for reading config file and producing customized struct. The ResConfig below is a customized struct.
```
resConfig := new(ResConfig)
err := (new(config.Config)).LoadConfig("/export/config/test.conf", resConfig)
if err != nil {
	return err
}
fmt.Println(fmt.Sprintf("%v", resConfig))
```
## III. Description
### read config file and produce map
```
func (c *Config) LoadFile(path string) (map[string]map[string]string, error)
```
- Parameter path is the absolute path for config file.
- The values' type in the map are all string type.
### read config file and produce struct
```
func (c *Config) LoadConfig(path string, v interface{}) error
```
- Parameter path is the absolute path for config file.
- Parameter v receives incoming struct. The config will load in the struct parameter when the function LoadConfig execute completed.
### config file example
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
### customized struct example
```
type ResConfig struct {
    StringSection  Strings
    IntSection     Ints
    FloatSection   Floats
    BooleanSection Booleans
    IPICMP         IPS `conf:"ip_icmp"`
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
type IPS struct {
    IPv4  string `conf:"ipv4"`
    IPv6  string `conf:"ipv6"`
    IPxx  string `conf:"ip_xx"`
    IPVxx string `conf:"ip_vxx"`
    IpRaw string
}
```
### the basic type supported by customized struct
- string
- bool
- int
- int8
- int16
- int32
- int64
- uint
- uint8
- uint16
- uint32
- uint64
- float32
- float64
### the relationship between keys in config file and fields in struct
- Keys in config file are snake format(this\_is\_section). Fields in struct are camel format(ThisIsSection).
- [section] is corresponding to the field belongs to the struct itself.
- key = value is corresponding to the filed belongs to the struct's children.
## IV. Example
- excute
```
cd $GOPATH/src/github.com/tpgzcyyao/config/test
go test
```
- you can get the result below
```
The config map is: 
map[boolean_section:map[boolean_first:false boolean_second:true] float_section:map[float_first:1.111111111111111111111111111111 float_second:2.1111111111111111111111111111111] int_section:map[int16_val:-32768 int32_val:-2147483648 int64_val:-9223372036854775808 int8_val:-128 int_val:-2147483648 uint16_val:65535 uint32_val:4294967295 uint64_val:18446744073709551615 uint8_val:255 uint_val:2147483647] ip_icmp:map[ip_raw:ip_raw ip_vxx:ip_vxx ip_xx:ip_xx ipv4:ipv4 ipv6:ipv6] string_section:map[string_first:this is string string_second:this is second = 2]]

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
                "int16_val": "-32768",
                "int32_val": "-2147483648",
                "int64_val": "-9223372036854775808",
                "int8_val": "-128",
                "int_val": "-2147483648",
                "uint16_val": "65535",
                "uint32_val": "4294967295",
                "uint64_val": "18446744073709551615",
                "uint8_val": "255",
                "uint_val": "2147483647"
        },
        "ip_icmp": {
                "ip_raw": "ip_raw",
                "ip_vxx": "ip_vxx",
                "ip_xx": "ip_xx",
                "ipv4": "ipv4",
                "ipv6": "ipv6"
        },
        "string_section": {
                "string_first": "this is string",
                "string_second": "this is second = 2"
        }
}

The config struct is: 
&{{this is string this is second = 2} {-2147483648 -128 -32768 -2147483648 -9223372036854775808 2147483647 255 65535 4294967295 18446744073709551615} {1.1111112 2.111111111111111} {false true} {ipv4 ipv6 ip_xx ip_vxx ip_raw}}
```
- the path for example config file
```
$GOPATH/src/github.com/tpgzcyyao/config/test/test.conf
```

