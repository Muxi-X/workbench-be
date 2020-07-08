# Feed Service & Subscribe Service

包含 feed service 和 subscribe service

subscribe service： feed数据写入数据库

### Start service locally

```shell
make

# feed service
./main

# subscribe service
./main -sub
```

### Feed Structure

```text
{
    "user": {   
        "name": "string",
        "id": "integer",
        "avatar_url": "string"
    },
    "action": "string", // 用户执行的动作，直接给出的就是可以使用的动作名，包括 "加入", "创建", "编辑", "删除", "评论", "移动" 
    "feed_id": "integer",
    "if_split": "boolean", // 是否有分割线
    "source": {
        "kind_id": "integer",
        "object_id": "integer",
        "object_name": "string",
        "project_id": "integer", // 没有则为-1，如进度就没有project_id
        "project_name": "string" // project名，没有为"noname"
    },
    "time_day": "2000/01/01",
    "time_hm": "00:30"
}

source: feed的来源. 其中kind_id的对应关系如下
1 -> 团队
2 -> 项目
3 -> 文档
4 -> 文件
5 -> ???
6 -> 进度  // 估计是当时写进度的那个人数错了 把6当成进度了，5没有内容  = =
```

#### 分割线

需要分割的情况：

1. 第一条数据
2. 不同日期
3. 不同项目，根据source.kind_id判断
