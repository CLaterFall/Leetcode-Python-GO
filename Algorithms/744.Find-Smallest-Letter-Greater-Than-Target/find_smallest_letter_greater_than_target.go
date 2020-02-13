func nextGreatestLetter(letters []byte, target byte) byte {
    length := len(letters)
    start := 0
    end := length - 1
    for start <= end {
        mid := (start + end) / 2
        if target < letters[mid] {
            end = mid -1
        } else {
            start = mid + 1
        }
    }
    if start >= length {
        return letters[0]
    } else {
        return letters[start]
    }
}
