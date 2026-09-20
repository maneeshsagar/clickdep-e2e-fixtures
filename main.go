package main
import ("encoding/json";"fmt";"log";"net/http";"os";"strings")
func main(){
  fmt.Println("STARTUP-MARKER go")
  http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
    if r.URL.Path=="/log"{log.Println("E2E-LOG-MARKER")}
    env:=map[string]string{}
    for _,e:=range os.Environ(){ if strings.HasPrefix(e,"E2E_"){kv:=strings.SplitN(e,"=",2);env[kv[0]]=kv[1]} }
    json.NewEncoder(w).Encode(map[string]any{"lang":"go","path":r.URL.Path,"env":env})
  })
  log.Fatal(http.ListenAndServe("0.0.0.0:8080",nil))
}
