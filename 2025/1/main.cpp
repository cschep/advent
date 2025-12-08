#include <fstream>
#include <iostream>
#include <ostream>

int main() {
    if (std::ifstream file("1/1.input"); file.is_open()) {
        int password = 0;
        int current = 50;
        int prev = current;

        // part 2
        const bool password_method_0x434C49434B = true;

        std::string line;
        while (getline(file, line)) {
            const char direction = line[0];
            const int num = stoi(line.substr(1, line.length()));

            const int positive_hundreds = abs(num / 100);
            const int new_num = num - (positive_hundreds * 100);

            // remove and deal with the even hundreds
            if (password_method_0x434C49434B && positive_hundreds > 0) {
                std::cout << "adding " << positive_hundreds << " to passwords" << std::endl;
                password += positive_hundreds;
            }

            std::cout << "starting at " << current << " moving " << direction << " " << num << " / " << new_num
                      << std::endl;

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
                    std::cout << "passed 0 right" << std::endl;
                    password++;
                }
                current -= 100;
            }

            // going around left
            if (current < 0) {
                if (password_method_0x434C49434B && prev != 0) {
                    std::cout << "passed 0 left" << std::endl;
                    password++;
                }
                current = 100 + current;
            }

            // if we hit zero on the nose, add it
            if (current == 0) {
                std::cout << "landed on 0!, password++" << std::endl;
                password++;
            }
        }

        std::cout << "password: " << password << std::endl;
    } else {
        std::cout << "Unable to open file!" << std::endl;
    }

    return 0;
}
