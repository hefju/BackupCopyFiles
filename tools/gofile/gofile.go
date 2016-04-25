package gofile
import (
    "fmt"
    "os"
    "path/filepath"
//    "strconv"
 //   "time"
    "io"
)

type  FileNode struct {
    FullPath ,Filename string
}
func GetAllFiles5(path string)(chan  *FileNode){
    fileArray:=make(chan *FileNode,50)
    go  WalkingFiles(path, fileArray)
    return  fileArray
}

//遍历文件
func WalkingFiles(inputPath string, fileArray chan *FileNode) {
    filepath.Walk(inputPath,
    func(path string, f os.FileInfo, err error) error {
        if f == nil {
            return err
        }
        if f.IsDir() {
            return nil
        }
        fn:=&FileNode{FullPath:path+f.Name(),Filename:f.Name()}
        fileArray<-fn
//         filepath:=path+f.Name()
//        fileArray<-filepath
        return nil
    })
    close(fileArray)
}
func CopyFiles(fileArray chan  *FileNode,targetpath string)  {
    fn:=<-fileArray
    dst:=targetpath+fn.Filename
    CopyFile(fn.FullPath,dst)
}
//复制文件
func CopyFile(src,dst string)(w int64,err error){
    fmt.Println("dst:",dst)
    panic("copyfile")
    srcFile,err := os.Open(src)
    if err!=nil{
        fmt.Println(err.Error())
        return
    }
    defer srcFile.Close()

    dstFile,err := os.Create(dst)

    if err!=nil{
        fmt.Println(err.Error())
        return
    }

    defer dstFile.Close()

    return io.Copy(dstFile,srcFile)
}



func GetAllFiles2(path string)[]string{
    fmt.Println(path)
    files:=make([]string,0)
    CalcFilesSync(path, files)
    return files
}

func CalcFilesSync(inputPath string,files []string) {
    filepath.Walk(inputPath,
        func(path string, f os.FileInfo, err error) error {
            if f == nil {
                return err
            }
            if f.IsDir() {
                return nil
            }
            files= append(files,f.Name())
            fmt.Println("CalcFilesSync "+f.Name())
            return nil
        })
}