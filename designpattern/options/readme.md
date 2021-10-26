# 功能选项模式（options）

扩展的构造函数和其他公共API中的可选参数

关键代码
- 定义一个Option接口，拥有一个设置参数的函数
- 定义一个optFunc实现Option接口
- 构造结构体时，接收可变类型的Option
- 遍历options，调用option中的设置参数方法

