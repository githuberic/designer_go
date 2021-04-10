# 接口隔离原则
接口隔离原则（Interface Segregation Principle, ISP）指用多个专门的接口，而不使用单一的总接口，客户端不应该依赖它不需要的接口
- 一个类对另一个类的依赖应该建立在最小接口上。
- 建立单一接口，不要建立庞大臃肿的接口。
- 尽量细化接口，接口中的方法尽量少。

# 场景
- 设计一个动物接口
- 不同动物可能有eat(), fly(), swim()等方法
- 设计实现动物接口的Bird类和Dog类

# Refer
https://studygolang.com/articles/33103