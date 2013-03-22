package utils

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"math"
	"net/smtp"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//分页计算函数
func Pages(results_count int, page int, pagesize int) (pages int, pageout int, beginnum int, endnum int, offset int) {
	//取得记录总数，计算总页数用
	//results_count,总共有results_count条记录

	//设定每一页显示的记录数 
	if pagesize < 0 || pagesize < 1 {
		pagesize = 10 //如无设置，则默认每页显示10条记录
	}

	//计算总页数
	pages = int(math.Ceil(float64(results_count) / float64(pagesize)))
	//返回pages

	//判断页数设置,否则，设置为第一页
	if page < 0 || page < 1 {
		page = 1
	}
	if page > pages {
		page = pages
	}
	//返回page

	beginnum = page - 4
	endnum = page + 5

	if page < 5 {
		beginnum = 1
		endnum = 10 //可用链接数，现在是当前页加前后两页共5页，if条件为可用链接数的一半
	}
	if page > pages-5 {
		beginnum = pages - 9
		endnum = pages
	}
	if beginnum < 1 {
		beginnum = 1
	}
	if endnum > pages {
		endnum = pages
	}
	//返回beginnum
	//返回endnum

	//计算记录偏移量
	offset = (page - 1) * pagesize
	return pages, page, beginnum, endnum, offset
}

func Pagesbar(keyword string, results_max int, pages int, page int, beginnum int, endnum int, style int) (output template.HTML) {
	/*
			"<div class='pagination'>"
		        "<span class='page-numbers'>共"+strconv.Itoa(pages)+"页</span>"
		        "<span class='page-numbers current'>"+strconv.Itoa(page)+"</span>"
		        "<a class='page-numbers' href='?page="+strconv.Itoa(beginnum)+"'>"+strconv.Itoa(beginnum)+"</a>"
		        "<a class="next page-numbers" href="?page="+strconv.Itoa(endnum)+">Next</a>"
		    "</div>"
	*/
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
	var raw string
	if keyword != "" {
		keyword = "keyword=" + keyword + "&"
	}
	switch {
	default:
		raw = "<div class='pagination'>"
		if results_max > 0 {
			raw = raw + "<span class='page-numbers'>找到相关结果" + strconv.Itoa(results_max) + "个，共" + strconv.Itoa(pages) + "页</span>"
			count := pages + 1
			for i := 1; i < count; i++ {
				if i == page {
					raw = raw + "<span class='page-numbers current'>" + strconv.Itoa(i) + "</span>"
				} else {
					raw = raw + "<a class='page-numbers' href='?" + keyword + "page=" + strconv.Itoa(i) + "'>" + strconv.Itoa(i) + "</a>"
				}
			}
			if (page != pages) && (page < pages) {
				raw = raw + "<a class='next page-numbers' href='?" + keyword + "page=" + strconv.Itoa(page+1) + "'>Next</a>"
			}

		} else {
			raw = raw + "<h2>没有数据啊，天要塌啦～</h2>"
			raw = raw + "<span class='page-numbers'>共0页</span>"
		}
		output = template.HTML(raw + "</div>")
	case style == 2:
		if results_max > 0 {
			raw = "<div class='pagination'><ul>"
			//raw = raw + "<span class='page-numbers'>找到相关结果" + strconv.Itoa(results_max) + "个，共" + strconv.Itoa(pages) + "页</span>"
			count := pages + 1
			for i := 1; i < count; i++ {
				//begin page
				if (page != beginnum) && (page != i) && (page > i) && (page > beginnum) {
					//raw = raw + "<a class='next page-numbers' href='?" + keyword + "page=" + strconv.Itoa(page+1) + "'>Next</a>"
					raw = raw + "<li><a href='?" + keyword + "page=" + strconv.Itoa(beginnum) + "'>&laquo;</a></li>"
				}
				//current page and loop pages
				if i == page {
					//raw = raw + "<span class='page-numbers current'>" + strconv.Itoa(i) + "</span>"
					raw = raw + "<li class='active'><a href='javascript:void();'>" + strconv.Itoa(i) + "</a></li>"
				} else {
					//raw = raw + "<a class='page-numbers' href='?" + keyword + "page=" + strconv.Itoa(i) + "'>" + strconv.Itoa(i) + "</a>"
					raw = raw + "<li><a href='?" + keyword + "page=" + strconv.Itoa(i) + "'>" + strconv.Itoa(i) + "</a></li>"
				}
				//end page
				if (page != endnum) && (page != i) && (page < i) && (page < endnum) {
					//raw = raw + "<a class='next page-numbers' href='?" + keyword + "page=" + strconv.Itoa(page+1) + "'>Next</a>"
					raw = raw + "<li><a href='?" + keyword + "page=" + strconv.Itoa(page+1) + "'>&raquo;</a></li>"
				}
			}
			raw = raw + "</ul></div>"
		}

		output = template.HTML(raw)
	}

	return output
}

