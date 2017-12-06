#bitmap 算法
1. 问题来源

+ hash set 或者 hashmap经常被用于获取一个元素是否在已经存在，比如：

    ``` go
    existedNumber := make(map[int] bool)
    
    if i,existed := existedNumber[i];existed{
        log.println("existed")
    }
    ```
+ 使用数组的方式

    ```go
    existedNumber := make([1<<32-1]bool)
    //遍历一遍给定的数据，然后对应下标的值设置为1
    //当给定一个数(假设是300)，判断对应数组下标 
    if b:= existedNumber[300];b{
  	    log.println("existed")
    }
    ```
    但是这种做法使用的内存比较大（int占用32位）；如果数据量比较大(40亿个int32,大约占用4G*4)，但内存有限(2G)；
    在这种情况下就无法使用上述方法
    
2. 原理

    针对上面的问题，bitmap采用了类似于第二种方法，但existedNumber数组所占用的内存为其1/32；
    使用某一个bit代表一个数，这样 int32 2^32个，数组只需要占用4G/32 125M的空间适用于数量很多40亿个，
    但是最大数不是很大(当前的条件下是2^32)
3. 示例