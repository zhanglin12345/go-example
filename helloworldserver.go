package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var (
	addr     = flag.String("addr", ":9999", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

func main() {
	p = &value
	flag.Parse()

	h := requestHandler
	if *compress {
		h = fasthttp.CompressHandler(h)
	}

	router := fasthttprouter.New()
	router.POST("/test", requestHandler)

	server := &fasthttp.Server{
		Handler:     h,
		IdleTimeout: 30 * time.Second,
	}
	server.ListenAndServe(*addr)

	// if err := fasthttp.ListenAndServe(*addr, h); err != nil {
	// 	log.Fatalf("Error in ListenAndServe: %s", err)
	// }
}

var value string = "{\"events\": [     {       \"elementName\": \"button01\",       \"actions\": [         {           \"type\": \"NAVIGATION_ACTIONS_ADD_AS_NAVIGATION_TARGET\",           \"arguments\": {             \"destination\": {               \"city\": \"北京市\",               \"title\": \"导航去天安门\",               \"lon\": 119.6696,               \"lat\": 31.298874,               \"street\": \"长安街\",               \"postalCode\": \"100000\",               \"houseNumber\": \"103\"             }           }         }       ]     },     {       \"elementName\": \"navigateTo_entry_013\",       \"actions\": [         {           \"type\": \"NAVIGATION_ACTIONS_ADD_AS_NAVIGATION_TARGET\",           \"arguments\": {             \"destination\": {               \"city\": \"北京市\",               \"title\": \"导航去首尔石锅饭\",               \"lon\": 116.481,               \"lat\": 39.99656,               \"street\": \"望京街道\",               \"postalCode\": \"100000\",               \"houseNumber\": \"1219\"             }           }         }       ]     },     {       \"elementName\": \"url01\",       \"actions\": [         {           \"type\": \"BROWSER_ACTIONS_SHOW_URL\",           \"arguments\": {             \"url\": \"https://lecakesp.tmall.com/index.htm?spm=a220o.1000855.w5001-17434081451.2.36bd7f137OgEXM&scene=taobao_shop\"           }         }       ]     },     {       \"elementName\": \"showd_entry_01\",       \"actions\": [         {           \"type\": \"NAVIGATION_ACTIONS_SHOW_DESTINATION\",           \"arguments\": {             \"destination\": {               \"city\": \"北京市\",               \"title\": \"蜜丝慕斯(望京SOHO店) \",               \"lon\": 116.481,               \"lat\": 39.99656,               \"street\": \"望京街道\",               \"postalCode\": \"100000\",               \"houseNumber\": \"1219\"             }           }         }       ]     },     {       \"elementName\": \"carouselEntry01\",       \"actions\": [         {           \"type\": \"ONLINE_UIACTIONS_UPDATE_WIDGET\",           \"arguments\": {             \"restRequest\": {               \"body\": null,               \"path\": \"/sequence/0/refresh?id=1\",               \"requestParameter\": null,               \"type\": \"Post\"             }           }         }       ]     }   ],   \"widgetData\": {     \"tabs\": [       {         \"title\": \"welcome\",         \"providedByText\": \"welcome 2\",         \"leftLogo\": null,         \"providedByLogo\": null,         \"emptyListText\": \"No Content\",         \"entries\": [           {             \"carouselEntry\": {               \"starsImage\": {                 \"url\": \"https://s3-media2.fl.yelpcdn.com/bphoto/MatIjPkgFG7NYaCYZ27d8w/ms.jpg\",                 \"uiImage\": null,                 \"base64\": null,                 \"asset\": null               },               \"imageFooter\": \"欢迎登录\",               \"title\": \"欢迎阿豪登录\",               \"subtitle1\": \"欢迎登录\",               \"image\": {                 \"url\": null,                 \"uiImage\": null,                 \"base64\": \"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\",                 \"asset\": null               },               \"voiceMessage\": \"some text to read out...\",               \"starsImageDescriptionKeepSize\": false,               \"reviews\": \"欢迎登录\",               \"imageHeader\": \"欢迎登录\",               \"starsImageEntryKeepSize\": false,               \"type\": \"locationBased\",               \"subtitle2\": \"欢迎登录\"             },             \"id\": \"entry01\"           }         ]       }     ],     \"optionConfig\": {       \"entries\": [         {           \"id\": \"weather_setting1\",           \"isActivated\": false,           \"name\": \"°C\"         },         {           \"id\": \"weather_setting2\",           \"isActivated\": true,           \"name\": \"°F\"         }       ],       \"title\": \"OPTIONS\"     }   } }"
var p *string

func requestHandler(ctx *fasthttp.RequestCtx) {
	time.Sleep(100 * time.Millisecond)
	fmt.Fprintf(ctx, *p)

	ctx.SetContentType("application/json")

	// Set arbitrary headers
	ctx.Response.Header.Set("X-My-Header", "my-header-value")

	// Set cookies
	var c fasthttp.Cookie
	c.SetKey("cookie-name")
	c.SetValue("cookie-value")
	ctx.Response.Header.SetCookie(&c)
}
