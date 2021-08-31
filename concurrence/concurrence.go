package concurrence

import (
	"errors"
	"reflect"
	"sync"
)

// ConcurrentIterate 并发迭代
// objects must be slice type
func ConcurrentIterate(objects interface{}, f func(i int, v interface{}) error) []error {
	var (
		wg   sync.WaitGroup
		lock sync.Mutex
		errs []error
	)
	keyArr, ok := interfaceToSlice(objects)
	if !ok {
		return []error{errors.New("objects param type invalid")}
	}
	// 并发执行
	for i, v := range keyArr {
		v1 := v
		i1 := i
		wg.Add(1)
		GoWithDefaultRecovery(func() {
			var err error
			defer func() {
				if e := recover(); e != nil {
					err = errors.New("ConcurrentIterate func panic")
				}
				if err != nil {
					lock.Lock()
					errs = append(errs, err)
					lock.Unlock()
				}
				wg.Done()
			}()
			err = f(i1, v1)
		})
	}
	wg.Wait()
	return errs
}

// ConcurrentExecMultiFunc 并发执行方法列表
func ConcurrentExecMultiFunc(funcList ...func() error) []error {
	var (
		lock    sync.Mutex
		errList []error
		wg      sync.WaitGroup
	)
	for _, v := range funcList {
		wg.Add(1)
		f := v
		GoWithDefaultRecovery(func() {
			var err error
			defer func() {
				if e := recover(); e != nil {
					err = errors.New("ConcurrentExecMultiFunc func panic")
				}
				if err != nil {
					lock.Lock()
					errList = append(errList, err)
					lock.Unlock()
				}
				wg.Done()
			}()
			err = f()
		})
	}
	wg.Wait()
	return errList
}

func interfaceToSlice(i interface{}) ([]interface{}, bool) {
	s := reflect.ValueOf(i)
	if s.Kind() != reflect.Slice {
		return nil, false
	}
	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret, true
}
