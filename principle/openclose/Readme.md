# 开闭原则
开闭原则（Open-Closed Principle, OCP）
- 指一个软件实体如类、模块和函数应该对扩展开放，对修改关闭。所谓开闭，也正是对扩展和修改两个行为的一个原则
- 实现开闭原则的核心思想就是面向抽象编程。

# 场景
- 某文玩电商平台, 提供系列文玩拍品(接口: IProduct)
- 每个拍品有id,name,price等属性
- 现在平台搞促销, 直播拍品(LivedProduct)分享10人后，可以打8折  
- 如何上架打折拍品? 是直接修改原拍品的价格, 还是增加折后拍品?

# 思路
- 开闭原则, 就是尽量避免修改, 改以扩展的方式, 实现系统功能的增加
- 增加"优惠折扣"接口 - IDiscount
- 增加"折后拍品" - DiscountedProduct, 同时实现拍品接口和折扣接口
- DiscountedProduct继承自Product, 添加实现折扣接口, 并覆盖IProduct.price()方法





















# Refer
- https://studygolang.com/articles/33100