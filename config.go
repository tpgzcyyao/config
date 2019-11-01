package config

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Config struct {
}

var (
	matchAllCap  = regexp.MustCompile(`([a-z0-9])([A-Z])`)
	matchSection = regexp.MustCompile(`(?:^\[)(\w+)(?:\]$)`)
)

/**
 * load config file
 * output map
 */
func (c *Config) LoadFile(path string) (map[string]map[string]string, error) {
	// open the file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// read by line
	result := make(map[string]map[string]string, 0)
	buffer := bufio.NewReader(file)
	section := ""
	for {
		lineBytes, _, err := buffer.ReadLine()
		if io.EOF == err {
			break
		}
		if err != nil {
			return nil, err
		}
		// comment in line
		strRealList := strings.SplitN(string(lineBytes), "#", 2)
		line := strings.Trim(strRealList[0], " ")
		// empty line
		if "" == line {
			continue
		}
		// section name line
		resRegex := matchSection.FindStringSubmatch(line)
		if len(resRegex) == 2 {
			section = resRegex[1]
			sectionMap := make(map[string]string, 0)
			result[section] = sectionMap
			continue
		}
		// common line
		if "" == section {
			continue
		}
		s := strings.SplitN(line, "=", 2)
		if len(s) < 2 {
			continue
		}
		result[section][strings.Trim(s[0], " ")] = strings.Trim(s[1], " ")
	}
	return result, nil
}

/**
 * load config file
 * output struct
 */
func (c *Config) LoadConfig(path string, v interface{}) error {
	// get config map
	resMap, err := c.LoadFile(path)
	if err != nil {
		return err
	}
	// resolve struct
	typ0 := reflect.TypeOf(v)
	if typ0.Kind() != reflect.Ptr {
		return errors.New("cannot map to non-pointor struct")
	}
	ele0 := reflect.ValueOf(v).Elem()
	typ0 = ele0.Type()
	len0 := typ0.NumField()
	// sections
	for i := 0; i < len0; i++ {
		val1 := ele0.Field(i)
		name1 := typ0.Field(i).Name
		typ1 := val1.Type()
		key1 := c.CamelToSnake(name1)
		len1 := typ1.NumField()
		// items in one section
		for j := 0; j < len1; j++ {
			val2 := val1.Field(j)
			name2 := typ1.Field(j).Name
			typ2 := typ1.Field(j).Type
			key2 := c.CamelToSnake(name2)
			if _, ok := resMap[key1][key2]; !ok {
				continue
			}
			valStr := resMap[key1][key2]
			switch typ2.Kind() {
			case reflect.String:
				val2.SetString(valStr)
			case reflect.Bool:
				valBool, err := strconv.ParseBool(valStr)
				if err != nil {
					return err
				}
				val2.SetBool(valBool)
			case reflect.Int:
				valInt, err := strconv.ParseInt(valStr, 10, 32)
				if err != nil {
					return err
				}
				val2.SetInt(valInt)
			case reflect.Int64:
				valInt, err := strconv.ParseInt(valStr, 10, 64)
				if err != nil {
					return err
				}
				val2.SetInt(valInt)
			case reflect.Uint:
				valUint, err := strconv.ParseUint(valStr, 10, 32)
				if err != nil {
					return err
				}
				val2.SetUint(valUint)
			case reflect.Uint64:
				valUint, err := strconv.ParseUint(valStr, 10, 64)
				if err != nil {
					return err
				}
				val2.SetUint(valUint)
			case reflect.Float32:
				valFloat, err := strconv.ParseFloat(valStr, 32)
				if err != nil {
					return err
				}
				val2.SetFloat(valFloat)
			case reflect.Float64:
				valFloat, err := strconv.ParseFloat(valStr, 64)
				if err != nil {
					return err
				}
				val2.SetFloat(valFloat)
			default:
				return errors.New(fmt.Sprintf("unsupport kind: %v", typ2.Kind()))
			}
		}

	}
	return nil
}

/**
 * format string from camel model to snake model
 * e.g. input: SectionNameXXXX output: section_name_xxxx
 */
func (c *Config) CamelToSnake(camel string) string {
	snake := matchAllCap.ReplaceAllString(camel, "${1}_${2}")
	return strings.ToLower(snake)
}
