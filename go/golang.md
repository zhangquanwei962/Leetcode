# GMP调度

GPM调度是Go语言中的协程调度器。G表示goroutine（协程），P表示处理器（Processor），M表示线程（Thread）。GPM调度器的主要目标是将系统中的goroutine（G）映射到可用的处理器（P）上，并将处理器映射到操作系统的线程（M）上。

GPM调度器的主要工作流程如下：

1. G的创建：当我们使用`go`关键字创建一个goroutine时，G会被创建，并且它的上下文（包括程序计数器、栈、寄存器等信息）会被保存。

2. P的创建和绑定：当G被创建时，调度器会将G绑定到一个处理器P上。处理器P维护了一个goroutine队列（称为runqueue）和一组系统线程M。初始时，调度器会根据系统的CPU核心数量创建一定数量的处理器P。

3. G的调度：当一个处理器P空闲时，它从自己的runqueue中选择一个G来执行。如果runqueue为空，则处理器会尝试从其他处理器的全局运行队列或全局等待队列中偷取一些G。这样可以确保负载均衡和公平性。

4. G的执行：被选中的G会在当前处理器P上执行，直到它主动让出处理器（如发生阻塞、时间片用完等情况）或执行完成。

5. M的管理：处理器P需要在操作系统的线程M上执行goroutine。调度器会根据需要创建或销毁线程M，并将处理器与线程绑定起来。如果某个G发生了系统调用阻塞，线程M还可以被释放给其他处理器使用。

6. 阻塞与唤醒：当一个G发生阻塞（如等待I/O、等待锁等）时，它会从当前处理器P上解绑，进入阻塞状态，此时P变为空闲状态。当阻塞的G可以继续执行时（如I/O完成、锁可用等），它会被重新唤醒，并重新加入到某个处理器P的队列中等待执行。

通过以上的调度过程，GPM调度器可以实现高效的并发编程，充分利用系统资源，并自动处理阻塞和唤醒的情况，保证程序的正确性和性能。值得注意的是，GPM调度器的具体实现和行为可能因Go语言版本而有所不同，以上是一种常见的调度机制描述。

# golang的协程
在Go语言中，协程又称为goroutine，是一种轻量级的并发机制。通过使用关键字go+函数或方法调用，可以创建一个goroutine。

下面是一些关于goroutine的特点和使用方式：

轻量级：goroutine的创建和销毁开销很小，可以创建大量的goroutine而不会造成太大的负担。

并发执行：多个goroutine可以同时执行，并且它们之间通过通信来同步和共享数据。

独立栈空间：每个goroutine都有独立的栈空间，相比于操作系统线程，goroutine的栈空间更小且可动态伸缩。

由调度器管理：Go语言的运行时系统会自动对goroutine进行调度，将其映射到可用的处理器上执行，并在需要时自动进行阻塞和唤醒。

使用goroutine的一般步骤如下：

使用go关键字加上函数或方法的调用，即可创建一个goroutine。例如：go funcName()、go methodName()。

如果需要在goroutine之间进行数据交互或同步，可以使用通道（channel）来实现。通道提供了一种安全且高效的方式来传递数据。

# sync.Map{}怎么用
`sync.Map` 是 Go 语言标准库中提供的一种并发安全的映射（Map）类型，可以在多个 Goroutine 之间安全地读写数据。与常规的 map 不同，`sync.Map` 的操作不需要加锁或使用其他同步机制。

`sync.Map` 的常用方法如下：

1. `Store(key, value interface{})`：向映射中存储一个键值对。
2. `Load(key interface{}) (value interface{}, ok bool)`：从映射中加载指定键对应的值。
3. `Delete(key interface{})`：从映射中删除指定键的值。
4. `Range(f func(key, value interface{}) bool)`：遍历映射中的所有键值对，并将其传递给指定的函数。

