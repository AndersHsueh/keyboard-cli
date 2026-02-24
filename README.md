# keyboard-cli

Virtual keyboard CLI - 通过命令行模拟键盘输入

## 功能

- **type**: 输入字符串 `keyboard-cli type "hello, world"`
- **key**: 发送按键/组合键 `keyboard-cli key ctrl+o`
- **list-keys**: 查看支持的按键列表

## 安全

⚠️ **Ctrl+Alt+Del 组合键被拦截**，无法执行

## 使用方法

```bash
# 安装
go install

# 或直接运行
./keyboard-cli

# 输入字符串
./keyboard-cli type "hello world"

# 发送组合键
./keyboard-cli key ctrl+o
./keyboard-cli key alt+f4
./keyboard-cli key ctrl+shift+esc

# 发送单键
./keyboard-cli key enter
./keyboard-cli key tab
./keyboard-cli key f5
```

## 权限

需要 root 权限或加入 `input` 用户组来访问 `/dev/uinput`：

```bash
# 方法1: 使用 sudo
sudo ./keyboard-cli key ctrl+c

# 方法2: 加入 input 组
sudo usermod -aG input $USER
# 然后重新登录
```

## 技术

- 使用 Linux uinput 内核模块创建虚拟键盘设备
- Go 1.21+

## License

MIT
