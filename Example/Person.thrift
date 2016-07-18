namespace swift MQ

include "Sex.thrift"

struct TPerson {
1:required string name
2:required TLanguage lang
3:required Sex.TSex sex
4:required list<string> strs
}

enum TLanguage {
    swift = 1
    java = 2
    objectc = 3
    golang = 4
}

enum TCode {
    swiftcode = 1
    javacode = 2
    objectc = 3
    golang = 4
}