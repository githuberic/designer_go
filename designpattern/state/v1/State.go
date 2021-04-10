package v1

type State interface {
	DoAction(context *Context)
	ToString() string
}
