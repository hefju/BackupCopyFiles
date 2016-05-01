package gofile
import (
    "fmt"
    "os"
    "path/filepath"
//    "strconv"
 //   "time"
    "io"
    "strings"
//    "io/ioutil"
)

//文件信息
type FileInfo struct {
    Path        string // 文件所在的路径
    Info   os.FileInfo        // 文件的 os.FileInfo 信息
}

//遍历文件夹
func GetAllFiles5(path string)(chan FileInfo){
    fileArray:=make(chan FileInfo,50)
    go  WalkingFiles(path, fileArray)
    return  fileArray
}

//遍历文件
func WalkingFiles(inputPath string, fileArray chan FileInfo) {

    filepath.Walk(inputPath,
    func(path string, f os.FileInfo, err error) error {
        if f == nil {
            return err
        }
        fileArray<- FileInfo{Path:path, Info:f}
//        if f.IsDir() {
//               return nil
//        }else{
//            fileArray<-path//+f.Name()
//        }
        return nil
    })
    close(fileArray)
}

//复制文件
func CopyFiles(fileArray chan FileInfo,sourcepath,targetpath string)  {
//    err:= os.MkdirAll(targetpath, 0777)
//    if err!=nil{
//        fmt.Println("创建文件夹失败!!:",targetpath)
//        return
//    }

    count,success:=1,1   //上面创建目标文件夹算是一个,遍历的第一个还是它所以重复计算了
    for elem:=range fileArray {
//        fmt.Println(elem.Path)
        dst:=GetTarget(elem.Path,sourcepath,targetpath)
//        fmt.Println(count," src:",elem.Path," dst:",dst)
        if elem.Info.IsDir(){
            err:= os.MkdirAll(dst, 0777)
            if err != nil {
                fmt.Println("创建文件夹失败!!:",dst)
            }else{
                success++
                fmt.Println(count,"复制成功:",elem.Path)
            }

        }else {
            _ , err :=  CopyFile(elem.Path,dst)//复制文件
            if err!=nil{
                fmt.Println("错误:",err)
            }else{
                success++
                fmt.Println(count,"复制成功:",elem.Path)
            }
        }
        count++
    }
    fmt.Println("总文件数量:",count-1," 成功复制:",success-1)//count要减去第一个文件夹
}
//组合目标文件路径, 主要是windows的斜杠问题
func  GetTarget(filepath,sourcepath,targetpath string)string{
    //在windows下, 斜杠反了.
    f:=func(r rune)rune{
        if r=='/'{
            return '\\'
        }
        return r
    }
    sourcepath=strings.Map(f,sourcepath)
    targetpath=strings.Map(f,targetpath)
//    fmt.Println(sourcepath)
//    fmt.Println(targetpath)

    filepath=strings.Replace(filepath,sourcepath,targetpath,1)
    return filepath
}

//复制文件
func CopyFile(src,dst string)(w int64,err error){
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

//判断文件夹是否存在
func IsDirExists(path string) bool {
    fi, err := os.Stat(path)

    if err != nil {
        return os.IsExist(err)
    } else {
        return fi.IsDir()
    }
    panic("not reached")
}

//创建文件夹
func CreatePath(path string) bool{
    err:= os.MkdirAll(path, 0777)
    if err!=nil{
        fmt.Println("创建文件夹失败!!:",err)
        return false
    }
    return true
}



