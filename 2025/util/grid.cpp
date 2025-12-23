#include <array>
#include <print>
#include <vector>

class Grid {
  public:
    void load_line(std::string line) { data.push_back(line); }
    int height() { return data.size(); }
    int width() { return data[0].size(); }
    Grid rotate_left() {
        Grid new_grid;
        for (int x = width() - 1; x >= 0; x--) {
            std::string new_line;
            for (int y = 0; y < height(); y++) {
                new_line += at(x, y);
            }

            new_grid.load_line(new_line);
        }

        return new_grid;
    };

    // removes the last char of every line and returns it as a string
    std::string pop_col() {
        std::string col;
        for (int y = 0; y < height(); y++) {
            std::string& line = data[y];
            col += line.back();
            line.pop_back();
        }

        return col;
    }

    char get(int x, int y) { return at(x, y); }
    char& operator[](int x, int y) { return data[y][x]; }
    // char operator[](int x, int y) { return at(x, y); }

    // Define begin() and end() member functions
    auto begin() { return data.begin(); }
    auto end() { return data.end(); }
    // Overloads for const objects (good practice)
    auto begin() const { return data.begin(); }
    auto end() const { return data.end(); }

    int number_surrounding(int x, int y, char c) {
        int result = 0;

        std::array<std::pair<int, int>, 8> directions = {{{0, -1},
                                                          {1, -1},
                                                          {1, 0},
                                                          {1, 1},
                                                          {0, 1},
                                                          {-1, 1},
                                                          {-1, 0},
                                                          {-1, -1}}};

        for (auto [dx, dy] : directions) {
            int nx = x + dx;
            int ny = y + dy;

            if (at(nx, ny) == c) {
                result++;
            }
        }

        return result;
    }

    void print() {
        std::println("{} x {} grid", width(), height());
        for (std::string line : data) {
            std::println("|{}|", line);
        }
        std::println();
    }

  private:
    std::vector<std::string> data;

    char at(int x, int y) {
        if (y < 0 || y >= height() || x < 0 || x >= width()) {
            return '0';
        }

        return data[y][x];
    }
};
