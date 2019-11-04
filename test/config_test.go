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
