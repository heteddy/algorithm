/*
@Copyright:

*/
/*
@Time : 2020/2/10 15:56
@Author : teddy
@File : pool_test.go
*/

package pool

import (
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type DebugTask struct {
	name string
}

func (d *DebugTask) Run() error {
	log.Printf("Debug task start running: %s", d.name)
	r := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	s := r % 15
	time.Sleep(time.Duration(s) * time.Second)
	log.Printf("Debug task running: %s,sleeping %d", d.name, s)
	return nil
}

func (d *DebugTask) RunError(err error) {
	log.Println("error of running", err)
}

var p *Pool

func TestPool_Put(t *testing.T) {
	for i := 0; i < 10; i++ {
		d := &DebugTask{name: "NO:" + strconv.Itoa(i)}
		p.Put(d)
	}
	time.Sleep(50 * time.Second)
	for i := 10; i < 20; i++ {
		d := &DebugTask{name: "NO:" + strconv.Itoa(i)}
		p.Put(d)
	}
	time.Sleep(10 * time.Second)
	p.Close(true)
}

func TestMain(m *testing.M) {
	cfg := Config{
		QSize:   10,
		Workers: 3,
		MaxIdle: time.Second * 3,
	}

	p = NewPool(&cfg)

	m.Run()
}
