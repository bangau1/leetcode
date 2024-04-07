// ((*()) -> true
// (()*) -> true
// (*)) -> true
// (()* -> true
// *(()))*
// **((
func checkValidString(s string) bool {
    balance := 0
    wildCard := 0

    for i:=0;i<len(s);i++{
        if s[i] == '(' {
            balance+=1
        }else if s[i] == ')' {
            if balance > 0 {
                balance -= 1
            }else if wildCard > 0 {
                wildCard -= 1
            }else{
                return false
            }
        }else if s[i] == '*' {
            wildCard += 1
        }
    }
    if balance > 0 && balance > wildCard {
        return false
    }

    balance = 0
    wildCard = 0

    for i:=len(s)-1;i>=0;i--{
        if s[i] == ')' {
            balance+=1
        }else if s[i] == '(' {
            if balance > 0 {
                balance -= 1
            }else if wildCard > 0 {
                wildCard -= 1
            }else{
                return false
            }
        }else if s[i] == '*' {
            wildCard += 1
        }
    }
    if balance > 0 && balance > wildCard {
        return false
    }
    return true
}