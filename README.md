# go-errors

Package go-errors provides details error wrapping.

## Adding details to an error

```go
	details := map[string]string{
		"name": "required",
	}
	
	err := errors.WrapDetails(originalError, details)
```


## Retrieving details of an error
```go
    details := Details(err)
```
