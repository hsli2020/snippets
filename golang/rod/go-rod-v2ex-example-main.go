package main

import (
	"flag"
	"log"
	"sync"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/robfig/cron/v3"
)

var clockInConf = flag.String("clockin", "0 12 * * *", "cron 语法的定时签到")
var topic = flag.String("topic", "", "要自动定时置顶的主题的 url，深夜不会触发 (10:00 - 24:00)")
var interval = flag.Duration("interval", 15*time.Minute, "自动点击置顶主题的间隔")
var locChina, _ = time.LoadLocation("Asia/Shanghai")
var lock = &sync.Mutex{}

func main() {
	flag.Parse()

	if *topic != "" {
		go func() {
			for {
				h := hour()
				if 9 < h && h < 24 {
					stickyTopic()
				}
				time.Sleep(*interval)
			}
		}()
	}

	if *clockInConf != "" {
		clockIn()

		scheduler := cron.New()
		_, err := scheduler.AddFunc(*clockInConf, func() {
			clockIn()
		})
		if err != nil {
			log.Fatal(err)
		}
		scheduler.Run()
	}
}

func stickyTopic() {
	if !isLoggedIn() {
		login()
	}

	lock.Lock()
	defer lock.Unlock()

	browser := newBrowser(true).Timeout(30 * time.Second)
	defer browser.Close()

	page := browser.MustPage(*topic)
	page.Element(".box")

	go page.HandleDialog(true, "")()
	wait := page.MustWaitRequestIdle()
	page.MustElementR(".box .fr a", "置顶 10 分钟").MustClick()
	wait()

	log.Println("置顶了", *topic)
}

func clockIn() {
	if !isLoggedIn() {
		login()
	}

	lock.Lock()
	defer lock.Unlock()

	browser := newBrowser(true)
	defer browser.Close()

	page := browser.Timeout(time.Minute).MustPage("https://www.v2ex.com/")

	el := page.MustElement(`[href="/mission/daily"]`, `.balance_area`)
	if el.MustMatches(`.balance_area`) {
		log.Println("已经签过到了")
		return
	}

	el.MustClick()

	page.MustElementR("input", "领取 X 铜币").MustClick()
	page.MustElementR(".message", "已成功领取每日登录奖励")
	log.Println("签到成功")
}

func isLoggedIn() bool {
	lock.Lock()
	defer lock.Unlock()

	browser := newBrowser(true)
	defer browser.Close()

	return browser.MustPage("https://www.v2ex.com/signin").MustWaitLoad().MustHasR("a", "登出|Sign Out")
}

func login() {
	lock.Lock()
	defer lock.Unlock()

	browser := newBrowser(false)
	defer browser.Close()

	browser.MustPage("https://www.v2ex.com/signin").MustElementR("a", "登出|Sign Out")
}

func newBrowser(headless bool) *rod.Browser {
	url := launcher.New().Headless(headless).UserDataDir("tmp/user").MustLaunch()
	return rod.New().ControlURL(url).MustConnect()
}

func hour() int {
	return time.Now().In(locChina).Hour()
}
