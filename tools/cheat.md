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



### 问题: No cheatsheet found for "command"
经常会出现这种提示,是因为没有设置全局cheatsheets
编辑配置文件 `vim ~/.config/cheat/conf.yml`

```bash
---
# The editor to use with 'cheat -e <sheet>'. Defaults to $EDITOR or $VISUAL.
editor: vim

# Should 'cheat' always colorize output?
colorize: true

# Which 'chroma' colorscheme should be applied to the output?
# Options are available here:
#   https://github.com/alecthomas/chroma/tree/master/styles
style: monokai

# Which 'chroma' "formatter" should be applied?
# One of: "terminal", "terminal256", "terminal16m"
formatter: terminal16m

# Through which pager should output be piped? (Unset this key for no pager.)
pager: less -FRX

# The paths at which cheatsheets are available. Tags associated with a cheatpath
# are automatically attached to all cheatsheets residing on that path.
#
# Whenever cheatsheets share the same title (like 'tar'), the most local
# cheatsheets (those which come later in this file) take precedent over the
# less local sheets. This allows you to create your own "overides" for
# "upstream" cheatsheets.
#
# But what if you want to view the "upstream" cheatsheets instead of your own?
# Cheatsheets may be filtered via 'cheat -t <tag>' in combination with other
# commands. So, if you want to view the 'tar' cheatsheet that is tagged as
# 'community' rather than your own, you can use: cheat tar -t community
cheatpaths:

  # Paths that come earlier are considered to be the most "global", and will
  # thus be overridden by more local cheatsheets. That being the case, you
  # should probably list community cheatsheets first.
  #
  # Note that the paths and tags listed below are placeholders. You may freely
  # change them to suit your needs.
  #
  # Community cheatsheets must be installed separately, though you may have
  # downloaded them automatically when installing 'cheat'. If not, you may
  # download them here:
  #
  # https://github.com/cheat/cheatsheets
  #
  # Once downloaded, ensure that 'path' below points to the location at which
  # you downloaded the community cheatsheets.
  - name: community
    path: /Users/wangkeji/.config/cheat/cheatsheets/community
    tags: [ community ]
    readonly: true

  # If you have personalized cheatsheets, list them last. They will take
  # precedence over the more global cheatsheets.
  - name: personal
    path: /Users/wangkeji/.config/cheat/cheatsheets/personal
    tags: [ personal ]
    readonly: false

  # 新增自己的cheats全局使用目录,后续都存放在本文件夹里面
  - name: cheats
    path: /Users/wangkeji/cheats
    tags: [ cheats ]
    readonly: false

  # While it requires no configuration here, it's also worth noting that
  # 'cheat' will automatically append directories named '.cheat' within the
  # current working directory to the 'cheatpath'. This can be very useful if
  # you'd like to closely associate cheatsheets with, for example, a directory
  # containing source code.
  #
  # Such "directory-scoped" cheatsheets will be treated as the most "local"
  # cheatsheets, and will override less "local" cheatsheets. Likewise,
  # directory-scoped cheatsheets will always be editable ('readonly: false').
```
