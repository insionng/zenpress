package helper

/*
	helper 模块是纯功能性质 辅助性质的代码
	对数据库直接操作的一切代码都不能写在此
*/

import (
	"bufio"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/disintegration/gift"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"html/template"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"os"
	"os/exec"
	fpath "path"
	"regexp"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf16"
)

//PrintError 打印错误
func PrintError() {
	if err := recover(); err != nil {
		log.Printf("%v", err)
		for i := 0; i < 10; i++ {
			funcName, file, line, ok := runtime.Caller(i)
			if ok {
				log.Printf("frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
			}
		}
	}
}

//FileMTime 返回文件的修改时间戳
func FileMTime(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

//FileSize 获取文件尺寸
func FileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

//Unlink 删除文件
func Unlink(file string) error {
	return os.Remove(file)
}

//Rename 重命名文件
func Rename(file string, to string) error {
	return os.Rename(file, to)
}

//FilePutContent 将指定内容写入指定文件
func FilePutContent(file string, content string) (int, error) {
	fs, e := os.Create(file)
	if e != nil {
		return 0, e
	}
	defer fs.Close()
	return fs.WriteString(content)
}

/*
//FilePutContent 追加模式
func FilePutContent(path string, content string) (int, error) {
	fs, e := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if e != nil {
		return 0, e
	}
	defer fs.Close()
	return fs.WriteString(content)
}
*/

//FileGetContent 获取指定文件的内容
func FileGetContent(file string) (string, error) {
	if !IsFile(file) {
		return "", os.ErrNotExist
	}
	b, e := ioutil.ReadFile(file)
	if e != nil {
		return "", e
	}
	return string(b), nil
}

//IsFile 当为目录或文件不存在时返回假
func IsFile(file string) bool {
	f, e := os.Stat(file)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

//IsExist 当文件或目录存在时返回真
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

//CreateFile 创建文件
func CreateFile(dir string, name string) (string, error) {
	src := dir + name + "/"
	if IsExist(src) {
		return src, nil
	}

	if err := os.MkdirAll(src, 0777); err != nil {
		if os.IsPermission(err) {
			fmt.Println("你不够权限创建文件")
		}
		return "", err
	}

	return src, nil
}

//FileRepos file operations
type FileRepos []repository

type repository struct {
	Name     string
	FileTime int64
}

func (r FileRepos) Len() int {
	return len(r)
}

func (r FileRepos) Less(i, j int) bool {
	return r[i].FileTime < r[j].FileTime
}

func (r FileRepos) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// DelFile : 获取所有文件，如果文件达到最上限，按时间删除
func DelFile(files []os.FileInfo, count int, fileDir string) {
	if len(files) <= count {
		return
	}

	result := new(FileRepos)

	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			*result = append(*result, repository{Name: file.Name(), FileTime: file.ModTime().Unix()})
		}
	}

	sort.Sort(result)
	deleteNum := len(files) - count
	for k, v := range *result {
		if k+1 > deleteNum {
			break
		}
		Unlink(fileDir + v.Name)
	}

	return
}

// CopyFile : copy the source ro dest
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}

// CopyDir : copy source directorie to dest
func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

// TouchFile : create file of blank
func TouchFile(path string) {
	if !IsExist(path) {
		WriteFile(path, "")
	}
}

// Pages :  分页计算函数
func Pages(totalRecords int64, page int64, pageSize int64) (pages int64, pageOutput int64, beginNum int64, endNum int64, offset int64) {
	//取得记录总数，计算总页数用
	//totalRecords,总共有totalRecords条记录

	//设定每一页显示的记录数
	if pageSize < 0 || pageSize < 1 {
		pageSize = 10 //如无设置，则默认每页显示10条记录
	}

	//计算总页数 结果逢余进一取整
	//pages = int64(math.Ceil(float64(totalRecords / pageSize)))
	pages = (totalRecords + pageSize - 1) / pageSize
	//返回pages

	//判断页数设置,否则，设置为第一页
	if page < 0 || page < 1 {
		page = 1
	}
	if page > pages {
		page = pages
	}
	//返回page

	beginNum = page - 4
	endNum = page + 5

	if page < 5 {
		beginNum = 1
		endNum = 10 //可用链接数，现在是当前页加前后两页共5页，if条件为可用链接数的一半
	}
	if page > pages-5 {
		beginNum = pages - 9
		endNum = pages
	}
	if beginNum < 1 {
		beginNum = 1
	}
	if endNum > pages {
		endNum = pages
	}
	//返回beginNum
	//返回endNum

	//计算记录偏移量
	offset = (page - 1) * pageSize
	return pages, page, beginNum, endNum, offset
}

// Pagesbar : is for the html template
func Pagesbar(url string, keyword string, resultsMax int64, pages int64, page int64, beginNum int64, endNum int64, style int64) (output template.HTML) {
	var raw string
	switch {
	case style == 0: //sdc 定制版pagesbar

		if keyword != "" {
			keyword = keyword + "/"
		}
		nextpage, prevpage := int64(0), int64(0)
		pagemindle := 2
		pagewidth := pagemindle * 2
		raw = `<div class="page-nav"><ul class="pagination">`
		if resultsMax > 0 {
			count := pages + 1
			//prev page
			if (page != beginNum) && (page > beginNum) {
				prevpage = page - 1
				raw = raw + `<li><a class="prev" href="` + url + keyword + "page" + strconv.FormatInt(prevpage, 10) + `/">上一页</a></li>`
			}

			//current page and loop pages
			j := 0
			for i := page; i < count; i++ {
				j++
				if i == page {
					raw = raw + `<li class="active"><a href="` + url + keyword + "page" + strconv.FormatInt(i, 10) + `/">` + strconv.FormatInt(i, 10) + "</a></li>"
				} else {

					raw = raw + `<li><a href="` + url + keyword + "page" + strconv.FormatInt(i, 10) + `/">` + strconv.FormatInt(i, 10) + "</a></li>"
				}
				if j > pagewidth {
					break
				}
			}

			raw = raw + "<li><span>共" + strconv.FormatInt(pages, 10) + "页</span></li>"

			//next page
			if (page != endNum) && (page < endNum) {
				nextpage = page + 1
				raw = raw + `<li><a class="next" href="` + url + keyword + "page" + strconv.FormatInt(nextpage, 10) + `/">下一页</a></li>`
			}
		} else {
			raw = raw + "<li><span>共0页</span></li>"

		}
		raw = raw + "</ul></div>"
		/*
			if nextpage == 0 && prevpage == 0 {
				output = template.HTML(raw)
			} else {

				if nextpage > 0 || prevpage > 0 {
					raw = raw + `<div id="pagenavi-fixed">`

					if prevpage > 0 {
						raw = raw + `
								<div class="pages-prev">
									<a href="` + url + keyword + `page` + strconv.FormatInt(prevpage, 10) + `/">上一页 &raquo;</a>
								</div>`
					}

					if nextpage > 0 {
						raw = raw + `
								<div class="pages-next">
									<a href="` + url + keyword + `page` + strconv.Itoa(nextpage) + `/">下一页 &raquo;</a>
								</div>`
					}
					raw = raw + "</div>"
				}
				output = template.HTML(raw)
			}
		*/
		output = template.HTML(raw)
	case style == 1:
		/*
			<ul class="pager">
				<li class="previous">
					<a href="#">&larr; Older</a>
				</li>
				<li class="next">
					<a href="#">Newer &rarr;</a>
				</li>
			</ul>
		*/
		if resultsMax > 0 {
			raw = "<ul class='pager'>"
			count := pages + 1
			//begin page
			if (page != beginNum) && (page > beginNum) {
				raw = raw + "<li class='previous'><a href='?" + keyword + "page=" + strconv.FormatInt((page-1), 10) + "'>&laquo;</a></li>"
			}

			for i := int64(1); i < count; i++ {
				//current page and loop pages
				if i == page {
					raw = raw + "<li class='active'><a href='javascript:void();'>" + strconv.FormatInt(i, 10) + "</a></li>"
				} else {
					raw = raw + "<li><a href='?" + keyword + "page=" + strconv.FormatInt(i, 10) + "'>" + strconv.FormatInt(i, 10) + "</a></li>"
				}
			}

			//next page
			if (page != endNum) && (page < endNum) {
				raw = raw + "<li class='next'><a href='?" + keyword + "page=" + strconv.FormatInt((page+1), 10) + "'>&raquo;</a></li>"
			}
			raw = raw + "</ul>"
		}

		output = template.HTML(raw)
	case style == 2:
		/*
			<div class="pagination"><ul>
					<li><a href="#">&laquo;</a></li>
					<li class="active"><a href="#">1</a></li>
					<li><a href="#">2</a></li>
					<li><a href="#">3</a></li>
					<li><a href="#">4</a></li>
					<li><a href="#">&raquo;</a></li>
			</ul></div>
		*/

		if resultsMax > 0 {
			raw = "<div class='pagination pagination-centered'><ul>"
			count := pages + 1
			//begin page
			if (page != beginNum) && (page > beginNum) {
				raw = raw + "<li><a href='?" + keyword + "page=" + strconv.FormatInt((page-1), 10) + "'>&laquo;</a></li>"
			}
			for i := int64(1); i < count; i++ {
				//current page and loop pages
				if i == page {
					raw = raw + "<li class='active'><a href='javascript:void();'>" + strconv.FormatInt(i, 10) + "</a></li>"
				} else {
					raw = raw + "<li><a href='?" + keyword + "page=" + strconv.FormatInt(i, 10) + "'>" + strconv.FormatInt(i, 10) + "</a></li>"
				}
				//next page
				if (page != endNum) && (page < endNum) && (i == pages) {
					raw = raw + "<li><a href='?" + keyword + "page=" + strconv.FormatInt((page+1), 10) + "'>&raquo;</a></li>"
				}
			}
			raw = raw + "</ul></div>"
		}

		output = template.HTML(raw)
	case style == 3:
		/*
			<div class="pagenav">
				<p>
					<a href="" class="on">1</a>
					<a href="">2</a>
					<a href="">3</a>
					<a href="">4</a>
				</p>
			</div>
		*/
		raw = "<div class=\"pagenav\">"
		if resultsMax > 0 {
			raw = raw + "<p>"
			count := pages + 1
			for i := int64(1); i < count; i++ {
				if i == page { //当前页
					raw = raw + "<a onclick=\"javascript:void();\" class=\"on\">" + strconv.FormatInt(i, 10) + "</a>"
				} else { //普通页码链接
					raw = raw + "<a href='?" + keyword + "page=" + strconv.FormatInt(i, 10) + "'>" + strconv.FormatInt(i, 10) + "</a>"
				}
			}
			if (page != pages) && (page < pages) { //下一页
				raw = raw + "<a class='next' href='?" + keyword + "page=" + strconv.FormatInt((page+1), 10) + "'>下一页</a>"
			}

		} else {
			raw = raw + "<h2>No Data!</h2>"
			raw = raw + "<span class='page-numbers'>共0页</span>"
		}
		raw = raw + "</p>"
		output = template.HTML(raw + "</div>")
	case style == 4: //mzr 定制版pagesbar

		if keyword != "" {
			keyword = keyword + "/"
		}
		nextpage, prevpage := int64(0), int64(0)
		pagemindle := 7
		pagewidth := pagemindle * 2
		raw = "<div class='pagination'>"
		if resultsMax > 0 {
			raw = raw + "<span class='page-numbers'>共" + strconv.FormatInt(pages, 10) + "页</span>"
			count := pages + 1
			//prev page
			if (page != beginNum) && (page > beginNum) {
				prevpage = page - 1
				raw = raw + "<a class='prev page-numbers' href='" + url + keyword + "" + strconv.FormatInt(prevpage, 10) + "/'>Prev</a>"
			}

			//current page and loop pages
			j := 0
			for i := page; i < count; i++ {
				j++
				if i == page {
					raw = raw + "<span class='page-numbers current'>" + strconv.FormatInt(i, 10) + "</span>"
				} else {

					raw = raw + "<a class='page-numbers' href='" + url + keyword + "" + strconv.FormatInt(i, 10) + "/'>" + strconv.FormatInt(i, 10) + "</a>"
				}
				if j > pagewidth {
					break
				}
			}

			//next page
			if (page != endNum) && (page < endNum) {
				nextpage = page + 1
				raw = raw + "<a class='next page-numbers' href='" + url + keyword + "" + strconv.FormatInt(nextpage, 10) + "/'>Next</a>"
			}
		} else {
			raw = raw + "<span class='page-numbers'>共0页</span>"
		}
		raw = raw + "</div>"

		if nextpage == 0 && prevpage == 0 {
			output = template.HTML(raw)
		} else {

			if nextpage > 0 || prevpage > 0 {
				raw = raw + `<div id="pagenavi-fixed">`

				if prevpage > 0 {
					raw = raw + `
							<div class="pages-prev">
								<a href="` + url + keyword + `` + strconv.FormatInt(prevpage, 10) + `/" >上一页 &raquo;</a>
							</div>`
				}

				if nextpage > 0 {
					raw = raw + `
							<div class="pages-next">
								<a href="` + url + keyword + `` + strconv.FormatInt(nextpage, 10) + `/" >下一页 &raquo;</a>
							</div>`
				}
				raw = raw + "</div>"
			}
			output = template.HTML(raw)
		}
	case style == 5: //yougam 定制版pagesbar
		/*
			 <ul class="pager">
					 <li class="previous disabled">
							 <a>上一页</a>
					 </li>
					 <li class="pager-nums">1 / 11</li>
					 <li class="next">
							 <a href="/?p=2">下一页</a>
					 </li>
			 </ul>
		*/
		if keyword != "" {
			keyword = keyword + "/"
		}

		if resultsMax > 0 {
			raw = "<ul class='pager'>"
			//begin page
			if (page != beginNum) && (page > beginNum) {
				raw = raw + "<li class='previous'><a href='" + url + keyword + "page" + strconv.FormatInt((page-1), 10) + "/'>上一页</a></li>"
			}

			raw = raw + `<li class="pager-nums">` + strconv.FormatInt(page, 10) + ` / ` + strconv.FormatInt(pages, 10) + `</li>`

			//next page
			if (page != endNum) && (page < endNum) {
				raw = raw + "<li class='next'><a href='" + url + keyword + "page" + strconv.FormatInt((page+1), 10) + "/'>下一页</a></li>"
			}
			raw = raw + "</ul>"
		}

		output = template.HTML(raw)

	}

	return output
}

