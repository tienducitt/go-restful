package goconf

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func Load(i interface{}, fc func(string) string) error {
	vi := reflect.ValueOf(i)
	if vi.Kind() != reflect.Ptr || vi.IsNil() {
		return fmt.Errorf("goconf: can not unmarshall nil or not pointer value")
	}
	v := vi.Elem()
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		tf := t.Field(i)
		vf := v.Field(i)

		// fieldName := tf.Name
		key := tf.Tag.Get("conf")
		required := tf.Tag.Get("required") == "true"
		defaultValue := tf.Tag.Get("default")
		// description := tf.Tag.Get("desc")
		if key == "" || key == "-" {
			continue
		}

		valueStr := fc(key)

		// check required & default value
		if valueStr == "" {
			if required {
				return fmt.Errorf("goconf: Config '%s' is required but not found", key)
			}
			if defaultValue != "" {
				valueStr = defaultValue
			}
		}

		// parse
		switch vf.Kind() {
		case reflect.String, reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64:
			value, err := getValue(vf.Kind(), valueStr)
			if err != nil {
				return fmt.Errorf("goconf: can not unmarshall `%s` into Go value of type %s", valueStr, vf.Kind().String())
			}
			vf.Set(value)
		case reflect.Slice:
			delim := ","
			strs := strings.Split(valueStr, delim)
			sli := reflect.MakeSlice(tf.Type, 0, len(strs))
			for _, str := range strs {
				vv, err := getValue(tf.Type.Elem().Kind(), str)
				if err != nil {
					return fmt.Errorf("goconf: can not unmarshall `%s` into Go value of type %s", str, tf.Type.Elem().Kind().String())
				}
				sli = reflect.Append(sli, vv)
			}
			vf.Set(sli)
		case reflect.Map:
			delim := ","
			strs := strings.Split(valueStr, delim)
			m := reflect.MakeMap(tf.Type)
			for _, str := range strs {
				mapStrs := strings.SplitN(str, "=", 2)
				if len(mapStrs) < 2 {
					return fmt.Errorf("goconf: %s, must be split by \"=\"", str)
				}
				vk, err := getValue(tf.Type.Key().Kind(), mapStrs[0])
				if err != nil {
					return fmt.Errorf("goconf: can not unmarshall `%s` into Go value of type %s", mapStrs[0], tf.Type.Key().Kind().String())
				}
				vv, err := getValue(tf.Type.Elem().Kind(), mapStrs[1])
				if err != nil {
					return fmt.Errorf("goconf: can not unmarshall `%s` into Go value of type %s", mapStrs[1], tf.Type.Elem().Kind().String())
				}
				m.SetMapIndex(vk, vv)
			}
			vf.Set(m)
		}
	}

	return nil
}

func parseBool(v string) bool {
	if v == "true" || v == "yes" || v == "1" || v == "y" || v == "enable" {
		return true
	} else if v == "false" || v == "no" || v == "0" || v == "n" || v == "disable" {
		return false
	} else {
		return false
	}
}

func getValue(t reflect.Kind, v string) (reflect.Value, error) {
	var vv reflect.Value
	switch t {
	case reflect.Bool:
		d := parseBool(v)
		vv = reflect.ValueOf(d)
	case reflect.Int:
		d, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return vv, err
		}
		vv = reflect.ValueOf(int(d))
	case reflect.Int8:
		d, err := strconv.ParseInt(v, 10, 8)
		if err != nil {
			return vv, err
		}
		vv = reflect.ValueOf(int8(d))
	case reflect.Int16:
		d, err := strconv.ParseInt(v, 10, 16)
		if err != nil {
			return vv, err
		}
		vv = reflect.ValueOf(int16(d))
	case reflect.Int32:
		d, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return vv, err
		}
		vv = reflect.ValueOf(int32(d))
	case reflect.Int64:
		d, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return vv, err
		}
		vv = reflect.ValueOf(int64(d))
	case reflect.Uint:
		d, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return vv, err
		}
		vv = reflect.ValueOf(uint(d))
	case reflect.Uint8:
		d, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return vv, err
		}
		vv = reflect.ValueOf(uint8(d))
	case reflect.Uint16:
		d, err := strconv.ParseUint(v, 10, 16)
		if err != nil {
			return vv, err
		}
		vv = reflect.ValueOf(uint16(d))
	case reflect.Uint32:
		d, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return vv, err
		}
		vv = reflect.ValueOf(uint32(d))
	case reflect.Uint64:
		d, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return vv, err
		}
		vv = reflect.ValueOf(uint64(d))
	case reflect.Float32:
		d, err := strconv.ParseFloat(v, 32)
		if err != nil {
			return vv, err
		}
		vv = reflect.ValueOf(float32(d))
	case reflect.Float64:
		d, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return vv, err
		}
		vv = reflect.ValueOf(float64(d))
	case reflect.String:
		vv = reflect.ValueOf(v)
	default:
		return vv, fmt.Errorf("unkown type: %s", t)
	}
	return vv, nil
}
