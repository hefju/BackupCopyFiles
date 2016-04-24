package main
import (
    "fmt"
    "github.com/hefju/BackupCopyFiles/tools/setting"
    "github.com/hefju/BackupCopyFiles/tools/gofile"
    "time"
)
func main(){

    setting.LoadProfile()//读取配置文件
    //fmt.Println(setting.AppConfig.ObjectPath)

//    done:=make(chan bool)
// go copyfiles(done)
//    <-done
  fs:=  gofile.GetAllFiles2(setting.AppConfig.ObjectPath)
    for _,v:=range fs{
        fmt.Println(v)
    }
   // time.Sleep(time.Second*5)
    fmt.Println("end")

}

func copyfiles(done2 chan bool){
    time.Sleep(time.Second*1)
  //  files:=make(chan string )
    files,done:=  gofile.GetAllFiles(setting.AppConfig.ObjectPath)
    for  {
        if <-done{
            done2<-true
            break;
        }
        fmt.Println(<-files)
    }
}