// TimeSince : 微博时间格式化显示
// timestamp，标准时间戳
func TimeSince(timestamp int64) string {

	//减去8小时
	//d, _ := time.ParseDuration("-8h")
	//t := timestamp.Add(d)
	//since := int(time.Since(t).Minutes())
	//since := math.Abs(float64(time.Now().UTC().Unix() - timestamp))

	since := (time.Now().Unix() - timestamp)
	output := ""
	switch {
	case (since < 60):
		output = "刚刚" //"小于 1 分钟"
	case (since < 60*60):
		output = fmt.Sprintf("%v 分钟之前", since/(60))
	case (since < 60*60*24):
		output = fmt.Sprintf("%v 小时之前", since/(60*60))
	case (since < 60*60*24*30):
		output = fmt.Sprintf("%v 天之前", since/(60*60*24))
	case (since < 60*60*24*365):
		output = fmt.Sprintf("%v 月之前", since/(60*60*24*30))
	default:
		output = fmt.Sprintf("%v 年之前", since/(60*60*24*365))
	}
	return output
}

// SmcTimeSince is format for time
func SmcTimeSince(timeAt time.Time) string {
	now := time.Now()
	since := math.Abs(float64(now.UTC().Unix() - timeAt.Unix()))

	output := ""
	switch {
	case since < 60:
		output = "刚刚"
	case since < 60*60:
		output = fmt.Sprintf("%v分钟前", math.Floor(since/60))
	case since < 60*60*24:
		output = fmt.Sprintf("%v小时前", math.Floor(since/3600))
	case since < 60*60*24*2:
		output = fmt.Sprintf("昨天%v", timeAt.Format("15:04"))
	case since < 60*60*24*3:
		output = fmt.Sprintf("前天%v", timeAt.Format("15:04"))
	case timeAt.Format("2006") == now.Format("2006"):
		output = timeAt.Format("1月2日 15:04")
	default:
		output = timeAt.Format("2006年1月2日 15:04")
	}
	// if math.Floor(since/3600) > 0 {
	//     if timeAt.Format("2006-01-02") == now.Format("2006-01-02") {
	//         output = "今天 "
	//         output += timeAt.Format("15:04")
	//     } else {
	//         if timeAt.Format("2006") == now.Format("2006") {
	//             output = timeAt.Format("1月2日 15:04")
	//         } else {
	//             output = timeAt.Format("2006年1月2日 15:04")
	//         }
	//     }
	// } else {
	//     m := math.Floor(since / 60)
	//     if m > 0 {
	//         output = fmt.Sprintf("%v分钟前", m)
	//     } else {
	//         output = "刚刚"
	//     }
	// }
	return output
}

// ThisHour 获取这个小时的开始点
func ThisHour() int64 {
	t := time.Now()
	year, month, day := t.Date()
	hour, _, _ := t.Clock()

	return time.Date(year, month, day, hour, 0, 0, 0, time.UTC).Unix()
}

// ThisDate 获取今天的开始点
func ThisDate() int64 {
	t := time.Now()
	year, month, day := t.Date()

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Unix()
}

//ThisWeek 获取这周的开始点
func ThisWeek() int64 {
	t := time.Now()
	year, month, day := t.AddDate(0, 0, -1*int(t.Weekday())).Date()

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC).Unix()
}

//ThisMonth 获取这月的开始点
func ThisMonth() int64 {
	t := time.Now()
	year, month, _ := t.Date()

	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Unix()
}

//ThisYear 获取今年的开始点
func ThisYear() int64 {
	t := time.Now()
	year, _, _ := t.Date()

	return time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
}

//FixedpathByNumber is fixed path via number
func FixedpathByNumber(n int, layer int) string {

	hash := md5.New()
	o := ""
	for i := 1; i < layer+1; i++ {

		s := strconv.Itoa(RangeRand(n^n/3+i) / 33)
		hash.Write([]byte(s))
		result := hex.EncodeToString(hash.Sum(nil))
		r := result[0:n]
		o += r + "/"
	}
	return o
}

//FixedpathByString is fixed path via string
func FixedpathByString(s string, layer int) string {

	hash := md5.New()
	output := ""
	for i := 1; i < layer+1; i++ {

		s += s + strconv.Itoa(i+i*i)
		hash.Write([]byte(s))
		result := hex.EncodeToString(hash.Sum(nil))
		r := result[0:2]
		output += r + "/"
	}
	return output
}

//StringNewRand : gen new random string
func StringNewRand(len int) string {

	u := make([]byte, len/2)

	// Reader is a global, shared instance of a cryptographically strong pseudo-random generator.
	// On Unix-like systems, Reader reads from /dev/urandom.
	// On Windows systems, Reader uses the CryptGenRandom API.
	_, err := io.ReadFull(crand.Reader, u)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x", u)
}

