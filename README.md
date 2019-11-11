# golang官网翻译
改用markdown保存pkg文档, 便于对比更新

## 最新 go version
go 1.13.4

## 进度
[已完成列表, 请见pkg.md](go_pkg/pkg.md).

## 工具
- markdown生成: `tools/pkg`
- 同步`go_pkg`和`go_pkg_origin`中的代码引用位置: `tools/source_def`, 未完成.

### 同步go version的步骤
1. 使用pkg工具生成的最新版本go_pkg_latest, 并放到根目录下
1. 使用source_def将go_pkg_latest的源码引用覆盖到go_pkg和go_pkg_origin
1. 使用`Beyond Compare`对比go_pkg_latest和go_pkg_origin, 将相关差异sync到go_pkg
1. 删除go_pkg_origin, 将go_pkg_latest重命名为go_pkg_origin

## 翻译原则
- 针对未翻译过的package, 先用go_pkg_orgin对应的文件覆盖一遍再翻译, 便于同步最新的文档.
- 碰到单行注释时, 直接在下一行添加翻译; 多行时, 换一行再添加翻译
- 接口有example时, 先补充该例子, 再翻译接口. 因为有例子帮助能更好地理解接口.
- 翻译完成后, 用git diff检查一遍翻译, 同时应删除git diff提示的无用空格.

## **注意**
1. pkg对应md里的examples需要手动添加, 因为原始pkg里的example是html格式.
1. `go_pkg_origin`是用于与最新版本对比时用到的原件.
1. `???`表示不理解的内容

## todo
- [x] tool : 更新pkg文件中的具体type/func的源码指向地址
- [ ] pkg/builtin : 生成的orgin文档不全

## 内容许可
除特别说明外，用户内容均采用[知识共享署名-非商业性使用-禁止演绎 4.0 国际许可协议](https://creativecommons.org/licenses/by-nc-nd/4.0/)进行许可
