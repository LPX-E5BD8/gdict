# gdict
A command line dictionary written in golang powered by multi engines.

### 快速使用指南
#### 安装

```bash
go get -u github.com/liipx/gdict
```

#### 升级
```bash
cd $GOPATH/src/github.com/liipx/gdict && git pull && make install
```

#### 使用
##### 简单使用案例
- 有道词典
```bash
➜  gdict git:(master) ✗ gdict peace

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
---
- 必应词典
```bash
➜  gdict git:(master) ✗ gdict peace -e bing

查询: peace

英: piːs    美: pis

Exps:

   n.  和平；平静；宁静；和睦;
  Web  平和；安宁；静谧;

例句:

 1. But even then we often do not know God well enough, and that's when doubt and anxiety can settle in, replacing serenity and trust and peace.
   不过，即使这样，我们对神的认识往往只是皮毛，结果是疑惑和担忧乘虚而入，叫我们失去平静、信任和平安。

 2. I need peace with you, God, and I need you to put your peace in my heart.
   我需要与你同在的安宁。上帝，我还需要你将你的安宁放在我的心中。

 3. The jury in Dearborn, home to one of the country's largest Muslim communities, said such a protest would disturb the peace.
   密西根的迪尔朋是美国最大的穆斯林社区之一。当地的陪审团表示，这种抗议会扰乱和平。

 4. 'I thought I would feel sadness and fear when I burned them, ' he told me. 'But I felt a great sense of release and peace.
   叶琛说，他事前曾认为自己烧这些日记时会感到悲伤和恐惧，但真烧时他感到的却只有放松与平和。

 5. He said the Maoists were fully committed to peace process and the constitution-drafting process.
   他说尼联共（毛）完全忠于和平进程和宪法的制定进程。

 6. Then David accepted from her hand what she had brought him and said, "Go home in peace. I have heard your words and granted your request. "
   大卫受了亚比该送来的礼物，就对她说：“我听了你的话，准了你的情面，你可以平平安安地回家吧！”

 7. Mr Dahal, who also goes by his nom de guerre Prachanda, or "the fierce one" , said he was stepping down to "save the peace process" .
   化名普拉昌达（意为凶猛的人）的达哈尔表示，他以下台“挽救和平进程”。

 8. We hope that this message has in some way brought those who needed to hear it peace and that it has in some way served you.
   我们希望这则信息在某种程度上带来了那些需要聆听平和的人并在某种方式中服务了你

 9. It will be a time of the sword, of warfare, not of peace, because there is so much corruption and decay.
   将会有一段刀兵相见、战火纷飞而不是和平的时期，因为存在着太多的腐朽和衰败了。

10. In so doing, we bring our hearts into harmony with the heart of God. We are invited into a joy and peace that words cannot describe.
   如此我们的心和神的灵便能彼此和谐协调，而我们亦得以享受一种非笔墨能形容的喜乐和平安了。
```
---
- 金山爱词霸
```bash
➜  gdict git:(master) ✗ gdict peace -e iciba

查询: peace

英: pi:s    美: pis

Exps:

   n.  和平；和睦；治安；安心；

例句:

 1. Adam Smith naturally understood under the word'peace'the'perpetual universal peace'of the Abb? St. Pierre.
    亚当·斯密所理解的 “ 和平”,当然是 象圣 皮埃尔神甫所说那样的“持久、普遍的和平 ”.

 2. The UN peace - keeping forces won the Nobel Peace Prize in 1988.
    1998年联合国维和部队获诺贝尔和平奖.

 3. But for peace her soul was yearning, And now peace laps her round.
    但她的灵魂在渴求安宁, 而现在安宁将她包笼.

 4. Street Battle In Heavy Shelling As Peace Talks Proceed As Peace Talks Proceeded.
    和平会谈进行之际巷战依然炮声隆隆.

 5. Money can't buy inward peace -- peace is the result of a constructive philosophy of life.
    金钱买不到快乐 --- 快乐是一种心态,人住在茅屋可以像住在大厦一样快乐.
```

### 说明
```text
$ gdict [options ...] word [options ...]
```
对参数做了处理，参数没有特定的位置限制【甚至于可以在句子中夹杂着参数..】

Options：
```text
Options:
  Style: -dark, -light              // 配色方案
   Read: -s, --say  (MacOS only)    // 魅惑发音
 Engine: -e <engine name>           // 词典引擎，目前支持Youdao、bing(默认为有道youdao)
   Help: -h, --help                 // 查看帮助
```

### TODO
1. 离线缓存
2. 词汇笔记本

### 其他说明
1. 有道云、爱词霸API所使用Key皆来自github，侵删。
2. MacOS支持鬼魅朗读，可以感受一下: `gdict xxx -s` ... 其实就是调用的say
