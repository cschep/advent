#include <fstream>
#include <iostream>
#include <ostream>

using namespace std;

int main() {
    if (ifstream file("../day01/input-1.txt"); file.is_open()) {
        int password = 0;
        int current = 50;
        int prev = current;

        // part 2
        const bool password_method_0x434C49434B = true;

        string line;
        while (getline(file, line)) {
            const char direction = line[0];
            const int num = stoi(line.substr(1, line.length()));

            const int positive_hundreds = abs(num / 100);
            const int new_num = num - (positive_hundreds * 100);

            // remove and deal with the even hundreds
            if (password_method_0x434C49434B && positive_hundreds > 0) {
                cout << "adding " << positive_hundreds << " to passwords" << endl;
                password += positive_hundreds;
            }

            cout << "starting at " << current << " moving " << direction << " " << num << " / " << new_num << endl;

            // which direction are we going?
            prev = current;
            if (direction == 'L') {
                current -= new_num;
            } else if (direction == 'R') {
                current += new_num;
            }

            // going around right
            if (current > 99) {
                if (password_method_0x434C49434B && current != 100) {
                    cout << "passed 0 right" << endl;
                    password++;
                }
                current -= 100;
            }

            // going around left
            if (current < 0) {
                if (password_method_0x434C49434B && prev != 0) {
                    cout << "passed 0 left" << endl;
                    password++;
                }
                current = 100 + current;
            }

            // if we hit zero on the nose, add it
            if (current == 0) {
                cout << "landed on 0!, password++" << endl;
                password++;
            }
        }

        cout << "password: " << password << endl;
    } else {
        cout << "Unable to open file!" << endl;
    }

    return 0;
}