//GUID 生成36位GUID
func GUID() string {
	b := make([]byte, 16)

	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return ""
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

//GUID32BIT generator a 32bit guid
func GUID32BIT() string {
	b := make([]byte, 16)

	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return ""
	}

	return fmt.Sprintf("%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

//StringNewUUID generates a new UUID based on version 4.
func StringNewUUID() string {

	u := make([]byte, 16)

	// Reader is a global, shared instance of a cryptographically strong pseudo-random generator.
	// On Unix-like systems, Reader reads from /dev/urandom.
	// On Windows systems, Reader uses the CryptGenRandom API.
	_, err := io.ReadFull(crand.Reader, u)
	if err != nil {
		panic(err)
	}

	// Set version (4) and variant (2).
	var version byte = 4 << 4
	var variant byte = 2 << 4
	u[6] = version | (u[6] & 15)
	u[8] = variant | (u[8] & 15)

	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

//Round 函数对浮点数进行四舍五入
//语法 round(val,prec) 参数 val 规定要舍入的数字。 prec 规定小数点后的位数
func Round(val float64, prec int) float64 {
	var t float64
	f := math.Pow10(prec)
	x := val * f
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return val
	}
	if x >= 0.0 {
		t = math.Ceil(x)
		if (t - x) > 0.50000000001 {
			t -= 1.0
		}
	} else {
		t = math.Ceil(-x)
		if (t + x) > 0.50000000001 {
			t -= 1.0
		}
		t = -t
	}
	x = t / f

	if !math.IsInf(x, 0) {
		return x
	}

	return t
}

//RangeRand 生成规定范围内的整数
//设置起始数字范围，0开始,n截止
func RangeRand(n int) int {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)

}

//Nrand 标准正态分布随机整数，n为随机个数,从0开始
func Nrand(n int64) float64 {
	//sample = NormFloat64() * desiredStdDev + desiredMean
	// 默认位置参数(期望desiredMean)为0,尺度参数(标准差desiredStdDev)为1.

	var i, sample int64 = 0, 0
	desiredMean := 0.0
	desiredStdDev := 100.0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i < n {
		rn := int64(r.NormFloat64()*desiredStdDev + desiredMean)
		sample = rn % n
		i++
	}

	return math.Abs(float64(sample))
}

//MD5 对字符串进行md5哈希,
// 返回32位小写md5结果
/*
func MD5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}
*/
func MD5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	result := hex.EncodeToString(hash.Sum(nil))
	return result
}

//MD5to16 对字符串进行md5哈希,
// 返回16位小写md5结果
func MD5to16(s string) string {
	return MD5(s)[8:24]
}

//SHA1 对字符串进行sha1哈希,
// 返回42位小写sha1结果
func SHA1(s string) string {

	hasher := sha1.New()
	hasher.Write([]byte(s))

	//result := fmt.Sprintf("%x", (hasher.Sum(nil)))
	result := hex.EncodeToString(hasher.Sum(nil))
	return result
}

//ZeroPadding 零式补码
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

//ZeroUnPadding 零式去补码
func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//PKCS5Padding PKCS5 补码
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//PKCS5UnPadding PKCS5 去除补码
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//PKCS7Padding PKCS7补码
func PKCS7Padding(data []byte) []byte {
	blockSize := 16
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)

}

//PKCS7UnPadding 去除PKCS7补码
func PKCS7UnPadding(data []byte) []byte {
	length := len(data)

	// 去掉最后一个字节 unpadding 次
	unpadding := int(data[length-1])

	if length < unpadding {
		panic("[PKCS7UnPadding Error:Data length < unpadding]")
	}
	return data[:(length - unpadding)]
}

// PKCS7Pad pads an byte array to be a multiple of 16
// http://tools.ietf.org/html/rfc5652#section-6.3
func PKCS7Pad(data []byte) []byte {
	dataLen := len(data)
	blockSize := 16
	var validLen int
	if dataLen%blockSize == 0 {
		validLen = dataLen
	} else {
		validLen = int(dataLen/blockSize+1) * blockSize
	}

	paddingLen := validLen - dataLen
	// The length of the padding is used as the byte we will
	// append as a pad.
	bitCode := byte(paddingLen)
	padding := make([]byte, paddingLen)
	for i := 0; i < paddingLen; i++ {
		padding[i] = bitCode
	}
	return append(data, padding...)
}

// PKCS7Unpad removes any potential PKCS7 padding added.
func PKCS7Unpad(data []byte) []byte {
	dataLen := len(data)
	blockSize := 16
	// Edge case
	if dataLen == 0 {
		return nil
	}
	// the last byte indicates the length of the padding to remove
	paddingLen := int(data[dataLen-1])

	// padding length can only be between 1-15
	if paddingLen < blockSize {
		return data[:dataLen-paddingLen]
	}
	return data
}

//AesCBCEncrypt aes cbc encrypt
//AES-128,设key为 16 bytes.
//AES-192,设key为 24 bytes.
//AES-256,设key为 32 bytes.
//AES256 加密CBC模式
func AesCBCEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	//origData = PKCS5Padding(origData, blockSize)
	//origData = PKCS7Pad(origData)
	origData = PKCS7Padding(origData)
	// origData = ZeroPadding(origData, block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//AesCBCDecrypt AES256 解密CBC模式
func AesCBCDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()

	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	//origData = PKCS5UnPadding(origData)
	origData = PKCS7UnPadding(origData)
	//origData = PKCS7Unpad(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

//AesCFBEncrypt AES256加密 CFB
func AesCFBEncrypt(content string, privateKey string, publicKey string) (string, error) {

	c, err := aes.NewCipher([]byte(privateKey))
	if err != nil {
		//fmt.Println("AesCFBEncrypt:", err)
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(c, []byte(publicKey))
	ciphertext := make([]byte, len(content))
	cfb.XORKeyStream(ciphertext, []byte(content))

	return string(ciphertext), err

}

//AesCFBDecrypt AES256解密 CFB
func AesCFBDecrypt(ciphertext string, privateKey string, publicKey string) (string, error) {
	c, err := aes.NewCipher([]byte(privateKey))
	if err != nil {
		return "", err
	}
	cipherz := []byte(ciphertext)
	cfbdec := cipher.NewCFBDecrypter(c, []byte(publicKey))
	contentCopy := make([]byte, len(cipherz))
	cfbdec.XORKeyStream(contentCopy, cipherz)

	return string(contentCopy), err
}

//RsaEncrypt RSA加密
func RsaEncrypt(origData []byte, publicKey []byte) ([]byte, error) {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub := pubInterface.(*rsa.PublicKey)
	return rsa.EncryptPKCS1v15(crand.Reader, pub, origData)
}

//RsaDecrypt RSA解密
func RsaDecrypt(ciphertext []byte, privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("Private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(crand.Reader, priv, ciphertext)
}

//Filehash get hash of file
func Filehash(pathOr string, fileOr *os.File) (string, error) {

	switch {
	case (pathOr != "" && fileOr == nil):

		file, err := os.Open(pathOr)
		if err != nil {
			return "", err
		}
		defer file.Close()
		h := sha1.New()

		_, erro := io.Copy(h, file)
		if erro != nil {
			return "", erro
		}
		//return fmt.Srintf("%x", h.Sum(nil))
		result := hex.EncodeToString(h.Sum(nil))
		//result := fmt.Sprintf("%d", h.Sum(nil))
		//result, _ := fmt.Printf("%d", h.Sum(nil))
		return result, nil
	case (pathOr == "" && fileOr != nil):
		h := sha1.New()
		_, erro := io.Copy(h, fileOr)
		if erro != nil {
			return "", erro
		}
		//return fmt.Srintf("%x", h.Sum(nil))
		result := hex.EncodeToString(h.Sum(nil))
		//result := fmt.Sprintf("%d", h.Sum(nil))
		//result, _ := fmt.Printf("%d", h.Sum(nil))
		return result, nil
	default:
		return "", errors.New("Error: 没有参数无法生成hash,请输入文件路径 或 *os.File!")
	}

}

//FilehashNumber 以数字形式表达文件希哈
func FilehashNumber(path string) (int, error) {

	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	h := sha1.New()

	_, err = io.Copy(h, file)
	if err != nil {
		return 0, err
	}

	result, err := strconv.Atoi(fmt.Sprintf("%d", h.Sum(nil)))
	//return fmt.Srintf("%x", h.Sum(nil))
	return result, err
}

//FilehashBlock 仅以区块范围作文件哈希
func FilehashBlock(path string, block int64) string {
	file, err := os.Open(path)
	defer file.Close()
	hash := ""

	if err != nil {
		return ""
	}

	data := make([]byte, block)
	for {
		n, err := file.Read(data)

		if n != 0 {
			//hash = MD5(string(data))
			hash = SHA1(string(data))
		} else {
			break
		}

		if err != nil && err != io.EOF {
			//panic(err)
			return ""
		}
	}

	return hash
}

//GetSensitiveInfoRemovedEmail 获取 sensitive 信息
func GetSensitiveInfoRemovedEmail(email string) string {
	const (
		mailSeparatorSign = "@"
		minMailIDLength   = 2
	)

	emailSepPos := strings.Index(email, mailSeparatorSign)

	if emailSepPos < 0 {
		return email
	}

	mailID, mailDomain := email[:emailSepPos], email[emailSepPos+1:]

	if mailIDLength := len(mailID); mailIDLength > minMailIDLength {
		firstChar, lastChar := string(mailID[0]), string(mailID[mailIDLength-1])
		stars := "***"
		switch mailIDLength - minMailIDLength {
		case 1:
			stars = "*"
		case 2:
			stars = "**"
		}
		mailID = firstChar + stars + lastChar
	}

	result := mailID + mailSeparatorSign + mailDomain
	return result
}

//Substr 截取字符
func Substr(strin string, start, length int, symbol string) string {
	rs := []rune(strin)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	rsout := string(rs[start:end])
	if strin == rsout {
		return rsout
	}

	return rsout + symbol

}

//GetFile 从因特网获取网页内容作为文件
func GetFile(fileURL string, filePath string, useragent string, referer string) error {

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Println("os.OpenFile errors:", err)
		return err
	}
	stat, err := f.Stat() //获取文件状态
	if err != nil {
		log.Println("f.Stat() errors:", err)
		return err
	}

	ss, _ := strconv.Atoi(fmt.Sprintf("%v", stat.Size))
	f.Seek(int64(ss), 0) //把文件指针指到文件末

	req, err := http.NewRequest("GET", fileURL, nil)
	if err != nil {
		log.Println("http.NewRequest errors:", err)
		return err
	}

	if useragent == "default" {
		useragent = "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/26.0.1410.64 Safari/537.31"
	}

	if referer != "" {
		req.Header.Set("Referer", referer)
	}

	req.Header.Set("User-Agent", useragent)
	req.Header.Set("Range", fmt.Sprintf("bytes=%v-", stat.Size))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("client.Do(req) errors:", err)
		return err
	}

	defer f.Close()
	defer resp.Body.Close()

	written, err := io.Copy(f, resp.Body)
	if err != nil {
		return err
	}

	fs, err := os.Stat(filePath)
	if err != nil {
		if err := os.Remove(filePath); err != nil {
			log.Println("Remove file error:", err)
		}
		return err
	}

	rh, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil || (fs.Size() != int64(rh)) {
		if rh != 0 {

			if fs.Size() != int64(rh) {

				err := errors.New(fileURL + " save failed!")
				log.Println(err)

				if err := os.Remove(filePath); err != nil {
					log.Println("Remove file error:", err)
				}
				return err

			}
			return err
		}
	}

	log.Println(fileURL+" download success!", "written: ", written)
	return err
}

