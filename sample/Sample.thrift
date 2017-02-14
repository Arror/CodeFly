namespace swift TR

struct Customer {
    1: string name
    2: Sex sex
    3: list<string> hobbies
}

struct Result {
    1: bool isSucceed
    2: string message
}

enum Sex {
    male = 1
    female = 2
}

service Greet {
    Result SayHallo(1:Customer customer)
}