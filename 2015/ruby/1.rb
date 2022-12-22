# frozen_string_literal: true

def part1
  input = File.open('1-input.txt').read
  result = 0
  input.each_char do |c|
    case c
    when '('
      result += 1
    when ')'
      result -= 1
    end
  end

  puts result
end

def part2
  input = File.open('1-input.txt').read
  result = 0
  input.each_char.with_index(1) do |c, i|
    case c
    when '('
      result += 1
    when ')'
      result -= 1
    end

    if result.negative?
      puts 'basement at:', i
      break
    end
  end
end

part2
