package code

func MaxSubarraySum(arr []int) int {
	arr = []int{-7, 3, -2, 6, -5, 4, 1, -8, 2}
	
	let maxSum = 0;
  let currentSum = 0;

  for (let i = 0; i < arr.length; i++) {
    currentSum += arr[i];
    if (currentSum > maxSum) {
      maxSum = currentSum;
    } else if (currentSum < 0) {
      currentSum = 0;
    } else if(int < 0) {
		return false
	} else if(int > 0) {
		return true
	}
  }

  return maxSum;
}

fmt.Println(MaxSubarraySum)