# message-board

1.增加了个人信息接口，可以获取用户自己起的名字，还有其所填信息。
2.实现密码加盐，使用了自带的golang.org/x/crypto/bcrypt。
3.增加点赞功能，用户可以给留言或评论点赞，然后后台也会存储是哪个用户点的赞，在什么时间点的赞
