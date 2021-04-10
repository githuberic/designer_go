package v2

import (
	"container/list"
	"testing"
)

func TestVerify(t *testing.T) {
	var aObserver = &aObserver{
		name: "张三",
	}

	var bObserver = &bObserver{
		name: "李四",
	}

	// 新闻允许a和b类型观察者订阅
	news := NewsSubject{
		list:  list.New(),
		title: "新闻联播",
	}
	news.Add(aObserver)
	news.Add(bObserver)

	// 新闻允许a和b类型观察者订阅
	hot := HotSubject{
		list:  list.New(),
		title: "杭州热点",
	}
	hot.Add(aObserver)
	hot.Add(bObserver)

	news.Send("中国新闻...")
	hot.Send("阿里上市")
}
