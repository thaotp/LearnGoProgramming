func factorial(n int32) int32 {
  var result int32 = 1
  for i:=n; i >= 1; i--{
    result = result*i
  }
  return result
}
