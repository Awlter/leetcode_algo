class ListNode
    attr_accessor :val, :next
    def initialize(val = 0, _next = nil)
        @val = val
        @next = _next
    end
end

def reverse_k_group(head, k)
  protected_tail = ListNode.new(nil, head)
  previous_tail = protected_tail
  loop do
    end_node = get_end_node(head, k)
    break if !end_node
    next_head = end_node.next
    reversed_head = reverse_list(head, k)
    head.next = next_head
    previous_tail.next = reversed_head
    previous_tail = head
    head = next_head
    break if !head
  end

  protected_tail.next
end

def get_end_node(head, k)
1.upto(k-1).each do
  head = head.next
  return nil if !head
end

head
end

def reverse_list(head, k)
  previous_node = nil
  1.upto(k).each do
    next_head = head.next
    head.next = previous_node
    previous_node = head
    head = next_head
  end
  previous_node
end

n5 =   ListNode.new(5, nil)

n4 =   ListNode.new(4, n5)

# n3 =   ListNode.new(3, n4)

# n2 =   ListNode.new(2, n3)

# n1 =   ListNode.new(1, n2)

p reverse_k_group(n4, 2)