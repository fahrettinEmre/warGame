package main

import (
	"fmt"
	"math/rand"
	"oyunmodulu/oyun"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Takımlar kaçar kişi olacak ?")
	var takım int
	dunyaliSlice := []oyun.Dunyali{}
	marslıSlice := []oyun.Marslı{}

	fmt.Scan(&takım)

	for i := 0; i < takım; i++ {
		dünyalı := new(oyun.Dunyali)
		dünyalı.Birlik = "Birinci Birlik"
		dünyalı.Can = 100
		dünyalı.MermiSayısı = 90
		dünyalı.Name = "Dünyalı-" + strconv.Itoa(i+1)
		dunyaliSlice = append(dunyaliSlice, *dünyalı)

		marslı := new(oyun.Marslı)
		marslı.Can = 100
		marslı.MermiSayısı = 90
		marslı.Name = "Marslı-" + strconv.Itoa(i+1)
		marslı.UzayGemisi = "ABCD-23"
		marslıSlice = append(marslıSlice, *marslı)

	}
	for {

		kontrolD := true
		kontrolM := true
		for i := 0; i < len(dunyaliSlice); i++ {
			if dunyaliSlice[i].Can > 0 {
				kontrolD = false
			}

		}
		if kontrolD {
			fmt.Println("Marslılar kazandı")
			break
		} else {
			for i := 0; i < len(marslıSlice); i++ {
				if marslıSlice[i].Can > 0 {
					kontrolM = false
				}
			}
			if kontrolM {
				fmt.Println("Dünyalılar kazandı")
				break

			}
		}
		random := rand.Intn(2)
		randomforSlice := rand.Intn(len(dunyaliSlice))
		if random == 0 {
			dunyaliSlice[randomforSlice].AtesEt()
			marslıSlice[randomforSlice].HasarAl()

		} else {

			marslıSlice[randomforSlice].AtesEt()
			dunyaliSlice[randomforSlice].HasarAl()
		}

	}
	fmt.Println(dunyaliSlice)
	fmt.Println(marslıSlice)
}
