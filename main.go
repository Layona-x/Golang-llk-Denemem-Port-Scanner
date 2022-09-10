package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := "8000"
        path := "massalahcokzekiyim"
	fmt.Printf("Oluşturulacak portu giriniz (8000 için boş bırakın): ")
	fmt.Scanf("%s", &port)
	if port == "" {
		port = "8000"
	} else if len(port) < 4 {
		port = "8000"
		fmt.Println("Portu 4 karakterden az girdiğiniz için 8000 olarak ayarlandı")
	} else if len(port) > 4 {
		port = "8000"
		fmt.Println("Portu 4 karakterden fazla girdiğiniz için 8000 olarak ayarlandı")
  }
    fmt.Printf("Port Oluşturulması İçin Klasöre İhtiyaç Var Lutfen Dosya Uzantısı Olmadan Dosya İsminizi Giriniz")
    fmt.Scanf("%s",&path)
       if path == "" {
          path="massalahcokzekiyim"
   }
	http.Handle("/", http.FileServer(http.Dir(path)))
	fmt.Printf("http://localhost:%s oluşturuldu\n", port)
	http.ListenAndServe(":"+port, nil)
}
