def merge(nums1, m, nums2, n)
    len = m + n;
    current_index = len - 1
  
    m -= 1
    n -= 1
  
    while current_index >= 0 do
      if (n < 0 || (nums1[m] > nums2[n] && m >= 0))
        p nums1[m]
        nums1[current_index] = nums1[m]
        m -= 1
      else
        p nums2[n]
        nums1[current_index] = nums2[n]
        n -= 1
      end
  
      current_index -= 1
    end
  p nums1
end

merge([1, 2, 3], 3, [], 0)