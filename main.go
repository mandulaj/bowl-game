package main

import (
  "fmt"
  "crypto/rand"
  "math/big"
)

func pseudoDiceMatch() int {
  max := big.NewInt(2)
  num, err := rand.Int(rand.Reader, max)
  if err != nil {
    panic(err)
  }
  return int(num.Int64())
}


func diceRoll() int {
  max := big.NewInt(6)
  num, err := rand.Int(rand.Reader, max)
  if err != nil {
    panic(err)
  }
  return int(num.Int64()) + 1
}

func diceMatch() int{
  for {
    first := diceRoll()
    second := diceRoll()

    if first > second {
      // fmt.Println("First won, he got ", first, " Second got ", second)
      // fmt.Println("End of match \n")
      return 1
    }
    if second > first {
      // fmt.Println("Second won, he got ", second, " First got ", first)
      // fmt.Println("End of match \n")
      return 0
    }
    // fmt.Println("Both got ", first)
  }
}

func makeGroup(length int) []int {
  group := make([]int, length)
  for i:= 0; i< length; i++ {
    group[i] = 1
  }
  return group
}

func bowlMatch(first, second int, diceMatchFunc func() int) (int, int){
  //fmt.Println("Match starting, first has ", first, " second has ", second)
  for {
    if first == 0 {
      return 2, second // second lost
    }
    if second == 0 {
      return 1, first // first lost
    }
    var bet = 0
    if first < second {
      bet = first
    } else {
      bet = second
    }


    if diceMatchFunc() == 0 {
      second -= bet
      first += bet
      //fmt.Println("Second won, first has ", first, " second has ", second)
    } else {
      second += bet
      first -= bet
      //fmt.Println("First won, first has ", first, " second has ", second)
    }
  }
}

func groupMatch(groupSize int) int{
  group := makeGroup(groupSize)
  bank := 0
  for i:=1; i<len(group); i++ {
    looser, bowls := bowlMatch(group[bank], group[i], pseudoDiceMatch)
    if (looser == 1) {
      group[i] = 0
    }
    if (looser == 2) {
      group[bank] = 0;
      bank = i
    }
    group[bank] = bowls
  }

  return bank
}

func testGroupSize(groupSize int) []int{
  var games = 1e5
  var group = make([]int, groupSize)

  for i:=int64(0); i< int64(games); i++ {
    looser := groupMatch(groupSize)
    group[looser] += 1
  }
  return group
}

func main() {
  for i:=1; i<50; i++ {
    looser := testGroupSize(i)
    fmt.Println("Group size: ", i, ": ", looser)
  }
}
