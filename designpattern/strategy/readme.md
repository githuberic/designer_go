# 策略模式 (strategy)
意图: 定义一系列的算法,把它们一个个封装起来, 并且使它们可相互替换。

关键代码: 实现同一个接口

### 应用实例:
- 主题的更换，每个主题都是一种策略
- 旅行的出游方式，选择骑自行车、坐汽车，每一种旅行方式都是一个策略。

## 步骤
- 首先定义接口，所有的策略都是基于一套标准，这样策略(类)才有可替换性。声明一个计算策略接口
- 接着两个接口实现类。复习golang语言实现接口是非侵入式设计。
- 声明一个策略类。
- 客户端调用