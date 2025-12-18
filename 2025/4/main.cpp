#include <fstream>
#include <print>
#include <string>

#include "../util/grid.cpp"

int main() {
    if (std::ifstream file("4/4.input"); file.is_open()) {
        Grid grid;

        std::string line;
        while (getline(file, line)) {
            grid.load_line(line);
        }

        int result = 0;
        for (int y = 0; y < grid.height(); y++) {
            for (int x = 0; x < grid.width(); x++) {
                if (grid[x, y] == '@') {
                    int paper_rolls = grid.number_surrounding(x, y, '@');
                    if (paper_rolls < 4) {
                        result++;
                    }
                }
            }
        }

        std::println("result: {}", result);
    } else {
        std::println("Unable to open file!");
    }
}
