type FreqCounter struct {
    digit [10]int
    lowerCase [26]int
    upperCase [26]int
}

func NewFreqCounter() *FreqCounter {
    return &FreqCounter{}
}

func (f *FreqCounter) Add(char byte){
    if char >= byte('a') && char <= byte('z') {
        l := int(char - 'a')
        f.lowerCase[l]++
    }else if char >= byte('A') && char <= byte('Z') {
        l := int(char - 'A')
        f.upperCase[l]++
    }else{
        l := int(char - '0')
        f.digit[l]++
    }
}

type Summary struct {
    char byte
    count int
}

func (f *FreqCounter) GetSummaries() []Summary {
    res := make([]Summary, 0)
    for i:=0;i<len(f.digit);i++{
        if f.digit[i] != 0 {
            res = append(res, Summary{byte(i)+byte('0'), f.digit[i]})
        }
    }

    for i:=0;i<len(f.lowerCase);i++{
        if f.lowerCase[i] != 0 {
            res = append(res, Summary{byte(i)+byte('a'), f.lowerCase[i]})
        }
    }

    for i:=0;i<len(f.upperCase);i++{
        if f.upperCase[i] != 0 {
            res = append(res, Summary{byte(i)+byte('A'), f.upperCase[i]})
        }
    }
    return res
}

func frequencySort(s string) string {
    counter := NewFreqCounter()
    for i:=0;i<len(s);i++{
        counter.Add(s[i])
    }

    summaries := counter.GetSummaries()
    sort.Slice(summaries, func(a, b int) bool {
        return summaries[a].count > summaries[b].count
    })

    res := make([]byte, 0)
    for _, summary := range summaries{
        for i:=0;i<summary.count;i++{
            res = append(res, summary.char)
        }
    }
    return string(res)

}