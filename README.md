# gdict
a command dict by golang power by YOUDAO.


> 这是一个为了自己做的词典，因为懒得用鼠标，python或js实现的词典依赖又太多。<br>
> 有需求可以给我留issue，我会尽量支持（如果有人用的话雾）

### 快速使用指南
#### 安装

```bash
go get -u github.com/liipx/gdict

```
#### 使用
##### 简单使用案例
```bash
# 有道词典
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
    
-----------------------------------------------------------------------------------------------------------------------
# 必应词典
➜  ~ gdict 和平 -e bing

查询: 和平

Exps:

 adj.  mild;
   n.  peace;
  Web  peaceful; Hoa Binh; Pax Americana;

例句:

 1. The jury in Dearborn, home to one of the country's largest Muslim communities, said such a protest would disturb the peace.
   密西根的迪尔朋是美国最大的穆斯林社区之一。当地的陪审团表示，这种抗议会扰乱和平。

 2. But the security forces themselves seem to be as much of a danger as the protesters.
   相反，安全力量本身构成了与抵抗者类似的对和平的威胁。

 3. We stand ready to work with the United States and other countries in the region to promote a peaceful and prosperous Asia-Pacific.
   我们愿与美国和亚太各国共同努力，为促进地区的和平与繁荣发挥作用。

 4. He said the Maoists were fully committed to peace process and the constitution-drafting process.
   他说尼联共（毛）完全忠于和平进程和宪法的制定进程。

 5. Mr Dahal, who also goes by his nom de guerre Prachanda, or "the fierce one" , said he was stepping down to "save the peace process" .
   化名普拉昌达（意为凶猛的人）的达哈尔表示，他以下台“挽救和平进程”。

 6. The earliest of these was the Black Ball Line opened in New York in 1816, only a year after the war.
   这些最早的是黑球线于纽约，在1816年，只有一年的和平。

 7. After the Agnew nomination was announced, what had been a peaceful gathering against poverty turned into a riot.
   当宣布提名阿格纽时，本来是反对贫穷的和平集会演变为一场骚乱。

 8. It will be a time of the sword, of warfare, not of peace, because there is so much corruption and decay.
   将会有一段刀兵相见、战火纷飞而不是和平的时期，因为存在着太多的腐朽和衰败了。

 9. The right to be brought up in a spirit of peace and brotherhood.
   在和平友爱的精神下成长的权利。

10. Greenpeace said the plans was "dramatically lacking in ambition" , a wasted effort and a failure from the outset.
   绿色和平组织（Greenpeace）表示，这项计划是“明显缺乏雄心壮志”，从一开始就是一场徒劳、一次失败。
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



### 其他说明
1. 使用有道云老版本API，Key来自GITHUB，侵删。
2. MacOS支持鬼魅朗读，可以感受一下: `gdict xxx -s` ... 其实就是调用的say

### TODO
1. 多词典引擎支持【必应、金山词霸等】
2. 多系统发音支持
3. ~~更多的配色方案选择~~✅
4. 配置持久化
5. 历史记录查询
7. 查询记录缓存
8. 想起来再加
