## 费曼笔记 ##

#### 数组和切片最重要的不同
> 数组类型的值（以下简称数组）的长度是固定的，而切片类型的值（以下简称切片）是可变长的。

> 切片看做是对数组的一层简单的封装，因为在每个切片的底层数据结构中，一定会包含一个数组。数组可以被叫做切片的底层数组，而切片也可以被看作是对数组的某个连续片段的引用。

#### 估算切片容量的增长?
1. 一旦一个切片无法容纳更多的元素，Go 语言就会想办法扩容。但它并不会改变原来的切片，而是会生成一个容量更大的切片，然后将把原有的元素和新元素一并拷贝到新切片中。在一般的情况下，你可以简单地认为新切片的容量（以下简称新容量）将会是原切片容量（以下简称原容量）的 2 倍。
2. 当原切片的长度（以下简称原长度）大于或等于1024时，Go 语言将会以原容量的1.25倍作为新容量的基准（以下新容量基准）。新容量基准会被调整（不断地与1.25相乘），直到结果不小于原长度与要追加的元素数量之和（以下简称新长度）。最终，新容量往往会比新长度大一些，当然，相等也是可能的.


#### 切片的底层数组什么时候会被替换？
1. 在无需扩容时，append函数返回的是指向原底层数组的新切片，而在需要扩容时，append函数返回的是指向新底层数组的新切片.
2. 只要新长度不会超过切片的原容量，那么使用append函数对其追加元素的时候就不会引起扩容。这只会使紧邻切片窗口右边的（底层数组中的）元素被新的元素替换掉。

#### 总结
1. 切片是基于数组的，可变长的，并且非常轻快。一个切片的容量总是固定的，而且一个切片也只会与某一个底层数组绑定在一起。
2. 切片的容量总会是在切片长度和底层数组长度之间的某一个值，并且还与切片窗口最左边对应的元素在底层数组中的位置有关系。那两个分别用减法计算切片长度和容量的方法你一定要记住。
3. 另外，append函数总会返回新的切片，而且如果新切片的容量比原切片的容量更大那么就意味着底层数组也是新的了。

#### 问题: 
1. 如果有多个切片指向了同一个底层数组，那么你认为应该注意些什么？
> 答: 当两个长度不一的切片使用同一个底层数组，并且两切片的长度均小于数组的容量时，对其中长度较小的一个切片进行append操作，但不超过底层数组容量，这时会影响长度较长切片中原来比较小切片多看到的值，因为底层数组被修改了。

2. 怎样沿用“扩容”的思想对切片进行“缩容”？请写出代码。
> 答: 可以截取切片的部分数据，然后创建新数组来缩容