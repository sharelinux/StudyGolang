## 费曼笔记 ##
#### 判断一个变量的类型

##### "类型断言"表达式:
value, ok := interface{}(container).([]string)

##### 类型断言:
1. 先把contarner的变量转化为空接口的值;
2. 一个用于判断前者的类型是否为切片类型.([]string)

##### 断言表达式的结果可以被赋值给两个变量value, ok,;
1. ok是布尔（bool）类型的，它将代表类型判断的结果，true或false;
2. 当ok值是true，那么被判断的值将会被自动转换为[]string类型的值，并赋给变量value，否则value将被赋予nil（即“空”）。

##### 类型断言说明
类型断言表达式的语法形式是x.(T)。其中的x代表要被判断类型的那个值。这个值当下的类型必须是接口类型的，不过具体是哪个接口类型其实是无所谓的。
所以，当这里的container变量类型不是任何的接口类型时，我们就需要先把它转成某个接口类型的值。


#### 类型转化

> 语法形式是T(x)
1. 其中的x可以是一个变量，也可以是一个代表值的字面量（比如1.23和struct{}），还可以是一个表达式。
2. x可以被叫做源值，它的类型就是源类型，而那个T代表的类型就是目标类型。

#### 类型转化规则注意哪些:
1. 首先，对于整数类型值、整数常量之间的类型转换，原则上只要源值在目标类型的可表示范围内就是合法的。
    - 整数在 Go 语言以及计算机中都是以补码的形式存储的。
    - 补码其实就是原码个位求反再加 1。
    - 一定要记住，当整数值的类型的有效范围由宽变窄时，只需在补码形式下截掉一定数量的高位二进制数即可。
    - 当把一个浮点数类型的值转换为整数类型值时，前者的小数部分会被全部截掉。
2. 第二，虽然直接把一个整数值转换为一个string类型的值是可行的，但值得关注的是，被转换的整数值应该可以代表一个有效的 Unicode 代码点，否则转换的结果将会是"�"（仅由高亮的问号组成的字符串值）。
3. 第三个知识点是关于string类型与各种切片类型之间的互转的。

注意: 正数的补码等于原码，负数的补码才是反码＋1

#### 类型转化
srcStr := "你好"
fmt.Printf("The string : %q\n", srcStr)                                // 字符串
fmt.Printf("The hex of %q: %x\n", srcStr, srcStr)                      // 十六进制
fmt.Printf("The byte slice of %q: % X\n", srcStr, []byte(srcStr))      // 字符串转化为byte类型切片
fmt.Printf("The string: %q\n", string([]byte{'\xe4', '\xbd', '\xa0'})) // byte切片转化为字符串
fmt.Printf("The rune slice of %q: %U\n", srcStr, []rune(srcStr))       // 字符串转为rune类型切片
fmt.Printf("The string: %q\n", string([]rune{'\u4F60', '\u597D'}))     // rune切片转化为字符串

#### 别名类型、类型再定义与潜在类型

- 别名类型
    type MyString = string
    MyString、string类型相同
    
- 类型再定义
    type MyString2 string
    MyString2和string是不同的类型
    
- 潜在类型
    type MyStrings []MyString2
    1. 对于这里的类型再定义来说，string可以被称为MyString2的潜在类型。
    2. 潜在类型的含义是某个类型在本质上是哪个类型或者是哪个类型的集合。
    3. 潜在类型相同的不同类型的值之间是可以进行类型转换的。