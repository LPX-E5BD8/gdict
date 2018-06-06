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

![](https://user-images.githubusercontent.com/39460745/41031204-b6a57410-69b2-11e8-97c7-029a75458e85.png)
---
- 必应词典

![](https://user-images.githubusercontent.com/39460745/41031258-d489d8c2-69b2-11e8-8301-2fd04e54cb9d.png)

---
- 金山爱词霸

![](https://user-images.githubusercontent.com/39460745/41031328-f0ae751c-69b2-11e8-8709-1498423d3720.png)

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
