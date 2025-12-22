#include <algorithm>
#include <fstream>
#include <print>
#include <string>
#include <vector>

// using Range = std::pair<int64_t, int64_t>;
class Range {
  public:
    int64_t first;
    int64_t second;
    bool merged = false;

    bool contains(int64_t id) { return id >= first && id <= second; }

    bool operator==(Range rhs) {
        return this->first == rhs.first && this->second == rhs.second;
    };
    //
    // TODO: wtf did we need const at the end?
    bool operator<(Range rhs) const { return this->first < rhs.first; }
};

// TODO: what is going on here
template <> struct std::formatter<Range> {
    constexpr auto parse(std::format_parse_context& ctx) { return ctx.begin(); }
    auto format(const Range& range, std::format_context& ctx) const {
        return std::format_to(ctx.out(), "{{{}, {}}}", range.first,
                              range.second);
    }
};

int main() {
    if (std::ifstream file("5/5.input"); file.is_open()) {
        std::string line;
        bool parse_ranges = true;
        std::vector<Range> ranges;
        int fresh_count = 0;
        while (getline(file, line)) {
            if (line == "") {
                parse_ranges = false;
                continue;
            }

            if (parse_ranges) {
                char delimiter = '-';
                size_t pos = line.find(delimiter);
                int64_t first = std::stoll(line.substr(
                    0, pos)); // Extract the part before the delimiter
                int64_t second = std::stoll(line.substr(
                    pos + 1, line.length())); // Erase the token and delimiter
                                              // from the original string

                ranges.push_back({first, second});
            } else {
                int64_t id = std::stoll(line);

                for (auto range : ranges) {
                    if (range.contains(id)) {
                        std::println("fresh! {} in {}", id, range);
                        fresh_count++;
                        break;
                    }
                }
            }
        }

        std::sort(ranges.begin(), ranges.end());

        std::println("sorted ranges!");
        for (auto range : ranges) {
            std::println("{}", range);
        }

        // PART 2
        // collapse ranges
        for (int i = 0; i < ranges.size(); i++) {
            Range& range = ranges[i];
            for (int j = i + 1; j < ranges.size(); j++) {
                Range& range2 = ranges[j];
                if (range.merged || range2.merged) {
                    std::println("skipping {} {}", range, range2);
                    continue;
                }
                std::println("range {} range2 {}", range, range2);

                if (range2.first <= range.second &&
                    range2.second > range.second) {
                    range.second = range2.second;
                    std::println("bumping to {}", range2.second);
                    range2.merged = true;
                } else if (range2.second <= range.second) {
                    std::println("range fully covered {}", range2);
                    range2.merged = true;
                }
            }
        }

        std::println("collapsed ranges!");
        int64_t total_fresh = 0;
        for (auto range : ranges) {
            if (range.merged) {
                continue;
            }
            int64_t diff = range.second - range.first + 1;
            std::println("{} {}", range, diff);
            if (diff < 0) {
                abort();
            }
            total_fresh += diff;
        }

        std::println("fresh count: {}", fresh_count);
        std::println("total fresh: {}", total_fresh);
    } else {
        std::println("Unable to open file!");
    }
}
