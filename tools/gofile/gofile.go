package gofile
import (
    "fmt"
    "os"
    "path/filepath"
//    "strconv"
 //   "time"
)

func GetAllFiles2(path string)[]string{
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


func GetAllFiles(path string)(chan string,chan bool){
    // file, _ := os.Getwd()
    // fmt.Println("current path:", file)
    fileArray:=make(chan string,10)

    done := make(chan bool)
  go  CalcFiles(path, done,fileArray)

   // println("loading files...")

   // time.Sleep(time.Second)
  //  println("loading finish!")
   // <-done
    return  fileArray,done
}


func CalcFiles2(inputPath string, fileArray chan string) {
    file := inputPath
    filepath.Walk(file,
    func(path string, f os.FileInfo, err error) error {
        if f == nil {
            return err
        }
        if f.IsDir() {
            return nil
        }
        fileArray<-f.Name()
        return nil
    })
}

func CalcFiles(inputPath string, done chan bool,fileArray chan string) {
    file := inputPath
//    var filesize int64 = 0
//    fileCount := 0
//    folderCount := 0
    filepath.Walk(file,
    func(path string, f os.FileInfo, err error) error {
        if f == nil {
            return err
        }
        if f.IsDir() {
            //println(path)
           // folderCount++
            return nil
        }
//        filesize += f.Size()
//        fileCount++
        fileArray<-f.Name()
        fmt.Println(f.Name())
        //println(f.Name())
        return nil
    })
//    fmt.Println("文件夹:", inputPath, "文件夹数量:", folderCount, "文件数量:", fileCount)
//    mb := strconv.FormatFloat(float64(filesize)/1024.0/1024.0, 'f', 2, 32)
//    gb := strconv.FormatFloat(float64(filesize)/1024.0/1024.0/1024.0, 'f', 2, 32)
//    fmt.Println("总大小:", mb, "MB ", gb, "GB")
    done <- true
}