#include <fstream>
#include <iostream>
#include <ostream>
#include <string>

int main() {
    if (std::ifstream file("1.input.small"); file.is_open()) {
		std::string line;
        while (getline(file, line)) {
			std::cout << line << std::endl;
		}
    } else {
		std::cout << "Unable to open file!" << std::endl;
    }

    return 0;
}
