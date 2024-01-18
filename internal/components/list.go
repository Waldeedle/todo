package components

import (
	"github.com/a-h/templ"
)

type ListComponentProps struct {
	Items        []string
	ToastMessage string
	ToastType    ToastType
}

func List(props ListComponentProps) templ.Component {
	return list(props)
}