下面是一个简单的示例代码，展示了如何使用 `sync.Map`：

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	// 存储键值对
	m.Store("name", "Alice")
	m.Store("age", 25)
	m.Store("city", "New York")

	// 加载特定键的值
	if value, ok := m.Load("name"); ok {
		fmt.Println("Name:", value)
	}

	// 遍历映射中的所有键值对
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
		return true
	})

	// 删除特定键的值
	m.Delete("age")

	// 再次遍历映射中的所有键值对
	m.Range(func(key, value interface{}) bool {
		fmt.Printf("Key: %s, Value: %v\n", key, value)
		return true
	})
}
```

这个示例展示了如何使用 `sync.Map` 进行存储、加载、删除和遍历操作。需要注意的是，在遍历 `sync.Map` 时，传递给 `Range` 方法的函数会按需多次调用，直到返回 `false` 为止。

请注意，虽然 `sync.Map` 提供了并发安全的操作，但它并不能替代所有的场景。在一些特定的应用场景中，还是需要自己手动使用锁来确保数据的并发安全性。
# go中三种以上拼接字符串，哪个性能更好？
在 Go 语言中，有多种方法可以拼接字符串。以下是三种常见的方法：使用 `+` 运算符、使用 `fmt.Sprintf()` 函数和使用 `strings.Join()` 函数。这些方法的性能会受到多种因素的影响，例如字符串长度、拼接次数等。

1. 使用 `+` 运算符：在 Go 中，可以使用 `+` 运算符连接两个字符串。

```go
str := "Hello, " + "World!"
```

2. 使用 `fmt.Sprintf()` 函数：`fmt.Sprintf()` 函数可以将格式化的字符串写入到一个变量中。

```go
str := fmt.Sprintf("Hello, %s!", "World")
```

3. 使用 `strings.Join()` 函数：`strings.Join()` 函数可以将一个字符串切片连接成一个字符串，并使用指定的分隔符进行分隔。

```go
str := strings.Join([]string{"Hello", "World!"}, ", ")
```

对于较小规模的字符串拼接，这三种方法的性能差异并不明显。然而，当需要拼接大量字符串或进行迭代拼接时，性能差异可能会显现。

从性能角度来看，使用 `strings.Join()` 函数通常是最高效的。这是因为 `strings.Join()` 函数在内部使用了更高效的字符串缓冲区来处理字符串拼接，并且可以避免重复创建中间字符串。

示例代码如下：

```go
import (
    "fmt"
    "strings"
)

func main() {
    // 使用 + 运算符拼接字符串
    result1 := ""
    for i := 0; i < 10000; i++ {
        result1 += "Hello"
    }
    
    // 使用 fmt.Sprintf() 函数拼接字符串
    result2 := ""
    for i := 0; i < 10000; i++ {
        result2 = fmt.Sprintf("%s%s", result2, "Hello")
    }
    
    // 使用 strings.Join() 函数拼接字符串
    var result3 []string
    for i := 0; i < 10000; i++ {
        result3 = append(result3, "Hello")
    }
    str := strings.Join(result3, "")
    
    fmt.Println(len(result1))
    fmt.Println(len(result2))
    fmt.Println(len(str))
}
```

在上述示例中，使用 `strings.Join()` 函数的性能最好，因为它避免了在每次迭代时创建新的中间字符串。而使用 `+` 运算符和 `fmt.Sprintf()` 函数则会频繁地创建中间字符串，性能较差。

尽管 `strings.Join()` 函数在性能上表现较好，但根据具体的使用场景和需求，也可以选择其他方法进行字符串拼接。对于小规模的字符串拼接，三种方法的性能差异可以忽略不计。选择合适的方法取决于代码的可读性、易用性和可维护性。因此，在编写代码时，应根据具体情况权衡利弊并做出合适的选择。

# go实现默认参数
在 Go 中，函数不支持默认参数的直接定义。但是，你可以使用可变参数和结构体等技术来实现类似默认参数的功能。

以下是两种常见的实现方式：

1. 使用可变参数：你可以定义一个接受可变参数的函数，并根据传入的参数个数来确定具体的行为。如果某个参数没有被指定，则使用默认值。

```go
func MyFunc(args ...int) {
    var param1 int
    var param2 string
    
    // 根据参数个数设置默认值
    if len(args) > 0 {
        param1 = args[0]
    } else {
        param1 = 10 // 默认值
    }
    
    if len(args) > 1 {
        param2 = args[1]
    } else {
        param2 = "default" // 默认值
    }
    
    // 函数逻辑
}
```

使用示例：

```go
MyFunc(20, "hello") // 使用传入的值作为参数
MyFunc(30)         // 使用默认值作为第二个参数
MyFunc()           // 使用默认值作为两个参数
```

2. 使用结构体：你可以定义一个结构体类型作为函数的参数，并在结构体中定义字段来表示不同的参数。然后，你可以通过创建结构体变量并设置字段的方式来调用函数。

```go
type Options struct {
    Param1 int
    Param2 string
}

