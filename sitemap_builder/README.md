## 练习项目5:
1. 提供一个域名，根据这个域名创建一个站点地图，站点地图是xml格式，如下：
```
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
  <url>
    <loc>http://www.example.com/</loc>
  </url>
  <url>
    <loc>http://www.example.com/dogs</loc>
  </url>
</urlset>
```
注意：
1. 提取a标签的链接时候，各页面之间的连接可能会产生循环嵌套，例如：
```
/about -> /contact
/contact -> /pricing
/pricing -> /testimonials
/testimonials -> /about
```

## 知识点：
1. html字符串的解析，golang.org/x/net/html 包的使用
2. xml格式字符串的转换，encoding/xml 包的使用
3. 命令行参数的解析，flag包的使用