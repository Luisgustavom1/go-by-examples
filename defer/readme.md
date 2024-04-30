## References
- https://rezakhademix.medium.com/defer-functions-in-golang-common-mistakes-and-best-practices-96eacdb551f0
- https://go.dev/blog/defer-panic-and-recover

## Use case
[json package](https://pkg.go.dev/encoding/json)

// https://go.dev/src/encoding/json/encode.go#:~:text=282-,//%20jsonError%20is%20an%20error%20wrapper%20type%20for%20internal%20use%20only.,%7D,-305%C2%A0%C2%A0
// jsonError is an error wrapper type for internal use only.
// Panics with errors are wrapped in jsonError so that the top-level recover
// can distinguish intentional panics from this package.
type jsonError struct{ error }

func (e *encodeState) marshal(v any, opts encOpts) (err error) {
	defer func() {
		if r := recover(); r != nil {
			if je, ok := r.(jsonError); ok {
				err = je.error
			} else {
				panic(r)
			}
		}
	}()
	e.reflectValue(reflect.ValueOf(v), opts)
	return nil
}

// error aborts the encoding by panicking with err wrapped in jsonError.
func (e *encodeState) error(err error) {
	panic(jsonError{err})
}
