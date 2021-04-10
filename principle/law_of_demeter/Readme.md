# 迪米特法则（Law of Demeter, LoD）
又叫作最少知道原则（Least KnowledgePrinciple, LKP），
指一个对象应该对其他对象保持最少的了解，尽量降低类与类之间的耦合。

- 场景
    - TeamLeader每天需要查看未完成的项目任务数
    - TeamLeader指派TeamMember进行任务统计
    - TeamMember提供对Task的汇总方法, 返回未完成的任务数
- 坏的设计:
    - Leader: 我需要统计未完成任务数
    - Member: 好的, 我可以统计, 但是任务清单在哪里呢?
    - Leader: ... 我稍后给你吧
- 好的设计:
    - Leader: 我需要统计未完成任务数
    - Member: 好的. 任务清单我知道在那里, 我会搞定的
    - Leader: 好兵!




































# Refer
- https://studygolang.com/articles/33104