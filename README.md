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