package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

/**
实例方法被调用时,实例成员进行值复制copy
*/
func (e Employee) StringV2() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d", e.Id, e.Name, e.Age)
}

/**
避免了内存拷贝
*/
func (e *Employee) String() string {
	// 查看地址是否有被复制
	fmt.Printf("Address is %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s/Name:%s/Age:%d", e.Id, e.Name, e.Age)
}

func TestVerifyV1(t *testing.T)  {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	t.Log(e.String())
}

func TestCreateEmployeeObj(t *testing.T) {
	e0 := Employee{"11007", "Bob", 20}
	t.Log(e0)
	t.Logf("e is %T", e0)
	t.Logf("e is %T", &e0)

	e1 := Employee{Id:"11008",Name: "Mike", Age: 30}
	t.Log(e1)
	t.Log(e1.Id)

	e2 := new(Employee) //返回指针
	e2.Id = "11009"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e2)
	t.Logf("e2 is %T", e2)
}


