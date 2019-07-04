package errors

import (
	"errors"
	"testing"

	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestWrapWithDetails(t *testing.T) {

	var originalErr = errors.New("oops")

	details := map[string]string{
		"name": "required",
	}

	t.Run("simple wrapped", func(t *testing.T) {
		wrappedErr := WrapDetails(originalErr, details)
		detailsUnwrapped := Details(wrappedErr)
		assert.Equal(t, "oops", wrappedErr.Error())
		assert.Equal(t, details, detailsUnwrapped)
		assert.Equal(t, originalErr, pkgerrors.Cause(wrappedErr))
	})

	t.Run("wrap wrapped with pkg", func(t *testing.T) {

		wrappedErr := pkgerrors.Wrap(originalErr, "pkg wrap")
		wrappedWrappedErr := WrapDetails(wrappedErr, details)

		detailsUnwrapped := Details(wrappedWrappedErr)
		assert.Equal(t, "pkg wrap: oops", wrappedWrappedErr.Error())
		assert.Equal(t, details, detailsUnwrapped)
		assert.Equal(t, originalErr, pkgerrors.Cause(wrappedErr))
	})

	t.Run("wrap wrapped wrapped", func(t *testing.T) {

		wrappedErr := pkgerrors.Wrap(originalErr, "pkg wrap")
		wrappedErr = WrapDetails(wrappedErr, details)
		wrappedErr = pkgerrors.Wrap(wrappedErr, "pkg wrap")

		detailsUnwrapped := Details(wrappedErr)
		assert.Equal(t, "pkg wrap: pkg wrap: oops", wrappedErr.Error())
		assert.Equal(t, details, detailsUnwrapped)
		assert.Equal(t, originalErr, pkgerrors.Cause(wrappedErr))
	})

	t.Run("unwrap the deepest details", func(t *testing.T) {

		deepDetails := map[string]string{
			"name": "can't be emtpy",
		}
		wrappedErr := pkgerrors.Wrap(originalErr, "pkg wrap")
		wrappedErr = WrapDetails(wrappedErr, deepDetails)
		wrappedErr = pkgerrors.Wrap(wrappedErr, "pkg wrap")
		wrappedErr = WrapDetails(wrappedErr, details)

		detailsUnwrapped := Details(wrappedErr)
		assert.Equal(t, "pkg wrap: pkg wrap: oops", wrappedErr.Error())
		assert.Equal(t, deepDetails, detailsUnwrapped)
		assert.Equal(t, originalErr, pkgerrors.Cause(wrappedErr))
	})

	t.Run("no details wrapped", func(t *testing.T) {
		assert.Equal(t, nil, Details(originalErr))
	})

	t.Run("nil err", func(t *testing.T) {
		assert.Equal(t, nil, Details(nil))
	})
}
