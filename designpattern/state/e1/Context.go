package e1

type Context struct {
	state State
}

//NewContext 实例化状态保存类
func NewContext() *Context {
	return &Context{
		state: nil,
	}
}

//SetState 设置状态保存类当前的状态
func (c *Context) SetState(s State) {
	c.state = s
}

//GetState 获取状态保存类当前的状态
func (c *Context) GetState() State {
	return c.state
}
