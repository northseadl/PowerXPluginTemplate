# 插件开发

## 快速开始

### 初始化

开始开发插件，首先需要初始化插件模板。您可以通过运行以下命令创建插件模板, 运行指令以后根据交互提示完成创建

```
# windows
./pluginctl.exe new

# linux
./pluginctl new

# mac
./pluginctl new
```

指令执行以后创建过程如以下示例:
```
$ [Step1]: Please input project name, only support hyphen format, must be lowercase, e.g. powerx-plugin
$ test-project-plugin
$ [Step2]: Please input project description, e.g. PowerX plugin
$ test
$ [Step3]: Please input menu name, e.g. PowerX plugin
$ test
$ [Step4]: Please input icon, e.g. icon-plus, refer to https://arco.design/vue/component/icon
$ icon-plus
$ [Step5]: Cloning template...
$ [Step6]: Preparing...
$ Done!
```
这将会在当前目录初始化出一个插件项目，您可以在该项目文件夹内找到以下几个主要文件和文件夹：

*   **backend**：后端代码目录
*   **frontend**：前端代码目录
*   **ctl**：插件指令目录（未来可能会移除）
*   **plugin.yaml**：插件描述文件

接下来，您需要验证**plugin.yaml**文件中的插件配置是否正确。以下是一个示例：

    name: plugin-example                   		# 插件名, 连缀符格式, 是插件的唯一标识 
    description: This is a plugin example  		# 描述, 插件的功能描述 
    version: 0.0.1 								# 版本号, 插件的版本号, 同样的插件仅运行一个版本运行 
    meta: 										# 插件元数据 
      icon: icon-plus 								# 插件在前端显示的icon, 目前仅支持arco内置icon 
      locale: 插件示例 								# 插件在左侧菜单栏显示的菜单名


到此, 插件项目的初始化已经完成

### 插件后端开发

插件的后端与 gozero单体应用的开发方式基本一致, 只需要定义 api, 然后运行 goctl 生成服务端代码, 进行开发即可, pkg包里提供了 PowerX 的 OpenAPI 用于和主程序交互, 以下是后端开发主要涉及的目录或文件

*   **plugin.api**: 后端api定义, 定义好后在项目目录执行 \`pluginctl build backend-api\`
*   **etc/etc.yaml**: 配置文件, 只能以{K,V}方式配置, 不允许嵌套

#### 接口调试:

1.  启动时指定 mode 为 dev, 将会在固定端口启动该服务, 并且读取etc目录下的配置
2.  从主程序获取token, 在http调试器设置同样的认证头即可调试接口

## 插件前端开发

插件前端主要有以下目录, 其他目录可以不管, 开发方式与 arco vue项目一致

*   **src/api**: api文件定义, 注意需要在后端api定义的基础上增加prefix, prefix为'/{plugin\_name}'
*   **src/router/routes/modules**: 路由配置
*   **src/views/{plugin\_name}**: 视图目录

## 构建插件

构建插件时，您可以使用ctl指令将插件代码打包为压缩包。需要注意的是，对应平台只能打包对应平台的插件压缩包。

```
# windows
./pluginctl.exe build plugin

# linux
./pluginctl build plugin

# mac
./pluginctl build plugin
```

## 使用插件

将构建得到的压缩包放在应用同级目录下的 plugins/.archive 文件夹(如果没有, 自行创建该目录), 然后运行主应用中的插件指令
```
// todo
```

等待构建完成以后, 重新启动 PowerX, 插件将会被加载到PowerX中

注意：前端插件将在另一个端口启动，并在启动后将其输出到日志中。请确保将该端口纳入安全组策略以实现正确的网络通信。

## 示例项目:

*   [微信公众号菜单](https://github.com/northseadl/WePublicMenu)

