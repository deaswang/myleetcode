func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    l1 := len(nums1)
    l2 := len(nums2)
    l := l1 + l2
    if l % 2 == 0{
        var a, b int
        ai := l/2 - 1
        bi := l/2
        for i, i1, i2 := 0, 0, 0 ; i <= bi; i++ {
            if i == ai {
                if i1 == l1 {
                    a = nums2[i2]
                    i2++
                } else if i2 == l2 {
                    a = nums1[i1]
                    i1++
                } else if nums1[i1] <= nums2[i2] {
                    a = nums1[i1]
                    i1 ++
                } else {
                    a = nums2[i2]
                    i2 ++
                }
            } else if i == bi {
                if i1 == l1 {
                    b = nums2[i2]
                    i2++
                } else if i2 == l2 {
                    b = nums1[i1]
                    i1++
                } else if nums1[i1] <= nums2[i2] {
                    b = nums1[i1]
                    i1 ++
                } else {
                    b = nums2[i2]
                    i2 ++
                }
            } else {
                if i1 == l1 {
                    i2++
                } else if i2 == l2 {
                    i1++
                } else if nums1[i1] <= nums2[i2] {
                    i1 ++
                } else {
                    i2 ++
                }
            }
        }
        return float64(a+b)/2
    } else {
        var n int
        ni := l/2
        for i, i1, i2 := 0, 0, 0 ; i <= ni; i++ {
            if i == ni {
                if i1 == l1 {
                    n = nums2[i2]
                    i2++
                } else if i2 == l2 {
                    n = nums1[i1]
                    i1++
                } else if nums1[i1] <= nums2[i2] {
                    n = nums1[i1]
                    i1 ++
                } else {
                    n = nums2[i2]
                    i2 ++
                }
            } else {
                if i1 == l1 {
                    i2++
                } else if i2 == l2 {
                    i1++
                } else if nums1[i1] <= nums2[i2] {
                    i1 ++
                } else {
                    i2 ++
                }
            }
        }
        return float64(n)
    }
}