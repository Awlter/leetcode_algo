func plusOne(digits []int) []int {
  promote := 1
  for i := len(digits) - 1; i >=0 && promote == 1; i-- {
    if digits[i] < 9 {
      promote = 0
      digits[i]++
    } else {
      digits[i] = 0
    }
  }

  if promote == 1 {
    return append([]int{1}, digits...)
  }

  return digits
}