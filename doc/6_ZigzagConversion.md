## 6. Z字形变换
将字符串 "PAYPALISHIRING" 以Z字形排列成给定的行数：
```
P   A   H   N
A P L S I I G
Y   I   R
```
之后从左往右，逐行读取字符："PAHNAPLSIIGYIR"

实现一个将字符串进行指定行数变换的函数:
```
string convert(string s, int numRows);
```
示例 1:
```
输入: s = "PAYPALISHIRING", numRows = 3
输出: "PAHNAPLSIIGYIR"
```
示例 2:
```
输入: s = "PAYPALISHIRING", numRows = 4
输出: "PINALSIGYAHRPI"
解释:

P     I    N
A   L S  I G
Y A   H R
P     I
```

#### 解题思路
通过各行字符对应在原字符串的下标位置找到如下规律：
```
if row = 0 or row = n - 1:
    i * (2 * n - 2) + row 
else:
    i * (2 * n - 2) - row
    i * (2 * n - 2) + row
    
row 表示当前行
n 表示要求的行数
i 表示第row行的第几个字符
``` 
特殊情况：
字符串的长度<=numRows or numRows = 1 时，字符串原样输出

