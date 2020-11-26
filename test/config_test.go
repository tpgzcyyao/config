package test

import (
    "encoding/json"
    "fmt"
    "os"
    "testing"

    "github.com/tpgzcyyao/config"
)

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
    IntVal    int
    Int8Val   int8
    Int16Val  int16
    Int32Val  int32
    Int64Val  int64
    UintVal   uint
    Uint8Val  uint8
    Uint16Val uint16
    Uint32Val uint32
    Uint64Val uint64
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

func TestLoadFile(t *testing.T) {
    file := "test.conf"
    str, _ := os.Getwd()
    resMap, err := (new(config.Config)).LoadFile(str + "/" + file)
    if err != nil {
        t.Errorf("LoadFile error: %v", err)
    }
    fmt.Println(fmt.Sprintf("The config map is: \n%v\n", resMap))
    strJson, err := json.MarshalIndent(resMap, "", "\t")
    if err != nil {
        t.Errorf("json.Marshl error: %v", err)
    }
    fmt.Println(fmt.Sprintf("The config map in json format is: \n%s\n", strJson))
}

func TestLoadConfig(t *testing.T) {
    file := "test.conf"
    str, _ := os.Getwd()
    resConfig := new(ResConfig)
    err := (new(config.Config)).LoadConfig(str+"/"+file, resConfig)
    if err != nil {
        t.Errorf("LoadConfig error: %v", err)
    }
    fmt.Println(fmt.Sprintf("The config struct is: \n%v\n", resConfig))
}

func TestIpIcmp(t *testing.T) {
    file := "test.conf"
    str, _ := os.Getwd()
    resConfig := new(ResConfig)
    err := (new(config.Config)).LoadConfig(str+"/"+file, resConfig)
    if err != nil {
        t.Errorf("LoadConfig error: %v", err)
    }
    if resConfig.IPICMP.IPv4 != "ipv4" {
        t.Errorf("IPv4: actual %s, expected %s.", resConfig.IPICMP.IPv4, "ipv6")
    }
    if resConfig.IPICMP.IPv6 != "ipv6" {
        t.Errorf("IPv6: actual %s, expected %s.", resConfig.IPICMP.IPv6, "ipv6")
    }
    if resConfig.IPICMP.IPxx != "ip_xx" {
        t.Errorf("IPxx: actual %s, expected %s.", resConfig.IPICMP.IPxx, "ip_xx")
    }
    if resConfig.IPICMP.IPVxx != "ip_vxx" {
        t.Errorf("IPVxx: actual %s, expected %s.", resConfig.IPICMP.IPVxx, "ip_vxx")
    }
    if resConfig.IPICMP.IpRaw != "ip_raw" {
        t.Errorf("IpRaw: actual %s, expected %s.", resConfig.IPICMP.IpRaw, "ip_raw")
    }
}
