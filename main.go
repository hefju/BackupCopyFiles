package main

import (
	"fmt"
	"github.com/hefju/BackupCopyFiles/tools/gofile"
	"github.com/hefju/BackupCopyFiles/tools/setting"
	"time"
)

func main() {

    var end string

	setting.LoadProfile() //读取配置文件
	//    fmt.Println(setting.AppConfig.SourcePath)
	//    fmt.Println(setting.AppConfig.TargetPath)
    if !gofile.IsDirExists(setting.AppConfig.SourcePath){
    fmt.Println("运行错误! 源文件夹不存在:",setting.AppConfig.SourcePath," 按任意键结束程序.")
        fmt.Scanln(&end)
        return
    }

    t1 := time.Now()
    tagrget:=setting.AppConfig.TargetPath+"/"+t1.Format("20060102")
    if !gofile.CreatePath(tagrget){
        fmt.Println("运行错误! 目标文件夹创建失败:",tagrget," 按任意键结束程序.")
        fmt.Scanln(&end)
        return
    }






	filelist := gofile.GetAllFiles5(setting.AppConfig.SourcePath)
	gofile.CopyFiles(filelist, setting.AppConfig.SourcePath, tagrget)

	fmt.Println("复制成功,按任意键结束. 耗时:", time.Now().Sub(t1)) //"press any key to continue")
	// bufio.NewReader(os.Stdin).ReadBytes('\n')

	//fmt.Scanln(&end)
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
