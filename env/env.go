package env

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	tag = "env"
)

var (
	ErrInputIsNil         = errors.New("input is nil")
	ErrInputPointerIsZero = errors.New("input pointer is zero")
	ErrNonStructInput     = errors.New("non struct input")
)

func Parse(input interface{}) error {
	if input == nil {
		return ErrInputIsNil
	}

	inputPointer := reflect.ValueOf(input)
	if inputPointer.IsZero() {
		return ErrInputPointerIsZero
	}

	inputValue := reflect.Indirect(inputPointer)

	inputType := inputValue.Type()
	if inputType.Kind() != reflect.Struct {
		return ErrNonStructInput
	}

	for i := 0; i < inputValue.NumField(); i++ {
		inputTypeField := inputType.Field(i)

		inputField := inputValue.Field(i)
		if !inputField.CanSet() {
			return fmt.Errorf("can not set `%s` field", inputTypeField.Name)
		}

		env, ok := inputTypeField.Tag.Lookup(tag)
		if !ok || env == "" {
			continue
		}

		value, ok := os.LookupEnv(env)
		if !ok {
			continue
		}

		switch inputField.Interface().(type) {
		case string:
			inputField.SetString(value)

		case []string:
			inputField.Set(reflect.ValueOf(strings.Split(value, ",")))

		case *string:
			inputField.Set(reflect.ValueOf(&value))

		case []*string:
			strs := strings.Split(value, ",")
			val := make([]*string, len(strs))

			for i := range strs {
				val[i] = &strs[i]
			}

			inputField.Set(reflect.ValueOf(val))

		case int, int64:
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("field `%s` failed to parse `%s` int: %w", inputTypeField.Name, value, err)
			}

			if inputField.OverflowInt(v) {
				return fmt.Errorf("field `%s` is oveflown by %d", inputTypeField.Name, v)
			}

			inputField.SetInt(v)

		case *int:
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("field `%s` failed to parse `%s` int: %w", inputTypeField.Name, value, err)
			}

			if !inputField.IsZero() && inputField.Elem().OverflowInt(v) {
				return fmt.Errorf("field `%s` is oveflown by %d", inputTypeField.Name, v)
			}

			i := int(v)
			inputField.Set(reflect.ValueOf(&i))

		case *int64:
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Errorf("field `%s` failed to parse `%s` int: %w", inputTypeField.Name, value, err)
			}

			if !inputField.IsZero() && inputField.Elem().OverflowInt(v) {
				return fmt.Errorf("field `%s` is oveflown by %d", inputTypeField.Name, v)
			}

			inputField.Set(reflect.ValueOf(&v))

		case uint, uint64:
			v, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return fmt.Errorf("field `%s` failed to parse `%s` uint: %w", inputTypeField.Name, value, err)
			}

			if inputField.OverflowUint(v) {
				return fmt.Errorf("field `%s` is oveflown by %d", inputTypeField.Name, v)
			}

			inputField.SetUint(v)

		case *uint:
			v, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return fmt.Errorf("field `%s` failed to parse `%s` uint: %w", inputTypeField.Name, value, err)
			}

			if !inputField.IsZero() && inputField.Elem().OverflowUint(v) {
				return fmt.Errorf("field `%s` is oveflown by %d", inputTypeField.Name, v)
			}

			i := uint(v)
			inputField.Set(reflect.ValueOf(&i))

		case *uint64:
			v, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return fmt.Errorf("field `%s` failed to parse `%s` uint: %w", inputTypeField.Name, value, err)
			}

			if !inputField.IsZero() && inputField.Elem().OverflowUint(v) {
				return fmt.Errorf("field `%s` is oveflown by %d", inputTypeField.Name, v)
			}

			inputField.Set(reflect.ValueOf(&v))

		case bool:
			v, err := strconv.ParseBool(value)
			if err != nil {
				return fmt.Errorf("field `%s` failed to parse `%s` bool: %w", inputTypeField.Name, value, err)
			}

			inputField.SetBool(v)

		case *bool:
			v, err := strconv.ParseBool(value)
			if err != nil {
				return fmt.Errorf("field `%s` failed to parse `%s` bool: %w", inputTypeField.Name, value, err)
			}

			inputField.Set(reflect.ValueOf(&v))

		case time.Duration:
			v, err := time.ParseDuration(value)
			if err != nil {
				return fmt.Errorf("field `%s` failed to parse `%s` duration: %w", inputTypeField.Name, value, err)
			}

			inputField.Set(reflect.ValueOf(v))

		case *time.Duration:
			v, err := time.ParseDuration(value)
			if err != nil {
				return fmt.Errorf("field `%s` failed to parse `%s` duration: %w", inputTypeField.Name, value, err)
			}

			inputField.Set(reflect.ValueOf(&v))
		}
	}

	return nil
}
