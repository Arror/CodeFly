namespace swift TR

struct Customer {
    1: required string name
    2: optional Sex sex
    3: required list<string> hobbies
}

struct Result {
    1: required bool isSucceed
    2: required string message
}

enum Sex {
    male = 1
    female = 2
}

service Greet {
    Result SayHallo(1:Customer customer)
}