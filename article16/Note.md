## 费曼笔记 ##
### go语句及其执行规则（上）

#### 基础概念

不要通过共享数据来通讯，恰恰相反，要以通讯的方式共享数据。

通道（也就是 channel）类型的值，可以被用来以通讯的方式共享数据。更具体地说，它一般被用来在不同的 goroutine 之间传递数据。


#### 进程与线程
进程，描述的就是程序的执行过程，是运行着的程序的代表。

线程总是在进程之内的，它可以被视为进程中运行着的控制流（或者说代码执行的流程）。

一个进程至少会包含一个线程。如果一个进程只包含了一个线程，那么它里面的所有代码都只会被串行地执行。每个进程的第一个线程都会随着该进程的启动而被创建，它们可以被称为其所属进程的主线程。

如果一个进程中包含了多个线程，那么其中的代码就可以被并发地执行。除了进程的第一个线程之外，其他的线程都是由进程中已存在的线程创建出来的。

主线程之外的其他线程都只能由代码显式地创建和销毁。这需要我们在编写程序的时候进行手动控制，操作系统以及进程本身并不会帮我们下达这样的指令，它们只会忠实地执行我们的指令。


#### Goroutine
Go 语言的运行时（runtime）系统会帮助我们自动地创建和销毁系统级的线程。

用户级线程的创建、销毁、调度、状态变更以及其中的代码和数据都完全需要我们的程序自己去实现和处理。

- goroutine 优势
  1. 创建和销毁速度很快;
  2. 调度、控制容易;
  
- goroutine 劣势
  1. 复杂度高;
  2. 全权负责与用户线程所有的实现;

#### Go调度器

调度器是 Go 语言运行时系统的重要组成部分，它主要负责统筹调配 Go 并发编程模型中的三个主要元素，即：G（goroutine 的缩写）、P（processor 的缩写）和 M（machine 的缩写）。

- M 指代的就是系统级线程
- P 指的是一种可以承载若干个 G，且能够使这些 G 适时地与 M 进行对接，并得到真正运行的中介
- G 仅相当于为需要并发执行代码片段服务的上下文环境而已


G 和 M 由于 P 的存在可以呈现出多对多的关系。当一个正在与某个 M 对接并运行着的 G，需要因某个事件（比如等待 I/O 或锁的解除）而暂停运行的时候，调度器总会及时地发现，并把这个 G 与那个 M 分离开，以释放计算资源供那些等待运行的 G 使用。

当一个 G 需要恢复运行的时候，调度器又会尽快地为它寻找空闲的计算资源（包括 M）并安排运行。另外，当 M 不够用时，调度器会帮我们向操作系统申请新的系统级线程，而当某个 M 已无用时，调度器又会负责把它及时地销毁掉。

M、P、G 之间的关系（简化版）

每条go语句一般都会携带一个函数调用，这个被调用的函数常常被称为go函数

#### 注意:
    go函数真正被执行的时间，总会与其所属的go语句被执行的时间不同。当程序执行到一条go语句的时候，Go 语言的运行时系统，会先试图从某个存放空闲的 G 的队列中获取一个 G（也就是 goroutine），它只有在找不到空闲 G 的情况下才会去创建一个新的 G。

#### 运行原理
1. 在拿到了一个空闲的 G 之后，Go 语言运行时系统会用这个 G 去包装当前的那个go函数（或者说该函数中的那些代码），然后再把这个 G 追加到某个存放可运行的 G 的队列中。
2. 这类队列中的 G 总是会按照先入先出的顺序，很快地由运行时系统内部的调度器安排运行。虽然这会很快，但是由于上面所说的那些准备工作还是不可避免的，所以耗时还是存在的。
3. 因此，go函数的执行时间总是会明显滞后于它所属的go语句的执行时间。当然了，这里所说的“明显滞后”是对于计算机的 CPU 时钟和 Go 程序来说的。我们在大多数时候都不会有明显的感觉。

#### 运行现象
这种原理下的表象。请记住，只要go语句本身执行完毕，Go程序完全不会等待go函数的执行，它会立刻去执行后边的语句。这就是所谓的异步并发地执行。

一个与主 goroutine 有关的重要特性，即：一旦主 goroutine 中的代码（也就是main函数中的那些代码）执行完毕，当前的 Go 程序就会结束运行。

如此一来，如果在 Go 程序结束的那一刻，还有 goroutine未得到运行机会，那么它们就真的没有运行机会了，它们中的代码也就不会被执行了。

#### 运行顺序
Go 语言并不会去保证这些 goroutine 会以怎样的顺序运行。由于主 goroutine 会与我们手动启用的其他 goroutine 一起接受调度，又因为调度器很可能会在 goroutine 中的代码只执行了一部分的时候暂停，以期所有的 goroutine 有更公平的运行机会。

所以哪个 goroutine 先执行完、哪个 goroutine 后执行完往往是不可预知的，除非我们使用了某种 Go 语言提供的方式进行了人为干预。


参考: https://segmentfault.com/a/1190000015352983