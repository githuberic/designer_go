package v2

import (
	"fmt"
	"testing"
)

//员工接口
type Employee struct {
	name string
	role string
}

//员工接口
type IEmployee interface {
	DoPBC()
}

// 工程师类
type Programmer struct {
	Employee
}

func NewProgrammer(name string) *Programmer {
	return &Programmer{Employee{name: name, role: "工程车"}}
}

// 工程师类实现Employee接口
func (p *Programmer) DoPBC() {
	fmt.Printf("Role=%s,name=%s,PBC=需求评审、写代码、测试、发布、处理故障...\n", p.role,p.name)
}

// TM类
type TM struct {
	Employee
}

func NewTM(name string) *TM {
	return &TM{Employee{name: name, role: "TM"}}
}

// TM类实现Employee接口
func (p *TM) DoPBC() {
	fmt.Printf("Role=%s,name=%s,,PBC=技术方案、任务分配、开会...\n", p.role, p.name)
}

// PM类
type PM struct {
	Employee
}

func NewPM(name string) *PM {
	return &PM{Employee{name: name, role: "PM"}}
}

// 老板类实现Employee接口
func (p *PM) DoPBC() {
	fmt.Printf("Role=%s,Name=%s,PBC=开会、领取任务、监督PBC落地...\n", p.role,p.name)
}

//工厂模式函数，根据传入工作不同动态返回不同类型
func Factory(role string) IEmployee {
	switch role {
	case "programmer":
		return NewProgrammer("刘**")
	case "tm":
		return NewTM("鲁**")
	case "pm":
		return NewPM("王**")
	default:
		panic("no such role")
	}
}

func TestVerifyV1(t *testing.T) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()

	person := Factory("tm")
	person.DoPBC()

	person = Factory("pm")
	person.DoPBC()
}
