import (
  "fmt"
  "strconv"
  "strings" 
)

type romanNumeral struct {
  literal string
  value int
}

var allRomanNumerals = []romanNumeral{
  {"M", 1000},
  {"CM", 900},
  {"D", 500},
  {"CD", 400},
  {"C", 100},
  {"XC", 90},
  {"L", 50},
  {"XL", 40},
  {"X", 10},
  {"IX", 9},
  {"V", 5},
  {"IV", 4},
  {"I", 1},  
}

func romToArab(input string) int {
  
  result := 0

  for i := 0; i < len(input); i++ {
    symbol := input[i:i+1]
    
    for _, numeral := range allRomanNumerals {
      if symbol == numeral.literal {
        result += numeral.value
        break
      }
    }
  }

  return result
}

func arabToRom(input int) string {

  var romanNumerals = []struct{
    value int
    literal string
  } {
    {1000, "M"},
    {900, "CM"},
    {500, "D"},
    {400, "CD"},
    {100, "C"}, 
    {90, "XC"},
    {50, "L"},
    {40, "XL"},
    {10, "X"},
    {9, "IX"},
    {5, "V"},
    {4, "IV"},
    {1, "I"},
  }

  var result string

  for _, numeral := range romanNumerals {
    for input >= numeral.value {
      result += numeral.literal
      input -= numeral.value 
    }
  }

  return result

}

func calculate(input string) {

  parts := strings.Split(input, " ")
  a, op, b := parts[0], parts[1], parts[2]   

  var x, y int
  var err error

  if strings.ContainsAny(a, "IVXLCDM") {
    x = romToArab(a)
  } else {  
    x, err = strconv.Atoi(a)
    if err != nil {
      fmt.Println(err)
      return
    }
  }

  if strings.ContainsAny(b, "IVXLCDM") {
    y = romToArab(b)
  } else {
    y, err = strconv.Atoi(b)
    if err != nil {
      fmt.Println(err)
      return
    }
  }

  var result int
  switch op {
    case "+":
      result = x + y
    case "-":
      result = x - y 
    case "*":
      result = x * y
    case "/":
      result = x / y
  }

  romanResult := arabToRom(result)
  fmt.Println(romanResult)  

}

func main() {

  calculate("V + VII")
  calculate("14 + 3") 
  calculate("X / V")
  calculate("15 - 4")

}
