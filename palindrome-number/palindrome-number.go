func isPalindrome(x int) bool {
    if x>=0&&x==Decode(x,0){
        return true
    }
    return false
}

func Decode(x int, s int) int {
    if x<10{
        return s*10+x
    }else{
        temp:=x%10
        return Decode(x/10,s*10+temp)
    }
}