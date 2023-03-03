func maxRepeating(sequence string, word string) int {
    var sr strings.Builder
    cnt := 0
    sr.WriteString(word)
    for strings.Contains(sequence, sr.String()){
        cnt++
        sr.WriteString(word)
    }
    return cnt
}