//PostFile 以POST方式发送内容到 web 端
func PostFile(filepath string, actionurl string, fieldname string) (*http.Response, error) {
	bodyBuf := bytes.NewBufferString("")
	bodyWriter := multipart.NewWriter(bodyBuf)

	// use the body_writer to write the Part headers to the buffer
	_, err := bodyWriter.CreateFormFile(fieldname, filepath)
	if err != nil {
		fmt.Println("error writing to buffer")
		return nil, err
	}

	// the file data will be the second part of the body
	fh, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error opening file")
		return nil, err
	}
	defer fh.Close()
	// need to know the boundary to properly close the part myself.
	boundary := bodyWriter.Boundary()
	closeString := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	closeBuf := bytes.NewBufferString(closeString)
	// use multi-reader to defer the reading of the file data until writing to the socket buffer.
	requestReader := io.MultiReader(bodyBuf, fh, closeBuf)
	fi, err := fh.Stat()
	if err != nil {
		fmt.Printf("Error Stating file: %s", filepath)
		return nil, err
	}

	req, err := http.NewRequest("POST", actionurl, requestReader)
	if err != nil {
		return nil, err
	}

	// Set headers for multipart, and Content Length
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.ContentLength = fi.Size() + int64(bodyBuf.Len()) + int64(closeBuf.Len())

	return http.DefaultClient.Do(req)

}

//WriteFile 按指定路径写入内容作为文件
func WriteFile(filepath string, content string) error {
	dirpath := fpath.Dir(filepath)

	os.MkdirAll(dirpath, 0644)
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(content)
	return writer.Flush()
}

//MoveFile 移动文件到目标路径
func MoveFile(frompath string, topath string) error {

	fromfile, err := os.Open(frompath)
	if err != nil {
		return err
	}

	tofile, err := os.OpenFile(topath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	io.Copy(tofile, fromfile)
	fromfile.Close()
	tofile.Close()
	os.Remove(frompath)
	/*
		io.Copy 在一般情况下拷贝不会出错，多个协程访问的时候可能会出现“read ./data/*.png: Access is denied.”的错误，
		造成这个错误的原因很可能是由于多个协程争抢打开文件导致，然而实际情况可能报错后却又删除成功。
		如果我们根据这个错误作出判断的话就会错上加错，所以在这里不做任何判断，完全由上帝决定好了。
	*/
	return nil

}

//CheckPassword 检测密码字符是否符合标准
func CheckPassword(password string) (b bool) {
	if ok, _ := regexp.MatchString(`^[\@A-Za-z0-9\!\#\$\%\^\&\*\~\{\}\[\]\.\,\<\>\(\)\_\+\=]{4,30}$`, password); !ok {
		return false
	}
	return true
}

//CheckUsername 检测用户名称是否符合标准
func CheckUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}A-Z0-9a-z_-]{2,30}$", username); !ok {
		return false
	}
	return true
}

//CheckEmail : 检测伊妹儿是否符合格式
func CheckEmail(email string) (b bool) {
	if ok, _ := regexp.MatchString(`^([a-zA-Z0-9._-])+@([a-zA-Z0-9_-])+((\.[a-zA-Z0-9_-]{2,3}){1,2})$`, email); !ok {
		return false
	}
	return true
}

/*
#gravity可用值有九个,分别是:

西北方 NorthWest：左上角为坐标原点，x轴从左到右，y轴从上到下，也是默认值。
北方   North：上部中间位置为坐标原点，x轴从左到右，y轴从上到下。
东北方 NorthEast：右上角位置为坐标原点，x轴从右到左，y轴从上到下。
西方   West：左边缘中间位置为坐标原点，x轴从左到右，y轴从上到下。
中央   Center：正中间位置为坐标原点，x轴从左到右，y轴从上到下。
东方   East：右边缘的中间位置为坐标原点，x轴从右到左，y轴从上到下。
西南方 SouthWest：左下角为坐标原点，x轴从左到右，y轴从下到上。
南方   South：下边缘的中间为坐标原点，x轴从左到右，y轴从下到上。
东南方 SouthEast：右下角坐标原点，x轴从右到左，y轴从下到上。

*/
/*
func Thumbnail(mode string, input_file string, output_file string, output_size string, output_align string, background string) error {
	//预处理gif格式
	if strings.HasSuffix(input_file, "gif") {
		if Exist(input_file) {

				//convert input_file -coalesce m_file

			cmd := exec.Command("convert", "-coalesce", input_file, input_file)
			err := cmd.Run()

			return err
		} else {
			return errors.New("需要被缩略处理的GIF图片文件并不存在!")
		}
	}

	switch {
	case mode == "resize":
		if Exist(input_file) {

				//convert -resize 256x256^ -gravity center -extent 256x256  src.jpg dest.jpg
				//详细使用格式 http://www.imagemagick.org/Usage/resize/
			cmd := exec.Command("convert", "-resize", output_size+"^", "-gravity", output_align, "-extent", output_size, "-background", background, input_file, output_file)
			err := cmd.Run()

			return err
		} else {
			return errors.New("需要被缩略处理的图片文件并不存在!")
		}
	case mode == "crop":
		if Exist(input_file) {
				 //convert -crop 300×400 center src.jpg dest.jpg 从src.jpg坐标为x:10 y:10截取300×400的图片存为dest.jpg
				 //convert -crop 300×400-10+10 src.jpg dest.jpg 从src.jpg坐标为x:0 y:10截取290×400的图片存为dest.jpg
				 //详细使用格式 http://www.imagemagick.org/Usage/crop/
			cmd := exec.Command("convert", "-gravity", output_align, "-crop", output_size+"+0+0", "+repage", "-background", background, "-extent", output_size, input_file, output_file)
			err := cmd.Run()

			return err
		} else {
			return errors.New("需要被缩略处理的图片文件并不存在!")
		}
	default:
		if Exist(input_file) {

			cmd := exec.Command("convert", "-thumbnail", output_size, "-background", background, "-gravity", output_align, "-extent", output_size, input_file, output_file)
			err := cmd.Run()

			return err
		} else {
			return errors.New("需要被缩略处理的图片文件并不存在!")
		}
	}

}
*/

//Thumbnail 对图片进行缩略处理
func Thumbnail(mode string, inputFile string, outputFile string, outputSize string, outputAlign string, background string) error {
	r, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer r.Close()

	w, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer w.Close()

	var width, height int
	if size := strings.Split(outputSize, "x"); len(size) >= 2 {
		width, _ = strconv.Atoi(size[0])
		height, _ = strconv.Atoi(size[1])
	} else {
		width, _ = strconv.Atoi(size[0])
		height = width
	}

	return GraphicsProcess(r, w, width, height, 100)

}

//Watermark 对输入的图片文件作水印处理
func Watermark(watermarkFile string, inputFile string, outputFile string, outputAlign string) error {
	r, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer r.Close()

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll：%v", err)
	}

	g, format, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("image.Decode(bytes.NewReader(b))：%v", err)
	}

	if (format == "png") || (format == "jpeg") {
		fr, err := os.Open(watermarkFile)
		if err != nil {
			return err
		}
		defer fr.Close()

		wk, _, err := image.Decode(fr)
		if err != nil {
			return (err)
		}

		dst := image.NewRGBA(g.Bounds())
		gift.New().Draw(dst, g)
		//水印放在右下角
		padding := 20
		x := (g.Bounds().Dx() - (wk.Bounds().Dx() + padding))
		y := (g.Bounds().Dy() - (wk.Bounds().Dy() + padding))
		gift.New().DrawAt(dst, wk, image.Point{x, y}, gift.OverOperator)

		w, err := os.Create(outputFile)
		if err != nil {
			return err
		}

		if format == "png" {
			err := png.Encode(w, dst)
			if err != nil {
				return err
			}
		} else {
			err := jpeg.Encode(w, dst, &jpeg.Options{Quality: 100})
			if err != nil {
				return err
			}
		}
	} else {
		return fmt.Errorf("Not support format:%s", format)
	}

	return nil
}

