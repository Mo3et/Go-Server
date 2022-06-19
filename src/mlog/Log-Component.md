# Log  Component

###### 需求:

实现 info debug error warn 类型的日志的打印 写入到log文件并输出到控制台

###### 思路：

1. 基础  实现原子类型

2. Logger 

```
struct {
    mu         //(atomic Write protect fields) 
    prefix     //前缀标识器
    flag
    out
    buf
    isDiscard
       }
```




