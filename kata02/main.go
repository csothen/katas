package kata02

func Execute(target int, arr []int) int {
	left := 0
	right := len(arr) - 1
	return binarySearch(arr, left, right, target)
}

func binarySearch(arr []int, left, right, target int) int {
	// Check if the pointers crossed
	if right < left {
		return -1
	}

	// Calculate the middle of the section
	mid := left + (right-left)/2

	// Check if the target was found
	if arr[mid] == target {
		return mid
	}

	// If the element we are in is larger than
	// the target then the target is to the left
	if arr[mid] > target {
		return binarySearch(arr, left, mid-1, target)
	}

	// Otherwise it is to the right
	return binarySearch(arr, mid+1, right, target)
}
