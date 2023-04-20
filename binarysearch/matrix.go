func searchMatrix(matrix [][]int, target int) bool {
	row := 0
    low := 0 
    high := len(matrix) - 1
    for (low <= high) {
        mid := low + (high - low) / 2
        if target < matrix[mid][0] {
            high = mid - 1
        } else if target > matrix[mid][len(matrix[0])-1] {
            low = mid + 1
        } else {
            row = mid
            break
        }
    }

    low = 0 
    high = len(matrix[0])-1
    for (low <= high) {
        mid := low + (high - low) / 2
        if target < matrix[row][mid] {
            high = mid - 1
        } else if target > matrix[row][mid] {
            low = mid + 1
        } else {
            return true
        }
    }

    return false
}