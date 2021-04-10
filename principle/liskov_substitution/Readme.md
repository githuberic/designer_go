# 里氏替换原则
里氏替换原则（Liskov Substitution Principle, LSP）
- 所有引用父类的地方
- 必须能透明地使用其子类对象
- 子类对象能够替换父类对象
- 而保持程序功能不变

## 里氏替换原则的优点:
- 约束继承泛滥，是开闭原则的一种体现
- 加强程序的健壮性，同时变更时可以做到非常好的兼容性

# 场景
- 某线上动物园系统, 定义了鸟类接口IBird和NormalBird类
- IBird接口定义了鸣叫 - Tweet(), 和飞翔 - Fly()方法
- 现需要增加一种"鸟类" - 鸵鸟: 鸵鸟只会跑 - Run(), 不会飞 - Fly()

## 坏的设计:
- 新增鸵鸟类 - OstrichBird, 从NormalBird继承
- 覆盖Fly方法, 并抛出错误
- 添加Run方法
- 调用方需要修改: 判断是否OstrichBird, 是则需要特别对待
- 存在问题: OstrichBird跟NormalBird已经有较大差异, 强行继承造成很多异味

## 好的设计:
- IBird接口保留鸣叫 - Tweet()方法
- NormalBird实现IBird接口, 移除Fly方法
- 新增IFlyableBird, 继承IBird接口, 并添加Fly()方法
- 新增FlyableBird, 继承NormalBird, 并实现IFlyableBird接口
- 新增IRunnableBird, 继承IBird接口, 并添加Run()方法
- 新增OstrichBird, 继承NormalBird, 并实现IRunnableBird
- 调用方判断是IFlyableBird, 还是IRunnableBird
































# Refer
- https://studygolang.com/articles/33105