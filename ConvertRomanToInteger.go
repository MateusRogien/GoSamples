func romanToInt(s string) int {
    
    //map values
    m := make(map[string]int)
    m["I"]=1
    m["V"]=5
    m["X"]=10
    m["L"]=50
    m["C"]=100
    m["D"]=500
    m["M"]=1000
    var solution = 0
    for i, c := range s {
        if i == 0{
            solution = solution + m[string(c)]
        }else if (string(c) == "V" || string(c) == "X" ) && (string(s[i-1]) == "I") {
            solution = solution - 2 +  m[string(c)] 
        }else if ((string(c) == "L" || string(c) == "C" ) && (string(s[i-1]) == "X")) {
            solution = solution - 20 +  m[string(c)] 
        }else if (string(c) == "D" || string(c) == "M" ) && (string(s[i-1]) == "C") {
            solution = solution - 200 +  m[string(c)] 
        }else{
            solution = solution + m[string(c)] 
        }
                 
        //test
        fmt.Printf("character %c starts at byte position %d\n", c, i)
    }
    
    fmt.Printf("Solution:%d",solution)
 return solution   
}
