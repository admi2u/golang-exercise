## 练习项目2：
1. 实现一个http 处理器（handler）,可以处理http请求；
2. 当请求的path和指定的path匹配时，就跳转到特定的url；
3. path映射规则可以由map变量类型提供，也可以由yaml格式的字符串提供

## 主要知识点：
1. net/http 包的使用，http.ServeMux，handler、HandleFunc，Handler、Handle的理解运用，自定义handler的创建；
2. yaml字符串的解析，yaml包的使用；