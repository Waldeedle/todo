package components

import (
	"github.com/a-h/templ"
)

type ToastType int

const (
	ToastTypeSuccess ToastType = iota + 1
	ToastTypeError
)

type ToastComponentProps struct {
	Message   string
	Type      ToastType
	className string
}

func Toast(props ToastComponentProps) templ.Component {
	if props.Type == ToastTypeSuccess {
		props.className = "alert alert-info"
		props.Message = "Successfully Created!"
	} else if props.Type == ToastTypeError {
		props.className = "alert alert-error"
	}
	return toast(props)
}
