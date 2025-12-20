#include <fstream>
#include <print>
#include <string>
#include <vector>

int main() {
    if (std::ifstream file("5/5.input"); file.is_open()) {

        std::string line;
        bool parse_ranges = true;
        std::vector<std::pair<int64_t, int64_t>> ranges;
        int fresh_count = 0;
        while (getline(file, line)) {
            if (line == "") {
                parse_ranges = false;
                continue;
            }

            if (parse_ranges) {
                char delimiter = '-';
                size_t pos = line.find(delimiter);
                int64_t first = std::stol(line.substr(0, pos)); // Extract the part before the delimiter
                int64_t second = std::stol(
                    line.substr(pos + 1, line.length())); // Erase the token and delimiter from the original string

                if (first > second) {
                    std::println("first > second {} {}", first, second);
                    std::abort();
                }

                ranges.push_back({first, second});
            } else {
                int64_t id = std::stol(line);

                for (auto range : ranges) {
                    if (id >= range.first && id <= range.second) {
                        std::println("fresh! {} in {}", id, range);
                        fresh_count++;
                        break;
                    }
                }
            }
        }

        std::println("fresh count: {}", fresh_count);

    } else {
        std::println("Unable to open file!");
    }
}
