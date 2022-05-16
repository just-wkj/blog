# cheat
cheat是一个简单的命令行笔记工具，可以帮助你快速查看某个命令或者笔记。

github地址: <https://github.com/cheat/cheat>.

- Show example usage of a command:
  > cheat command

- Edit the cheat sheet for a command:
  > cheat -e command

- List the available cheat sheets:
  > cheat -l

- Search available the cheat sheets for a specified command name:
  > cheat -s command

- Get the current cheat version:
  > cheat -v
 
### 新增cheat笔记
- 在 `~/.cheat` 或者 `~/.config/cheat/cheatsheets/personal`  文件夹下创建文件，文件名为你想要查看的命令名称，文件内容为你想要查看的命令的笔记。
- 直接使用 `cheat -e command` 即可编辑你想要查看的命令的笔记。 默认会存放在`~/.config/cheat/cheatsheets/personal`
比如 创建 ~/.cheat/redistest 文件
### 使用cheat笔记
`cheat redistest` 就可以查看到redistest的笔记了。
