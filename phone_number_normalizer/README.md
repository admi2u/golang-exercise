## 练习项目8：
该项目的目的是学习golang关系型数据库相关包的使用，具体步骤：
1. 在数据库中创建表，写入以下数据：
```
1234567890
123 456 7891
(123) 456 7892
(123) 456-7893
123-456-7894
123-456-7890
1234567892
(123)456-7892
```
2. 使用golang连接数据库，读取数据并将以上数据做格式化处理，处理后的数据格式如下：
```
1234567890
1234567891
1234567892
1234567893
1234567894
```
即，移除字符串中的非数字字符，并且相同的字符串只保留一个

3. 将处理后的数据写回数据库，更新数据

## 其他：
1. 数据库连接可以使用的包有：
   - 官方自带的 `database/sql`
   - 第三方包 `sqlx`
   - 使用第三方ORM，如 `gorm`（本例中使用gorm）
  
## 知识点：
1. gorm包的使用，database/sql、和sqlx包的了解
2. regexp正则表达式包的使用