#include <cstdint>
#include <fstream>
#include <iostream>
#include <ostream>
#include <sstream>

using namespace std;

int main() {
    if (ifstream file("2/2.input.small"); file.is_open()) {
        string line;
        while (getline(file, line)) {
            vector<int64_t> invalid_ids;
            stringstream ss(line);
            string range;
            while (getline(ss, range, ',')) {
                size_t split_at = range.find('-');
                int64_t bottom = stoll(range.substr(0, split_at));
                int64_t top = stoll(range.substr(split_at + 1, range.length()));
                // cout << range << endl;

                for (int64_t id = bottom; id <= top; id++) {
                    string str = to_string(id);
                    string first_half = str.substr(0, str.length() / 2);
                    string second_half = str.substr(str.length() / 2, str.length());
                    if (first_half == second_half) {
                        invalid_ids.push_back(id);
                    }
                }
            }

            int64_t total = 0;
            for (const int64_t id : invalid_ids) {
                total += id;
            }
            cout << "total: " << total << endl;
        }
    } else {
        cout << "Unable to open file!" << endl;
    }
}
