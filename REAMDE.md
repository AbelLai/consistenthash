# consistenthash
> English: This repo is based on this [article](https://medium.com/@sent0hil/consistent-hashing-a-guide-go-implementation-fe3421ac3e8f#.upih8qaqy). And you can find the source code of that article [here](https://github.com/sent-hil/consistenthash). Since there is no virtual slot implementation descriped at that article, so I read [Groupcache's implementation](https://github.com/golang/groupcache/blob/master/consistenthash/consistenthash.go) and add the stuff about virtual slot in this repo.

> 中文: 此repo实现是基于[这篇文章](https://medium.com/@sent0hil/consistent-hashing-a-guide-go-implementation-fe3421ac3e8f#.upih8qaqy)的内容(它的代码在[这里](https://github.com/sent-hil/consistenthash))以及[Groupcache的一致性Hash算法实现](https://github.com/golang/groupcache/blob/master/consistenthash/consistenthash.go)的一个合并，并非个人的原创。因为该文章没有虚拟节点的实现，因此在参考Groupcache的实现，把虚拟节点部分加上。算是对该文章内容的一个小小的升级吧。

## TODO
> Example.