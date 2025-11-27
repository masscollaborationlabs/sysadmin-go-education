package main

import (
	"fmt"
	"time"
)

func main() {
	demoBreak()
	demoCase()
	demoChan()
	demoConst()
	demoContinue()
	demoDefault()
	demoDefer()
	demoElse()
	demoFallthrough()
	demoFor()
	demoFunc()
	demoGo()
	demoGoto()
	demoIf()
	demoInterface()
	demoMap()
	demoRange()
	fmt.Println(demoReturn())
	demoSelect()
	demoStruct()
	demoSwitch()
	demoType()
	demoVar()
}

func demoBreak() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("break döngüyü 2'de durdurur")
			break
		}
	}
}

func demoCase() {
	value := 1
	switch value {
	case 1:
		fmt.Println("case değeri 1 ile eşleşti")
	case 2:
		fmt.Println("case değeri 2 ile eşleşti")
	}
}

func demoChan() {
	ch := make(chan string, 1)
	ch <- "chan veriyi taşıyor"
	fmt.Println(<-ch)
}

func demoConst() {
	const greeting = "const değerleri değişmez tutar"
	fmt.Println(greeting)
}

func demoContinue() {
	for i := 0; i < 3; i++ {
		if i == 1 {
			continue
		}
		fmt.Println("continue değeri atlar", i)
	}
}

func demoDefault() {
	value := 42
	switch value {
	default:
		fmt.Println("default uyuşmayan değerleri yakalar")
	}
}

func demoDefer() {
	defer fmt.Println("defer en son çalışır")
	fmt.Println("defer işi sıraya koyar")
}

func demoElse() {
	number := 5
	if number > 5 {
		fmt.Println("sayı daha büyük")
	} else {
		fmt.Println("else kalan durumları ele alır")
	}
}

func demoFallthrough() {
	value := 1
	switch value {
	case 1:
		fmt.Println("ilk dal")
		fallthrough
	case 2:
		fmt.Println("fallthrough sonraki dalı zorlar")
	}
}

func demoFor() {
	for i := 0; i < 2; i++ {
		fmt.Println("for döngüsü", i, "değerini geziyor")
	}
}

func demoFunc() {
	inner := func(name string) string {
		return fmt.Sprintf("func literalleri çalışır, merhaba %s", name)
	}
	fmt.Println(inner("Gopher"))
}

func demoGo() {
	done := make(chan struct{})
	go func() {
		fmt.Println("go anahtar kelimesi goroutine başlatır")
		close(done)
	}()
	<-done
}

func demoGoto() {
	count := 0
start:
	count++
	if count < 2 {
		goto start
	}
	fmt.Println("goto etiketlere atlar")
}

func demoIf() {
	if 1 < 2 {
		fmt.Println("if koşulları değerlendirir")
	}
}

type interfaceWrapper struct {
	data string
}

func (w interfaceWrapper) Describe() string {
	return w.data
}

func demoInterface() {
	type describer interface {
		Describe() string
	}
	var d describer = interfaceWrapper{data: "interface çok biçimliliğe izin verir"}
	fmt.Println(d.Describe())
}

func demoMap() {
	m := map[string]int{"apples": 3}
	fmt.Println("map araması:", m["apples"])
}

func demoRange() {
	values := []int{1, 2}
	for idx, val := range values {
		fmt.Printf("range indeks %d değer %d\n", idx, val)
	}
}

func demoReturn() string {
	return "return değerleri geri döner"
}

func demoSelect() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() {
		time.Sleep(10 * time.Millisecond)
		ch2 <- "select ch2 kanalını seçti"
	}()
	ch1 <- "select ch1 kanalını seçti"

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	default:
		fmt.Println("select default dalına düştü")
	}
}

func demoStruct() {
	type point struct {
		X int
		Y int
	}
	p := point{X: 1, Y: 2}
	fmt.Println("struct alanları:", p.X, p.Y)
}

func demoSwitch() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("switch hafta sonunu yakaladı")
	default:
		fmt.Println("switch hafta içini yakaladı")
	}
}

func demoType() {
	type localInt int
	var value localInt = 7
	fmt.Println("type yeni takma tipler tanımlar:", value)
}

func demoVar() {
	var count int = 5
	fmt.Println("var değişken tanımlar:", count)
}
