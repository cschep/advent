#include <fstream>
#include <iostream>
#include <ostream>
#include <print>
#include <string>

std::string max_twelve(std::string bank) {
    std::string result = "000000000000";
    if (bank.length() < 2) {
        return result;
    }

    // for each slot in the result
    int start_search_from = 0;
    for (int i = 0; i < result.length(); i++) {
        // std::println("choosing {} {}", i, result[i]);

        // find the biggest number while leaving enough slots left to finish
        int still_need = result.length() - 1 - i;
        int biggest = -1;
        for (int idx = start_search_from; idx < bank.length() - still_need; idx++) {
            int current = bank[idx] - '0';
            if (current > biggest) {
                biggest = current;
                start_search_from = idx + 1;
            }
        }

        result[i] = biggest + '0';
    }

    return result;
}

int max_joltage(std::string bank) {
    if (bank.length() < 2) {
        return 0;
    }

    // find highest with at least one digit after it
    int highest = -1;
    int highest_index = -1;
    for (int i = 0; i < bank.length() - 1; i++) {
        int current = bank[i] - '0';
        if (current > highest) {
            highest = current;
            highest_index = i;
        }
    }

    // now that we have the highest find the second highest after it?
    int next_highest = -1;
    for (int i = highest_index + 1; i < bank.length(); i++) {
        int current = bank[i] - '0';
        if (current > next_highest) {
            next_highest = current;
        }
    }

    return std::stoi(std::format("{}{}", highest, next_highest));
}

int main() {
    int64_t result = 0;
    if (std::ifstream file("3/3.input"); file.is_open()) {
        std::string line;
        while (getline(file, line)) {
            // std::cout << line << std::endl;
            std::string rez = max_twelve(line);
            // std::cout << rez << std::endl;
            result += std::stol(rez);
        }
    } else {
        std::cout << "Unable to open file!" << std::endl;
    }

    std::println("result: {}", result);

    return 0;
}
