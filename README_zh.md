# Nunu — A CLI tool for building go aplication.


Nunu是一个基于Golang的应用脚手架，它的名字来自于英雄联盟中的游戏角色，一个骑在雪怪肩膀上的小男孩。和努努一样，该项目也是站在巨人的肩膀上，它是由Golang生态中各种非常流行的库整合而成的，它们的组合可以帮助你快速构建一个高效、可靠的应用程序。

[英文介绍](https://github.com/go-nunu/nunu/blob/main/README.md)

![Nunu](https://github.com/go-nunu/nunu/blob/main/.github/assets/banner.png)

## 文档
* [使用指南](https://github.com/go-nunu/nunu/blob/main/docs/zh/guide.md)
* [分层架构](https://github.com/go-nunu/nunu/blob/main/docs/zh/architecture.md)
* [上手教程](https://github.com/go-nunu/nunu/blob/main/docs/zh/tutorial.md)
* [高效编写单元测试](https://github.com/go-nunu/nunu/blob/main/docs/zh/unit_testing.md)

## 许可证

Nunu是根据MIT许可证发布的。有关更多信息，请参见[LICENSE](LICENSE)文件。

```shell
PS C:\Users\Administrator\Desktop> nunu new go-gravatar
? Please select a layout: Advanced
git clone https://github.com/go-nunu/nunu-layout-advanced.git
go mod tidy
go install github.com/google/wire/cmd/wire@latest
go: downloading github.com/google/wire v0.6.0
go: downloading github.com/google/subcommands v1.2.0
go: downloading golang.org/x/tools v0.17.0
go: downloading golang.org/x/mod v0.14.0

 _   _
| \ | |_   _ _ __  _   _
|  \| | | | | '_ \| | | |
| |\  | |_| | | | | |_| |
|_| \_|\__,_|_| |_|\__,_|

A CLI tool for building go aplication.

🎉 Project go-gravatar created successfully!

Done. Now run:

› cd go-gravatar
› nunu run
```
