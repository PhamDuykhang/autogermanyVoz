package app

import (
	"crypto/md5"
	"fmt"
	"github.com/PhamDuyKhang/autogermanyVoz/conf"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func LoginVOZ(useName, passWord string) ([]*http.Cookie, error) {
	md5Cry := md5.New()
	_, err := io.WriteString(md5Cry, passWord)
	if err != nil {
		logrus.Error("can't encode your password", err)
		return nil, err
	}
	passWord = fmt.Sprintf("%x", md5Cry.Sum(nil))
	form := url.Values{}
	form.Add("vb_login_username",useName)
	form.Add("vb_login_password","")
	form.Add("s","")
	form.Add("securitytoken","guest")
	form.Add("do","login")
	form.Add("cookieuser","1")
	form.Add("login","Log in")
	form.Add("vb_login_md5password",passWord)
	form.Add("vb_login_md5password_utf",passWord)
	client := GetHTTPClient()
	rq, err := http.NewRequest("POST",conf.Login_End_Point,strings.NewReader(form.Encode()))
	if err != nil {
		logrus.Error("can't encode request body", err)
		return nil, err
	}
	rq.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	rq.Header.Add("Content-Type","application/x-www-form-urlencoded")
	//rq.Header.Add("Referer", "https://forums.voz.vn")
	res, err := client.Do(rq)
	if err != nil {
		logrus.Println(err)
		return nil, err
	}
	err = res.Body.Close()
	if err!= nil{

	}
	return res.Cookies(), nil
}
func FindNewThread(threadURL chan <- string,boxNumber int)  {

}
