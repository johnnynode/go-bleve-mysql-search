# bleve-mysql-search

测试中

### 参考

- 解决Cannot find package "golang.org/x/net/context" 方案：https://github.com/blevesearch/bleve/issues/624
- 解决cannot find package "golang.org/x/text/unicode/norm" 方案：https://github.com/blevesearch/bleve-explorer/issues/11

### 相关参考文档：

- 关于做全文检索
 * [bleve](https://github.com/blevesearch/bleve)
 * [wukong 悟空](https://github.com/huichen/wukong)
 * http://www.blevesearch.com/
 * http://studygolang.com/articles/2537
 * http://www.blevesearch.com/docs/Building/
 * https://medium.com/developers-writing/full-text-search-and-indexing-with-bleve-part-1-bd73599d82ef
 * http://www.infoq.com/cn/news/2015/03/bleve-couchbase-go
 * http://www.tuicool.com/articles/nauiEv2
 * https://github.com/yanyiwu/gojieba

- beego框架利用bee api创建api框架：

  ```
   bee api hello -conn=root:root@tcp(127.0.0.1:3306)/test
   bee run -downdoc=true -gendoc=true

  ```

- go语言中的路由设置：https://beego.me/docs/mvc/controller/router.md

- beego api 自动化：https://beego.me/docs/advantage/docs.md

- go语言高级查询： https://beego.me/docs/mvc/model/query.md#update

- go语言实战系列：http://blog.csdn.net/column/details/16165.html

- beego API开发以及自动化文档：https://my.oschina.net/astaxie/blog/284072

- beego增删改查代码实现：http://studygolang.com/articles/7241

- go语言实战系列：http://blog.csdn.net/column/details/16165.html

- go语言数值与字符串的替换： http://blog.csdn.net/jiangwei0910410003/article/details/73699649

- go语言jwt: http://blog.csdn.net/wangshubo1989/article/details/74529333

- https://github-trending.com/

- https://beego.me/video

- beego orm中的关联查询: http://www.cnblogs.com/vipzhou/p/5893568.html

### 常用go语言库网站：

- https://golanglibs.com
- https://golang.org/pkg/index/suffixarray/
- http://www.yiibai.com/go/go_type.html
- http://www.runoob.com/go/go-structures.html
- http://go-search.org/
- https://gowalker.org/

### 搜索流程

- 前台根据不同需求，构建不同的搜索条件
- 后台根据不同的搜索条件，处理不同的数据，包括排序，逻辑关系，查找关键字
- 后台通过检索引擎检索，处理搜索结果，将数据拿到，然后处理搜索时间，得分情况，通过创建uuid集合从新从数据库中查询出来
- 整合数据资源，返回给前台。