//Rex 如果文本符合正则表达式则返回真
func Rex(text string, iregexp string) (b bool) {
	if ok, _ := regexp.MatchString(iregexp, text); !ok {
		return false
	}
	return true
}

//Exist : if file exist return true
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

//RsaAesSendPacket 发送报文 是否加密 HTTP状态 动作URL 数据内容 RSA公匙
func RsaAesSendPacket(encrypt bool, status string, actionurl string, content string, aesKey string, aesPublicKey string, rsaPublicKey []byte) (*http.Response, error) {
	/*
	 1.对数据进行AES加密
	 2.对AES密匙KEY进行RSA加密
	 3.POST的时候,把RSA密码串放置于URL发送
	 4.POST的时候,把AES密码串放置于BODY发送
	*/
	//只有公钥则只能加密  公钥私钥都有才能解密 所以私匙不能在客户端公开  客户端获取的内容由服务端的权限控制决定
	var bodyBuf io.Reader
	if encrypt {
		// AES对内容进行加密
		aesEncryptContent, err := AesCFBEncrypt(content, aesKey, aesPublicKey)
		if err != nil {
			return nil, err
		}
		bodyBuf = bytes.NewBufferString(aesEncryptContent)

		// 对AES密匙aesKey进行RSA加密
		rsaEncryptContent, err := RsaEncrypt([]byte(aesKey), rsaPublicKey)
		if err != nil {
			return nil, err
		}
		//转换RSA密文BYTE编码为16进制字符串
		aesKey = fmt.Sprintf("%x", rsaEncryptContent)
	} else {
		//无需加密
		bodyBuf = bytes.NewBufferString(content)

	}

	//hash就是各种内容的混合体加key的hash值,验证这个hash是否一致来保证内容不被非法更改
	createtime := strconv.Itoa(int(time.Now().UnixNano()))
	//hash+createtime+aeskey
	actionurl = actionurl + "?hash=" + EncryptHash(status+createtime+string(content)+string(rsaPublicKey), nil) + "-" + createtime + "-" + aesKey

	req, err := http.NewRequest(status, actionurl, bodyBuf)
	if err != nil {
		return nil, err
	}
	hd, err := http.DefaultClient.Do(req)
	return hd, err
}

//RsaAesReceivingPacket 接受加密数据包
func RsaAesReceivingPacket(decrypt bool, hash string, status string, content []byte, aesPublicKey string, rsaPublicKey []byte, rsaPrivateKey []byte) ([]byte, error) {

	//防擅改校验数据
	if hash != "" {
		/*
		 1.对AES数据进行AES解密得出内容
		 2.对RSA数据进行RSA解密得出AES密匙KEY
		*/

		//分解hash+createtime+aeskey
		s := strings.Split(hash, "-")
		hash = s[0]
		createtime := s[1]
		aseKey := s[2]

		//若 decrypt为真则进行解密
		if decrypt {
			if aseKey != "" {

				//对16进制字符串aseKey进行解码
				if x, err := hex.DecodeString(aseKey); err == nil {

					//RSA解密  得出 AES KEY
					rsaDecryptContent, err := RsaDecrypt(x, rsaPrivateKey)
					if err != nil {
						return nil, err
					}

					//还原  aseKey
					aseKey = string(rsaDecryptContent)

					//对AES数据进行AES解密得出内容
					aesDecryptContent, err := AesCFBDecrypt(string(content), aseKey, aesPublicKey)
					if err != nil {
						return nil, err
					}

					content = []byte(aesDecryptContent)

				} else {
					//16进制解码错误
					return nil, err
				}

			} else {
				return nil, errors.New("AES KEY为空无法进行解密")
			}
		}

		if (hash != "") && (createtime != "") {

			if ValidateHash(hash, status+createtime+string(content)+string(rsaPublicKey)) {
				//返回数据明文
				return content, nil
			}

			return nil, errors.New("报文无法通过数据校验")

		}
	}

	return nil, errors.New("数据校验HASH值为空")

}

//SendPacket 发送报文
func SendPacket(status string, actionurl string, content string, ckJar *cookiejar.Jar) (*http.Response, error) {

	client := &http.Client{
		Jar: ckJar,
	}
	var bodyBuf io.Reader
	bodyBuf = bytes.NewBufferString(content)
	req, err := http.NewRequest(status, actionurl, bodyBuf)
	if err != nil {
		return nil, err
	}
	hd, err := client.Do(req)
	return hd, err
}

//GetBanner : Get one of images
func GetBanner(content string) (string, error) {

	if imgs, num := GetImages(content); num > 0 {

		for _, v := range imgs {
			// 只获取本地图片,外部图片不太可靠
			if IsLocal(v) {
				if Exist(URL2local(v)) {

					return v, nil
				}
			}
			return v, errors.New("GetBanner没有图片链接")
		}
	}
	return "", errors.New("GetBanner没有图片链接")
}

//IsLocal : if is local path then return true
func IsLocal(path string) bool {
	if path != "" {
		/*
			把本地路径的无点形式转为有点形式
			转换之后,如果之前传入的是恰好是一个网址而不是本地路径,则在后面的分拣中会把它列入并非本地路径的行列
			因为本地路径在本系统中是预设想必为当前网站项目文件夹范围内的 ./root/path  而不能跳出 到 ../root/path外,
			所以跳出到 ../root/path 外的路径必定不是本地路径!
		*/
		path = URL2local(path)

		//检查带点的本地路径
		s := strings.SplitN(path, ".", -1)
		if len(s) > 1 && len(s) < 4 {
			//通过路径的前缀是否为"/"判断是不是本地文件
			if s[1] != "" {
				if strings.HasPrefix(s[1], "/") {
					return true
				}
				// 第一轮次检查的时候碰上"/"开头的本地路径会判断不出来,需要再进行第2次判断"/"开头的情况
				ss := strings.SplitN("."+s[1], ".", -1)
				if len(ss) > 1 && len(ss) < 4 {
					//通过路径的前缀是否为"/"判断是不是本地文件
					if ss[1] != "" {
						if strings.HasPrefix(ss[1], "/") {
							return true
						}
						return false
					}
					return false
				}
				return false

			}
			return false
		}
	}
	return false
}

//Local2url is turn the local path to url
func Local2url(path string) string {
	/*
		if strings.HasPrefix(path, "./") {
			path = strings.Replace(path, "./", "/", -1)
		}
	*/
	if (!strings.HasPrefix(path, "http") || !strings.HasPrefix(path, "ftp")) && (!strings.HasPrefix(path, "/")) {
		path = "/" + path
	} else if strings.HasPrefix(path, "./") {
		path = strings.Replace(path, "./", "/", 1)
	}
	return path
}

//URL2local is turn the url to local path
func URL2local(path string) string {

	if (!strings.HasPrefix(path, "http") || !strings.HasPrefix(path, "ftp")) && (!strings.HasPrefix(path, "/")) {
		path = "./" + path
	} else if strings.HasPrefix(path, "/") {
		path = strings.Replace(path, "/", "./", 1)
	}
	return path
}

//SetSuffix 设置后缀
func SetSuffix(content string, str string) string {

	content = URL2local(content)
	if len(content) > 0 {
		s := strings.SplitN(content, ".", -1)
		if len(s) > 1 && len(s) < 4 {
			// 判断是不是本地文件或本地路径
			if s[1] != "" && IsLocal(s[1]) {
				return Local2url(s[1] + str)
			}
			return Local2url(content)
		}
	}
	return ""
}

func staticProcess(g image.Image, format string, w io.Writer, width, height, quality int) error {

	filter := gift.New(
		// high-quality resampling with pixel mixing
		//gift.ResizeToFit(width, height, gift.LinearResampling),
		//gift.ResizeToFill(width, height, gift.LanczosResampling, gift.CenterAnchor),
		//gift.ResizeToFill(width, height, gift.LinearResampling, gift.CenterAnchor),

		//gift.ResizeToFit(width, height, gift.LanczosResampling),

		//gift.Resize(width, height, gift.LinearResampling),
		gift.Resize(width, height, gift.LanczosResampling),
		//gift.CropToSize(width, height, gift.CenterAnchor),
	)

	//tmp := image.NewNRGBA(g.Bounds())
	//gift.New().DrawAt(tmp, g, g.Bounds().Min, gift.OverOperator)
	//gift.New().Draw(tmp, g)
	//dst := image.NewRGBA(filter.Bounds(tmp.Bounds()))
	//filter.Draw(dst, tmp)

	dst := image.NewRGBA(filter.Bounds(g.Bounds()))
	filter.Draw(dst, g)

	var err error
	if format == "png" {
		err = png.Encode(w, dst)
		if err != nil {
			return err
		}
	} else {
		err = jpeg.Encode(w, dst, &jpeg.Options{Quality: quality})
		if err != nil {
			return err
		}
	}

	return nil
}

