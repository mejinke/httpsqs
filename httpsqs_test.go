package httpsqs

import "testing"

var host = ""

func TestPrivate(t *testing.T) {
    mq := New()
    if query:=mq.makeQuery("opt=1");query != "http://localhost:1218/?auth=&charset=utf-8&opt=1" {
        t.Log("Make Url Failed: " + query)
        t.Fail()
    }
    mq = New(host)
    if query:=mq.makeQuery("opt=1");query != "http://" + host + ":1218/?auth=&charset=utf-8&opt=1" {
        t.Log("Make Url Failed: " + query)
        t.Fail()
    }
    mq = New(host, "6413")
    if query:=mq.makeQuery("opt=1");query != "http://" + host + ":6413/?auth=&charset=utf-8&opt=1" {
        t.Log("Make Url Failed: " + query)
        t.Fail()
    }
    mq = New(host, "6413", "pwd")
    if query:=mq.makeQuery("opt=1");query != "http://" + host + ":6413/?auth=pwd&charset=utf-8&opt=1" {
        t.Log("Make Url Failed: " + query)
        t.Fail()
    }
    mq = New(host, "6413", "pwd", "gbk")
    if query:=mq.makeQuery("opt=1");query != "http://" + host + ":6413/?auth=pwd&charset=gbk&opt=1" {
        t.Log("Make Url Failed: " + query)
        t.Fail()
    }
}

func TestPublic(t *testing.T) {

    var err error
    var value = "testing"
    mq := New(host)

    _, err = mq.SyncTime(0)
    if err != nil {
        t.Log(err)
        t.Fail()
    }

    _, err = mq.Status("test")
    if err != nil {
        t.Log(err)
        t.Fail()
    }

    _, err = mq.Put("test", value)
    if err != nil {
        t.Log(err)
        t.Fail()
    }

    res := ""
    res, _, err = mq.PGet("test")
    if err != nil {
        t.Log(err)
        t.Fail()
    }
    if res != value {
        t.Log("Error get")
        t.Fail()
    }

    _, err = mq.Reset("test")
    if err != nil {
        t.Log(err)
        t.Fail()
    }
}
