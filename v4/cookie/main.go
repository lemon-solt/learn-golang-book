package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:  "f1",
		Value: "value-1",
	}

	c2 := http.Cookie{
		Name:  "f2",
		Value: "value-2",
	}

	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())

}

/*参照渡し setCooki関数*/
func deepSetCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:  "f1-deep",
		Value: "value-1",
	}

	c2 := http.Cookie{
		Name:  "f2-deep",
		Value: "value-2",
	}

	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)

}

func 取得(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
}

func getCookies(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("f1")
	if err != nil {
		fmt.Fprintln(w, "該当cookieはありません")
	}

	cs := r.Cookies()
	fmt.Println("cookie: ", c1)
	fmt.Println("cookies: ", cs)
}

func flashSetCookies(w http.ResponseWriter, r *http.Request) {
	msg := []byte("愛はあるんか")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}

	http.SetCookie(w, &c)
	fmt.Println("set cookie")
}

func showFlashCookies(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "クッキーが無い")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func main() {
	server := http.Server{
		Addr: ":7001",
	}
	fmt.Printf("run serve")

	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/deepSetCookie", deepSetCookie)
	http.HandleFunc("/取得", 取得)
	http.HandleFunc("/getCookiess", getCookies)
	http.HandleFunc("/flashSetCookies", flashSetCookies)
	http.HandleFunc("/showFlashCookies", showFlashCookies)
	server.ListenAndServe()
}
