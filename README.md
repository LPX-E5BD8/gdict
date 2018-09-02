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
3. 支持Alfred workflow 格式的输出，可以按照下面的步骤定制自己的workflow
4. 或者直接使用alfred目录下我导出的文件进行安装，安装后记得修改路径

```text
# 1. 打开 workflow
# 2. 添加 script filter
# 3. 填写 `Keyword` 等基本信息
# 4. language选择/bin/bash，内容从下文`query=$1`到最后
# 5. 输出追加一个`copy to clipboard`即可

# 脚本内粘贴以下内容
    
query=$1
# 将下面的信息修改为 gdict binary 的存放路径
BINPATH="/like/your/gopath/bin"
     
# 参数一定要指定-w，输出特定格式的值
# 按照以下配置进行引擎切换
# $BINPATH/gdict -w -e bing $query
# $BINPATH/gdict -w -e iciba $query
     
$BINPATH/gdict -w -e youdao $query`
```

#### 效果图
![](https://user-images.githubusercontent.com/39460745/44953434-f790ff80-aec7-11e8-82cf-271f5dbeccd1.png)

![](https://user-images.githubusercontent.com/39460745/44953473-48085d00-aec8-11e8-813e-f9fe3ea32558.png)