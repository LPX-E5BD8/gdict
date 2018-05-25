# gdict
a command dict by golang power by YOUDAO.

### 快速使用指南
#### 安装

```bash
go get -u github.com/liipx/gdict

```
#### 使用
##### 英译汉
```bash
➜  ~ gdict peace

查询: peace

英: piːs    美: pis

Exps:

    n.  和平；平静；和睦；秩序
    n.  (Peace)人名；(英)皮斯

翻译:

  和平

网络释义:

   1.peace:
    和平, 和平, 平静

   2.Peace River:
    皮斯里弗 (艾伯塔省), 皮斯河, 皮斯河

   3.World peace:
    世界和平, 世界和平, 世界和平

```
##### 汉译英
```bash
➜  ~ gdict 和平

查询: 和平

拼音: hé píng

Exps:

 1.peace


翻译:

  peace

网络释义:

   1.和平:
    peace, Olga, Irene

   2.和平队:
    Peace Corps, The Peace Corps, PeaceCorps

   3.绿色和平:
    Greenpeace, Greenpeace East Asia, Greenpeace International
```
### 说明
```text
$ gdict word [options ...]
```
Options：
```text
Options:
  Style: -dark, -light              // 配色方案
   Read: -s, --say  (MacOS only)    // 魅惑发音
 Engine: -youdao                    // 词典引擎，目前仅支持Youdao
   Help: -h, --help                 // 查看帮助
```

> 这是一个为了自己做的词典，因为懒得用鼠标，python或js实现的词典依赖又太多。<br>
> 有需求可以给我留issue，我会尽量支持（如果有人用的话雾）

1. 使用有道云老版本API，Key来自GITHUB，侵删。
2. MacOS支持鬼魅朗读，可以感受一下: `gdict xxx -s` ... 其实就是调用的say
3. dark友好配色方案【light后续支持】

### TODO
1. 多词典引擎支持【必应、金山词霸等】
2. 多系统发音支持
3. 更多的配色方案选择
4. 配置持久化
5. 历史记录查询
7. 查询记录缓存
8. 想起来再加