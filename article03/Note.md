## 费曼笔记 ##
#### 库源码文件学习

#### 看完本文加深的关键点:
    1.同一个文件夹下，包的声明语句需要相同，代表同一个包。
    2.包名不需要和其所在的文件夹名相同。
    3.首字母大小写来代表可见性，大写:public, 小写:private
    4.导入包时，import的是相对src的相对文件路径
    5.使用包内的函数时，其限定符是：包名.函数名(),压根与程序的文件名没有啥关系.
    4.模块级私有internal的使用姿势(这个厉害了，之前还不知道这玩意儿)


#### 从go语言的规范来看有几种避免的方法：
Import declaration Local name of Sin
    1. import "lib/math" math.Sin
    2. import m "lib/math" m.Sin
    3. import . "lib/math" Sin
    4. import _ "lib/math"