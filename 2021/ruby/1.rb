# frozen_string_literal: true

def part1(file)
  count = 0
  file.each_cons(2) do |pair|
    first = pair[0].strip.to_i
    second = pair[1].strip.to_i

    count += 1 if second > first
  end
  puts count
end

def part2(file)
  count = 0
  previous_sum = 0
  file.each_cons(3) do |pair|
    sum = pair[0].strip.to_i + pair[1].strip.to_i + pair[2].strip.to_i
    count += 1 if sum > previous_sum && !previous_sum.zero?
    previous_sum = sum
  end

  puts count
end

File.open('1-input.txt') do |f|
  part1(f)
  f.rewind
  part2(f)
end
