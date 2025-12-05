#include <fstream>
#include <iostream>
#include <ostream>

using namespace std;

int main() {
    if (ifstream file("../day01/input-1.txt"); file.is_open()) {
        int password = 0;
        int current = 50;

        string line;
        while (getline(file, line)) {
            const char direction = line[0];
            const int num = stoi(line.substr(1, line.length()));

            if (direction == 'L') {
                current -= num;
                if (current < 0) {
                    const int unit = current % 10;
                    const int tens = (current / 10) % 10;
                    const int rest = tens*10 + unit;
                    current = 100 + rest;
                }
                // TODO: understand better why this doesn't happen on 'R' ?
                if (current == 100) { current = 0; }
            } else if (direction == 'R') {
                current += num;
                if (current > 99) {
                    const int unit = current % 10;
                    const int tens = (current / 10) % 10;
                    const int rest = tens*10 + unit;
                    current = rest;
                }
            }

            if (current == 0) {
                password++;
            }
            cout << direction << " " << num << " " << current << endl;
        }

        cout << "password: " << password << endl;
    } else {
        cout << "Unable to open file!" << endl;
    }
    return 0;
}

