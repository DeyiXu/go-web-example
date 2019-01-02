# go-web-example
go web example.

# 下载例子项目
```
git clone https://github.com/DeyiXu/go-web-example.git

cd go-web-example
```
# 删除 git文件
```
rm -rf .git
```
# 替换内容 `github.com/DeyiXu/go-web-example` 为 `book`
```
find -f './' | xargs perl -pi -e 's|github.com/DeyiXu/go-web-example|book|g'
```

# 项目引用

> WEB框架：[gin](https://github.com/gin-gonic/gin)

> ORM：[gorm](https://github.com/jinzhu/gorm)

> 后台管理：[AdminLTE](https://github.com/almasaeed2010/AdminLTE)

