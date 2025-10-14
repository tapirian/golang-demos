package main

import (
    "fmt"
    "log"
    "os"

    "github.com/fsnotify/fsnotify"
)

func main() {
    // 创建新的 watcher
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    done := make(chan bool)

    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                fmt.Println("事件:", event)

                if event.Op&fsnotify.Create == fsnotify.Create {
                    fmt.Println("创建文件:", event.Name)
                }
                if event.Op&fsnotify.Write == fsnotify.Write {
                    fmt.Println("修改文件:", event.Name)
                }
                if event.Op&fsnotify.Remove == fsnotify.Remove {
                    fmt.Println("删除文件:", event.Name)
                }
                if event.Op&fsnotify.Rename == fsnotify.Rename {
                    fmt.Println("重命名文件:", event.Name)
                }
                if event.Op&fsnotify.Chmod == fsnotify.Chmod {
                    fmt.Println("修改权限:", event.Name)
                }

            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                fmt.Println("错误:", err)
            }
        }
    }()

    watchDir := "./watchdir"

    // 创建一个用于测试的目录
    if _, err := os.Stat(watchDir); os.IsNotExist(err) {
        os.Mkdir(watchDir, 0755)
    }

    // 添加目录到 watcher
    err = watcher.Add(watchDir)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("开始监听目录:", watchDir)
    fmt.Println("你可以在该目录下创建、修改、删除文件来测试")

    // 阻塞程序运行
    <-done
}

