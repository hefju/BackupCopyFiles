package setting
import (
    "os"
    "log"
    "encoding/json"
)
var AppConfig Config
type Config struct {
    OriginalPath string
    TargetPath string
}
func LoadProfile(){
    file:="conf.json"
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
    log.Println("load setting from "+file)
}