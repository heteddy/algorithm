# 前缀树

Trie树，又称字典树、前缀树，是一种树形结构，是哈希树的变种，是一种用于快速检索的多叉树结构。
典型应用是用于统计和排序大量的字符串（但不仅限于字符串），所以经常被搜索引擎系统用于文本词频统计。
它的优点是：最大限度地减少无谓的字符串比较，查询效率比哈希表高。
Trie的核心思想是空间换时间。利用字符串的公共前缀来降低查询时间的开销以达到提高效率的目的。
Trie树也有它的缺点,Trie树的内存消耗非常大。

我们使用辞典或者是搜索引擎的时候，输入appl，后面会自动显示一堆前缀是appl的东东吧。
那么有可能是通过字典树实现的，前面也说了字典树可以找到公共前缀，我们只需要把剩余的后缀遍历显示出来即可。

## 统计和排序大量的字符串（但不仅限于字符串）
1 找有多少个字符串是以给定的字符串为前缀
2 添加过多少个某个给定的字符串
3 文本词频统计
…

