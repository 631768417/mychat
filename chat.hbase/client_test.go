package hbase

import (
	"fmt"
	"testing"

	"chat.utils"
	"chat.logger"
)

func TestClient(t *testing.T) {
	Init()
	//testscan()
	//testInsert()
	//testdeleterow()
}

func testscan() {
	bean := new(Bean)
	bean.Family = "index"
	bean.Qualifier = "2690A7FC70FE18FF2637763D910DB840"
	beans := []*Bean{bean}
	results, err := ScansFromRow("my_offline", beans, 0, true)
	if err == nil {
		for _, result := range results {
			my_offline := new(Tim_offline)
			Result2object(result, my_offline)
			logger.Debug("==========>", my_offline.Mid, " ", my_offline.Createtime, " ", my_offline.Fromuser, " ")
		}
	} else {
		logger.Error("error:", err.Error())
	}
}

func testInsert() {
	fmt.Println("------------>testInsert")
	for i := 0; i < 1; i++ {
		my_message := new(Tim_message)
		my_message.Chatid = utils.TimeMills()
		my_message.Createtime = utils.NowTime()
		my_message.Stamp = my_message.Createtime
		my_message.Fromuser = fmt.Sprint("wuxiaodong_", i)
		my_message.Gname = fmt.Sprint("wu_", i)
		my_message.Large = "1"
		my_message.Msgmode = "1"
		my_message.Msgtype = "1"
		my_message.Small = "0"
		my_message.Stanza = fmt.Sprint("aaaaaaaaaaaaaaaaaa_", i)
		my_message.Touser = fmt.Sprint("dong_", i)
		my_message.IndexChatid = fmt.Sprint(my_message.Chatid)
		row, err := my_message.Insert()
		fmt.Println("timdomain=========>", row, err)
	}
}

func testdeleterow() {
	DeleteRow("my_message", 1)
}

func _Benchmark_client(b *testing.B) {
	Init()
	b.SetBytes(1024 * 1024 * 100)
	b.SetParallelism(8)
	for i := 0; i < b.N; i++ {
		testscan()
	}
}

func _BenchmarkClientParallel(b *testing.B) {
	b.SetBytes(1024 * 1024 * 100)
	b.SetParallelism(8)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			testscan()
		}
	})
}
