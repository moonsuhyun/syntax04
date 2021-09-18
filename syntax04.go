package main

import (
	"fmt"
	"math/rand"
	"time"
)

//ver 0.4
func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
		}
		fmt.Print(i, " ")
	}

	if isPrime {
		fmt.Println(number, "은 소수입니다.")
	} else {
		fmt.Println(number, "은 소수가 아닙니다.")
	}

}

//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
////ver 0.3
//func main() {
//	seed := time.Now().Unix()
//	rand.Seed(seed)
//
//	isPrime := true
//	number := rand.Intn(150) + 2
//	fmt.Println("임의로 추출된 수 : ", number)
//
//	for i := 2; i < number; i++ {
//		if number%i == 0 {
//			isPrime = false
//		}
//	}
//
//	if isPrime {
//		fmt.Println(number, "은 소수입니다.")
//	} else {
//		fmt.Println(number, "은 소수가 아닙니다.")
//	}
//
//}

//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
////ver 0.2
//func main() {
//	seed := time.Now().Unix()
//	rand.Seed(seed)
//
//	count := 0
//	number := rand.Intn(150) + 2
//	fmt.Println("임의로 추출된 수 : ", number)
//
//	for i := 2; i < number; i++ {
//		if number%i == 0 {
//			count++
//		}
//	}
//
//	if count == 0 {
//		fmt.Println(number, "은 소수입니다.")
//	} else {
//		fmt.Println(number, "은 소수가 아닙니다.")
//	}
//
//}

//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
////ver 0.1
//func main()  {
//	seed := time.Now().Unix()
//	rand.Seed(seed)
//
//	count := 0
//	number := rand.Intn(150) + 2
//	fmt.Println("임의로 추출된 수 : ", number)
//
//	for i := 1; i <= number; i++ {
//		if number % i == 0 {
//			count++
//		}
//	}
//
//	if count == 2 {
//		fmt.Println(number, "은 소수입니다.")
//	} else {
//		fmt.Println(number, "은 소수가 아닙니다.")
//	}
//
//}

//package main
//
//import (
//	"fmt"
//	"math/rand"
//	"time"
//)
//
//func main()  {
//	seed := time.Now().Unix()
//	rand.Seed(seed)
//
//	for i := 1; i < 6; i++ {
//		dice := rand.Intn(6)+1
//		fmt.Println(dice)
//	}
//}
