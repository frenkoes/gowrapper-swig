#ifdef  __cplusplus
#include "number.h"

Number::Number() {

}

Number::~Number() {

}

void Number::SetNumber(int x){
    this->x = x;
}

void Number::Print() {
    std::cout << this->x << std::endl;
}
#endif