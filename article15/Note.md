## 费曼笔记 ##
### 关于指针的有限操作

#### 基础概念

##### 以下列表中的值都是不可寻址的。
* 常量的值。
* 基本类型值的字面量。
* 算术操作的结果值。
* 对各种字面量的索引表达式和切片表达式的结果值。不过有一个例外，对切片字面量的索引结果值却是可寻址的。
* 对字符串变量的索引表达式和切片表达式的结果值。
* 对字典变量的索引表达式的结果值。
* 函数字面量和方法字面量，以及对它们的调用表达式的结果值。
* 结构体字面量的字段值，也就是对结构体字面量的选择表达式的结果值。
* 类型转换表达式的结果值。
* 类型断言表达式的结果值。
* 接收表达式的结果值。

##### 其中常用的包括以下几种。(临时结果)
* 用于获得某个元素的索引表达式。
* 用于获得某个切片（片段）的切片表达式。
* 用于访问某个字段的选择表达式。
* 用于调用某个函数或方法的调用表达式。
* 用于转换值的类型的类型转换表达式。
* 用于判断值的类型的类型断言表达式。
* 向通道发送元素值或从通道那里接收元素值的接收表达式。

##### 可寻址的:
- 第一个关键词：不可变的。
- 第二个关键词：临时结果。
- 第三个关键词：不安全的。

##### 总结一下。
1. 不可变的值不可寻址。常量、基本类型的值字面量、字符串变量的值、函数以及方法的字面量都是如此。其实这样规定也有安全性方面的考虑。
2. 绝大多数被视为临时结果的值都是不可寻址的。算术操作的结果值属于临时结果，针对值字面量的表达式结果值也属于临时结果。但有一个例外，对切片字面量的索引结果值虽然也属于临时结果，但却是可寻址的。
3. 若拿到某值的指针可能会破坏程序的一致性，那么就是不安全的，该值就不可寻址。由于字典的内部机制，对字典的索引结果值的取址操作都是不安全的。另外，获取由字面量或标识符代表的函数或方法的地址显然也是不安全的。
最后说一句，如果我们把临时结果赋给一个变量，那么它就是可寻址的了。如此一来，取得的指针指向的就是这个变量持有的那个值了。


#### 转换规则
1. 一个指针值（比如*Dog类型的值）可以被转换为一个unsafe.Pointer类型的值，反之亦然。
2. 一个uintptr类型的值也可以被转换为一个unsafe.Pointer类型的值，反之亦然。
3. 一个指针值无法被直接转换成一个uintptr类型的值，反过来也是如此。
所以，对于指针值和uintptr类型值之间的转换，必须使用unsafe.Pointer类型的值作为中转。


#### 总结
可寻址的， 关于这方面你需要记住三个关键词：不可变的、临时结果和不安全的。只要一个值符合了这三个关键词中的任何一个，它就是不可寻址的。