func MyFunc(options Options) {
    if options.Param1 == 0 {
        options.Param1 = 10 // 默认值
    }
    
    if options.Param2 == "" {
        options.Param2 = "default" // 默认值
    }
    
    // 函数逻辑
}
```

使用示例：

```go
MyFunc(Options{Param1: 20, Param2: "hello"}) // 使用传入的值作为参数
MyFunc(Options{Param1: 30})                 // 使用默认值作为第二个参数
MyFunc(Options{})                          // 使用默认值作为两个参数
```

使用结构体作为参数的方式更加灵活，可以支持多个参数且易于扩展。

虽然 Go 不直接支持默认参数，但通过这些技术，你可以模拟实现相似的功能。请根据具体情况选择合适的实现方式。

# go语言的tag有什么用处
在 Go 语言中，Tag 是一种用于在结构体字段上附加元数据的机制。Tag 是一个字符串，可以在字段声明中使用反引号包围起来，并放置在字段的后面，如下所示：

```go
type MyStruct struct {
    Field1 int `tag1:"value1" tag2:"value2"`
    Field2 string `tag3:"value3"`
}
```

Go 的反射机制可以通过结构体的 Tag 来获取与字段相关的信息，这些信息可以在运行时被利用。Tag 在很多场景中都有实际的用途，下面是一些常见的用途：

1. 序列化和反序列化：Tag 经常用于定义结构体字段与 JSON、XML 或其他格式数据之间的映射关系。例如，使用 JSON 标签可以指定字段在 JSON 编码和解码时的名称、忽略某些字段等。

```go
type Person struct {
    Name string `json:"name"`
    Age int `json:"age"`
}
```

2. 数据校验：Tag 可以用于验证结构体字段的合法性。通过在字段上设置特定的验证规则，可以在运行时对字段进行校验。常见的数据校验库（如 `go-playground/validator`）使用 Tag 来定义验证规则。

```go
type User struct {
    Username string `validate:"required,min=5,max=20"`
    Password string `validate:"required,min=8,max=50"`
}
```

3. 数据库映射：Tag 可以用于指定结构体字段与数据库表格的映射关系。对于使用 ORM（对象关系映射）的应用程序，可以使用 Tag 来定义表格名称、列名、主键等信息。

```go
type Product struct {
    ID int `gorm:"primaryKey"`
    Name string `gorm:"column:name"`
}
```

4. 文档生成：在一些文档生成工具中，会使用结构体字段的 Tag 作为字段的注释或描述信息。这样可以通过反射将结构体的 Tag 提取出来，并生成相应的文档。

总之，Tag 在 Go 语言中是非常有用的机制，它允许我们在结构体字段上附加额外的元数据，用于实现序列化、验证、映射和文档等功能。通过反射，可以在运行时获取并解析这些 Tag，从而实现更加灵活和可扩展的代码。
# 如何一次性获取一个结构体里所有的 tag
要一次性获取一个结构体中所有字段的 Tag，可以使用 Go 语言的反射机制。下面是一个示例代码，演示如何获取结构体中所有字段的 Tag：

```go
package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	Field1 int    `tag1:"value1" tag2:"value2"`
	Field2 string `tag3:"value3"`
}

func main() {
	s := MyStruct{Field1: 10, Field2: "hello"}

	t := reflect.TypeOf(s)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field name: %s\n", field.Name)
		fmt.Printf("Tag value: %s\n", field.Tag)
		fmt.Println()
	}
}
```

在上面的示例中，我们定义了一个名为 `MyStruct` 的结构体，并给其中的字段添加了不同的标签。在 `main` 函数中，我们创建了一个结构体实例 `s`，然后使用反射获取其类型信息。

通过 `reflect.TypeOf(s)`，我们获取到 `MyStruct` 的类型。使用 `t.NumField()` 获取该结构体中字段的数量，然后使用循环遍历每个字段。

在循环中，我们调用 `t.Field(i)` 获取第 `i` 个字段的信息，其中包括字段的名称和标签。使用 `field.Name` 获取字段名称，使用 `field.Tag` 获取字段的标签。

运行以上代码，你将会看到类似下面的输出：

```
Field name: Field1
Tag value: tag1:"value1" tag2:"value2"

