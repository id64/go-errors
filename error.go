package errors

// WrapDetails wrap error with additional data.
func WrapDetails(err error, data interface{}) error {
	return errorWithDetails{cause: err, data: data}
}

// Details extract data from an error.
func Details(err error) interface{} {
	if errDetails, ok := cause(err).(errorDetails); ok {
		return errDetails.Details()
	}
	return nil
}

type errorWithDetails struct {
	cause error
	data  interface{}
}

// Error comply interface.
func (e errorWithDetails) Error() string {
	return e.cause.Error()
}

// Details returns additional info.
func (e errorWithDetails) Details() interface{} {
	return e.data
}

// Cause get internal error.
func (e errorWithDetails) Cause() error {
	return e.cause
}

type errorDetails interface {
	Error() string
	Details() interface{}
}

func cause(err error) error {
	type causer interface {
		Cause() error
	}

	var detErrs []errorDetails

	for err != nil {
		if detailed, ok := err.(errorDetails); ok {
			detErrs = append(detErrs, detailed)
		}
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}

	if len(detErrs) > 0 {
		return detErrs[len(detErrs)-1]
	}

	return err
}
