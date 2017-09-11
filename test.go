package main
 
import (
    "fmt"
    "log"
    "net/http"
)
 
func sayHello(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hello World!")
 
}
func main() {
    http.HandleFunc("/hello", sayHello)　　//注册URI路径与相应的处理函数
    er := http.ListenAndServe(":9090", nil)  // 监听9090端口，就跟javaweb中tomcat用的8080差不多一个意思吧
    if er != nil {
        log.Fatal("ListenAndServe: ", er)
    }
}