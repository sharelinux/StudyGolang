## 费曼笔记 ##

#### 源代码文件:
    1. 命令源码文件
    2. 库源码文件
    3. 测试源码文件

#### 命令源码文件的用途:
    命令源码文件是程序的运行入口，是每个可独立运行的程序必须拥有的。我们可以通过构建或安装生成与其对应的可执行文件，后者一般会与该命令源码文件的直接父目录同名。

#### 命令源码文件如何编写:
    如果一个源码文件声明属于包，并且包含一个无参数声明且无结果声明的main函数，那么就是命令源码文件。

    当需要模块化编程时，我们往往会将代码拆分到多个文件，甚至拆分到不同的代码包中。但无论怎样，对于一个独立的程序来说，命令源码文件永远只会也只能有一个。如果有与命令源码文件同包的源码文件，那么它们也应该声明属于.


#### flag官方文档
https://golang.google.cn/pkg/flag/