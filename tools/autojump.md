# autojump

1. mac 安装

   - `brew install autojump`
   - 修改 .zshrc文件
     - plugin中添加 autojump
     - 结尾处添加一行 `[[ -s $(brew --prefix)/etc/profile.d/autojump.sh ]] && . $(brew --prefix)/etc/profile.d/autojump.sh`
2. centos 安装
    - `yum install autojump`
    - `echo "/usr/share/autojump/autojump.bash" >> ~/.bashrc`
    - `chmod 777 /usr/share/autojump/autojump.bash` 

