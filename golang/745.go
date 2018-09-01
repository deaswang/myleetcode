type WordFilter struct {
    Words []string
    Length int
}


func Constructor(words []string) WordFilter {
    var ret = WordFilter{}
    ret.Words = words
    ret.Length = len(words)
    return ret
}


func (this *WordFilter) F(prefix string, suffix string) int {
    for i:= this.Length-1;i>=0;i-- {
        if strings.HasPrefix(this.Words[i], prefix) && strings.HasSuffix(this.Words[i], suffix) {
            return i
        }
    }
    return -1
}


/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
 