func animateProcess(g *gif.GIF, w io.Writer, width, height int) error {

	filter := gift.New(
		// high-quality resampling with pixel mixing
		//gift.ResizeToFit(width, height, gift.LanczosResampling),
		//gift.ResizeToFill(width, height, gift.LanczosResampling, gift.CenterAnchor),
		//gift.Resize(width, height, gift.LinearResampling),
		gift.Resize(width, height, gift.LanczosResampling),
		//gift.CropToSize(width, height, gift.CenterAnchor),
	)

	var ng = make([]*image.Paletted, 0)
	tmp := image.NewNRGBA(g.Image[0].Bounds())
	for _, i := range g.Image {
		gift.New().DrawAt(tmp, i, i.Bounds().Min, gift.OverOperator)
		dst := image.NewPaletted(filter.Bounds(tmp.Bounds()), i.Palette)
		filter.Draw(dst, tmp)
		ng = append(ng, dst)
	}

	d := &gif.GIF{
		Image:     ng,
		Delay:     g.Delay,
		LoopCount: g.LoopCount,
	}

	err := gif.EncodeAll(w, d)
	if err != nil {
		return err
	}

	return nil
}

//GraphicsProcess 图像处理过程
func GraphicsProcess(r io.Reader, w io.Writer, width, height, quality int) error {

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll：%v", err)
	}

	g, format, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("image.Decode(bytes.NewReader(b))：%v", err)
	}

	if (format == "png") || (format == "jpeg") {
		return staticProcess(g, format, w, width, height, quality)
	} else if format == "gif" {
		g, err := gif.DecodeAll(bytes.NewReader(b))
		if err != nil {
			return fmt.Errorf("gif.DecodeAll(bytes.NewReader(b))：%v", err)
		}
		return animateProcess(g, w, width, height)
	} else {
		return errors.New("bad format")
	}
}

//GetBannerThumbnail 从内容获取banner
func GetBannerThumbnail(content string) (string, error) {
	//开始提取img
	s, e := GetBanner(content)
	if e == nil {
		//配置缩略图
		inputFile := URL2local(s)
		outputFile := URL2local(SetSuffix(s, "_banner.jpg"))
		outputSize := "680x300"
		outputAlign := "center"
		background := "black"

		//处理缩略图
		err := Thumbnail("crop", inputFile, outputFile, outputSize, outputAlign, background)
		if err == nil {
			return Local2url(outputFile), err
		}
		log.Println("GetBannerThumbnail生成缩略图出错:", err)

		if e := os.Remove(outputFile); e != nil {
			fmt.Println("GetBannerThumbnail清除残余缩略图文件出错:", e)
			return "", e
		}
		return "", err
	}

	return "", e
}

//GetThumbnails get thumbnails from the content
func GetThumbnails(content string) (thumbnails string, thumbnailslarge string, thumbnailsmedium string, thumbnailssmall string, err error) {
	/*
		Thumbnails        string //Original remote file , 请在外部处理成200x200再输入进来
		ThumbnailsLarge   string //100x100
		ThumbnailsMedium  string //48x48
		ThumbnailsSmall   string //32x32
	*/
	//开始提取img 默认所有图片为本地文件
	originalFile, e := GetBanner(content)
	if e == nil {

		//配置缩略图
		inputFile := URL2local(originalFile)
		outputFileLarge := URL2local(SetSuffix(originalFile, "_large.jpg"))
		outputFileMedium := URL2local(SetSuffix(originalFile, "_medium.jpg"))
		outputFileSmall := URL2local(SetSuffix(originalFile, "_small.jpg"))
		outputSizeLarge := "100x100"
		outputSizeMedium := "48x48"
		outputSizeSmall := "32x32"
		outputAlign := "center"
		background := "#ffffff"

		//处理缩略图
		//原始文件
		thumbnails = originalFile

		//大缩略图
		if err := Thumbnail("resize", inputFile, outputFileLarge, outputSizeLarge, outputAlign, background); err == nil {

			thumbnailslarge = Local2url(outputFileLarge)
		} else {

			fmt.Println("GetThumbnails生成thumbnailslarge缩略图出错:", err)

			if e := os.Remove(outputFileLarge); e != nil {
				fmt.Println("GetThumbnails清除残余thumbnailslarge缩略图文件出错:", e)

			}
		}

		//中缩略图
		if err := Thumbnail("resize", inputFile, outputFileMedium, outputSizeMedium, outputAlign, background); err == nil {

			thumbnailsmedium = Local2url(outputFileMedium)
		} else {

			fmt.Println("GetThumbnails生成output_file_Medium缩略图出错:", err)

			if e := os.Remove(outputFileMedium); e != nil {
				fmt.Println("GetThumbnails清除残余output_file_Medium缩略图文件出错:", e)

			}
		}

		//小缩略图
		if err := Thumbnail("resize", inputFile, outputFileSmall, outputSizeSmall, outputAlign, background); err == nil {

			thumbnailssmall = Local2url(outputFileSmall)
		} else {

			log.Println("GetThumbnails生成outputFileSmall缩略图出错:", err)
			if e := os.Remove(outputFileSmall); e != nil {
				fmt.Println("GetThumbnails清除残余outputFileSmall缩略图文件出错:", e)

			}
		}
		return thumbnails, thumbnailslarge, thumbnailsmedium, thumbnailssmall, nil
	}
	return "", "", "", "", e
}

//MakeThumbnails 对指定图片文件进行缩略处理
func MakeThumbnails(localpath string) (thumbnails string, thumbnailslarge string, thumbnailsmedium string, thumbnailssmall string, err error) {
	/*
		Thumbnails        string //Original remote file
		ThumbnailsLarge   string //200x300
		ThumbnailsMedium  string //200x150
		ThumbnailsSmall   string //70x70
	*/
	//开始提取img 默认所有图片为本地文件
	if originalFile := URL2local(localpath); originalFile != "" {

		//配置缩略图
		inputFile := URL2local(originalFile)
		outputFileLarge := URL2local(SetSuffix(originalFile, "_large.jpg"))
		outputFileMedium := URL2local(SetSuffix(originalFile, "_medium.jpg"))
		outputFileSmall := URL2local(SetSuffix(originalFile, "_small.jpg"))
		outputSizeLarge := "200x300"
		outputSizeMedium := "200x150"
		outputSizeSmall := "70x70"
		outputAlign := "center"
		background := "#ffffff"

		//处理缩略图
		//原始文件 也缩略处理以适合内容框大小
		if err := Thumbnail("thumbnail", inputFile, originalFile, "696x", outputAlign, background); err == nil {
			watermarkFile := "./static/yougam/img/watermark.png"
			Watermark(watermarkFile, originalFile, originalFile, "SouthEast")
			thumbnails = Local2url(originalFile)
		} else {

			fmt.Println("MakeThumbnails生成thumbnails缩略图出错:", err)

			if e := os.Remove(originalFile); e != nil {
				fmt.Println("MakeThumbnails清除残余thumbnails缩略图文件出错:", e)

			}
		}

		//大缩略图
		if err := Thumbnail("resize", inputFile, outputFileLarge, outputSizeLarge, outputAlign, background); err == nil {

			thumbnailslarge = Local2url(outputFileLarge)
		} else {
			log.Println("MakeThumbnails生成thumbnailslarge缩略图出错:", err)
			if e := os.Remove(outputFileLarge); e != nil {
				fmt.Println("MakeThumbnails清除残余thumbnailslarge缩略图文件出错:", e)

			}
		}

		//中缩略图
		if err := Thumbnail("resize", inputFile, outputFileMedium, outputSizeMedium, outputAlign, background); err == nil {

			thumbnailsmedium = Local2url(outputFileMedium)
		} else {

			fmt.Println("MakeThumbnails生成outputFileMedium缩略图出错:", err)

			if e := os.Remove(outputFileMedium); e != nil {
				fmt.Println("MakeThumbnails清除残余outputFileMedium缩略图文件出错:", e)

			}
		}

		//小缩略图
		if err := Thumbnail("resize", inputFile, outputFileSmall, outputSizeSmall, outputAlign, background); err == nil {

			thumbnailssmall = Local2url(outputFileSmall)
		} else {
			log.Println("MakeThumbnails生成outputFileSmall缩略图出错:", err)
			if e := os.Remove(outputFileSmall); e != nil {
				log.Println("MakeThumbnails清除残余outputFileSmall缩略图文件出错:", e)
			}
		}
		return thumbnails, thumbnailslarge, thumbnailsmedium, thumbnailssmall, nil
	}
	return "", "", "", "", errors.New("Error: 输入的图片路径为空!")
}

//GetImages  返回 图片url列表集合
func GetImages(content string) (imgs []string, num int) {

	//替换HTML的空白字符为空格
	ren := regexp.MustCompile(`\s`) //ns*r
	bodystr := ren.ReplaceAllString(content, " ")

	//匹配所有图片
	//rem := regexp.MustCompile(`<img.*?src="(.+?)".*?`) //匹配最前面的图
	rem := regexp.MustCompile(`<img.+?src="(.+?)".*?`) //匹配最前面的图
	imgUrls := rem.FindAllSubmatch([]byte(bodystr), -1)

	for _, bv := range imgUrls {
		if m := string(bv[1]); m != "" {

			if !IsContainsSets(imgs, m) {
				imgs = append(imgs, m)
			}
		}
	}

	return imgs, len(imgs)
}

//Base64Encoding 对内容进行标准base64编码
func Base64Encoding(s string) string {
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	defer encoder.Close()
	encoder.Write([]byte(s))
	return buf.String()
}

//GetPage 返回获得的网页内容
func GetPage(url string) (string, error) {

	//ua := "Mozilla/5.0 (Windows; U; Windows NT 8.8; en-US) AppleWebKit/883.13 (KHTML, like Gecko) Chrome/88.3.13.87 Safari/883.13"
	ua := "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.92 Safari/537.1 yougam.Com"
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", ua)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err

}

