package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/dlclark/regexp2"
)

func main() {
	list, _ := ListDir("./", ".php")
	for _, name := range list {
		b, err := ioutil.ReadFile(name)

		if err != nil {
			log.Fatal(err)
		}
		str := string(b)
		str, _ = Preg_replace(`<\?php /\*(((?!\*\/)[\s\S])*)\*. \?>`, "{{/*$1*/}}", str)

		str, _ = Preg_replace(`<\?php (((?!\?>)[\s\S])*)\?>`, `{{$1}}`, str)
		str, _ = Preg_replace(`{{if\((((?!\):}})[\s\S])*)\):}}`, "{{if $1}}", str)
		str, _ = Preg_replace(`{{endif;}}`, "{{end}}", str)

		str, _ = Preg_replace(`\$this->app->loadConfig\('([^']+)'\)`, `loadConfig . "$1"`, str)
		str, _ = Preg_replace(`\$config->([A-z]+)->([A-z]+)->([A-z]+)->([A-z]+)->([A-z]+)`, `.Config.$1.$2.$3.$4.$5`, str)
		str, _ = Preg_replace(`\$config->([A-z]+)->([A-z]+)->([A-z]+)->([A-z]+)`, `.Config.$1.$2.$3.$4`, str)
		str, _ = Preg_replace(`\$config->([A-z]+)->([A-z]+)->([A-z]+)`, `.Config.$1.$2.$3`, str)
		str, _ = Preg_replace(`\$config->([A-z]+)->([A-z]+)`, `.Config.$1.$2`, str)
		str, _ = Preg_replace(`\$config->([A-z]+)`, `.Config.$1`, str)

		str, _ = Preg_replace(`\$lang->([A-z]+)->([A-z]+)->([A-z]+)`, `.Lang.$1.$2.$3`, str)
		str, _ = Preg_replace(`\$lang->([A-z]+)->([A-z]+)`, `.Lang.$1.$2`, str)
		str, _ = Preg_replace(`\$app->([A-z]+)->([A-z]+)->([A-z]+)`, `.App.$1.$2.$3`, str)
		str, _ = Preg_replace(`\$app->([A-z]+)->([A-z]+)`, `.App.$1.$2`, str)
		str, _ = Preg_replace(`printf\(([^,;]+),\s*([^);]+)\)`, `printf $1 $2`, str)

		str, _ = Preg_replace(`foreach\(([^ ]+) as ([^=]+)=>([^);]+)\):`, `range $2,$3 :=$1`, str)

		res, _ := Preg_match_result(`(html|css|js|common)::([A-z]+)\(([^,;]+),([^,;]+),([^,;]+),([^);]+)\);?`, str, -1)
		for _, r := range res {
			for k, v := range r[3:] {
				v = strings.Trim(v, " ")
				r[3:][k], _ = Preg_replace(`^'(.*)'$`, `"$1"`, v)
			}
			replace := fmt.Sprintf("%s_%s %s %s %s %s", r[1], r[2], r[3], r[4], r[5], r[6])
			str = strings.Replace(str, r[0], replace, 1)
		}
		res, _ = Preg_match_result(`(html|css|js|common)::([A-z]+)\(([^,;]+),([^,;]+),([^);]+)\);?`, str, -1)
		for _, r := range res {
			for k, v := range r[3:] {
				v = strings.Trim(v, " ")
				r[3:][k], _ = Preg_replace(`^'(.*)'$`, `"$1"`, v)
			}
			replace := fmt.Sprintf("%s_%s %s %s %s", r[1], r[2], r[3], r[4], r[5])
			str = strings.Replace(str, r[0], replace, 1)
		}
		res, _ = Preg_match_result(`(html|css|js|common)::([A-z]+)\(([^,;]+),([^);]+)\);?`, str, -1)
		for _, r := range res {
			for k, v := range r[3:] {
				v = strings.Trim(v, " ")
				r[3:][k], _ = Preg_replace(`^'(.*)'$`, `"$1"`, v)
			}
			replace := fmt.Sprintf("%s_%s %s %s", r[1], r[2], r[3], r[4])
			str = strings.Replace(str, r[0], replace, 1)
		}
		res, _ = Preg_match_result(`(html|css|js|common)::([A-z]+)\(([^);]+)\);?`, str, -1)
		for _, r := range res {
			for k, v := range r[3:] {
				v = strings.Trim(v, " ")
				r[3:][k], _ = Preg_replace(`^'(.*)'$`, `"$1"`, v)
			}
			replace := fmt.Sprintf("%s_%s %s", r[1], r[2], r[3])
			str = strings.Replace(str, r[0], replace, 1)
		}
		str = strings.ReplaceAll(str, ";}}", "}}")
		str = strings.ReplaceAll(str, "echo ", "")
		str = strings.ReplaceAll(str, "{{endforeach}}", "{{end}}")
		str, _ = Preg_replace(`\$lang->`, `.Lang.`, str)
		str = strings.Replace(str, "$this->moduleName", ".ModuleName", -1)
		str = strings.ReplaceAll(str, "{{include '../../common/view/footer.html.php'}}", `{{template "footer.html" .}}`)
		str = strings.ReplaceAll(str, "{{include '../../common/view/header.html.php'}}", `{{template "header.html" .}}`)
		str = strings.ReplaceAll(str, "common_printOrderLink ", "common_printOrderLink . ")
		str = strings.ReplaceAll(str, "{{html::submitButton()}}", "{{html_submitButton .}}")
		newname := strings.Replace(name, ".php", "", 1)
		newname = strings.Replace(newname, ".hook", "", 1)
		os.Remove(newname)
		f, err := os.Create(newname)
		if err != nil {
			log.Fatal(err)
		}
		f.Write([]byte(str))
		f.Close()
	}
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func DEBUG(v ...interface{}) {

	_, file, line, ok := runtime.Caller(1)
	if ok {
		v = append([]interface{}{fmt.Sprintf("%s,line %d:", file, line)}, v...)
	}
	fmt.Println(v...)

}
func Log(format string, v ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		v = append([]interface{}{fmt.Sprintf("%s,line %d:", file, line)}, v...)
	}
	fmt.Printf("%s "+format+"\r\n", v...)
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := "/"
	if os.IsPathSeparator('\\') { //前边的判断是否是系统的分隔符
		PthSep = "\\"
	}
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}

//返回匹配结果,n=次数
func Preg_match_result(regtext string, text string, n int) ([][]string, error) {

	r, err := regexp2.Compile(regtext, 0)
	if err != nil {
		return nil, err
	}

	m, err := r.FindStringMatch(text)
	if err != nil {
		return nil, err
	}
	var result [][]string
	for m != nil && n != 0 {
		var res_v []string
		for _, v := range m.Groups() {
			res_v = append(res_v, v.String())
		}

		m, _ = r.FindNextMatch(m)
		result = append(result, res_v)
		n--
	}

	return result, nil
}

//正则替换
func Preg_replace(regtext string, text string, src string) (string, error) {
	r, err := regexp2.Compile(regtext, 0)
	if err != nil {
		return "", err
	}
	return r.Replace(src, text, -1, -1)

}
