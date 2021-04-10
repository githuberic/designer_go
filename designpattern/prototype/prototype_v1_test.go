package prototype

import (
	"fmt"
	"testing"
)

func TestVerifyV1(t *testing.T)  {
	ComputerStart4()
}

func ComputerStart1() {
	Pc1 := Computer{
		SystemName: "Windows",
		UseNumber:  1000,
		Memory:     Memory{Count: 4, MemorySize: []int{32, 32, 32, 32}},
		Fan:        map[string]FanSpeed{"left": {2500}, "right": {2000}},
		Money:      Money{123.45},
	}

	//浅拷贝
	Pc2:=Pc1
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", Pc1, Pc2)
	//修改切片内容以及map信息影响Pc1
	Pc2.SystemName ="MacOs"
	Pc2.UseNumber =100
	Pc2.Memory.Count =2
	Pc2.Memory.MemorySize[0]=8
	Pc2.Memory.MemorySize[1]=8
	Pc2.Memory.MemorySize[2]=0
	Pc2.Memory.MemorySize[3]=0
	Pc2.Fan["left"]=FanSpeed{2000}
	Pc2.Fan["right"]=FanSpeed{1500}
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", Pc1, Pc2)
}

func ComputerStart2() {
	Pc1 := Computer{
		SystemName: "Windows",
		UseNumber:  1000,
		Memory:     Memory{Count: 4, MemorySize: []int{32, 32, 32, 32}},
		Fan:        map[string]FanSpeed{"left": {2500}, "right": {2000}},
		Money:      Money{123.45},
	}

	//浅拷贝
	Pc2:=Pc1
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", Pc1, Pc2)

	//切片以及map新空间互不影响
	Pc2.SystemName ="MacOs"
	Pc2.UseNumber =100
	Pc2.Memory =Memory{Count: 2, MemorySize: []int{8, 8}}
	Pc2.Fan =map[string]FanSpeed{"left": {2000}, "right": {1500}}
	Pc2.Money =Money{1000.45}
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", Pc1, Pc2)
}

func ComputerStart3() {
	Pc1 := Computer{
		SystemName: "Windows",
		UseNumber:  1000,
		Memory:     Memory{Count: 4, MemorySize: []int{32, 32, 32, 32}},
		Fan:        map[string]FanSpeed{"left": {2500}, "right": {2000}},
		Money:      Money{123.45},
	}

	//浅拷贝
	Pc2:=Pc1
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", Pc1, Pc2)

	ModifyCat(Pc2)
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", Pc1, Pc2)
}

func ComputerStart4() {
	Pc1 := &Computer{
		SystemName: "Windows",
		UseNumber:  1000,
		Memory:     Memory{Count: 4, MemorySize: []int{32, 32, 32, 32}},
		Fan:        map[string]FanSpeed{"left": {2500}, "right": {2000}},
		Money:      Money{123.45},
	}

	//深拷贝
	Pc2:= new(Computer)
	if err:= deepCopy(Pc2,Pc1);err!=nil{
		panic(err.Error())
	}
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", Pc1, Pc2)

	ModifyCat(*Pc2)
	fmt.Printf("PcInfo Pc1:%v, Pc2:%v\n", Pc1, Pc2)
}

func ModifyCat(pc Computer) {
	fmt.Printf("PcInfo Pc1:%v\n", pc)
	pc.SystemName ="MacOs"
	pc.UseNumber =100
	pc.Memory.Count =2
	pc.Memory.MemorySize[0]=8
	pc.Memory.MemorySize[1]=8
	pc.Memory.MemorySize[2]=0
	pc.Memory.MemorySize[3]=0
	pc.Fan["left"]=FanSpeed{2000}
	pc.Fan["right"]=FanSpeed{1500}
	fmt.Printf("PcInfo Pc1:%v\n", pc)
}
