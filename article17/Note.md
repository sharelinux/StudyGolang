## 费曼笔记 ##
### go语句及其执行规则 (下)

#### 基础概念

##### 空结构体类型
声明通道sign的时候是以chan struct{}作为其类型的.它代表了既不包含任何字段也不拥有任何方法的空结构体类型。

struct{}类型值的表示法只有一个，即：struct{}{}。并且，它占用的内存空间是0字节。确切地说，这个值在整个 Go 程序中永远都只会存在一份。虽然我们可以无数次地使用这个值字面量，但是用到的却都是同一个值。

仅仅把通道当作传递某种简单信号的介质的时候，用struct{}作为其元素类型是再好不过的了;


##### 思考问题 
问题 1：怎样才能让主 goroutine 等待其他 goroutine？
答：

  方法1: 最简单粗暴的办法就是让主 goroutine“小睡”一会儿。
    
  > `time.Sleep(time.Millisecond * 500)`

  方法2: 使用通道;
  - 先创建一个通道，它的长度应该与我们手动启用的 goroutine 的数量一致。在每个手动启用的 goroutine即将运行完毕的时候，我们都要向该通道发送一个值。
  
  - 注意，这些发送表达式应该被放在它们的go函数体的最后面。对应的，我们还需要在main函数的最后从通道接收元素值，接收的次数也应该与手动启用的 goroutine 的数量保持一致。
  
  方法3: 使用sync.WaitGroup()


问题 2：怎样让我们启用的多个 goroutine 按照既定的顺序运行？

答：设置一个全局变量counter，判断counter是否与参数i相同，如果相同就调用fn函数，然后把conter++，显示的退出当前循环；具体如代码: demo40.go

trigger函数实现了一种自旋（spinning）。除非发现条件已满足，否则它会不断地进行检查。

参考：https://studygolang.com/articles/16480

参考：https://zhuanlan.zhihu.com/p/27608263