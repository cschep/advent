#include "../util/grid.cpp"
#include <algorithm>
#include <charconv>
#include <fstream>
#include <print>
#include <sstream>
#include <string>
#include <vector>

void part1() {
    if (std::ifstream file("6/6.input.small"); file.is_open()) {
        std::string line;
        std::vector<std::vector<int>> grid;
        std::vector<std::string> ops;
        while (getline(file, line)) {
            std::stringstream ss(line);
            std::vector<int> numbers;
            std::string token;
            while (ss >> token) {
                int num;
                auto [ptr, ec] = std::from_chars(
                    token.data(), token.data() + token.size(), num);
                if (ec == std::errc{} && ptr == token.data() + token.size()) {
                    numbers.push_back(num);
                } else {
                    ops.push_back(token);
                }
            }
            if (numbers.size() > 0) {
                grid.push_back(numbers);
            }
        }

        int64_t result = 0;
        std::println("{} {} {}", grid[0][1], grid[0].size(), ops.size());
        for (int i = 0; i < grid[0].size(); i++) {
            std::vector<int> operands;
            for (int j = 0; j < grid.size(); j++) {
                int number = grid[j][i];
                operands.push_back(number);
            }

            int64_t col_result = 0;
            if (ops[i] == "*") {
                col_result = 1;
            }
            for (const int number : operands) {
                if (ops[i] == "+") {
                    col_result += number;
                } else if (ops[i] == "*") {
                    col_result *= number;
                }
            }
            // std::println("{} {} {}", operands, ops[i], col_result);
            result += col_result;
        }

        std::println("{}", result);
    } else {
        std::println("Unable to open file!");
    }
}

void part2() {
    if (std::ifstream file("6/6.input"); file.is_open()) {
        Grid grid;
        std::string line;
        while (getline(file, line)) {
            grid.load_line(line);
        }

        grid.print();
        std::println();

        Grid new_grid = grid.rotate_left();
        new_grid.load_line(" ");
        new_grid.print();

        std::string op_str = new_grid.pop_col();
        new_grid.print();

        std::stringstream ss(op_str);
        std::string token;
        std::vector<std::string> ops;
        while (ss >> token) {
            ops.push_back(token);
        }
        std::println("operators: {}", ops);

        int64_t total_result = 0;

        int64_t current_result = 0;
        int op_idx = 0;
        std::string current_operator = ops[op_idx];
        if (current_operator == "*") {
            current_result = 1;
        }

        for (const std::string& line : new_grid) {
            if (std::ranges::all_of(line, ::isspace)) {
                std::println("{} ing makes {}", current_operator,
                             current_result);
                total_result += current_result;

                current_result = 0;
                op_idx++;
                current_operator = ops[op_idx];
                if (current_operator == "*") {
                    current_result = 1;
                }
                std::println();

                continue;
            }

            int64_t number = std::stol(line);
            std::println("{}", number);
            if (current_operator == "*") {
                current_result *= number;
            } else {
                current_result += number;
            }
        }

        std::println("total result: {}", total_result);
    } else {
        std::println("Unable to open file!");
    }
}

int main() {
    part1();
    part2();
}