//DifferenceSets 差集
func DifferenceSets(a []string, b []string) []string {

	var f = make([]string, 0)

	for _, v := range a {
		//如果a集合某元素存在于b集合中
		var in bool
		for _, vv := range b {
			if v == vv {
				in = true
				break
			}
		}
		if !in {
			f = append(f, v)
		}
	}
	return f
}

//IntersectionSets 交集
func IntersectionSets(fora []string, forb []string) []string {

	i, c, d := []string{}, []string{}, []string{}
	if len(fora) > len(forb) {

		c = forb
		d = fora

	} else {

		c = fora
		d = forb
	}
	for _, v := range c {

		//如果c集合中某元素v存在于d集合中
		for _, vv := range d {
			if v == vv {
				i = append(i, v)
				break
			}
		}
	}
	return i
}

//SymmetricDifferenceSets 对称差=并集-交集  即是 并集和交集的差集就是对称差
func SymmetricDifferenceSets(fora []string, forb []string) []string {

	return DifferenceSets(UnionSets(fora, forb), IntersectionSets(fora, forb))
}

//UnionSets 并集
func UnionSets(fora []string, forb []string) []string {
	uvalue := []string{}
	//求两个字符串数组的并集
	for _, v := range fora {
		if IsContainsSets(uvalue, v) {
			continue
		} else {
			uvalue = append(uvalue, v)
		}

	}
	for _, v := range forb {
		if IsContainsSets(uvalue, v) {
			continue
		} else {
			uvalue = append(uvalue, v)
		}
	}

	return uvalue
}

//IsContainsSets if is contains then return true
func IsContainsSets(values []string, ivalue string) bool {
	for _, v := range values {

		if v == ivalue {
			return true
		}
	}
	return false
}

//DelLostImages delete the lost images
func DelLostImages(oldz string, newz string) {

	oldfiles, onum := GetImages(oldz)
	newfiles, nnum := GetImages(newz)

	//初步过滤门槛,提高效率,因为下面的操作太多循环,能避免进入则避免
	if (onum > 0 && nnum > 0) || (onum > 0 && nnum < 1) || (onum == nnum) {

		oldfilesLocal := []string{}
		newfilesLocal := []string{}

		for _, v := range oldfiles {
			if IsLocal(v) {
				oldfilesLocal = append(oldfilesLocal, v)
				//如果本地同时也存在banner缓存文件,则加入旧图集合中,等待后面一次性删除
				if p := URL2local(SetSuffix(v, "_banner.jpg")); Exist(p) {
					oldfilesLocal = append(oldfilesLocal, p)
				}
			}
		}
		//fmt.Println("旧图集合:", oldfilesLocal)

		for _, v := range newfiles {
			if IsLocal(v) {
				newfilesLocal = append(newfilesLocal, v)
			}
		}
		//fmt.Println("新图集合:", newfiles_local)

		//旧图集合-新图集合 =待删图集
		for k, v := range DifferenceSets(oldfilesLocal, newfilesLocal) {
			if p := URL2local(v); Exist(p) { //如若文件存在,则处理,否则忽略
				log.Println("删除文件:", p)
				if err := os.Remove(p); err != nil {
					log.Println("#", k, ",DEL FILE ERROR:", err)
				}
			}
		}
	}

}

//StringToUTF16 字符串转换来unit16
func StringToUTF16(s string) []uint16 {
	return utf16.Encode([]rune(s + "\x00"))
}

//VerifyUserfile verify the user file
func VerifyUserfile(path string, usr string) bool {
	fname := fpath.Base(path)[0:48]
	if fhashed, e := Filehash(URL2local(path), nil); e == nil {
		return ValidateHash(fname, fhashed+usr)
	}
	return false
}

//AtUsers 获取文本中 @user 中的用户名集合
func AtUsers(content string) (usrs []string) {
	// 新浪微博中的用户名格式为是“4-30个字符，支持中英文、数字、"_"或减号”
	//也就是说，支持中文、字母、数字、下划线及减号，并且是4到30个字符,这里 汉字作为一个字符

	rx := regexp.MustCompile("@([\\x{4e00}-\\x{9fa5}A-Z0-9a-z_-]+)")
	//^[\\x{4e00}-\\x{9fa5}]+$
	//rx := regexp.MustCompile("@[^,，：:\\s@]{4,30}")
	atusr := rx.FindAllSubmatch([]byte(content), -1)
	for _, v := range atusr {
		if m := string(v[1]); m != "" {
			//usrs = append(usrs, m)
			if IsContainsSets(usrs, m) {
				continue
			} else {
				usrs = append(usrs, m)
			}
		}
	}

	return usrs
}

//AtPages 获取文本中 @urls 的网址集合 ###AtPages函数的作用是提取@后面的url网址,并不是提取图片,请注意!
func AtPages(content string) ([]string, string) {
	urls := []string{}
	rxs := "@([a-zA-z]+://[^\\s]*)"
	rx := regexp.MustCompile(rxs)

	aturl := rx.FindAllSubmatch([]byte(content), -1)

	if len(aturl) > 0 {

		for _, v := range aturl {
			if m := string(v[0]); m != "" {

				if !IsContainsSets(urls, m[1:]) {
					//替换@link链接
					content = strings.Replace(content, m,
						"<a href='/url/?localtion="+m[1:]+"' target='_blank' rel='nofollow'><span>@</span><span>"+m[1:]+"</span></a>", -1)

					urls = append(urls, m[1:])
				}
			}
		}
	}

	return urls, content
}

//AtWhois get user by at
func AtWhois(content string) []string {
	whoisz := []string{}
	//@[^,，：:\s@]+
	//@[\u4e00-\u9fa5a-zA-Z0-9_-]{4,30}
	rxs := "@([\u4e00-\u9fa5a-zA-Z0-9_-]+[^,，：:\\s@]*)"
	rx := regexp.MustCompile(rxs)

	atWhois := rx.FindAllSubmatch([]byte(content), -1)

	if len(atWhois) > 0 {

		for _, v := range atWhois {
			if m := string(v[0]); m != "" {

				if !IsContainsSets(whoisz, m[1:]) {

					whoisz = append(whoisz, m[1:])
				}
			}
		}
	}

	return whoisz
}

//AtPagesGetImages 从网址的内容中提取图片
func AtPagesGetImages(content string) ([]string, string) {
	imgs := []string{}
	links, content := AtPages(content)
	for _, v := range links {

		page, _ := GetPage(v)
		imgz, n := GetImages(page)

		if n > 0 {
			for _, vv := range imgz {
				//vv为单图url 相对路径的处理较为复杂,这里暂时抛弃相对路径的图片 后续再修正
				if strings.HasPrefix(vv, "/") {

					if strings.HasSuffix(v, "/") {
						vv = v + vv[1:]
					} else {
						vv = v + vv
					}

					//vv = v + vv[1:]
				} else {
					vv = FixURL(v, vv)
				}
				if !strings.HasPrefix(vv, "../") {

					if !strings.HasSuffix(v, "js") {
						if !IsContainsSets(imgs, vv) {
							imgs = append(imgs, vv)
						}
					}
				}

			}
		}
	}
	return imgs, content
}

//FixURL 修正url
func FixURL(currentURL, url string) string {

	re1, _ := regexp.Compile("http[s]?://[^/]+")
	destrooturl := re1.FindString(currentURL)

	//当url为：//wah/x/t.png
	if strings.HasPrefix(url, "//") {
		url = "http:" + url
	} else if strings.HasPrefix(url, "/") {
		// re1,_ := regexp.Compile("http[s]?://[^/]+")
		// destrooturl := re1.FindString(currentURL)
		url = destrooturl + url
	}

	//当url为："../wahaha/xiaoxixi/tupian.png"、"./wahaha/xiaoxixi/tupian.png"、"wahaha/xiaoxixi/tupian.png"
	if !strings.HasPrefix(url, "/") && !strings.HasPrefix(url, "http") && !strings.HasPrefix(url, "https") {
		// current_url = strings.TrimSuffix(current_url, "/")
		if destrooturl == currentURL {
			url = currentURL + "/" + url
		} else {
			re2, _ := regexp.Compile("[^/]+?$")
			url = re2.ReplaceAllString(currentURL, "") + url
		}

	}

	return url
}

//PingFile 检测远端文件i是否可用
func PingFile(url string) bool {
	r, e := http.Head(url)
	if e != nil {
		return false
	}
	if (r.Status == "404") || (r.Status != "200") {
		return false
	}
	return true
}

//Split 分割tags imgs等自行用特定符号组合的字符串为列表
func Split(content string, str string) []string {

	if content != "" && str != "" {

		ss := []string{}
		s := strings.SplitN(content, str, -1)
		for _, v := range s {
			if v != "" {
				ss = append(ss, v)
			}
		}

		return ss

	}
	return nil
}

//Metric 返回数字带数量级单位 千对k 百万对M  京对G
func Metric(n int64) string {

	switch {
	case n >= 1000:
		return fmt.Sprint(n/1000, "k")
	case n >= 1000000:
		return fmt.Sprint(n/1000000, "m")
	default:
		return fmt.Sprint(n)
	}
}

//Gravatar 根据用户邮箱显示Gravatar头像
func Gravatar(email string, height int) string {
	if email != "" && height != 0 {
		// 将邮箱转换成MD5哈希值，并设置图像的大小为height像素
		usergravatar := `http://www.gravatar.com/avatar/` + MD5(email) + `?s=` + strconv.Itoa(height)
		return usergravatar
	}
	return ""
}

