# autojump

1. 安装

   - `brew install autojump`
   - 修改 .zshrc文件
     - plugin中添加 autojump
     - 结尾处添加一行 `[[ -s $(brew --prefix)/etc/profile.d/autojump.sh ]] && . $(brew --prefix)/etc/profile.d/autojump.sh`

2. 使用

   - `j -a 目录昵称 目标目录`

   - `j 目录昵称即可`

   - 比如: 

     > `j -a blog /home/blog`
     >
     > `j blog` 此时即可跳转到 /home/blog 文件目录中 配合编辑器使用,简直效率神器

