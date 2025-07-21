# Skip List

Skip List 是一个基于概率的数据结构，提供高效的增删查操作，平均时间复杂度为 O(log n)。它通过多层级的指针结构来实现快速的有序数据访问。

## 特性

- **高效的操作**: 插入、删除、查找操作的平均时间复杂度为 O(log n)
- **有序存储**: 自动维护键值对的有序性
- **范围查询**: 支持高效的范围查询操作
- **Go 版本兼容**: 支持 Go 1.21+ 和 Go 1.23+ 的特性
- **类型安全**: 使用 Go 泛型确保类型安全
- **Unicode 支持**: 完全支持 Unicode 字符串作为键

## 接口设计

### 基本操作

```go
type Interface[K cmp.Ordered, V any] interface {
    // 基础操作
    Len() int                           // 返回元素数量
    Get(key K) (V, bool)               // 获取值
    GetMutable(key K) (*V, bool)       // 获取可变值的指针
    Set(key K, value V)                // 设置键值对
    Delete(key K) bool                 // 删除键值对
    Has(key K) bool                    // 检查键是否存在
    Clear()                            // 清空所有元素

    // 批量操作
    Keys() []K                         // 获取所有键（有序）
    Values() []V                       // 获取所有值（按键排序）
    Pairs() []pair.Pair[K, V]         // 获取所有键值对（有序）

    // 范围操作
    Range(fn func(key K, value V) bool)                    // 遍历所有元素
    RangeFrom(start K, fn func(key K, value V) bool)       // 从指定键开始遍历
    RangeBetween(start, end K, fn func(key K, value V) bool) // 在指定范围内遍历
}
```

### Go 1.23+ 迭代器支持

```go
// Go 1.23+ 版本额外提供迭代器支持
All() iter.Seq2[K, V]                    // 迭代所有元素
AllFrom(start K) iter.Seq2[K, V]         // 从指定键开始迭代
AllBetween(start, end K) iter.Seq2[K, V] // 在指定范围内迭代
```

## 使用示例

### 基本用法

```go
package main

import (
    "fmt"
    "github.com/feepwang/br/container/skip_list"
)

func main() {
    // 创建 Skip List
    sl := skip_list.NewOrderedSkipList[int, string]()
    
    // 插入数据
    sl.Set(5, "five")
    sl.Set(2, "two")
    sl.Set(8, "eight")
    sl.Set(1, "one")
    
    // 查询数据
    if value, exists := sl.Get(5); exists {
        fmt.Printf("Key 5: %s\n", value) // 输出: Key 5: five
    }
    
    // 检查键是否存在
    fmt.Printf("Has key 3: %t\n", sl.Has(3)) // 输出: Has key 3: false
    
    // 获取有序的键
    fmt.Printf("All keys: %v\n", sl.Keys()) // 输出: All keys: [1 2 5 8]
}
```

### 范围查询

```go
// 遍历所有元素
sl.Range(func(key int, value string) bool {
    fmt.Printf("%d: %s\n", key, value)
    return true // 继续遍历
})

// 范围查询
sl.RangeBetween(2, 6, func(key int, value string) bool {
    fmt.Printf("%d: %s\n", key, value) // 只输出键在 [2, 6] 范围内的元素
    return true
})
```

### 字符串键

```go
strSL := skip_list.NewOrderedSkipList[string, int]()

strSL.Set("banana", 1)
strSL.Set("apple", 2)
strSL.Set("cherry", 3)

// 自动按字典序排序
fmt.Printf("Keys: %v\n", strSL.Keys()) // 输出: Keys: [apple banana cherry]
```

### Go 1.23+ 迭代器

```go
// 使用迭代器遍历所有元素
for k, v := range sl.All() {
    fmt.Printf("%d: %s\n", k, v)
}

// 从指定位置开始迭代
for k, v := range sl.AllFrom(5) {
    fmt.Printf("%d: %s\n", k, v)
    if k >= 10 {
        break // 提前结束
    }
}

// 范围迭代
for k, v := range sl.AllBetween(2, 8) {
    fmt.Printf("%d: %s\n", k, v)
}
```

### 自定义比较器（Go 1.23+）

```go
// 创建反向排序的 Skip List
reverseSL := skip_list.NewSkipList[int, string](func(a, b int) int {
    return cmp.Compare(b, a) // 反向比较
})

reverseSL.Set(1, "one")
reverseSL.Set(2, "two")
reverseSL.Set(3, "three")

fmt.Printf("Keys: %v\n", reverseSL.Keys()) // 输出: Keys: [3 2 1]
```

## 性能特征

| 操作 | 平均时间复杂度 | 最坏时间复杂度 |
|------|----------------|----------------|
| 插入 | O(log n) | O(n) |
| 删除 | O(log n) | O(n) |
| 查找 | O(log n) | O(n) |
| 范围查询 | O(log n + k) | O(n) |

其中 k 是结果集的大小，n 是元素总数。

**空间复杂度**: O(n)

## 实现细节

- **多层级结构**: 使用概率为 0.5 的几何分布确定节点高度
- **最大层级**: 限制为 32 层，防止过度内存使用
- **头节点**: 使用哨兵头节点简化边界条件处理
- **线程安全**: 当前实现不是线程安全的，需要外部同步

## 适用场景

1. **有序数据集合**: 需要维护数据有序性的场景
2. **范围查询**: 频繁进行范围查询的应用
3. **动态插入删除**: 数据集合大小经常变化的场景
4. **替代平衡树**: 作为红黑树或 AVL 树的简单替代方案

## 与其他数据结构的比较

| 数据结构 | 插入/删除/查找 | 实现复杂度 | 内存开销 |
|----------|----------------|------------|----------|
| Skip List | O(log n) 平均 | 简单 | 中等 |
| 红黑树 | O(log n) 保证 | 复杂 | 低 |
| AVL 树 | O(log n) 保证 | 复杂 | 低 |
| 哈希表 | O(1) 平均 | 中等 | 中等 |

Skip List 的主要优势是实现简单且性能稳定，虽然最坏情况下可能退化到 O(n)，但在实际应用中很少发生。