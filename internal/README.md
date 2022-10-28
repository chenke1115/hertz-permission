<!--
 * @Author: changge <changge1519@gmail.com>
 * @Date: 2022-08-18 13:55:55
 * @LastEditTime: 2022-08-18 14:20:23
 * @Description: Do not edit
-->

# 私有应用程序和库代码
    - 这种布局模式是由 Go 编译器本身强制执行的
    - 不仅限于顶级internal目录。可以在项目树的任何级别拥有多个internal目录。
    - /internal/apiserver：该目录中存放真实的应用代码。这些应用的共享代码存放在/internal/pkg 目录下。
    - /internal/pkg：存放项目内可共享，项目外不共享的包。这些包提供了比较基础、通用的功能，例如工具、错误码、用户验证等功能。

