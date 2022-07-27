# @param {String} s
# @return {Integer}
def calculate(s)
  tokens = []
  ops = []
  number = ""
  s << ' '
  need_fit_zero = true
  s.each_char do |c|
    if c <= '9' && '0' <= c
      number << c
      need_fit_zero = false
    elsif !number.empty?
      tokens << number 
      number = ''
    end

    if c == "("
      ops << c
      next
    end

    if c == ')'
      # while ops.last != '('
      #   op = ops.pop
      #   p op
      #   p ops
      #   p tokens
      #   tokens << op
      # end
      # p ops
      # ops.pop

      loop
        op = ops.pop
        p op
        p ops
        p tokens
        break if op == '('
        tokens << ops.pop
      end

      p ops

      next
    end



    if get_op_rank(c) > 0
      tokens << '0' if need_fit_zero
      need_fit_zero = true

      while !ops.empty? && get_op_rank(c) <= get_op_rank(ops.last)
        tokens << ops.pop
      end
      ops.push(c)
    end
  end

  ops.reverse_each do |op|
    tokens << op
  end

  eval_rpn(tokens)
end

def get_op_rank(op)
  case op
  when '*', '/'
    2
  when '+', '-'
    1
  else
    0
  end
end

def eval_rpn(tokens)
  operands = ['+', '-', '*', '/']
  operation_stacks = []

  tokens.each do |token|
    if !operands.include?(token)
      operation_stacks.push(token.to_i)
    else
      first, second = operation_stacks.pop(2)
      result =  case token
                when '+'
                  first + second
                when '-'
                  first - second
                when '*'
                  first * second
                when '/'
                  (first.to_f / second).to_i
                end
      operation_stacks.push(result)
    end
  end

  operation_stacks[0]
end

p calculate("(1+(4+5+2)-3)+(6+8)")
