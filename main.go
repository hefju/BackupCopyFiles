package main
import (
    "fmt"
    "github.com/hefju/BackupCopyFiles/tools/setting"
    "github.com/hefju/BackupCopyFiles/tools/gofile"
  //  "time"
)
func main(){

    setting.LoadProfile()//读取配置文件

    count:=0
  list:=  gofile.GetAllFiles5(setting.AppConfig.OriginalPath)
    for elem:=range list{
       // time.Sleep(time.Microsecond*10)
      //  fmt.Println(elem)
        gofile.CopyFiles(elem,setting.AppConfig.TargetPath)
        count++
    }
    fmt.Println("总文件数量:",count)
//list:=gofile.GetAllFiles2(setting.AppConfig.ObjectPath)
//    for _,v:=range list{
//        fmt.Println(v)
//    }

   // time.Sleep(time.Second*5)
    fmt.Println("press any key to continue")
   // bufio.NewReader(os.Stdin).ReadBytes('\n')
    var end string
    fmt.Scanln(&end)
}

//func copyfiles(done2 chan bool){
//    time.Sleep(time.Second*1)
//  //  files:=make(chan string )
//    files,done:=  gofile.GetAllFiles(setting.AppConfig.ObjectPath)
//    for  {
//        if <-done{
//            done2<-true
//            break;
//        }
//        fmt.Println(<-files)
//    }
//}
