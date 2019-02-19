## 费曼笔记 ##

#### Go语言规范中map键的约束
1. 字典的键类型不可以是函数类型、字典类型和切片类型;
2. 如果键类型是接口，那么键实际类型不能是函数类型、字典类型、切片类型；否则会引起panic;
3. 如果键的类型是数组类型，那么还要确保该类型的元素类型不是函数类型、字典类型或切片类型，这里会引发编译错误;
4. 如果键的类型是结构体类型，那么还要保证其中字段的类型的合法性。不能是函数类型、字典类型或切片类型;
5. 如果键的类型是数组类型，那么还要确保该类型的元素类型不是函数类型、字典类型或切片类型; 无论埋的多深，编译器都会找出来;
6. 如果键的类型是结构体类型，那么还要保证其中字段的类型的合法性。不能是函数类型、字典类型或切片类型, 无论埋的多深，编译器都会找出来;

##### map键规范汇总
- Go 语言的字典类型其实是一个哈希表（hash table）的特定实现，在这个实现中，键和元素的最大不同在于，键的类型是受限的，而元素却可以是任意类型的。
- Go 语言规范规定，在键类型的值之间必须可以施加操作符==和!=。换句话说，键类型的值必须要支持判等操作。由于函数类型、字典类型和切片类型的值并不支持判等操作，所以字典的键类型不能是这些类型。


#### map增、删、改、查操作
##### 在nil map上操作
- 在nil map上读取不会发生panic;
- 在nil map上删除不会发生panic;
- 在nil map的长度为0;
- 在nil map生新增键值对会发生panic;


#### 注意
- 字典并不是并发安全的
- 在同一时间段内但在不同的 goroutine中对同一个值进行操作是不安全的, 需要对其进行上读写锁;


参考链接1: https://tour.golang.org/moretypes/19