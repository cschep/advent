#include <fstream>
#include <iostream>
#include <ostream>
#include <print>
#include <string>

int max_joltage(std::string bank) {
    if (bank.length() < 2) {
        return 0;
    }
    // find highest with at least one digit after it
    //
    int highest = -1;
    int highest_index = -1;
    for (int i = 0; i < bank.length() - 1; i++) {
        int current = bank[i] - '0';
        if (current > highest) {
            std::println("new highest: {}", current);
            highest = current;
            highest_index = i;
        }
    }

    // now that we have the highest find the second highest after it?
    //
    int next_highest = -1;
    for (int i = highest_index + 1; i < bank.length(); i++) {
        int current = bank[i] - '0';
        if (current > next_highest) {
            std::println("new next highest: {}", current);
            next_highest = current;
        }
    }

    int joltage = std::stoi(std::format("{}{}", highest, next_highest));

    std::println("{}", joltage);

    return joltage;
}

int main() {
    int result = 0;
    if (std::ifstream file("3/3.input"); file.is_open()) {
        std::string line;
        while (getline(file, line)) {
            std::cout << line << std::endl;
            result += max_joltage(line);
        }
    } else {
        std::cout << "Unable to open file!" << std::endl;
    }

    std::println("result: {}", result);

    return 0;
}