/** 微博时间格式化显示
 * @param timestamp，标准时间戳
 */
func SmcTimeSince(timeAt time.Time) string {
	now := time.Now()
	since := math.Abs(float64(now.Unix() - timeAt.Unix()))

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

//获取这个小时的开始点
func ThisHour() time.Time {
	t := time.Now()
	year, month, day := t.Date()
	hour, _, _ := t.Clock()

	return time.Date(year, month, day, hour, 0, 0, 0, time.UTC)
}

//获取今天的开始点
func ThisDate() time.Time {
	t := time.Now()
	year, month, day := t.Date()

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

//获取这周的开始点
func ThisWeek() time.Time {
	t := time.Now()
	year, month, day := t.AddDate(0, 0, -1*int(t.Weekday())).Date()

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

//获取这月的开始点
func ThisMonth() time.Time {
	t := time.Now()
	year, month, _ := t.Date()

	return time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
}

//获取今年的开始点
func ThisYear() time.Time {
	t := time.Now()
	year, _, _ := t.Date()

	return time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
}

func GetEmailRegexp() (*regexp.Regexp, error) {
	return regexp.Compile(`^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((\.[a-zA-Z0-9_-]{2,3}){1,2})$`)
}

// 对字符串进行md5哈希,
// 返回32位小写md5结果
func MD5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 对字符串进行md5哈希,
// 返回16位小写md5结果
func MD5_16(s string) string {
	return MD5(s)[8:24]
}

/** 
* user : example@example.com login smtp server user
* password: xxxxx login smtp server password
* host: smtp.example.com:port   smtp.163.com:25
* to: example@example.com;example1@163.com;example2@sina.com.cn;...
* subject:The subject of mail
* body: The content of mail
* mailtype: mail type html or text
 */
func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func GetSensitiveInfoRemovedEmail(email string) string {
	const (
		mail_separator_sign = "@"
		min_mail_id_length  = 2
	)

	emailSepPos := strings.Index(email, mail_separator_sign)

	if emailSepPos < 0 {
		return email
	}

	mailId, mailDomain := email[:emailSepPos], email[emailSepPos+1:]

	if mailIdLength := len(mailId); mailIdLength > min_mail_id_length {
		firstChar, lastChar := string(mailId[0]), string(mailId[mailIdLength-1])
		stars := "***"
		switch mailIdLength - min_mail_id_length {
		case 1:
			stars = "*"
		case 2:
			stars = "**"
		}
		mailId = firstChar + stars + lastChar
	}

	result := mailId + mail_separator_sign + mailDomain
	return result
}

//截取字符
func Substr(str string, start, length int) string {
	rs := []rune(str)
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

	return string(rs[start:end])
}

func Writefile(path string, filename string, content string) error {
	//path = path[0 : len(path)-len(filename)]
	filename = path + filename

	os.MkdirAll(path, 0777)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(content)
	writer.Flush()
	return nil
}

func Htmlquote(text string) string {
	//HTML编码为实体符号
	/*
	   Encodes `text` for raw use in HTML.
	       >>> htmlquote("<'&\\">")
	       '&lt;&#39;&amp;&quot;&gt;'
	*/

	text = strings.Replace(text, "&", "&amp;", -1) // Must be done first!
	text = strings.Replace(text, "<", "&lt;", -1)
	text = strings.Replace(text, ">", "&gt;", -1)
	text = strings.Replace(text, "'", "&#39;", -1)
	text = strings.Replace(text, "\"", "&quot;", -1)
	text = strings.Replace(text, "“", "&ldquo;", -1)
	text = strings.Replace(text, "”", "&rdquo;", -1)
	text = strings.Replace(text, " ", "&nbsp;", -1)
	return text
}

func Htmlunquote(text string) string {
	//实体符号解释为HTML
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
	text = strings.Replace(text, "&#39;", "'", -1)
	text = strings.Replace(text, "&gt;", ">", -1)
	text = strings.Replace(text, "&lt;", "<", -1)
	text = strings.Replace(text, "&amp;", "&", -1) // Must be done last!
	return text
}
