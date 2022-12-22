# frozen_string_literal: true

def part1(file)
  result = {
    horizonal_position: 0,
    depth: 0
  }

  file.each_line do |line|
    command, val = line.split
    operator = command == 'up' ? :- : :+
    target = command == 'forward' ? :horizonal_position : :depth

    result[target] = result[target].public_send(operator, val.to_i)
  end

  puts result[:horizonal_position] * result[:depth]
end

Result = Struct.new(:horizonal_position, :depth, :aim)
def part2(file)
  result = Result.new(0, 0, 0)

  file.each_line do |line|
    command, val = line.split
    val = val.to_i

    case command
    when 'up'
      result.aim -= val
    when 'down'
      result.aim += val
    when 'forward'
      result.horizonal_position += val
      result.depth += (val * result.aim)
    end
  end

  puts result.horizonal_position * result.depth
end

File.open('2-input.txt') do |f|
  part2(f)
end
