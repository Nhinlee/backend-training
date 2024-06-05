package main

import (
	"fmt"
	"training/golibs"
)

type Account struct {
	customerID int
	accountID  int
}

// func main() {
// 	// input
// 	customerIds := []int{1, 1, 2, 3, 4, 3, 4, 5, 5, 6}
// 	accountIds := []int{10, 11, 13, 11, 14, 10, 13, 10, 11, 13}

// 	accs := []*Account{}
// 	for i, v := range accountIds {
// 		accs = append(accs, &Account{
// 			accountID:  v,
// 			customerID: customerIds[i],
// 		})
// 	}

// 	sort.Slice(accs, func(i, j int) bool {
// 		return accs[i].accountID < accs[j].accountID
// 	})

// 	accountIDsByCustomerId := map[int]string{}

// 	for _, v := range accs {
// 		accountIDsByCustomerId[v.customerID] += strconv.Itoa(v.accountID)
// 	}

// 	rs := map[string][]int{}
// 	for key, value := range accountIDsByCustomerId {
// 		if _, ok := rs[value]; !ok {
// 			rs[value] = []int{}
// 		}

// 		rs[value] = append(rs[value], key)
// 	}

// 	for _, value := range rs {
// 		if len(value) > 1 {
// 			for _, v := range value {
// 				fmt.Printf("%d ,", v)
// 			}
// 			fmt.Printf("\n")
// 		}
// 	}
// }

func main() {
	fileCount, err := golibs.CountFilesInFolder("./algo")
	if err != nil {
		fmt.Errorf("err: %v", err)
	}
	fmt.Printf("file count: %d", fileCount)
}
