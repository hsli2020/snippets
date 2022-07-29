package twiutil		// https://github.com/ookkoouu/twiutil

import (
	urlmod "net/url"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
)

func GetTweet(client *twitter.Client, id int64) (tweet *twitter.Tweet, err error) {
	tweet, _, err = client.Statuses.Show(id, &twitter.StatusShowParams{TweetMode: "extended"})
	return
}

func FindUrlAll(text string) (urls []string) {
	regexTwitterUrl := regexp.MustCompile(`https?://twitter\.com(/\w+)?/status(es)?/\d+`)
	return regexTwitterUrl.FindAllString(text, -1)
}

func FindIdAll(s string) (ids []int64) {
	urlsStr := FindUrlAll(s)
	if len(urlsStr) == 0 {
		return
	}

	for _, uStr := range urlsStr {
		url, _ := urlmod.Parse(uStr)
		idStr := url.Path[strings.LastIndex(url.Path, "/")+1:]
		id, _ := strconv.ParseInt(idStr, 10, 64)
		ids = append(ids, id)
	}
	return
}

func FindId(s string) (id int64) {
	u := FindUrlAll(s)
	if len(u) == 0 {
		return
	}
	url, _ := urlmod.Parse(u[0])
	idStr := url.Path[strings.LastIndex(url.Path, "/")+1:]
	id, _ = strconv.ParseInt(idStr, 10, 64)
	return
}

func GetVideoUrl(media twitter.MediaEntity) (url string) {
	if len(media.VideoInfo.Variants) == 0 {
		return
	}

	variants := make([]twitter.VideoVariant, 0, 4)
	for _, v := range media.VideoInfo.Variants {
		if v.ContentType == "video/mp4" {
			variants = append(variants, v)
		}
	}
	sort.Slice(variants, func(h, f int) bool { return variants[h].Bitrate > variants[f].Bitrate })

	url = variants[0].URL
	return
}

func GetMediaUrls(tweet *twitter.Tweet) (urls []string) {
	if !HasMedia(tweet) {
		return
	}

	medias := tweet.ExtendedEntities.Media
	for _, media := range medias {
		switch {
		// photo
		case len(media.VideoInfo.Variants) == 0:
			urls = append(urls, media.MediaURLHttps)

		// video animated_gif
		case len(media.VideoInfo.Variants) > 0:
			urls = append(urls, GetVideoUrl(media))
		}
	}
	return urls
}

func GetMediaTypes(tweet *twitter.Tweet) (types []string) {
	if !HasMedia(tweet) {
		return
	}

	medias := tweet.ExtendedEntities.Media
	for _, media := range medias {
		types = append(types, media.Type)
	}
	return types
}

func HasMedia(tweet *twitter.Tweet) bool {
	return tweet.ExtendedEntities != nil && len(tweet.ExtendedEntities.Media) > 0
}

func HasQuotedTweet(tweet *twitter.Tweet) bool {
	if tweet.QuotedStatus == nil {
		return false
	} else {
		return true
	}
}

func GetQuotedTweetUrl(tweet *twitter.Tweet) string {
	if !HasQuotedTweet(tweet) {
		return ""
	}
	return "https://twitter.com/i/status/" + tweet.QuotedStatusIDStr
}

func GetSource(tweet *twitter.Tweet) (source string) {
	source = strings.Split(strings.Split(tweet.Source, ">")[1], "<")[0]
	return
}