//Markdown 转换markdown为html模板对象
func Markdown(md string) template.HTML {
	return template.HTML(Markdown2String(md))
}

//Markdown2String md内容转为string
func Markdown2String(md string) string {
	return string(Markdown2Byte(md))
}

//Markdown2Byte md内容转为bytes
func Markdown2Byte(md string) []byte {
	unsafe := blackfriday.MarkdownCommon([]byte(md))
	return bluemonday.UGCPolicy().SanitizeBytes(unsafe)
}

//Markdown2Text md内容转为text
func Markdown2Text(md string) string {
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	/*
		text := re.ReplaceAllString(string(blackfriday.MarkdownCommon([]byte(md))), "")
		text = strings.Replace(text, "&amp;#34;", "&#34;", -1) // &#34; """
		text = strings.Replace(text, "&amp;#39;", "&#39;", -1) //&#39; '''
		text = strings.Replace(text, "&amp;lt;", "&lt;", -1)   // <
		text = strings.Replace(text, "&amp;gt;", "&gt;", -1)   // >
		text = strings.Replace(text, "&hellip;", "…", -1)  // 省略号
	*/

	text := re.ReplaceAllString(Markdown2String(md), "")
	return text
}

//String2Time 时间字符串转换为时间类型
func String2Time(strtime string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05.999999999", strtime, time.Local)
}

//String2UnixNano 时间字符串转换为纳秒级的时间戳
func String2UnixNano(strtime string) (int64, error) {
	t, e := time.ParseInLocation("2006-01-02 15:04:05.999999999", strtime, time.Local)
	if e != nil {
		return -1, e
	}
	return t.UnixNano(), nil
}

//GetVideoAddress 调用you-get命令对输入的视频网站网址进行抽取真实视频源网址
func GetVideoAddress(input string) string {

	cmd := exec.Command("you-get", "-F", "mp4", "-u", input)
	b, _ := cmd.Output()
	cmd.Run()

	rx := regexp.MustCompile(`\['(.+?)'\]`)
	urls := rx.Find(b)
	output := strings.Replace(string(urls), "[", "", -1)
	output = strings.Replace(output, "]", "", -1)
	output = strings.Replace(output, "'", "", -1)

	return output

}

//Tag4Video 获取视频标签对内的视频网址
func Tag4Video(input string) string {

	rx := regexp.MustCompile(`\[v\](.*)\[/v\]`)
	tmpurl := rx.Find([]byte(input))
	url := strings.Replace(string(tmpurl), "[v]", "", -1)
	url = strings.Replace(url, "[/v]", "", -1)

	return url
}

//VideoTags 调用you-get命令对视频标签进行解析
func VideoTags(in string) string {
	rx := regexp.MustCompile(`\[video\](.+?)\[/video\]`)
	text := string(rx.Find([]byte(in)))

	if text != "" {
		text = strings.Replace(text, "[video]", "", -1)
		text = strings.Replace(text, "[/video]", "", -1)
		text = GetVideoAddress(text)

		if text != "" {
			s := Split(text, ",")

			if len(s) >= 1 {
				text = ""
				for k, v := range s {
					text = text + `
						{
													src : [
														'` + v + `'
													],
													poster : '',
													title : 'Video ` + strconv.Itoa(k) + `'
												}`
					if k != 0 || k != (len(s)-1) {
						text = text + `,`
					}
				}
			}

		}
	}

	return text
}

//Htmlquote HTML编码为实体符号
func Htmlquote(text string) string {
	/*
		 Encodes `text` for raw use in HTML.
				 >>> htmlquote("<'&\\">")
				 '&lt;&#39;&amp;&quot;&gt;'
	*/

	text = strings.Replace(text, "&", "&amp;", -1) // Must be done first!
	text = strings.Replace(text, "…", "&hellip;", -1)
	text = strings.Replace(text, "<", "&lt;", -1)
	text = strings.Replace(text, ">", "&gt;", -1)
	text = strings.Replace(text, "'", "&#39;", -1)
	text = strings.Replace(text, "\"", "&#34;", -1)
	text = strings.Replace(text, "\"", "&quot;", -1)
	text = strings.Replace(text, "“", "&ldquo;", -1)
	text = strings.Replace(text, "”", "&rdquo;", -1)
	text = strings.Replace(text, " ", "&nbsp;", -1)
	return text
}

//Htmlunquote 实体符号解释为HTML
func Htmlunquote(text string) string {
	/*
		 Decodes `text` that's HTML quoted.
				 >>> htmlunquote('&lt;&#39;&amp;&quot;&gt;')
				 '<\\'&">'
	*/

	// strings.Replace(s, old, new, n)
	// 在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换

	text = strings.Replace(text, "&nbsp;", " ", -1)
	text = strings.Replace(text, "&rdquo;", "”", -1)
	text = strings.Replace(text, "&ldquo;", "“", -1)
	text = strings.Replace(text, "&quot;", "\"", -1)
	text = strings.Replace(text, "&#34;", "\"", -1)
	text = strings.Replace(text, "&#39;", "'", -1)
	text = strings.Replace(text, "&gt;", ">", -1)
	text = strings.Replace(text, "&lt;", "<", -1)
	text = strings.Replace(text, "&hellip;", "…", -1)
	text = strings.Replace(text, "&amp;", "&", -1) // Must be done last!
	return text
}

//HTML2str 过滤html文本内容中的特殊标签或内容
func HTML2str(html string) string {
	src := string(html)

	//替换HTML的空白字符为空格
	re := regexp.MustCompile(`\s`) //ns*r
	src = re.ReplaceAllString(src, " ")

	//将HTML标签全转换成小写
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	return strings.TrimSpace(src)
}

// Str2HTML 转换html文本为html模板对象
func Str2HTML(raw string) template.HTML {
	return template.HTML(raw)
}

// ObjPolicy 对用户生成内容进行过滤
func ObjPolicy() *bluemonday.Policy {
	p := bluemonday.UGCPolicy()
	// Permits the "dir", "id", "lang", "title" attributes globally
	p.AllowStandardAttributes()

	// Permits the "img" element and it's standard attributes
	p.AllowImages()

	// Permits ordered and unordered lists, and also definition lists
	p.AllowLists()

	// Permits HTML tables and all applicable elements and non-styling attributes
	p.AllowTables()

	p.AllowDataURIImages()
	p.AllowAttrs("style").OnElements("img", "span")
	p.AllowStyling()

	//视频音频标签过滤及支持
	p.AllowAttrs("controls", "autoplay", "loop", "preload", "autoplay").OnElements("video", "audio")
	p.AllowAttrs("width", "height").OnElements("video", "object", "embed")
	p.AllowAttrs("align").OnElements("object", "embed")
	p.AllowAttrs("type").Matching(regexp.MustCompile("(application/x-shockwave-flash|video/webm|video/ogg|video/mp4|video/mpeg|audio/wav|audio/mpeg|audio/x-midi)")).OnElements("object", "embed")
	p.AllowAttrs("quality", "allowFullScreen", "allowScriptAccess").OnElements("embed")
	p.AllowAttrs("src", "data").Matching(regexp.MustCompile("[a-zA-z]+://[^\\s]*(.swf|.flv|.webm|.ogg|.mp4|.mpeg|.mpg|.mpe|.wav|.mp3|.midi|.mid)")).OnElements("video", "audio", "embed", "object")
	p.AllowAttrs("wmode").OnElements("object", "embed")

	//<param name='allowFullScreen' value='true'/> <param name="" value="opaque"/>
	p.AllowAttrs("name").Matching(regexp.MustCompile("(allowFullScreen)")).OnElements("param")
	p.AllowAttrs("value").Matching(regexp.MustCompile("(true)")).OnElements("param")
	return p
}

//StandardURLsPolicy 返回标准URL Policy对象
func StandardURLsPolicy() *bluemonday.Policy {
	p := bluemonday.NewPolicy()
	p.AllowStandardURLs()
	return p
}

//PurePolicy 返回纯Policy对象
func PurePolicy() *bluemonday.Policy {
	return bluemonday.NewPolicy()
}

//Htm2Str 将html字符内容以纯Policy方式过滤内容后返回字符串
func Htm2Str(htm string) string {
	p := PurePolicy()
	return p.Sanitize(htm)
}

//UnixNS2Time 转换时间戳为时间字符串 单位为纳秒
func UnixNS2Time(ns int64, timeLayout string) string {
	return time.Unix(0, ns).Format(timeLayout)
	//上面那行代码实际上并没神马卵用，
	//因在格式化的时候接受不了太高精度，显示会有问题，是个bug！
}

//Unix2Time 转换时间戳为时间字符串 单位为秒
func Unix2Time(s int64, timeLayout string) string {
	return time.Unix(s, 0).Format(timeLayout)
}

//OrderKey 对key进行排序拼接
func OrderKey(FUid, SUid int64) string {
	min := FUid
	max := SUid
	if min > max {
		min = SUid
		max = FUid
	}
	return fmt.Sprintf("%v:%v", min, max)
}

//Compare 在模板中比较变量
func Compare(f, s, t string) bool {

	fInt, e := strconv.ParseInt(f, 10, 64)
	if e != nil {
		return false
	}

	tInt, e := strconv.ParseInt(t, 10, 64)
	if e != nil {
		return false
	}

	switch {
	case s == "<":
		return fInt < tInt
	case s == "<=":
		return fInt <= tInt
	case s == ">":
		return fInt > tInt
	case s == ">=":
		return fInt >= tInt
	case s == "!=":
		return fInt != tInt
	case s == "==":
		return fInt == tInt
	}

	return false

}
