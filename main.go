package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/tidwall/gjson"
	"github.com/yanyiwu/gojieba"
)

func GetHotHttpJson(url string) string {
	client := &http.Client{}
	apiURL := url

	req, err := http.NewRequest("GET", apiURL, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:99.0) Gecko/20100101 Firefox/99.0")
	if err != nil {
		fmt.Printf("post failed, err:%v\n\n", err)
		return ""
	}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed,err:%v\n\n", err)
		return ""
	}
	return string(b)
}

//https://v1.hitokoto.cn/?encode=json&c=d&c=j&c=k&c=k&c=i&lang=cn

const letters = `abcdefghjkmnpqrstuvwxyzABDEFGHJMNQRTZ23456789+!@#.,$%=/`

func randStr(n int) string {
	seeds := rand.New(rand.NewSource(time.Now().UnixNano()))
	n = seeds.Intn(n)
	b := make([]byte, n+1)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}

func parseJson(json string, jsonPath string) string {

	value := gjson.Get(json, jsonPath) //"data.#.title")
	// for _, name := range value.Array() {
	// 	println(name.String())
	// }
	//flag:
	seeds := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := seeds.Intn(len(value.Array()))
	arr := value.Array()
	keyword := arr[i].String()
	return keyword

}
func removeElement(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func parseWords(keyword string) string {
	var words []string
	use_hmm := true
	x := gojieba.NewJieba()
	defer x.Free()

	//words = x.CutAll(keyword)
	// fmt.Println(keyword)
	// fmt.Println("全模式:", strings.Join(words, "/"))

	words = x.Cut(keyword, use_hmm)
	//fmt.Println("精确模式:", strings.Join(words, "/"))
	s := ""
	for i := 0; i < len(words); i++ {
		seeds := rand.New(rand.NewSource(time.Now().UnixNano()))
		j := seeds.Intn(len(words))
		s = s + words[j] //+ randStr(len(words)-1)
		words = removeElement(words, j)
		i = i + j%6
	}
	//println(s)
	return s

}
func main() {
	//keywords := parseJson(GetHttpJson())
	json := `{"code":200,"data":[{"title":"习近平：以金砖担当开创美好未来","index":1,"hotValue":"499.9万","link":"https://www.baidu.com/s?wd=习近平：以金砖担当开创美好未来"},{"title":"日本核污染水已进入大海","index":2,"hotValue":"494.3万","link":"https://www.baidu.com/s?wd=日本核污染水已进入大海"},{"title":"海关总署：全面暂停进口日本水产品","index":3,"hotValue":"486.4万","link":"https://www.baidu.com/s?wd=海关总署：全面暂停进口日本水产品"},{"title":"解码“金砖＋”吸引力","index":4,"hotValue":"473.0万","link":"https://www.baidu.com/s?wd=解码“金砖＋”吸引力"},{"title":"核污水排海现场：海水呈两种颜色","index":5,"hotValue":"468.7万","link":"https://www.baidu.com/s?wd=核污水排海现场：海水呈两种颜色"},{"title":"瓦格纳中心大楼亮起十字架","index":6,"hotValue":"459.6万","link":"https://www.baidu.com/s?wd=瓦格纳中心大楼亮起十字架"},{"title":"杜海涛吃蜈蚣在嘴里八进八出","index":7,"hotValue":"440.4万","link":"https://www.baidu.com/s?wd=杜海涛吃蜈蚣在嘴里八进八出"},{"title":"片仔癀化妆品原董事长林进生被双开","index":8,"hotValue":"430.2万","link":"https://www.baidu.com/s?wd=片仔癀化妆品原董事长林进生被双开"},{"title":"广东：市民不必效仿海外囤盐","index":9,"hotValue":"428.4万","link":"https://www.baidu.com/s?wd=广东：市民不必效仿海外囤盐"},{"title":"福岛周边最高辐射值为东京200倍","index":10,"hotValue":"415.7万","link":"https://www.baidu.com/s?wd=福岛周边最高辐射值为东京200倍"},{"title":"吃下核污染食品后果有多严重","index":11,"hotValue":"407.6万","link":"https://www.baidu.com/s?wd=吃下核污染食品后果有多严重"},{"title":"外媒:普里戈任尸体已被初步确认","index":12,"hotValue":"391.2万","link":"https://www.baidu.com/s?wd=外媒:普里戈任尸体已被初步确认"},{"title":"女生地铁给大爷让座被赠蓝莓","index":13,"hotValue":"385.5万","link":"https://www.baidu.com/s?wd=女生地铁给大爷让座被赠蓝莓"},{"title":"张亮送天天出国留学","index":14,"hotValue":"379.6万","link":"https://www.baidu.com/s?wd=张亮送天天出国留学"},{"title":"瓦格纳几乎所有高层都在失事飞机上","index":15,"hotValue":"367.8万","link":"https://www.baidu.com/s?wd=瓦格纳几乎所有高层都在失事飞机上"},{"title":"东京电力：今天预计排放核废水200吨","index":16,"hotValue":"358.8万","link":"https://www.baidu.com/s?wd=东京电力：今天预计排放核废水200吨"},{"title":"佩斯科夫：普京已知普里戈任遇难","index":17,"hotValue":"342.8万","link":"https://www.baidu.com/s?wd=佩斯科夫：普京已知普里戈任遇难"},{"title":"宁静公开抨击行业乱象","index":18,"hotValue":"338.8万","link":"https://www.baidu.com/s?wd=宁静公开抨击行业乱象"},{"title":"美驻日大使要访问福岛尝海鲜","index":19,"hotValue":"326.0万","link":"https://www.baidu.com/s?wd=美驻日大使要访问福岛尝海鲜"},{"title":"女子称熬夜致免疫力低下确诊紫癜","index":20,"hotValue":"312.6万","link":"https://www.baidu.com/s?wd=女子称熬夜致免疫力低下确诊紫癜"}],"msg":"请求成功"}`
	url := "https://api.codelife.cc/api/top/list?lang=cn&id=Jb0vmloB1G&size=50"
	url = "https://v1.hitokoto.cn/?encode=json&c=k&c=i&lang=cn"
	json = GetHotHttpJson(url)
	keywords := strings.ReplaceAll(strings.ReplaceAll(parseJson(json, "hitokoto"), "。", randStr(5)), "，", randStr(5))
	//println(keywords)
	keywords = parseWords(keywords)
	println("\n\t" + randStr(6) + keywords + randStr(5) + "\n")
}
