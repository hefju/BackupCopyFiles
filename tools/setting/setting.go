package setting
import (
    "os"
    "log"
    "encoding/json"
)
var AppConfig Config
type Config struct {
    SourcePath string//源文件夹
    TargetPath string//目标文件夹
}
func LoadProfile(){
   // path,_:=os.Getwd()
    file:="C:\\Intel\\BackupCopyFiles\\BackupCopyFiles.json"
    r,err:=os.Open(file)
    if err!=nil{
        log.Fatalln(err)
    }
    decoder:=json.NewDecoder(r)
    var c Config
    err=decoder.Decode(&c)
    if err!=nil{
        log.Fatalln(err)
    }
    AppConfig=c
  //  log.Println("load setting from "+file)
}