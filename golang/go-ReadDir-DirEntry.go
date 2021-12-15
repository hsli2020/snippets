// Go 1.16 即将到来的函数：ReadDir 和 DirEntry 
// 
// 为什么需要ReadDir？ 简短的答案是：性能。
// 
// 当调用读取文件夹路径的系统函数时，操作系统一般会返回文件名_和_它的类型（在Windows下，
// 还包括如文件大小和最后修改时间等的stat信息）。然而，原始版本的Go和Python接口会丢掉这
// 些额外信息，这就需要在读取每个路径时再多调用一个stat。系统调用的性能较差 ，stat 可能
// 从磁盘、或至少从磁盘缓存读取信息。
// 
// 在循环遍历目录树时，你需要知道一个路径是文件还是文件夹，这样才可以知道循环遍历的方式。
// 因此即使一个简单的目录树遍历，也需要读取文件夹路径并获取每个路径的stat信息。但如果使
// 用操作系统提供的文件类型信息，就可以避免那些stat系统调用，同时遍历目录的速度也将提高
// 几倍（在网络文件系统上甚至可以快十几倍）。具体信息可以参考Python版本的基准测试。

// 在Go语言中，调用os.ReadDir(path)，将会返回一个os.DirEntry对象的切片，如下所示：

type DirEntry interface {
    Name() string
    IsDir() bool
    Type() FileMode
    Info() (FileInfo, error)
}

// 在Go中（一旦1.16发布），GetTreeSize对应的函数如下所示：

func GetTreeSize(path string) (int64, error) {
    entries, err := os.ReadDir(path)
    if err != nil {
        return 0, err
    }
    var total int64
    for _, entry := range entries {
        if entry.IsDir() {
            size, err := GetTreeSize(filepath.Join(path, entry.Name()))
            if err != nil {
                return 0, err
            }
            total += size
        } else {
            info, err := entry.Info()
            if err != nil {
                return 0, err
            }
            total += info.Size()
        }
    }
    return total, nil
}

//Go （pre-1.16版本）语言中有一个相似的函数，filepath.Walk，但不幸的是 FileInfo 接口的设计
//无法支持各种方法调用时的错误报告。为了获得性能提升， filepath.Walk 的调用需要改成 
//filepath.WalkDir ——尽管非常相似，但遍历函数的参数是DirEntry 而不是 FileInfo。

//下面的代码是Go版本的使用现有filepath.Walk 函数的list_non_dot ：

func ListNonDot(path string) ([]string, error) {
    var paths []string
    err := filepath.Walk(path, func(p string, info os.FileInfo,
                                    err error) error {
        if strings.HasPrefix(info.Name(), ".") {
            if info.IsDir() {
                return filepath.SkipDir
            }
            return err
        }
        if !info.IsDir() {
            paths = append(paths, p)
        }
        return err
    })
    return paths, err
}

//当然，在Go 1.16中这段代码也可以运行，但如果你想得到性能收益就需要做少许修改——在上面的代码
//中仅需要把 Walk 替换为 WalkDir，并把 os.FileInfo 替换成 os.DirEntry：

err := filepath.WalkDir(path, func(p string, info os.DirEntry,
