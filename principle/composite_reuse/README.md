# 合成复用原则
合成复用原则（Composite/Aggregate Reuse Principle, CARP）指尽量使用对象组合（has-a）或对象聚合（contanis-a）的方式实现代码复用，而不是用继承关系达到代码复用的目的。

合成复用原则可以使系统更加灵活，降低类与类之间的耦合度，一个类的变化对其他类造成的影响相对较小。

- 继承，又被称为白箱复用，相当于把所有实现细节暴露给子类。
- 组合/聚合又被称为黑箱复用，对类以外的对象是无法获取实现细节的。


# 场景
某订单业务系统, 需要连接数据库对产品信息进行CRUD操作

## 不好的设计:
 - 定义DBConnection类, 实现对数据库的连接和SQL执行
 - 定义ProductDAO类, 继承DBConnection类, 并封装对产品资料的增删改查

问题: ProductDAO对DBConnection的继承仅仅是为了代码复用, 不符合合成复用原则

## 更好的设计:
- 定义IDBConnection接口
- 定义MysqlConnection类, 实现对mysql数据库的连接和SQL执行
- 定义ProductDAO类, 通过Setter方法注入IDBConnection实例






























// https://studygolang.com/articles/33106