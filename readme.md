# Ferret
Open-source Vulkan binding for Windows and Linux-based systems.  
This library provides full functional from creating window to rendering Vulkan-based applications.

### Windows
Package `windows` provides binding to a part of Win32 API for simple creating windows from go-code without direct work 
with unsafe and cgo.

### Linux
Package `linux` supports X11 window system and linux opengl impl.

### MacOS
Maybe in the future... General problem is difference between old Mac OS and new Mac OS for M1 chips. Apple have full 
control under used in Mac OS and iOS technologies and can stop Vulkan support in favor of Metal framework.