#include "../util/grid.cpp"
#include <fstream>
#include <print>
#include <string>

void part2() {}

int main() {
    if (std::ifstream file("7/7.input"); file.is_open()) {
        int total_splits = 0;
        Grid grid;
        std::string line;
        while (getline(file, line)) {
            grid.load_line(line);
        }

        grid.print();

        for (int y = 1; y < grid.height(); y++) {
            for (int x = 0; x < grid.width(); x++) {
                auto above = grid[x, y - 1];
                auto current = grid[x, y];
                if (above == 'S' || above == '|') {
                    if (current == '^') {
                        grid[x - 1, y] = '|';
                        grid[x + 1, y] = '|';
                        total_splits++;
                    } else {
                        grid[x, y] = '|';
                    }
                }
            }
        }

        grid.print();

        std::println("total result: {}", total_splits);
    } else {
        std::println("Unable to open file!");
    }
}
