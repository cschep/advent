#include <cstdint>
#include <fstream>
#include <iostream>
#include <ostream>
#include <print>
#include <sstream>
#include <vector>

using namespace std;

/*
 * for 1 to the half the length of the string
 *   make a substring of that length, and a substring of the rest
 *   loop through rest in chunks of size
 *   	if they !eq the main chunk, bail on this id
 *   	if you get all the way through the loop store the id as invalid
 */

const bool part1 = false;

bool check_invalid(string id, int size) {
    if (part1 && id.length() % 2 != 0) {
        return false;
    }
    if (id.length() % size != 0) {
        return false;
    }

    string new_id = id.substr(0, size);
    string rest = id.substr(size, id.length());

    for (int i = 0; i < rest.length(); i += size) {
        string chunk = rest.substr(i, size);
        if (new_id != chunk) {
            return false;
        }
    }

    return true;
}

int main() {
    if (ifstream file("2/2.input"); file.is_open()) {
        string line;
        while (getline(file, line)) {
            vector<int64_t> invalid_ids;
            stringstream ss(line);
            string range;
            while (getline(ss, range, ',')) {
                size_t split_at = range.find('-');
                int64_t bottom = stoll(range.substr(0, split_at));
                int64_t top = stoll(range.substr(split_at + 1, range.length()));
                cout << range << endl;

                for (int64_t id = bottom; id <= top; id++) {
                    string id_str = to_string(id);
                    bool already_invalid = false;
                    for (int size = 1; !already_invalid && size <= id_str.length() / 2; size++) {
                        // println("checking {} with size {}", id_str, size);
                        if (check_invalid(id_str, size)) {
                            println("{} invalid!", id_str);
                            already_invalid = true;
                            invalid_ids.push_back(id);
                        }
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
        cout << "unable to open file!" << endl;
    }
}
