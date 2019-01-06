## 费曼笔记 ##
#### container/list包用法
> list是一个双向链表。该结构具有链表的所有功能。

```golang
type Element
type Element struct {
        Value interface{}   //在元素中存储的值
}
```

```golang
func (e *Element) Next() *Element  //返回该元素的下一个元素，如果没有下一个元素则返回nil
func (e *Element) Prev() *Element  //返回该元素的前一个元素，如果没有前一个元素则返回nil。
type List
func New() *List //返回一个初始化的list
func (l *List) Back() *Element  //获取list l的最后一个元素
func (l *List) Front() *Element //获取list l的第一个元素
func (l *List) Init() *List     //list l初始化或者清除list l
func (l *List) InsertAfter(v interface{}, mark *Element) *Element  //在list l中元素mark之后插入一个值为v的元素，并返回该元素，如果mark不是list中元素，则list不改变。
func (l *List) InsertBefore(v interface{}, mark *Element) *Element //在list l中元素mark之前插入一个值为v的元素，并返回该元素，如果mark不是list中元素，则list不改变。
func (l *List) Len() int //获取list l的长度
func (l *List) MoveAfter(e, mark *Element)  //将元素e移动到元素mark之后，如果元素e或者mark不属于list l，或者e==mark，则list l不改变。
func (l *List) MoveBefore(e, mark *Element) //将元素e移动到元素mark之前，如果元素e或者mark不属于list l，或者e==mark，则list l不改变。
func (l *List) MoveToBack(e *Element)    //将元素e移动到list l的末尾，如果e不属于list l，则list不改变。
func (l *List) MoveToFront(e *Element)    //将元素e移动到list l的首部，如果e不属于list l，则list不改变。
func (l *List) PushBack(v interface{}) *Element  //在list l的末尾插入值为v的元素，并返回该元素。
func (l *List) PushBackList(other *List)  //在list l的尾部插入另外一个list，其中l和other可以相等。
func (l *List) PushFront(v interface{}) *Element //在list l的首部插入值为v的元素，并返回该元素。
func (l *List) PushFrontList(other *List)        //在list l的首部插入另外一个list，其中l和other可以相等。
func (l *List) Remove(e *Element) interface{}    //如果元素e属于list l，将其从list中删除，并返回元素e的值。
```


#### container/ring包用法
> ring包实现了环形链表

type Ring  //Ring类型代表环形链表的一个元素，同时也代表链表本身。环形链表没有头尾；指向环形链表任一元素的指针都可以作为整个环形链表看待。Ring零值是具有一个（Value字段为nil的）元素的链表。

```golang
type Ring struct {
     Value interface{} // 供调用者使用，本包不会对该值进行操作
     // 包含未导出字段
}
```

```golang
func New(n int) *Ring                   //创建一个长度为n的环形链表
func (r *Ring) Do(f func(interface{}))  //对链表中任意元素执行f操作，如果f改变了r，则该操作造成的后果是不可预期的。
func (r *Ring) Len() int            //求环长度，返回环中元素数量
func (r *Ring) Link(s *Ring) *Ring  //Link连接r和s，并返回r原本的后继元素r.Next()。r不能为空。
```
##### 注意
1. 如果r和s指向同一个环形链表，则会删除掉r和s之间的元素，删掉的元素构成一个子链表，返回指向该子链表的指针（r的原后继元素）；
2. 如果没有删除元素，则仍然返回r的原后继元素，而不是nil。
3. 如果r和s指向不同的链表，将创建一个单独的链表，将s指向的链表插入r后面，返回s原最后一个元素后面的元素（即r的原后继元素）。

```golang
func (r *Ring) Unlink(n int) *Ring //删除链表中n % r.Len()个元素，从r.Next()开始删除。如果n % r.Len() == 0，不修改r。返回删除的元素构成的链表，r不能为空。
func (r *Ring) Move(n int) *Ring   //返回移动n个位置（n>=0向前移动，n<0向后移动）后的元素，r不能为空。
func (r *Ring) Next() *Ring  //获取当前元素的下个元素
func (r *Ring) Prev() *Ring  //获取当前元素的上个元素
```

##### 注意
1. 向后移动n％r.Len() 元素向后移动（n <0）或向前移动（n> = 0）
2. Unlink从ring r中删除n％r.Len() 元素


#### Ring与List的区别
container/ring包中的Ring类型实现的是一个循环链表，也就是我们俗称的环。其实List在内部就是一个循环链表。它的根元素永远不会持有任何实际的元素值，而该元素的存在就是为了连接这个循环链表的首尾两端。
List的零值是一个只包含了根元素，但不包含任何实际元素值的空链表。那么，既然Ring和List在本质上都是循环链表，那它们到底有什么不同呢？

##### 最主要的不同有下面几种。
1. Ring类型的数据结构仅由它自身即可代表，而List类型则需要由它以及Element类型联合表示。这是表示方式上的不同，也是结构复杂度上的不同。
2. 一个Ring类型的值严格来讲，只代表了其所属的循环链表中的一个元素，而一个List类型的值则代表了一个完整的链表。这是表示维度上的不同。
3. 在创建并初始化一个Ring值的时候，我们可以指定它包含的元素的数量，但是对于一个List值来说却不能这样做（也没有必要这样做）。循环链表一旦被创建，其长度是不可变的。这是两个代码包中的New函数在功能上的不同，也是两个类型在初始化值方面的第一个不同。
4. 仅通过var r ring.Ring语句声明的r将会是一个长度为1的循环链表，而List类型的零值则是一个长度为0的链表。别忘了List中的根元素不会持有实际元素值，因此计算长度时不会包含它。这是两个类型在初始化值方面的第二个不同。
5. Ring值的Len方法的算法复杂度是 O(N) 的，而List值的Len方法的算法复杂度则是 O(1) 的。这是两者在性能方面最显而易见的差别。
其他的不同基本上都是方法方面的了。比如，循环链表也有用于插入、移动或删除元素的方法，不过用起来都显得更抽象一些，等等。


#### 思考

1. 思考题 container/ring包中的循环链表的适用场景都有那些？
答:
  - list可以作为queue和stack的基础数据结构
  - list典型应用场景是构造FIFO队列
  - ring可以用来保存固定数量的元素，例如保存最近100条日志，用户最近10次操作
  - ring典型应用场景是构造定长环回队列

