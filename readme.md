# 反转视频颜色处理工具

这是一个用Go语言编写的命令行工具,可以使用FFmpeg反转视频文件的颜色。

功能特点:

- 可以处理单个视频文件或整个文件夹
- 支持多种视频格式(.mp4, .avi, .mov, .mkv, .flv, .wmv)
- 并发处理以提高性能
- 处理目录时保留文件夹结构

使用前提:

- Go 1.15或更高版本
- 已安装FFmpeg并可在系统PATH中使用

安装方法:

1. 克隆此仓库:
   git clone https://github.com/你的用户名/reverse-color-video.git

2. 进入项目目录:
   cd reverse-color-video

3. 构建可执行文件:
   go build -o reverse-color

使用方法:

`reverse-color [文件/文件夹...]`

您可以提供一个或多个文件或文件夹路径作为参数。例如:

`reverse-color video1.mp4 video2.avi folder1 folder2`

- 处理单个文件时,输出将保存在同一目录下,文件名后会添加"_reversed"。
- 处理文件夹时,将创建一个新文件夹,名称为原文件夹名加上"_reversed",并保持原始文件夹结构。

工作原理:

该工具使用FFmpeg对视频应用颜色反转效果。它使用以下视频滤镜:

`lutrgb='r=255*0.9-val*0.9+25:g=255*0.9-val*0.9+25:b=255*0.9-val*0.9+25',eq=brightness=-0.05:contrast=1.1`

此滤镜反转颜色,同时保持一些亮度和对比度调整,以获得更好的视觉效果。

贡献:

欢迎贡献!请随时提交Pull Request。

许可证:

本项目采用MIT许可证 - 详情请参阅LICENSE文件。