Field name: Field2
Tag value: tag3:"value3"
```

这样，你就可以一次性获取结构体中所有字段的 Tag，并进行进一步的处理。请注意，`field.Tag` 返回的是一个字符串，在需要解析具体的标签内容时，你可以使用相应的解析方法来提取所需的信息。

希望这个示例对你有帮助！

# %v %+v %#v 有什么区别
`%v`、`%+v` 和 `%#v` 是 Go 语言中用于格式化输出的 verb（占位符），用于在字符串中插入变量的值。

这些 verb 在功能上有一些区别：

1. `%v`：`%v` 是最常用的占位符，使用默认的格式输出变量的值。对于结构体类型，它会递归打印结构体的字段名和值。

示例代码：

```go
type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 25}
	fmt.Printf("%v\n", p)
}
```

输出结果：

```
{Alice 25}
```

2. `%+v`： `%+v` 会在 `%v` 的基础上，额外输出结构体的字段名和值。对于结构体类型，在输出时会包含字段的名称。

示例代码：

```go
type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 25}
	fmt.Printf("%+v\n", p)
}
```

输出结果：

```
{Name:Alice Age:25}
```

3. `%#v`：`%#v` 会输出 Go 语法格式的值，包括变量的类型和其源代码表示。对于复杂的数据类型，例如结构体和切片，在输出时也会包含完整的类型信息。

示例代码：

```go
type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 25}
	fmt.Printf("%#v\n", p)
}
```

输出结果：

```
main.Person{Name:"Alice", Age:25}
```

综上所述，`%v`、`%+v` 和 `%#v` 在格式化输出时有一些不同的行为。你可以根据需要选择适当的占位符来打印变量的值，并根据具体的场景来决定使用哪个占位符。

# 哪些能作为 map 的 key 哪些不能
在 Go 语言中，可以作为 map 的 key 的类型必须满足以下要求：

1. 可以进行相等性比较的类型：map 的 key 必须是可比较的类型。这包括基本类型（如整数、浮点数、布尔值、字符串等）和一些引用类型（如指针、数组、结构体等），只要它们的值可以通过 `==` 运算符进行比较即可。

2. 不可变的类型：map 的 key 必须是不可变的类型，因为如果 key 的值发生了修改，可能导致无法正确查找或删除关联的值。例如，切片、函数和包含切片的结构体都不能作为 map 的 key。

3. Hashable 的类型：map 的 key 必须是可哈希的类型，也就是说，key 的值能够通过哈希函数（`hash.Hash()` 方法）转换为一个唯一的哈希值。可哈希的类型包括所有基本类型和一些引用类型（如字符串、数组、结构体等），只要它们的值不会发生变化。

以下是一些常见的可作为 map 的 key 的类型示例：

- 整数类型（int、int64、uint32 等）
- 字符串类型（string）
- 浮点数类型（float64）
- 数组类型（例如 [2]int、[3]bool）
- 结构体类型（只要结构体的字段都是可比较的类型）
- 指针类型（指向可比较的类型）

以下是一些不能作为 map 的 key 的类型示例：

- 切片类型（[]int、[]string 等）
- 函数类型
- 接口类型
- 包含切片的结构体类型

需要注意的是，如果尝试使用不符合上述要求的类型作为 map 的 key，编译器会在编译时报错。

希望对你有所帮助！
# 什么时候map会panic
在使用 `map` 时，以下情况可能导致 panic 的出现：

1. 试图从一个为 nil 的 map 中读取或删除元素：如果一个 map 变量没有初始化（未分配内存），或者它的值被显式设置为 nil，那么在对该 map 进行读取或删除操作时会触发 panic。

   ```go
   var m map[string]int  // 未初始化的 map
   fmt.Println(m["key"]) // 读取操作会导致 panic

   var n map[string]int = nil  // 设置为 nil 的 map
   fmt.Println(n["key"])      // 读取操作会导致 panic
   ```

