def part1(file)
  total = 0
  file.each_line do |line|
    l, w, h = line.split('x')
    l = l.to_i
    w = w.to_i
    h = h.to_i

    side1 = l * w
    side2 = w * h
    side3 = l * h

    smallest = [side1, side2, side3].min

    total += (2 * side1) + (2 * side2) + (2 * side3) + smallest
  end
  puts total
end

File.open('2-input.txt') do |f|
  part1(f)
end
