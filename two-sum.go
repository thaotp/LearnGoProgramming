func twoSum(nums []int, target int) []int {
  for i, num := range nums {
    for j:=i+1; j <= len(nums) - 1; j++ {
      if(num + nums[j] == target){
        return []int{i, j}
      }
    }
  }
  return []int{0,0}
}
