#include <fstream>
#include <print>
#include <string>

#include "../util/grid.cpp"

int removal_pass(Grid *grid) {

    std::vector<std::pair<int, int>> to_remove;

    for (int y = 0; y < grid->height(); y++) {
        for (int x = 0; x < grid->width(); x++) {
            if (grid->get(x, y) == '@') {
                int paper_rolls = grid->number_surrounding(x, y, '@');
                if (paper_rolls < 4) {
                    to_remove.push_back({x, y});
                }
            }
        }
    }

    for (auto [x, y] : to_remove) {
        (*grid)[x, y] = 'x';
    }

    return to_remove.size();
}

int main() {
    if (std::ifstream file("4/4.input"); file.is_open()) {
        Grid grid;

        std::string line;
        while (getline(file, line)) {
            grid.load_line(line);
        }

        // PART 2
        grid.print();

        int result = 0;
        int curr = 1;
        while (curr > 0) {
            curr = removal_pass(&grid);
            result += curr;
        }
        grid.print();

        std::println("result: {}", result);
    } else {
        std::println("Unable to open file!");
    }
}
