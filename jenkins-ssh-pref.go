package main // import "jenkins-ssh-pref"
import (
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	// GET 호출
	resp, err := http.Get("http://myserver/ip-info.php")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	// IP 취득
	urlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// 환경변수 읽기 - 경로 준비
	userProfile := os.Getenv("UserProfile")
	userProfile = strings.Replace(userProfile, "\\", "/", -1) + "/.jenkins"
	filePath := userProfile + "/" + "jenkins.plugins.publish_over_ssh.BapSshPublisherPlugin.xml"

	// 파일 읽기
	// bytes, err := ioutil.ReadFile("./possh.xml")
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	re := regexp.MustCompile("<hostname>([0-9|.]+)</hostname>")

	modified := re.ReplaceAllString(string(bytes), "<hostname>"+string(urlData)+"</hostname>")

	// 파일 쓰기
	// err = ioutil.WriteFile("./possh_mod.xml", []byte(modified), 0)
	err = ioutil.WriteFile(filePath, []byte(modified), 0)
	if err != nil {
		panic(err)
	}

}
