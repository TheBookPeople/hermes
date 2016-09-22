package hermes

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"reflect"
)

func validStruct(s interface{}) error {
	t := reflect.TypeOf(s)
	k := t.Kind()
	v := reflect.ValueOf(s)
	if k != reflect.Struct {
		panic(fmt.Sprintf("Expected Struct, found %v.", k))
	}
	for i := 0; i < v.NumField(); i++ {
		//		k := v.Field(i).Kind()
		f := v.Field(i)
		if f.CanInterface() {
			err := valid(f.Interface())
			if err != nil {
				return err
			}
		} else {
			//			fmt.Println("skipping field", f)
		}
	}
	_, err := govalidator.ValidateStruct(s)
	return err
}

func validPtr(p interface{}) error {
	t := reflect.TypeOf(p)
	k := t.Kind()
	if k != reflect.Ptr {
		panic(fmt.Sprintf("Expected Ptr, found %v.", k))
	}

	//	fmt.Printf("validPtr(%v)\n", p)
	ptrVal := reflect.ValueOf(p)
	v := ptrVal.Elem()
	if ptrVal.IsNil() {
		return nil
	}

	return validStruct(v.Interface())
	/*
		fmt.Println(v.NumField())
		for i := 0; i < v.NumField(); i++ {
			//		k := v.Field(i).Kind()
			f := v.Field(i)
			if f.CanInterface() {
				fmt.Println("ptr processing field", f)
				err := valid(f.Interface())
				if err != nil {
					return err
				}
			} else {
				fmt.Println("ptr skipping field", f)
			}
		}
		color.Cyan(fmt.Sprintf("here2 %v %v %v %v", t, k, v, ptrVal))
		_, err := govalidator.ValidateStruct(v.Interface())
		return err
	*/
}
func validSlice(s interface{}) error {
	k := reflect.TypeOf(s).Kind()
	if k != reflect.Slice {
		panic(fmt.Sprintf("Expected slice, found %v.", k))
	}
	//	fmt.Println("slice")
	sv := reflect.ValueOf(s)
	for i := 0; i < sv.Len(); i++ {
		iv := sv.Index(i)
		//		fmt.Println("slice index value", iv, reflect.TypeOf(iv), iv.Kind(), iv.Interface())
		err := valid(iv.Interface())
		if err != nil {
			return err
		}
	}
	return nil
}

func valid(i interface{}) error {
	t := reflect.TypeOf(i)
	//v := reflect.ValueOf(i)
	k := t.Kind()
	switch k {
	case reflect.Ptr:
		//		fmt.Println("Ptr:", i)
		return validPtr(i)
	case reflect.Struct:
		//		fmt.Println("struct:", t, k, v)
		return validStruct(i)
	case reflect.Slice:
		//		fmt.Println("slice:", t, k, v)
		return validSlice(i)
	default:
		return nil
	}
}