2. 试图向一个为 nil 的 map 中添加元素：类似于上述情况，在一个为 nil 的 map 上进行添加操作也会引发 panic。

   ```go
   var m map[string]int  // 未初始化的 map
   m["key"] = 10         // 添加操作会导致 panic

   var n map[string]int = nil  // 设置为 nil 的 map
   n["key"] = 10                // 添加操作会导致 panic
   ```

3. 并发访问和修改同一个 map：当多个 goroutine 并发地对同一个 map 进行读写操作时，会导致竞态条件。这可能导致 map 数据结构被损坏，进而触发 panic。为了避免这种情况，应使用适当的并发控制机制（如互斥锁）来保护 map 的访问。

   ```go
   m := make(map[string]int)
   var wg sync.WaitGroup
   wg.Add(2)
   go func() {
       defer wg.Done()
       // 并发写操作
       m["key1"] = 10
   }()
   go func() {
       defer wg.Done()
       // 并发读操作
       fmt.Println(m["key1"])
   }()
   wg.Wait()
   ```

4. 删除一个不存在的 key：尝试删除一个 map 中不存在的 key 值会触发 panic。为了避免这种情况，可以使用 `delete` 函数前先检查 key 是否存在。

   ```go
   m := make(map[string]int)
   delete(m, "key")  // 删除一个不存在的 key 会触发 panic
   ```

总之，在使用 map 时，特别是在并发环境下，并且需要注意处理 nil map、并发安全和 key 的存在性等问题，以避免触发 panic。
# go中反引号`
在 Go 中，反引号（`）用作字符串字面量的界定符。使用反引号可以创建原始字符串字面量，这种字符串字面量中的特殊字符（包括换行符和引号）都会保持原样，不需要进行转义处理。

反引号被用于以下情况：

1. 创建多行字符串：使用反引号可以方便地创建多行字符串，而无需在每一行的末尾添加显式的换行符。这对于存储文本块或代码模板非常有用。

```go
str := `这是一个
多行字符串
示例`
```

2. 避免转义字符：反引号内的字符串不会进行转义，这意味着你可以直接插入特殊字符、引号或其他转义字符，而无需使用反斜杠进行转义操作。

```go
str := `这是一个包含特殊字符的字符串：\r \n " '`
```

在这种情况下，反引号使得字符串的书写更加清晰和简洁。

需要注意的是，反引号字符串可以包含任意字符，包括换行符和控制字符。这使得它在处理包含大量格式化内容的字符串时非常有用。

要注意的是，反引号字符串不能包含反引号本身，因为反引号是字符串的界定符。如果需要在反引号字符串中表示反引号，可以使用其他字符串拼接的方式来实现。

总之，Go 中的反引号（`）在字符串字面量中允许原始文本的使用，不需要转义特殊字符，使得字符串的书写更加方便和清晰。

# 如果我们把一个 map 的 key 删除了，他的内存会释放吗？
在 Go 中，当从 map 中删除一个 key-value 对时，并不会立即释放 key 和 value 所占用的内存空间。Go 的垃圾回收机制（Garbage Collector）负责管理和释放不再使用的内存。

当你从 map 中删除一个 key-value 对时，实际上只是将该 key 标记为已删除，而不是立即释放其所占用的内存。垃圾回收机制会周期性地检查和回收不再被引用的对象，包括被标记为删除的 key 所关联的 value。

垃圾回收机制的工作是自动进行的，并且由 Go 运行时环境自主管理。它会根据一些策略和算法来确定何时释放被标记为删除的 key 所关联的 value 占用的内存空间。

具体的垃圾回收机制实现细节可以参考 Go 官方文档中关于垃圾回收的介绍。

需要注意的是，虽然删除 map 中的 key 并不会立即释放内存，但通过删除无用的 key-value 对可以帮助垃圾回收机制更快地释放不再使用的内存资源，从而减少内存的占用并提高程序的性能。

总结起来，当从 map 中删除一个 key-value 对时，并不会立即释放它们所占用的内存空间，而是由垃圾回收机制负责管理和释放不再使用的内存。