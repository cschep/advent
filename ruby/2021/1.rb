# frozen_string_literal: true

File.open('1-input.txt') do |f|
  count = 0
  f.each_cons(2) do |pair|
    first = pair[0].strip.to_i
    second = pair[1].strip.to_i

    count += 1 if second > first
  end
  puts count
end
