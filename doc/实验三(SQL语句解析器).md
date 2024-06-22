# 实验三 SQL 解析器设计实现

## 主要内容

SQL 语句，包括 Select 语句， Insert 语句和 Update，创建表语句的解析并对接实验2中对应的创建表、写入和查询函数实现数据表创建、写入和查询操作。

1. 构建语法解析器实现部分 SQL 语句，包括 Select 语句， Insert 语句和 Update，创建表语句的语法解析。
2. 构建语义解析器对 SQL 语句进行语义解析
3. 将解析的语义对接底层实验 2 中实现的各个数据操作函数

## 提示

1. 掌握掌握语法分析和语义分析知识语法分析和语义分析知识
2. 运用语法和语义分析知识实现简单运用语法和语义分析知识实现简单SQL 解析器解析器

## 程序设计思路

1. 将输入的字符串以 “;”分割成短的字符串，存入队列 Queue
2. 取队首字符串，对其进行词法分析，将字符串变为小写（ SQL不区分大小写），除去多余的空格，将其中的关键字、标识符、常数、界符以及运算符等进行编码和分类，并按照顺序存储在 ArrayList中。
3. 对 ArrayList中的 tokens进行语法分析，将 token从左到右逐个匹配，如果能找到一条路线完匹配 tokens，则完成语法分析，调用相对应的函数。
4.  回到 (2)，直到队列为空。

## 算法分析与设计

1. 预处理
   将输入的字符串全部转换为小写（SQL语句不区分大小写），并以 “；；”分割，存入 sentences队列。![image-20240623011544363](D:\Desktop\myfile\system\doc\assets\image-20240623011544363.png)

2. 词法分析
   逐个获取字符串中的字符，跳过空格，若为单个字符则判断其是否为“*”、“=”和分隔符号，若为多个字母或数字则将其连接起来，并判断其组合是标识符、常数还是关键字，并将这些tokens对应。直到读完整个字符串：![image-20240623012015364](D:\Desktop\myfile\system\doc\assets\image-20240623012015364.png)

​	Tokens按照读取顺序存入：![image-20240623011702382](D:\Desktop\myfile\system\doc\assets\image-20240623011702382.png)

​	其中，word的数据结构为：![image-20240623011711494](D:\Desktop\myfile\system\doc\assets\image-20240623011711494.png)

3. 语法分析
   Token流存储在words中后，sentenceAnalyze类中的checkSentence会对token流进行语法分析。这里使用的是递归下降分析法，这一过程基于match()函数实现，Match函数对每个token进行正则匹配：

   ![image-20240623011723726](D:\Desktop\myfile\system\doc\assets\image-20240623011723726.png)其中，regex为正则表达式，p为int类型的index，初始值为0，若token匹配正则表达式，则p自增，return true，否	return false。与match函数同理的还有用于匹配标识符的id()函数，用于匹配一个或多个以“,”分隔的标识符的idList()函数等。在上述匹配函数的基础上，checkSentence函数得以实现。以SELECT语句为例，使用一系列的match函数对token流进行匹配。![image-20240623011733641](D:\Desktop\myfile\system\doc\assets\image-20240623011733641.png)其中，readAll会读取指定表中所有列的数据，而read则会根据从idList中传入的列读取这些列中的数据。两个函数都会连接到数据库，对数据库中的数据进行操作。UPDATE、DELETE、INSERT与SELECT同理，此处不再赘述。若无匹配，则会输出“grammarerror”。根据类型对标识符、常数、数据类型进行匹配：![image-20240623011744708](D:\Desktop\myfile\system\doc\assets\image-20240623011744708.png)一个或者多个以逗号分隔的标识符：![image-20240623011753980](D:\Desktop\myfile\system\doc\assets\image-20240623011753980.png)（List类的boolean函数同理，此处不再赘述。)

   对token流进行检查，如果:![image-20240623011804263](D:\Desktop\myfile\system\doc\assets\image-20240623011804263.png)执行对应的read函数。![image-20240623011813171](D:\Desktop\myfile\system\doc\assets\image-20240623011813171.png)执行write函数。![image-20240623011820550](D:\Desktop\myfile\system\doc\assets\image-20240623011820550.png)执行创建表createTable函数。![image-20240623011831128](D:\Desktop\myfile\system\doc\assets\image-20240623011831128.png)执行对应的updateTable函数。

## 数据结构

### Token

每个token都有名字String name以及类型int type

![image-20240623011908096](D:\Desktop\myfile\system\doc\assets\image-20240623011908096.png)

###  Sentences队列

断句后，句子字符串按照顺序存储在sentences队列中`Queue<String> sentences = new LinkedList<String>();`
每次取队首句子进行词法分析以及语法分析，直到队空：

![image-20240623011922317](D:\Desktop\myfile\system\doc\assets\image-20240623011922317.png)

## 实现

用户在界面上方的输入框内输入单个SQL语句（包括 SELECT、INSERT、 UPDATE、 DELETE语句），或者输入多个 SQL语句，以 “;”间隔，点击 “解析运行 ”按钮， sql解析器将对输入的语句进行解析，若语法没有错误则会执行对应的增删查改操作 。例，输入 select * from goodslist;查询语句，在控制台有词法分析结果，客户端有查询结果。插入、建表、更新语句同 select语句。

![image-20240623011939323](D:\Desktop\myfile\system\doc\assets\image-20240623011939323.png)